package orametrics

import (
	_ "fmt"
	. "github.com/blacked/go-zabbix"
	"time"
)

func send(zabbixData map[string]string, zabbixHost string, zabbixPort int, hostName string) {
	var metrics []*Metric
	for k, v := range zabbixData {
		metrics = append(metrics, NewMetric(hostName, k, v, time.Now().Unix()))
	}
	// Create instance of Packet class
	packet := NewPacket(metrics)

	// Send packet to zabbix
	z := NewSender(zabbixHost, zabbixPort)
	z.Send(packet)
}

func sendD(j string, k string, zabbixHost string, zabbixPort int, hostName string) {
	var metrics []*Metric
	metrics = append(metrics, NewMetric(hostName, k, string(j), time.Now().Unix()))

	// Create instance of Packet class
	packet := NewPacket(metrics)

	// Send packet to zabbix
	z := NewSender(zabbixHost, zabbixPort)
	z.Send(packet)
}
