package api

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the landing page")
	fmt.Println("End point")
}
