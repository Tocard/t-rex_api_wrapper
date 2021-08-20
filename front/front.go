package front

import (
	"log"
	"net/http"
)

func InitFrontEnd() {

	http.Handle("/", http.FileServer(http.Dir("static")))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
