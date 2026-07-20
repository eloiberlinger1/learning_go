package products_test

import (
	"context"
	"testing"

	repo "ecom-local/internal/adapters/postgresql/sqlc"
	"ecom-local/internal/products"
)

// 1. Définition de notre FAUX Repo (le Mock)
type mockRepo struct {
	// On "embarque" l'interface SQLC pour ne pas avoir à écrire
	// toutes les autres méthodes (comme CreateProduct etc...)
	repo.Querier
}

// 2. On surcharge UNIQUEMENT la fonction dont on a besoin pour ce test
func (m *mockRepo) ListProductById(ctx context.Context, id int64) (repo.Product, error) {
	// Au lieu de taper dans la BDD, on invente un faux produit
	// On lui donne le même ID que celui qui a été demandé en paramètre
	return repo.Product{
		ID:           id,
		Name:         "Produit de test",
		PriceInCents: 9900, // 99.00
		Quantity:     10,
	}, nil
}

// 3. Voici le test officiel
func TestListProductById(t *testing.T) {
	// Étape 1 : On crée une instance de notre faux repo
	fauxRepo := &mockRepo{}

	// Étape 2 : L'injection magique !
	// On passe notre faux repo au constructeur de ton VRAI service
	// Le service n'y verra que du feu.
	vraiService := products.NewService(fauxRepo)

	// Étape 3 : On exécute la logique
	var expectedID int64 = 42
	produit, err := vraiService.ListProductById(context.Background(), expectedID)

	// Étape 4 : On vérifie que tout s'est bien passé
	if err != nil {
		t.Fatalf("On s'attendait à un succès, mais on a reçu l'erreur : %v", err)
	}

	if produit.ID != expectedID {
		t.Errorf("On s'attendait à l'ID %d, mais on a reçu %d", expectedID, produit.ID)
	}

	if produit.Name != "Produit de test" {
		t.Errorf("On s'attendait au nom 'Produit de test', mais on a reçu '%s'", produit.Name)
	}
}
