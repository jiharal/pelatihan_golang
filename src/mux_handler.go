package src

import (
	authHandlerRouter "pelatihan_golang/src/auth/router"

	"github.com/gorilla/mux"
)

func InitRouterMux(r *mux.Router) {
	authHandlerRouter.NewAuthHandler(r)
}
