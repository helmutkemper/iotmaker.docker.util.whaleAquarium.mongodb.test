package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", testMongoDB)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("http server error: %v", err.Error())
	}
}

func testMongoDB(w http.ResponseWriter, r *http.Request) {
	var err error

	query := r.URL.Query()
	dbAddress := query.Get("db")

	err = mongoDbTest(dbAddress)
	if err != nil {
		log.Printf("mongodb error: %v\naddress: %v", err.Error(), dbAddress)
		_, err = fmt.Fprintf(w, `{"ok": false, "error": "%v"}`, err.Error())
		if err != nil {
			log.Fatalf("http server write response error: %v", err.Error())
		}
		return
	}

	_, err = fmt.Fprint(w, `{"ok": true, "error": ""}`)
	if err != nil {
		log.Fatalf("http server write response error: %v", err.Error())
	}
}

func mongoDbTest(address string) (err error) {
	var ctx context.Context
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	var client *mongo.Client
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://" + address))
	if err != nil {
		return
	}

	err = client.Connect(ctx)
	if err != nil {
		return
	}

	defer func(ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("mongodb disconnect error: %v", err.Error())
		}
	}(ctx)

	err = client.Ping(ctx, nil)

	return
}
