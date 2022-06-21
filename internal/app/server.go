package app

import (
	"context"
	"flag"
	"github.com/ancarebeca/PortDomainService/internal/adapter"
	"github.com/ancarebeca/PortDomainService/internal/controller"
	"github.com/ancarebeca/PortDomainService/internal/domain/entity"
	"github.com/ancarebeca/PortDomainService/internal/domain/repository"
	"github.com/ancarebeca/PortDomainService/internal/reader"
	service "github.com/ancarebeca/PortDomainService/internal/service"
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

	file, err := os.Open("./ports.json") // TODO: The path could be stored inside `/resources/values.yml` and load it from there
	if err != nil {
		log.Fatalf(err.Error())
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", timeout, "the duration for which the server gracefully wait for existing connections to finish")
	flag.Parse()

	repo := repository.Repository{
		DB: adapter.MemDB{DB: make(map[string]entity.Port)},
	}

	portLoader := service.Loader{
		ObjectReader: reader.JSONReader{},
		Repository:   repo,
	}

	if err := portLoader.Load(file); err != nil {
		log.Fatalf("cannot load ports from file: %v", err)
	}

	portController := controller.PortController{
		PortsService: &service.Ports{
			Repository: repo,
		},
	}

	router := loadRoutes(portController)

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
		log.Fatalf("srv.Shutdown: %v", err)
	}
}

func loadRoutes(portHandler controller.PortController) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ports", portHandler.GetAllPorts).Methods("GET")
	router.HandleFunc("/ports/{id}", portHandler.GetPort).Methods("GET")
	router.HandleFunc("/ports", portHandler.AddOrUpdatesPort).Methods("POST")
	return router
}
