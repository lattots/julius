package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/lattots/julius/pkg/event"
	"log"
	"net/http"
	"os"
)

const port = 3000

func main() {
	err := godotenv.Load("data/.env")
	if err != nil {
		log.Fatalln("error loading environment variables:", err)
	}

	fmt.Println("Starting server...")

	r := http.NewServeMux()

	r.HandleFunc("GET /event", handleGetEvent)
	r.HandleFunc("POST /event", handleNewEvent)

	fmt.Printf("Server started on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		log.Fatalln(err)
	}
}

func handleGetEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET parameters:", r.URL.Query())

	eventID := r.URL.Query().Get("event-id")
	if eventID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("no event-id in request")
		return
	}

	fmt.Println("Request event ID:", eventID)
	_, err := fmt.Fprintf(w, "Event ID in request: %s\n", eventID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error writing response:", err)
		return
	}
}

func handleNewEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event.Event
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error decoding request body:", err)
		return
	}

	db, err := sql.Open("mysql", os.Getenv("DATABASE_APP"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error opening database connection:", err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("error closing database connection:", err)
		}
	}()
	newEvent.DB = db

	eventID, err := newEvent.Insert()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error inserting event to database:", err)
		return
	}

	_, err = fmt.Fprintln(w, "Successfully inserted event to database. Event ID:", eventID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error writing response", err)
		return
	}
}
