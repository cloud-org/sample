package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"net/http"
)

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	return app.H1().Text("Hello World~")
}

func main() {
	app.Route("/", &hello{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World Example",
	})

	if err := http.ListenAndServe(":8777", nil); err != nil {
		log.Fatal(err)
	}

}
