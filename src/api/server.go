package api

import (
	"fmt"
	"log"
	"net/http"
	"src/blogos/src/api/router"
	"src/blogos/src/auto"
	"src/blogos/src/config"
)

func init() {
	config.Load()
	auto.Load()
}

func Run() {
	fmt.Printf("Listenning [::] ... %d\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
