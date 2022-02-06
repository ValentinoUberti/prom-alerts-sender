package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ValentinoUberti/prom-alerts-sender/internal/icingautils"

	"github.com/gorilla/websocket"
	"github.com/vshn/go-icinga2-client/icinga2"
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

type WsClientConnection struct {
	Socket *websocket.Conn // shared websocket connection
	mu     sync.Mutex
}

func (p *WsClientConnection) send(v interface{}) error {

	p.mu.Lock()
	b, err := json.Marshal(v)
	p.Socket.WriteMessage(websocket.TextMessage, b)
	//err := p.Socket.WriteJSON(v)
	log.Println(err)
	p.mu.Unlock()
	return err
}

func (p *WsClientConnection) read() (msgType int, msg []byte, err error) {
	msgType, msg, err = p.Socket.ReadMessage()
	return
}

func (p *WsClientConnection) close() {
	_ = p.Socket.Close()
	return
}

type MsgWsType struct {
	Command string       `json:"command"`
	Result  bool         `json:"result"`
	Data    AlertingRule `json:"data"`
}

/*
func (m *MsgWsType) PrepareMessage(Command string, Result bool, Fase string, data []string) {

	/* Input command
	       FIRE_ALERT
		   RESOLVE_ALERT


	m.Command = Command
	m.Result = Result
	m.Data = data

}*/

func sendAlert(alert AlertingRule) (error, int) {

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
		return nil, 0
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil, resp.StatusCode

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

//Channel for sending message through websocket connection
//Messages are sent to all the websocket clients
var msgWsChannel = make(chan MsgWsType)

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint

func prepareAndSendAlertToIcinga(alert AlertingRule) (error, int) {

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
	return sendAlert(alert)

}

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

type alertsStates struct {
	Alerts      AlertingRule
	IcingaState string
}

type AlertsFiringInIcinga struct {
	mutex          sync.Mutex
	firingServices []icinga2.Service
}

func (alertsFiringInIcinga *AlertsFiringInIcinga) addAlerts(firingServices []icinga2.Service) error {
	alertsFiringInIcinga.mutex.Lock()
	alertsFiringInIcinga.firingServices = firingServices
	alertsFiringInIcinga.mutex.Unlock()
	return nil
}

func (alertsFiringInIcinga AlertsFiringInIcinga) listAlerts() error {

	log.Printf("+v\n", alertsFiringInIcinga.firingServices)
	return nil
}

func checkAlertsOnIcinga(firingServices *AlertsFiringInIcinga) {

	services, err := icingautils.AlertFiringChecker("api.ocp2.lab.seeweb")

	if err != nil {
		log.Println(err)
		return
	}

	firingServices.addAlerts(services)
	firingServices.listAlerts()

}

func startIcingaServicesPool(firingServices *AlertsFiringInIcinga, interval int) {

	log.Println("Icinga services pool started")
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			checkAlertsOnIcinga(firingServices)
			break

		}
	}

}

