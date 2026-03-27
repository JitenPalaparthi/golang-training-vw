package main

import (
	"flag"
	"http-demo/handlers"
	"log/slog"
	"net/http"
	"os"
	"runtime"
)

var (
	PORT     string
	FILENAME string
)

const (
	DEFAULT_PORT          = "8089"
	DEFAULT_USER_FILENAME = "users.db"
)

func init() {
	PORT = os.Getenv("APP_PORT")
	FILENAME = os.Getenv("FILE_NAME")
}

func main() {

	if PORT == "" {
		flag.StringVar(&PORT, "port", DEFAULT_PORT, "--port=8089")
	}
	if FILENAME == "" {
		flag.StringVar(&FILENAME, "filename", DEFAULT_USER_FILENAME, "--filename=8089")
	}
	flag.Parse()

	slog.Info("Server started and listening on:" + PORT)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World! Welcome to Http service"))
	})

	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/greet", handlers.Greet("Jiten"))
	http.HandleFunc("/healthz", handlers.Healthz)

	userhandler := handlers.NewUser(FILENAME)

	http.HandleFunc("/users", userhandler.Create)

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		slog.Error(err.Error())
		runtime.Goexit() //it would ensure that all other goroutines would complete their executuion
	} // starts the server
}

// go does not require any separate webserver

// self hosted applications
