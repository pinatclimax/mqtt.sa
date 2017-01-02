package mqtt

import (
	"fmt"
	"os"
	"sync"

	"climax.com/mqtt.sa/etcd"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func SubTopics(slaveHostIP string) {
	fmt.Println("slaveHostIP:" + slaveHostIP)
	var wg sync.WaitGroup
	opts := MQTT.NewClientOptions().AddBroker("tcp://10.15.8.129:1883")
	opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	resp := etcd.Select("/mqtt/sa/connected/" + slaveHostIP)
	wg.Add(int(resp.Count))

	for _, mac := range resp.Kvs {
		go subTestTopic(c, string(mac.Value), &wg)
	}

	wg.Wait()
	etcd.ConnectedWatcher(slaveHostIP)
}

func subTestTopic(c MQTT.Client, topic string, wg *sync.WaitGroup) {
	if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Println("topic:", topic)
	wg.Done()
}
