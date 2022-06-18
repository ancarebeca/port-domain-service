package app

import (
	"context"
	"flag"
	"github.com/ancarebeca/PortDomainService/internal/httpservice"
	"github.com/ancarebeca/PortDomainService/internal/port"
	"github.com/ancarebeca/PortDomainService/internal/reader"
	"github.com/ancarebeca/PortDomainService/internal/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const timeout = 15 * time.Second

func Run() {

	file, err := os.Open("/usr/src/app/fixtures/ports.json") // TODO: The path could be stored inside `/resources/values.yml` and load it from there
	if err != nil {
		log.Fatalf(err.Error())
	}

	//file, err := os.Open("/Users/rebeca/workspace/port-domain-service/fixtures/ports.json") // TODO: The path could be stored inside `/resources/values.yml` and load it from there
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", timeout, "the duration for which the server gracefully wait for existing connections to finish")
	flag.Parse()

	repo := loadRepository()
	err = repo.LoadPortsFromFile(file)
	if err != nil {
		log.Fatalf(err.Error())
	}

	ctx := context.Background()
	portHandler := httpservice.PortsHandler{
		Repository: repo,
	}

	router := loadRoutes(portHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Server running ...")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("srv.Shutdown: %w", err)
	}
}

func loadRoutes(portHandler httpservice.PortsHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ports", portHandler.GetAllPorts).Methods("GET")
	router.HandleFunc("/ports/{id}", portHandler.GetPort).Methods("GET")
	router.HandleFunc("/ports", portHandler.AddOrUpdatesPort).Methods("POST")
	return router
}

func loadRepository() repository.Repository {
	db := make(map[string]port.Port)
	repo := repository.Repository{
		DB:     repository.MemDB{DB: db},
		Reader: reader.Reader{},
	}
	return repo
}
