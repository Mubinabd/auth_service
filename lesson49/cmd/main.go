package main

import (
	"net/http"
)

func main() {
	//cfg := config.Load()
	//
	//services, err := client.NewGrpcClients(cfg)
	//if err != nil {
	//	log.Fatalf("error while connecting clients. err: %s", err.Error())
	//}
	//
	//engine := api.NewGin(handler.NewHandler(services))
	//
	//err = engine.Run()
	//if err != nil {
	//	log.Fatalf("error while running server. err: %s", err.Error())
	//}

	http.Handle("/api", messageHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))

	})))

	http.Handle("/", messageHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})))

	http.ListenAndServe(":8080", nil)
}

func messageHandler(message http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("And"))
		if r.URL.Path != "/api" {
			w.Write([]byte("WRONG"))
		} else {
			message.ServeHTTP(w, r)
		}
	})
}
