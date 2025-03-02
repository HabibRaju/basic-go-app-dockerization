package routes

import (
	"database/sql"
	"net/http"

	"sample-health/internal/handler"
	"sample-health/internal/repository"
	"sample-health/internal/service"
)

func SetupRoutes(db *sql.DB) http.Handler {
	// Initialize repository, service, and handler
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Create a new ServeMux (router)
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/get", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUserByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
