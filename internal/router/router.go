package router

import (
	"context"
	"flag"
	configuration "main/configs"
	"main/internal/cache"
	"main/internal/check"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
)

// Router with gracefull shutdown
func Router(cfg *configuration.Configuration) {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "duration time 15s")
	flag.Parse()

	router := mux.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	srv := &http.Server{
		Addr:         cfg.Addr + cfg.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Cache initialized
	cache := cache.NewCache()
	handlerRepo := newHandler(cache, cfg)
	router.HandleFunc("/", handlerRepo.handleConnection)
	http.Handle("/", router)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

type handler struct {
	cache *cache.Cache
	cfg   *configuration.Configuration
}

func newHandler(cache *cache.Cache, cfg *configuration.Configuration) *handler {
	return &handler{
		cache: cache,
		cfg:   cfg,
	}
}

func (h *handler) handleConnection(w http.ResponseWriter, r *http.Request) {
	// Check json data in cache
	cache := h.cache.Get()
	if cache == nil {
		result := check.New(h.cfg).CheckResult()
		h.cache.DataSet(result)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(cache)
	}
}
