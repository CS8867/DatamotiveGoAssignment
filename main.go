package main

import (
	"log"
	"net/http"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

// func close(client *mongo.Client, ctx context.Context,
// 	cancel context.CancelFunc) {

// 	defer cancel()

// 	defer func() {
// 		if err := client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()
// }

// func connect(uri string) (*mongo.Client, context.Context,
// 	context.CancelFunc, error) {

// 	ctx, cancel := context.WithTimeout(context.Background(),
// 		30*time.Second)

// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	return client, ctx, cancel, err
// }

// func ping(client *mongo.Client, ctx context.Context) error {

// 	if err := client.Ping(ctx, readpref.Primary()); err != nil {
// 		return err
// 	}
// 	fmt.Println("connected successfully")
// 	return nil
// }

func main() {

	// client, ctx, cancel, err := connect("mongodb+srv://admin-chaitanya:7NXy6m0WV3rNeP9R@cluster0.qpvmw.mongodb.net/?retryWrites=true&w=majority")
	// if err != nil {
	// 	panic(err)
	// }

	// defer close(client, ctx, cancel)

	// ping(client, ctx)

	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
