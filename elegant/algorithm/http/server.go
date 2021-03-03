package main

import (
	"fmt"
	"log"
	"net/http"
)

// curl http://localhost:8080
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header)
	fmt.Fprintf(w, "Hi, This is an example of http service in golang.")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	http.HandleFunc("/", handler)
	log.Println("listening http server")
	http.ListenAndServe(":8080", nil)
}

/*
cat /proc/sys/net/ipv4/ip_local_port_range
sudo sysctl -w net.ipv4.ip_local_port_range="32768 32769"
sudo sysctl -w net.ipv4.ip_local_port_range="32768 60000"
*/

/*
netstat -atnp
netstat -atnp | grep "ESTABLISHED" | grep "8080"
netstat -atnp | grep "54710"

netstat -atnp | wc -l
netstat -atnp | grep "ESTABLISHED" | wc -l
netstat -atnp | grep "LISTEN" | wc -l
netstat -atnp | grep "TIME_WAIT" | wc -l
netstat -atnp | grep "CLOSE_WAIT" | wc -l
*/

/*
sudo sysctl -a | grep "tcp_fin_timeout"
sudo sysctl -a | grep "tcp_timestamps"
sudo sysctl -a | grep "tcp_tw_recycle"
sudo sysctl -a | grep "tcp_tw_reuse"
sudo sysctl -a | grep "ip_local_port_range"
sudo vim /etc/sysctl.conf

ss -s
*/
