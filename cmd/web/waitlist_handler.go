package web

import (
	"fmt"
	"net/http"
)

// post
func AddToWaitListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	email := r.FormValue("email")

	fmt.Fprintf(w, "Email %s adicionado com sucesso!", email)

	// component := HelloPost(name)
	// err = component.Render(r.Context(), w)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	// }
}
