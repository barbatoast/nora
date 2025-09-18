package main

import (
	"log"
	"net/http"

	"vapi-dashboard/server/config"
	"vapi-dashboard/server/internal/db"
	"vapi-dashboard/server/internal/handlers"
	"vapi-dashboard/server/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Env()

	dbconn, err := db.Open(cfg.DSN)
	if err != nil { log.Fatal(err) }
	if err := db.Migrate(dbconn); err != nil { log.Fatal(err) }

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: cfg.CORSOrigins,
		AllowedMethods: []string{"GET","POST","PUT","PATCH","DELETE","OPTIONS"},
		AllowedHeaders: []string{"Authorization","Content-Type"},
		AllowCredentials: true,
	}))

	ah := handlers.AuthHandler{DB: dbconn, Cfg: cfg}
	fh := handlers.FirmHandler{DB: dbconn}
	ph := handlers.ProviderHandler{DB: dbconn}
	asth := handlers.AssistantHandler{DB: dbconn}
	nh := handlers.NumberHandler{DB: dbconn}

	r.Post("/api/login", ah.Login)

	r.Group(func(pr chi.Router) {
		pr.Use(middleware.Auth(cfg, "admin", "manager"))
		pr.Get("/api/firms", fh.List)
		pr.Get("/api/providers", ph.List)
		pr.Post("/api/providers", ph.Create)
		pr.Get("/api/assistants", asth.List)
		pr.Post("/api/assistants", asth.Create)
		pr.Get("/api/numbers", nh.List)
		pr.Post("/api/numbers", nh.Create)
	})

	port := ":" + cfg.Port
	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
