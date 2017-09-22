package main

import (
	"log"
	"net/http"
	"os"

	"github.com/k8s-community/codegen/pkg/config"
	"github.com/k8s-community/codegen/pkg/handlers"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var config config.Config
	err := envconfig.Process("codegen", &config)
	if err != nil {
		log.Fatalf("Couldn't get service config: %s", err)
	}

	directory := "/tmp/archive"
	err = os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		log.Fatalf("Сannot create archive dir: %s", err)
	}

	http.HandleFunc("/", handlers.Root)
	http.HandleFunc("/generate", handlers.GenerateCode)

	http.Handle("/archive/", http.StripPrefix("/archive/", http.FileServer(http.Dir(directory))))
	http.Handle("/static/", http.FileServer(http.Dir("/")))

	err = http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
