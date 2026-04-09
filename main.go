package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    env := os.Getenv("ENV")
    if env == "" {
        env = "unknown"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "🚀 Invaders Web\nEnvironment: %s\n", env)
	fmt.Fprintf(w, "🚀 Invaders Web (Fork 2 maforits)\nEnvironment: %s\n", env)
    })

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })

    log.Printf("Starting web server on :%s (env=%s)", port, env)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
