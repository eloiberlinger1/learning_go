package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	var prettyLogs = flag.Bool("pretty-logs", false, "use --pretty-logs to disable slog package")
	flag.Parse()

	app := application{
		config: cfg,
	}

	if !*prettyLogs {
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

		slog.SetDefault(logger)
	} else {
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey || a.Key == slog.LevelKey {
					return slog.Attr{}
				}
				return a
			},
		}))
		slog.SetDefault(logger)
	}

	h := app.mount()
	err := app.run(h)
	if err != nil {
		log.Println("Server has failed to start")
		os.Exit(1)
	}
}
