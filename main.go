package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
	"log"
	"net/http"
	DB "websocketServer/databaseSchemes"
	"websocketServer/structs"
)

var dbConn *memdb.MemDB

func main() {
	setupRoutes()
	setupDatabase()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createSessionId() (sessionId string) {
	return uuid.New().String()
}

/*
Trigger new session id generation
*/
func registerSession(w http.ResponseWriter, r *http.Request) {
	var sessionId = createSessionId()
	fmt.Printf("SessionID: %s", sessionId)
	userInfo := &structs.ChannelHost{Id: 1, URI: r.RemoteAddr, SessionID: sessionId, Alive: true}
	DB.Write(userInfo)
	w.Header().Set("session-Id", sessionId)
	w.WriteHeader(http.StatusOK)
}

/*
Check if session id existent and connect to channel host
*/
func joinSession(w http.ResponseWriter, r *http.Request) {

}

/*
Terminate session and clear session id
*/
func terminateSession(w http.ResponseWriter, r *http.Request) {
	var sessionId = r.URL.Query().Get("sessionId")
	DB.DeleteBySessionId(sessionId)
}

func setupDatabase() {
	dbConn = DB.InitDatabaseScheme()
	DB.SetDbConn(dbConn)
	fmt.Println("Database successfully initialized")
}

func setupRoutes() {
	fmt.Println("Setting routes")
	http.HandleFunc("/register", registerSession)
	http.HandleFunc("/join", joinSession)
	http.HandleFunc("/terminate", terminateSession)
}
