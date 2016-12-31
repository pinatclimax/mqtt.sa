package dispatch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"climax.com/mqtt.sa/etcd"
	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

//etcdctl put /mqtt/sa/host/10.0.1.23 0
//etcdctl put /mqtt/sa/host/10.0.1.24 0
//etcdctl put /mqtt/sa/host/10.0.1.25 0

var dispatchCount = 0
var h HostInfo

//HostInfo stores the information of node
type HostInfo struct {
	Count         int64
	HostsInfoList []string
}

//Dispatch ...
func Dispatch(ctx context.Context, cli *clientv3.Client, panelInfo string) {
	hostCount := GetHostsCount(ctx, cli)

	if int64(dispatchCount) < hostCount {
		host := h.HostsInfoList[dispatchCount]
		//set to h.HostsInfoList to etcd
		//get value from etcd
		resp, err := cli.Get(ctx, host)

		var connectedValue string
		for _, ev := range resp.Kvs {
			connectedValue = string(ev.Value)
			fmt.Println("key: ", string(ev.Key))
			fmt.Println("connectedValue: ", connectedValue)
		}

		if err != nil {
			log.Fatal(err)
		}

		connectedValueToInt, err := strconv.Atoi(connectedValue)
		connectedValueToInt++
		connectedValue = strconv.Itoa(connectedValueToInt)

		_, err2 := cli.Put(ctx, host, connectedValue)

		if err2 != nil {
			log.Fatal(err2)
		}

		//update /mqtt/panel/001d940361c0 10.0.1.xx
		hostSplit := strings.Split(host, "/")
		hostIP := hostSplit[len(hostSplit)-1]
		_, err3 := cli.Put(ctx, panelInfo, hostIP)
		if err3 != nil {
			log.Fatal(err)
		}

		dispatchCount++

		mac := getPanelMac(panelInfo)
		key := "/mqtt/sa/connected/" + hostIP + "/" + mac
		etcd.Upsert(ctx, cli, key, mac)

	} else {
		dispatchCount = 0
		fmt.Println(h.HostsInfoList[dispatchCount])

		host := h.HostsInfoList[dispatchCount]
		//set to h.HostsInfoList to etcd
		//get value from etcd
		resp, err := cli.Get(ctx, host)

		var connectedValue string
		for _, ev := range resp.Kvs {
			connectedValue = string(ev.Value)
			fmt.Println("key: ", string(ev.Key))
			fmt.Println("connectedValue: ", connectedValue)
		}

		if err != nil {
			log.Fatal(err)
		}

		connectedValueToInt, err := strconv.Atoi(connectedValue)
		connectedValueToInt++
		connectedValue = strconv.Itoa(connectedValueToInt)

		_, err2 := cli.Put(ctx, host, connectedValue)

		if err2 != nil {
			log.Fatal(err2)
		}

		//update /mqtt/panel/001d940361c0 10.0.1.xx
		hostSplit := strings.Split(host, "/")
		hostIP := hostSplit[len(hostSplit)-1]
		_, err3 := cli.Put(ctx, panelInfo, hostIP)
		if err3 != nil {
			log.Fatal(err)
		}

		dispatchCount++
		fmt.Println(dispatchCount)
	}

}

// GetMqttPanel function
func GetMqttPanel(ctx context.Context, cli *clientv3.Client) {

	resp, err := cli.Get(ctx, "/mqtt/panel/", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))

	for _, ev := range resp.Kvs {
		k := string(ev.Key)
		v := string(ev.Value)

		if v == "undefined" {
			fmt.Println("undefined")

			fmt.Println(k)
			fmt.Println(v)

			Dispatch(ctx, cli, k)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}

//GetHostsCount function
func GetHostsCount(ctx context.Context, cli *clientv3.Client) int64 {
	resp, err := cli.Get(ctx, "/mqtt/sa/host/", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	h.HostsInfoList = make([]string, resp.Count)
	for i, ev := range resp.Kvs {
		h.HostsInfoList[i] = string(ev.Key) //get host information from etcd
	}

	h.Count = resp.Count //get hosts count

	// for _, value := range h.HostsInfoList {
	// 	fmt.Println("value: " + value)
	// }

	if err != nil {
		log.Fatal(err)
	}

	return resp.Count
}

func getPanelMac(panelInfo string) string {
	str := strings.Split(panelInfo, "/")
	return str[len(str)-1]
}
