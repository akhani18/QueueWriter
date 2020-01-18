package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/akhani18/QueueWriter/queue"

	"github.com/gorilla/mux"
)

func writeMessage(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Printf("Payload: %s", string(reqBody))
	log.Printf("Payload Size: %d", r.ContentLength)

	msgId, err := queue.Send(reqBody)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	log.Printf("Successfully sent messageId: %s", msgId)
}

func main() {
	qURL := queue.InitQueue()
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Queue Writer Service")
	})

	router.HandleFunc("/message", writeMessage).Methods(http.MethodPost)

	log.Printf("Server listening at port %d, for queue %s.", 8000, qURL)

	log.Fatal(http.ListenAndServe(":8000", router))
}
