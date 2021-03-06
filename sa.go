package main

import (
	"runtime"
	"time"

	"climax.com/mqtt.sa/dispatch"
	"climax.com/mqtt.sa/slave"

	"os"
)

func main() {

	runtime.GOMAXPROCS(4)

	runType := os.Args[1]
	if runType == "master" {
		runType = "master"
	} else {
		runType = "slave"
	}

	switch runType {
	case "master":
		go masterGo()
		go slaveGo()
	case "slave":
		go slaveGo()

	}

	<-make(chan int)

}

func masterGo() {
	for {
		dispatch.GetMqttPanel()
		time.Sleep(1 * time.Second)
	}
}

func slaveGo() {
	slave.SubTopic()
}

// func health(ctx context.Context, cli *clientv3.Client) {
// 	for {
// 		go healthz.Check()
// 		time.Sleep(1 * time.Second)
// 	}
// }
