package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"gitlab.com/pedrojhrossi/golang-pa/config"
	"gitlab.com/pedrojhrossi/golang-pa/internal/adapters/handler"
	"gitlab.com/pedrojhrossi/golang-pa/internal/adapters/repository"
	"gitlab.com/pedrojhrossi/golang-pa/internal/core/services"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	cfg := config.LoadConfig()

	// Connecto to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	tenantRepo := repository.NewMongoTenantRepository(client, cfg.DBName)

	tenantService := services.NewTenantService(tenantRepo)

	tenantHandler := handler.NewTenantHandler(tenantService)

	r := chi.NewRouter()
	r.Post("/tenants", tenantHandler.Create)

	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
