package main

import (
	"fmt"
	"net/http"

	// Import the internal/translations package, so that its init()
	// function is called.
	"bookstore.example.com/internal/localizer"
	_ "bookstore.example.com/internal/translations"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Initialize a new Localizer based on the locale ID in the URL
	l, ok := localizer.Get(r.URL.Query().Get(":locale"))
	if !ok {
		http.NotFound(w, r)
		return
	}

	var totalBookCount = 1_252_794

	// Update these to use the new Translate() method.
	fmt.Fprintln(w, l.Translate("Welcome!"))
	fmt.Fprintln(w, l.Translate("%d books available", totalBookCount))

	// Add an additional "Launching soon!" message.
	fmt.Fprintln(w, l.Translate("Launching soon!"))
}
