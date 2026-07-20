package main

import (
	"log"
	"log/slog"
	"net/http"
	"time"

	repo "ecom-local/internal/adapters/postgresql/sqlc"
	customMiddleware "ecom-local/internal/middleware"
	"ecom-local/internal/services/orders"
	"ecom-local/internal/services/products"
	"ecom-local/internal/services/users"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	slogchi "github.com/samber/slog-chi"
)

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

type application struct {
	config config
	// logger
	db *pgx.Conn
}

func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(slogchi.New(slog.Default()))
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Rate limiting middleware
	r.Use(customMiddleware.RateLimiter)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All good"))
	})

	productsService := products.NewService(repo.New(app.db))
	productsHandler := products.NewHandler(productsService)

	r.Get("/products", productsHandler.ListProducts)
	r.Get("/product/{id}", productsHandler.ListProductById)

	ordersHandler := orders.NewHandler(nil)
	r.Group(func(r chi.Router) {
		r.Use(customMiddleware.AuthMiddleware)
		r.Post("/orders", ordersHandler.CreateOrder)
	})

	userStore := users.NewStore(app.db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(r)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at addr %s", app.config.addr)

	return srv.ListenAndServe()
}
