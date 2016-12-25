package main

import (
	"log"
	"time"

	"climax.com/mqtt.sa/dispatch"

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

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.TODO()
	// _, err = cli.Put(ctx, "/panel/001d940361c0", "climax")

	// resp, err := cli.Get(ctx, "/nodes/", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	// for _, ev := range resp.Kvs {
	// 	// fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	// 	k := ev.Key
	// 	v := ev.Value

	// 	fmt.Println(string(k), string(v))

	// }

	// fmt.Println(resp.Count)
	go timer(ctx, cli)
	//time.Sleep(10 * time.Second)

	<-make(chan int)

}

func timer(ctx context.Context, cli *clientv3.Client) {
	for true {
		go dispatch.GetMqttPanel(ctx, cli)
		time.Sleep(1 * time.Second)
	}
}
