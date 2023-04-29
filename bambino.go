package main

import (
	"github.com/joegasewicz/bambino/views"
	"github.com/joegasewicz/gomek"
	"log"
)

func main() {
	// config
	c := gomek.Config{
		BaseTemplateName: "layout",
		BaseTemplates: []string{
			"./templates/layout.gohtml",
		},
	}
	// app
	app := gomek.New(c)
	// views
	app.Route("/health").Methods("GET").Resource(&views.Health{})
	app.Route("/files").Methods("GET").Resource(&views.File{})
	// middleware
	app.Use(gomek.Logging)
	app.Use(gomek.CORS)
	app.Listen(4444)
	if err := app.Start(); err != nil {
		log.Printf("Error running Gomek: %e", err)
	}
}
