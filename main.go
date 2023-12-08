package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port env is Required")
	}

	instanceID := os.Getenv("INSTANCE_ID")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "the http request method is not allowed", http.StatusMethodNotAllowed)
			return
		}

		text := "Hello World"
		if instanceID != "" {
			text = text + ". from " + instanceID
		}
		w.Write([]byte(text))

	})

	server := new(http.Server)
	server.Handler=mux
	server.Addr=":" + port


	log.Println("web server is starting at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}


}