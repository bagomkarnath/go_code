package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bagomkarnath/go_code/assignments/assignmentthree/routers"
	"github.com/bagomkarnath/go_code/assignments/assignmenttwo/mapstore"

	"github.com/gorilla/mux"
)

func initializeRoutes(h *routers.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/customers", h.GetAll).Methods("GET")
	r.HandleFunc("/api/customers/{id}", h.Get).Methods("GET")
	r.HandleFunc("/api/customers", h.Create).Methods("POST")
	r.HandleFunc("/api/customers/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/api/customers/{id}", h.Delete).Methods("DELETE")
	return r
}

func main() {
	logger := log.New(os.Stdout, "customer-api", log.LstdFlags)

	h := &routers.Handler{
		Repository: mapstore.NewMapStore(), // Injecting dependency
	}
	router := initializeRoutes(h)

	s := http.Server{
		Addr:         ":9090",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		logger.Println("Listening at port 9090 ....")
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Recieved terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
