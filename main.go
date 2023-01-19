/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http"
	http2 "test/http"
	"test/utils"
)

func main() {

	container := &utils.Helper{}
	r := http2.InitializeRoutes(container)
	server := &http.Server{
		Addr:    "localhost:9090",
		Handler: r,
	}
	log.Println("Running on " + server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Could not start: %v", err))
	}
}
