package main

import (
	"context"
	"fmt"

	"github.com/HunCoding/golang-asynchronous/syncVsAsync/mongodb"
)

func main() {
	err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		panic(err)
	}

	chanUser := make(chan mongodb.ChanUser)

	go mongodb.InsertUserAsyncMongoDB(
		context.Background(),
		mongodb.User{
			Name: "huncoding",
			Age:  20,
		},
		chanUser,
	)

	chanUserValue := <-chanUser
	if chanUserValue.Err != nil {
		panic(chanUserValue.Err)
	}

	fmt.Println(chanUserValue.User)
}
