package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	db := client.Database(cfg.DBName)

	//** Initialize Repositories (Secondary Adapters)
	tenantRepo := repository.NewMongoTenantRepository(db)
	appointmentRepo := repository.NewMongoAppointmentRepository(db)

	//** Initialize Services (Core logic)
	tenantService := services.NewTenantService(tenantRepo)
	appointmentService := services.NewAppointmentService(appointmentRepo)

	//** Initialize Handlers (Primary Adapters)
	tenantHandler := handler.NewTenantHandler(tenantService)
	appointmentHandler := handler.NewAppointmentHandler(appointmentService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/tenants", func(r chi.Router) {
		r.Post("/", tenantHandler.Create)
		r.Get("/", tenantHandler.List)
		r.Get("/{id}", tenantHandler.Get)
		r.Delete("/{id}", tenantHandler.Delete)

		r.Route("/{tenantID}/appointments", func(r chi.Router) {
			r.Use(handler.TenantContext(tenantService))

			r.Post("/", appointmentHandler.Schedule)
			r.Get("/", appointmentHandler.List)
			r.Get("/{id}", appointmentHandler.GetByID)
			r.Delete("/{id}/cancel", appointmentHandler.Cancel)
		})
	})

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
