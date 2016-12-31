package etcd

import (
	"context"
	"log"

	"github.com/coreos/etcd/clientv3"
)

//Select pass the key and get value or values
func Select(ctx context.Context, cli *clientv3.Client, key string) clientv3.GetResponse {
	resp, err := cli.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))

	if err != nil {
		log.Fatal(err)
	}

	return *resp

}

//Upsert represents insert or update
func Upsert(ctx context.Context, cli *clientv3.Client, key string, value string) {
	_, err := cli.Put(ctx, key, value)

	if err != nil {
		log.Fatal(err)
	}
}

//Delete the key and value
func Delete(ctx context.Context, cli *clientv3.Client, key string) {

}
