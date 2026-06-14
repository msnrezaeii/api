package router

import (
	"log"
	"net/http"

	"github.com/msnrezaeii/api/topic"
)

func Router() {

	router := http.NewServeMux()
	server := &http.Server{
		Addr:    ":4040",
		Handler: router,
	}

	router.HandleFunc("GET /topic/{id}", topic.GetTopic)
	router.HandleFunc("POST /topic/", topic.CreateTopic)
	router.HandleFunc("PUT /topic/{id}", topic.Update)
	router.HandleFunc("DELETE /topic/{id}", topic.DeleteTopic)

	log.Println("Starting Server on port :4040")
	err := server.ListenAndServe()
	log.Println(err)

}
