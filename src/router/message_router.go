package router

import (
	"net/http"
)

func MessageRouter() {
	http.HandleFunc("/ws/message/getAll", func(w http.ResponseWriter, r *http.Request) {

	})

}
