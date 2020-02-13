package app

import (
	"log"
	"os"
	"os/signal"
	"net/http"
	"syscall"
	"time"
	"context"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/AlexFrazer/films-server/app/movies"
	"github.com/AlexFrazer/films-server/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func initializeLogging() {
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}
}

func (a *App) Initialize(config *config.Config) {
	initializeLogging()
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Database connection failure")
	}
	a.DB = movies.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/movies", a.handleRequest(movies.GetAll)).Methods(http.MethodGet)
	a.Router.HandleFunc("/movies", a.handleRequest(movies.Create)).Methods(http.MethodPost)
}

func (a *App) Run(host string) {
	srv := &http.Server{
		Handler:      a.Router,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	waitForShutdown(srv)
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}

// Listens for a graceful shutdown.
func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}