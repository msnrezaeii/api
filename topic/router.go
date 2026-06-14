package topic

import (
	"log"
	"net/http"
)

func Router() {

	router := http.NewServeMux()
	server := &http.Server{
		Addr:    ":4040",
		Handler: router,
	}

	router.HandleFunc("GET /topic/{id}", GetTopic)
	router.HandleFunc("POST /topic/", CreateTopic)
	router.HandleFunc("PUT /topic/{id}", Update)
	router.HandleFunc("DELETE /topic/{id}", DeleteTopic)

	log.Println("Starting Server on port :4040")
	err := server.ListenAndServe()
	log.Println(err)

}
