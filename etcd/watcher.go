package etcd

import (
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

func ConnectedWatcher(hostIP string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.0.1.11:2379", "10.0.1.12:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	rch := cli.Watch(context.Background(), "/mqtt/sa/connected/"+hostIP, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
