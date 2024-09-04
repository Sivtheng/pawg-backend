package routes

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	SetupUserRoutes(r)
	SetupGetInTouchRoutes(r)
	SetupAppointmentRoutes(r)
	SetupAdoptionApplicationRoutes(r)
}
