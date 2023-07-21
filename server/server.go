package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// https://tutorialedge.net/golang/go-websocket-tutorial/
func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func createSessionId() (sessionId string) {
	return uuid.New().String()
}

func registerSession(w http.ResponseWriter, r *http.Request) {
	var sessionId = createSessionId()
	fmt.Fprintf(w, "Session ID returned: "+sessionId)
}

func joinSession(w http.ResponseWriter, r *http.Request) {

}

func terminateSession(w http.ResponseWriter, r *http.Request) {

}

func setupRoutes() {
	fmt.Println("Setting routes")
	http.HandleFunc("/register", registerSession)
	http.HandleFunc("/join", joinSession)
	http.HandleFunc("/terminate", terminateSession)
}
