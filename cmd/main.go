package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/pmcatominey/sketch-game/pkg/game"
)

var (
	listen   = flag.String("listen", ":8080", "http listen address")
	openCors = flag.Bool("openCors", false, "allow all cors origins")

	gameServer http.Handler
)

func main() {
	flag.Parse()

	gameServer = game.NewServer(*openCors)

	log.Printf("[INFO] main: listening on %s\n", *listen)
	checkFatal(http.ListenAndServe(*listen, http.HandlerFunc(handle)))
}

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] http: %s %s\n", r.Method, r.URL.Path)

	if *openCors {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	gameServer.ServeHTTP(w, r)
}
