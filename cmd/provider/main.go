// cmd/provider/main.go
package main

import (
	"log"
	"net/http"
	"os"

	postgre "github.com/Microservices-Gopher-Hex/provider-service/adapter/repository/postgres"
	postgreentity "github.com/Microservices-Gopher-Hex/provider-service/adapter/repository/postgres/entity"
	httpHandler "github.com/Microservices-Gopher-Hex/provider-service/api/http/handler"
	appProvider "github.com/Microservices-Gopher-Hex/provider-service/application/provider"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	port := getenv("HTTP_PORT", "8082")
	dsn := mustEnv("DB_DSN")

	// Conexión GORM
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migración de la entidad de persistencia (solo dev)
	if err := db.AutoMigrate(&postgreentity.Provider{}); err != nil {
		log.Fatal(err)
	}

	// Wiring hexagonal
	repository := postgre.New(db)
	createUC := &appProvider.CreateUseCase{Repo: repository}
	listUC := &appProvider.ListUseCase{Repo: repository}
	h := &httpHandler.ProviderHandler{Create: createUC, List: listUC}

	// Router
	r := chi.NewRouter()
	r.Get("/health", h.Health)
	r.Post("/providers", h.Post)
	r.Get("/providers", h.GetList)

	log.Printf("provider-service :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func mustEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("missing env %s", k)
	}
	return v
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
