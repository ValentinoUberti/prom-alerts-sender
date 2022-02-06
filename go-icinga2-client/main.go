package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"

	"github.com/vshn/go-icinga2-client/icinga2"
)

func main() {

	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		log.Println("rootCAs")
		log.Fatalln(err)

	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            rootCAs,
	}

	icinga, err := icinga2.New(icinga2.WebClient{
		URL:               "https://icinga2.lab.seeweb:5665",
		Username:          "root",
		Password:          "57d1002d81b69197",
		Debug:             true,
		DisableKeepAlives: true,
		TLSConfig:         tlsConfig})

	if err != nil {
		log.Println("Icinga client")
		log.Fatalln(err)

	}
	filter := icinga2.QueryFilter{}
	//filter.Filter = `service.problem==1 && match("ClusterOperatorDown",service.display_name)`
	filter.Filter = `service.problem==1 && match("ocp3.lab.seeweb",service.host_name)`
	//filter.Filter = "service.name==''"
	//filter.Filter = "service.state==state && match(pattern,service.attrs.host_name)"
	services, err := icinga.ListServices(filter)

	if err != nil {
		log.Println("Services")
		log.Fatalln(err)

	}

	for _, service := range services {
		fmt.Printf("%s <-> %s\n", service.HostName, service.DisplayName)
	}
	

}
