package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"program/handlers"
	"program/storage/mongostorage"
	"program/users"
	"syscall"
	"time"
)

func main() {
	var userServer *users.UserServer

	mongoStorage, err := mongostorage.NewDatabaseStorage("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Error during connect...", "error", err)
	} else {
		userServer = users.NewUserServer(mongoStorage)
		fmt.Println("Connected to MongoDB database")
	}
	router := handlers.HandleRequest(handlers.RetHandler(userServer))
	s := http.Server{
		Addr:         ":9090",
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	sig := <-signalCh
	fmt.Printf("got signal:%v", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = s.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
