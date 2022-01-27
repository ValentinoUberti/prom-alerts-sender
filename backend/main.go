package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type AlertingLabel struct {
	Alertname string `json:"alertname"`
	Service   string `json:"service,omitempty"`
	Severity  string `json:"severity"`
	//Instance    string `json:"instance"`
	//Prometheus  string `json:"prometheus"`
	//Namespace   string `json:"namespace"`
	//Pod         string `json:"pod"`
	//Integration string `json:"integration"`
	//Container   string `json:"container"`
	//Endpoint    string `json:"endpoint"`
	//Job         string `json:"job"`
}

type AlertingAnnotations struct {
	Summary     string `json:"summary"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

type AlertingRule struct {
	Status       string              `json:"status"`
	Labels       AlertingLabel       `json:"labels"`
	Annotations  AlertingAnnotations `json:"annotations"`
	GeneratorURL string              `json:"generatorURL"`
	StartsAt     time.Time           `json:"startsAt,omitempty"`
	EndsAt       time.Time           `json:"endsAt,omitempty"`
}

func sendAlert(alert AlertingRule) {

	alertArray := []AlertingRule{}

	alertArray = append(alertArray, alert)
	b, err := json.MarshalIndent(alertArray, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))

	jsonAlert, err := json.Marshal(alertArray)

	fmt.Printf("%+v", alertArray)

	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}

	url := "http://localhost:9093/api/v1/alerts"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonAlert))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {

		// print out that message for clarity

		alert := AlertingRule{}

		err := conn.ReadJSON(&alert)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", alert)

		/*

		 */

		//alert.Status = "firing"
		//alert.Labels.Alertname = msg
		//alert.Labels.Service = "alertmanager-main"
		//alert.Labels.Severity = "warning"
		//alert.Annotations.Summary = "My special summary"
		//alert.Annotations.Description = "My special description"
		//alert.Annotations.Message = "My special message"
		//alert.Labels.Instance = "my-test-service.ocp2.lab.seeweb"
		//alert.Labels.Prometheus = "openshift-monitoring/k8s"
		//alert.Labels.Namespace = "openshift-monitoring"

		//alert.Labels.Container = "alertmanager-proxy"
		//alert.Labels.Endpoint = "end"
		//alert.Labels.Instance = "10.131.0.23:9095"
		//alert.Labels.Integration = "webhook"
		//alert.Labels.Job = "alertmanager-main"
		//alert.Labels.Pod = "alertmanager-main-0"

		//alert.StartsAt = startAt
		//alert.EndsAt = endAt
		alert.GeneratorURL = "https://bjlovers.bj/" + alert.Labels.Alertname
		alert.Labels.Service = "Service - " + alert.Labels.Alertname

		if alert.Status == "resolved" {
			alert.EndsAt = time.Now()
		} else {
			now := time.Now()
			count := 130
			alert.StartsAt = now.Add(time.Duration(-count) * time.Minute)
			alert.EndsAt = now.Add(time.Duration(count) * time.Minute)
		}
		sendAlert(alert)

		/*
			if err = conn.WriteJSON(alert); err != nil {
				fmt.Println(err)
			}*/

	}
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
