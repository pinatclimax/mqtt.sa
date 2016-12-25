package dispatch

import (
	"fmt"
	"log"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

//etcdctl put /mqtt/sa/host/10.0.1.23 0
//etcdctl put /mqtt/sa/host/10.0.1.24 0
//etcdctl put /mqtt/sa/host/10.0.1.25 0

var dispatchCount = 0

//Dispatch ...
func Dispatch(ctx context.Context, cli *clientv3.Client) {
	hostCount := GetHostsCount(ctx, cli)

	if int64(dispatchCount) < hostCount {
		//do dispatch
		dispatchCount++
		fmt.Println(dispatchCount)
	} else {
		dispatchCount = 0
		//do dispatch
		dispatchCount++
		fmt.Println(dispatchCount)
	}

}

// GetMqttPanel function
func GetMqttPanel(ctx context.Context, cli *clientv3.Client) {

	Dispatch(ctx, cli)

	resp, err := cli.Get(ctx, "/mqtt/panel/", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))

	for _, ev := range resp.Kvs {
		k := ev.Key
		v := ev.Value

		fmt.Println(string(k), string(v))
	}

	if err != nil {
		log.Fatal(err)
	}
}

//GetHostsCount function
func GetHostsCount(ctx context.Context, cli *clientv3.Client) int64 {
	resp, err := cli.Get(ctx, "/mqtt/sa/host/", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))

	if err != nil {
		log.Fatal(err)
	}

	return resp.Count
}
