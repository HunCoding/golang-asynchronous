package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/HunCoding/golang-asynchronous/syncVsAsync/mongodb"
)

func TestMain(m *testing.M) {
	err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func BenchmarkAsyncFlow(b *testing.B) {

	ctx := context.Background()

	user := mongodb.User{
		Name: "HunCoding",
		Age:  22,
	}

	for i := 0; i < b.N; i++ {
		chanUser := make(chan mongodb.ChanUser)
		go mongodb.InsertUserAsyncMongoDB(
			ctx,
			user,
			chanUser,
		)

		chanUserValue := <-chanUser
		if chanUserValue.Err != nil {
			b.FailNow()
			return
		}

		fmt.Println(chanUserValue.User)
	}
}

func BenchmarkSyncFlow(b *testing.B) {

	ctx := context.Background()

	user := mongodb.User{
		Name: "HunCoding",
		Age:  22,
	}

	for i := 0; i < b.N; i++ {
		user_returned, err := mongodb.InsertUserSyncMongoDB(
			ctx,
			user,
		)

		if err != nil {
			b.FailNow()
			return
		}

		fmt.Println(user_returned)
	}
}
