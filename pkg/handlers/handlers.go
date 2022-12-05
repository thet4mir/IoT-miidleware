package handlers

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func DeviceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var ctx = context.Background()

		red := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		res, err := red.Get(ctx, "device").Result()

		if err != nil {
			w.Write([]byte("Error!"))
		}
		w.Write([]byte(res))
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("Error!"))
		}
		r.Body.Close().Error()
		var ctx = context.Background()

		red := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		_, err = red.Set(ctx, "device", body, 0).Result()
		if err != nil {
			w.Write([]byte("Error!"))

		}
		w.Write([]byte(body))
	}
}
