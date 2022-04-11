package main

import (
	"io/ioutil"
	"log"
	"time"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "test-pub")
	if err != nil {
		log.Println(err)
	}
	defer sc.Close()

	// sc.Publish("l0", []byte("{\"customer_id\": \"Session 1\"}"))
	data1, _ := ioutil.ReadFile("model1.json")
	sc.Publish("l0", data1)
	time.Sleep(5 * time.Second)
	data2, _ := ioutil.ReadFile("model2.json")
	sc.Publish("l0", data2)
	time.Sleep(5 * time.Second)
	data3, _ := ioutil.ReadFile("model3.json")
	sc.Publish("l0", data3)
	time.Sleep(5 * time.Second)
	data4, _ := ioutil.ReadFile("model4-broken.json")
	sc.Publish("l0", data4)
}
