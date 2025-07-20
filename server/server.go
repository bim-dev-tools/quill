package server

import (
	"fmt"
	"net/http"
	"spagen/config"
	"sync"

	"github.com/gorilla/websocket"
)

func Start(restart chan string) {
	buildDir := config.Get().BuildDir
	port := config.Get().Server.Port

	var clientsMu sync.Mutex
	var clients = make(map[*websocket.Conn]struct{})

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	http.HandleFunc("/livereload", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		clientsMu.Lock()
		clients[ws] = struct{}{}
		clientsMu.Unlock()
		defer func() {
			clientsMu.Lock()
			delete(clients, ws)
			clientsMu.Unlock()
			ws.Close()
		}()
		for {
			if _, _, err := ws.NextReader(); err != nil {
				break
			}
		}
	})

	http.Handle("/", http.FileServer(http.Dir(buildDir)))

	go func() {
		for range restart {
			clientsMu.Lock()
			for ws := range clients {
				ws.WriteMessage(websocket.TextMessage, []byte("reload"))
			}
			clientsMu.Unlock()
		}
	}()

	fmt.Printf("Dev server listening on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != http.ErrServerClosed {
		panic(fmt.Errorf("error starting server: %w", err))
	}
}
