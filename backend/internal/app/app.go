package app

import (
	"net/http"

	"github.com/Leonid-Yakovlev63/mouse-tracker/backend/internal/service"
	"github.com/Leonid-Yakovlev63/mouse-tracker/backend/internal/transport/ws"
)

func Run() {
	cordsService := service.NewCordsService()
	cordsHandler := ws.NewCordsHandler(cordsService)

	mux := http.NewServeMux()
	cordsHandler.RegisterRoutes(mux)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}
