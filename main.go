/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http"
	"test/database"
	http2 "test/http"
)

func main() {

	container := &database.Container{}
	r := http2.InitializeRoutes(container)
	server := &http.Server{
		Addr:    ":9091",
		Handler: r,
	}
	log.Println("Running...")

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Could not start: %v", err))
	}
}
