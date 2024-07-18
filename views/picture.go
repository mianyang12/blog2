package views

import (
	"net/http"
)

func Mux(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/tupian/index.html")

}
