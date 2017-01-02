package slave

import (
	"fmt"
	"net"
	"os"

	"climax.com/mqtt.sa/etcd"
	"climax.com/mqtt.sa/slave/mqtt"
)

//Start to start slave host
func SlaveStart() {
	slaveIP := getSlaveHostIP()
	mqtt.SubTopics(slaveIP)
}

func getSlaveHostIP() string {

	var addr string

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				resp := etcd.Select("/mqtt/sa/connected/" + ipnet.IP.String())
				if resp.Count != 0 {
					addr = ipnet.IP.String()
				}
			}
		}
	}
	return addr
}
