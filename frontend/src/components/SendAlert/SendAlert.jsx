import React, { Component } from "react";
import { Button, TextField } from '@material-ui/core';
import store from '../../store'
import createAlertJson from "../../alertHelper.js"

//createAlertJosn(status,alertname,severity,summary,message,description)


class SendAlert extends Component {

    constructor(props) {
        super(props);
        this.state = {
            client: props.websocket,
            alertDetails: {
                alertName: "TestAlert",
                alertPriority: "warning"
            }
        }
        this.handleChangeAlertName = this.handleChangeAlertName.bind(this);
        this.sendAlert = this.sendAlert.bind(this);

    }


    sendAlert() {

        var msg = createAlertJson("firing",this.state.alertDetails.alertName,
        "warning","summary","message","description");

        console.log(msg)
        this.state.client.send(msg);

        store.dispatch({ type: 'alert/firing', payload: JSON.parse('{ "alertName": "' + this.state.alertDetails.alertName + '", "alertPriority": "warning", "alertState": "firing", "resolved": false}') })
        var state = { ...this.state }
        state.alertDetails.alertName = "";
        this.setState({ state });
    }

    handleChangeAlertName(event) {
        var state = { ...this.state }
        state.alertDetails.alertName = event.target.value;
        this.setState({ state });
    }

    render() {
        return (

          
                <div>
                    <TextField id="outlined-basic"  value={this.state.alertDetails.alertName} onChange={this.handleChangeAlertName} />
                    <Button onClick={this.sendAlert}>Fire alert</Button>
                </div>
          

        )
    }

}


export default SendAlert;