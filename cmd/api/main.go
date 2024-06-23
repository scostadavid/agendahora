package main

import (
	"agendahora/cmd/web"
	"agendahora/database"
	"agendahora/domain/handlers"
	"agendahora/domain/repositories"
	"agendahora/domain/services"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

func main() {

	db := database.New()

	waitlistRepository := repositories.NewWaitlistRepository(db)
	waitlistService := services.NewWaitlistService(*waitlistRepository)
	waitlistHandler := handlers.NewWaitlistHandler(*waitlistService)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	mux := http.NewServeMux()

	// mux.HandleFunc("GET /hello-world", web.HelloWorldHandler)
	// mux.HandleFunc("POST /hello", web.HelloWebHandler)
	mux.HandleFunc("POST /add", waitlistHandler.AddToWaitList)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("GET /assets/", fileServer)
	mux.Handle("GET /", templ.Handler(web.LandingPage()))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
