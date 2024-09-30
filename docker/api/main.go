package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

// create http server
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ticker := time.NewTicker(5 * time.Second)

	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				err := rdb.Set(context.TODO(), "key", rand.IntN(1000), 0).Err()
				if err != nil {
					panic(err)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		val, err := rdb.Get(context.TODO(), "key").Result()
		if err != nil {
			panic(err)
		}
		s := fmt.Sprintf("%v", val)
		fmt.Println(s)
		fmt.Fprintf(w, s)
	})
	http.ListenAndServe(":8080", nil)
}
