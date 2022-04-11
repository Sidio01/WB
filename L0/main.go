package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	stan "github.com/nats-io/stan.go"
)

var cache = make(map[string]interface{})

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		w.Write([]byte("{\"error\": \"enter correct id\"}"))
		return
	}

	order, ok := cache[id]
	if !ok {
		w.Write([]byte("{\"error\": \"there is no order with such id\"}"))
		return
	}

	json.NewEncoder(w).Encode(order)
}

func main() {
	cache, _ = dbGet()

	go func() {
		http.HandleFunc("/", handler)
		http.ListenAndServe(":8080", nil)
	}()

	sc, err := stan.Connect("test-cluster", "test-recv")
	if err != nil {
		log.Println(err)
	}
	defer sc.Close()

	sc.Subscribe("l0", func(m *stan.Msg) {
		order := Order{}
		json.Unmarshal(m.Data, &order)
		err := validate(order)
		if err != nil {
			log.Printf("validation error: %v. skipping message.\n", err)
		} else {
			orderId, _ := updateDb(order)
			cache[strconv.Itoa(orderId)] = order
			log.Printf("message processed successfully. assigned id - %v.\n", orderId)
		}
	}, stan.DurableName("l0-recv"))

	for {
		time.Sleep(1 * time.Second)
	}
}
