package icingautils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"

	"github.com/vshn/go-icinga2-client/icinga2"
)

type Vars map[string]interface{}
type Service struct {
	Name               string  `json:"name,omitempty"`
	DisplayName        string  `json:"display_name"`
	HostName           string  `json:"host_name"`
	CheckCommand       string  `json:"check_command"`
	EnableActiveChecks bool    `json:"enable_active_checks"`
	Notes              string  `json:"notes"`
	NotesURL           string  `json:"notes_url"`
	ActionURL          string  `json:"action_url"`
	Vars               Vars    `json:"vars"`
	Zone               string  `json:"zone,omitempty"`
	CheckInterval      float64 `json:"check_interval"`
	RetryInterval      float64 `json:"retry_interval"`
	MaxCheckAttempts   float64 `json:"max_check_attempts"`
	CheckPeriod        string  `json:"check_period,omitempty"`
	State              float64 `json:"state,omitempty"`
	LastStateChange    float64 `json:"last_state_change,omitempty"`
}

func AlertFiringChecker(hostname string) ([]icinga2.Service, error) {

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
	filter.Filter = fmt.Sprintf("service.problem==1 && match(\"%v\",service.host_name)", hostname)
	//filter.Filter = `service.problem==1 && match("ocp2.lab.seeweb",service.host_name)`
	//filter.Filter = `service.problem==1 && match("ocp3.lab.seeweb",service.host_name)`
	//filter.Filter = "service.name==''"
	//filter.Filter = "service.state==state && match(pattern,service.attrs.host_name)"
	services, err := icinga.ListServices(filter)

	if err != nil {
		log.Println("Error retrieving firing alerts from Icinga2")
		log.Println(err)
		return nil, err

	}

	/*for _, service := range services {
		fmt.Printf("%s <-> %s\n", service.HostName, service.DisplayName)
	}*/

	return services, nil
}
