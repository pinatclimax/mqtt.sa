package etcd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"golang.org/x/net/context"
)

//ConnectedWatcher ...
func ConnectedWatcher(hostIP string, c MQTT.Client) {
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
			switch ev.Type {
			case clientv3.EventTypePut:
				addMqttTopics(string(ev.Kv.Value), c)
			case clientv3.EventTypeDelete:

			}
		}
	}
}

func addMqttTopics(topic string, c MQTT.Client) {
	if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Println("topic:", topic)
}
