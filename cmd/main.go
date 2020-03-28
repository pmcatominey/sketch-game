package main

import (
	"flag"
	"log"
	"mime"
	"net/http"
	"path"
	"strings"

	"github.com/pmcatominey/sketch-game/pkg/game"
	_ "github.com/pmcatominey/sketch-game/pkg/ui"
	"github.com/rakyll/statik/fs"
)

var (
	listen   = flag.String("listen", ":8080", "http listen address")
	openCors = flag.Bool("openCors", false, "allow all cors origins")

	gameServer http.Handler
	uiServer   http.Handler
)

func main() {
	flag.Parse()

	gameServer = game.NewServer(*openCors)

	statikFs, err := fs.New()
	checkFatal(err)
	uiServer = http.FileServer(statikFs)

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

	if !strings.HasPrefix(r.URL.Path, "/api") {
		mimeType := mime.TypeByExtension(path.Ext(r.URL.Path))
		if mimeType == "" {
			mimeType = "text/html"
		}
		w.Header().Set("Content-Type", mimeType)
		uiServer.ServeHTTP(w, r)
		return
	}

	if *openCors {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	gameServer.ServeHTTP(w, r)
}
