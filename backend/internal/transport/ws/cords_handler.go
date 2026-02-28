package ws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Leonid-Yakovlev63/mouse-tracker/backend/internal/domain"
	"github.com/Leonid-Yakovlev63/mouse-tracker/backend/internal/service"
	"github.com/gorilla/websocket"
)

type (
	CordsHandler interface {
		HandleCords(w http.ResponseWriter, r *http.Request)
		RegisterRoutes(mux *http.ServeMux)
	}

	cordsHandler struct {
		upgrader     websocket.Upgrader
		clients      map[*websocket.Conn]bool
		cordsService service.CordsService
	}
)

func NewCordsHandler(cordsService service.CordsService) CordsHandler {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &cordsHandler{upgrader: upgrader, clients: make(map[*websocket.Conn]bool), cordsService: cordsService}
}

func (h *cordsHandler) HandleCords(w http.ResponseWriter, r *http.Request) {

	conn, _ := h.upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	h.clients[conn] = true
	defer delete(h.clients, conn)

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break
		}

		var cords domain.Cords

		err = json.Unmarshal(message, &cords)
		if err != nil {
			fmt.Println("error unmarshalling JSON: %v\n", err)
			continue
		}

		go h.cordsService.ProcessCords(cords)
	}

}

func (h *cordsHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ws", h.HandleCords)
}
