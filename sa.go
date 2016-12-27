package main

import (
	"log"
	"time"

	"climax.com/mqtt.sa/dispatch"
	"climax.com/mqtt.sa/healthz"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.0.1.11:2379", "10.0.1.12:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx := context.TODO()

	go timer(ctx, cli)
	go health(ctx, cli)

	<-make(chan int)

}

func timer(ctx context.Context, cli *clientv3.Client) {
	for {
		go dispatch.GetMqttPanel(ctx, cli)
		time.Sleep(1 * time.Second)
	}
}

func health(ctx context.Context, cli *clientv3.Client) {
	for {
		go healthz.Check()
		time.Sleep(1 * time.Second)
	}
}
