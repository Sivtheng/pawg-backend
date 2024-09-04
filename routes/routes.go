package routes

import (
	"backend/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	// Public routes
	SetupLoginRoutes(r)

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuthMiddleware)

	SetupUserRoutes(api)
	SetupGetInTouchRoutes(api)
	SetupAppointmentRoutes(api)
	SetupAdoptionApplicationRoutes(api)
}