func checkIfAlertIsShowingInIcinga(firingServices *AlertsFiringInIcinga, interval int) {

	log.Println("Icinga services pool started")
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			checkAlertsOnIcinga(firingServices)
			break

		}
	}

}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	var wsClient = new(WsClientConnection)
	var err error
	// upgrade this connection to a WebSocket
	// connection
	wsClient.Socket, err = upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection

	// Start Icinga Services Pool

	firingServices := AlertsFiringInIcinga{}
	interval := 10

	go startIcingaServicesPool(&firingServices, interval)

	go func(socket *WsClientConnection) {
		for {

			_, msg, err := socket.read()
			if err != nil {
				log.Println(err)
				return

			} else {
				var jsonMessage MsgWsType
				//fmt.Println("here")
				err = json.Unmarshal(msg, &jsonMessage)
				if err != nil {
					log.Println(err)
					return
				}
				/* Input command
				       FIRE_ALERT
					   RESOLVE_ALERT
				*/
				//fmt.Printf("%+v\n", jsonMessage)
				switch jsonMessage.Command { //Comando ricevuto dal webclient

				case "FIRE_ALERT":
					log.Println("FIRE_ALERT")

					//Send confirmation "ALERT_RECEIVED_IN_WS_SERVER"
					confirmation := MsgWsType{}
					confirmation.Command = "ALERT_RECEIVED_IN_WS_SERVER"
					confirmation.Result = true
					confirmation.Data = jsonMessage.Data
					msgWsChannel <- confirmation

					// Changing state test
					go func() {
						confirmation.Command = "ALERT_SENT_TO_ALERTMANAGER"

						err, status := prepareAndSendAlertToIcinga(jsonMessage.Data)
						log.Println(status)
						if (err == nil) && status == 200 {
							msgWsChannel <- confirmation

							// WAITING_FOR_ICINGA_CONFIRMATION

							time.Sleep(4 * time.Second)
							confirmation.Command = "WAITING_FOR_ICINGA_CONFIRMATION"
							msgWsChannel <- confirmation

							go func() {

								time.Sleep(4 * time.Second)
								confirmation.Command = "ALERT_FIRING_ON_ICINGA"
								ticker := time.NewTicker(time.Duration(1) * time.Second)
								for {
									select {
									case <-ticker.C:

										for _, service := range firingServices.firingServices {
											log.Println("HERE 2    -----")
											service_vars := service.GetVars()
											log.Printf("%v\n", service_vars)

											if service.DisplayName == confirmation.Data.Labels.Alertname {
												log.Println("HERE 3    -----")
												log.Println(service.DisplayName)
												//log.Println(confirmation.Data.Labels.Severity)
												//log.Println(service_vars["label_severity"])

												if service_vars["label_severity"] == confirmation.Data.Labels.Severity {
													log.Println("HERE 4    -----")
													log.Printf("State : %v", service.State)

													/*
														Service state:
														0: OK
														1: Warning
														2: Critical
														3: Unknow
													*/
													if (service.State == 1) || (service.State == 2) || (service.State == 3) {

														log.Println("HERE 5    -----")

														msgWsChannel <- confirmation

														//
														return

													}
												}

											}

										}

									}
								}

								//

							}()

						}
						// TODO send error message to REACT

					}()

				case "RESOLVE_ALERT":
					log.Println("RESOLVING_ALERT")
					confirmation := MsgWsType{}
					confirmation.Command = "WAITING_FOR_ICINGA_RESOLVED_CONFIRMATION"
					confirmation.Result = true
					confirmation.Data = jsonMessage.Data

					err, status := prepareAndSendAlertToIcinga(jsonMessage.Data)
					log.Println(status)
					if (err == nil) && status == 200 {

						msgWsChannel <- confirmation

						go func() {

							time.Sleep(4 * time.Second)

							ticker := time.NewTicker(time.Duration(1) * time.Second)
							for {
								select {
								case <-ticker.C:

									// If a service is not in the list, means that is resolved
									// because services are filterd
									// filter.Filter = fmt.Sprintf("service.problem==1 && match(\"%v\",service.host_name)", hostname)
									found := false

									for _, service := range firingServices.firingServices {
										log.Println("RESOLVED HERE 2    -----")
										service_vars := service.GetVars()
										log.Printf("%v\n", service_vars)

										if service.DisplayName == confirmation.Data.Labels.Alertname {
											log.Println("RESOLVED  3    -----")
											log.Println(service.DisplayName)
											log.Println(confirmation.Data.Labels.Severity)
											log.Println(service_vars["label_severity"])

											if service_vars["label_severity"] == confirmation.Data.Labels.Severity {
												log.Println("RESOLVED  HERE 4    -----")
												log.Printf("State : %v", service.State)

												found = true

											}

										}

									} // for

									if found != true {

										log.Println("HERE 5    -----")
										confirmation.Command = "ALERT_RESOLVED_IN_ICINGA"
										msgWsChannel <- confirmation

										//
										return

									}

								}
							}

							//

						}()
					}

				} // switch jsonMessage.Command

			} // else
		} // for

	}(wsClient) // go fun Reader
	//reader(ws)

	go func(socket *WsClientConnection, mymsgChannel chan MsgWsType) {

		var msg MsgWsType

		for {
			msg = <-mymsgChannel
			//log.Println(msg.Command)
			//log.Println("Received message from channel")
			//log.Println(msg)
			//b, err := json.Marshal()
			//socket.mu.Unlock()
			err = socket.send(msg)
			if err != nil {
				//socket.close()
				fmt.Println("Errore in send ", err)
				//mymsgChannel <- msg
				//break
			}
			//log.Println(string(b))
			log.Println("Channel message sent to websocket client")

		}
	}(wsClient, msgWsChannel)

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
