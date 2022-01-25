import React, { Component } from "react";
import { Button } from '@material-ui/core';
import store from '../../store'

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
        this.state.client.send(this.state.value);
        var state = {...this.state}
        state.alertDetails.alertName="";
        this.setState({ state });
        store.dispatch({ type: 'alert/firing', payload: JSON.parse('{ "alertName": "a", "alertPriority": "b", "alertState": "c", "resolved": true}') })
      }

    handleChangeAlertName(event) {
        var state = {...this.state}
        state.alertDetails.alertName=event.target.value;
        this.setState({ state });
    }

    render() {
        return (
            <div>

                <input value={this.state.alertDetails.alertName} onChange={this.handleChangeAlertName} />
                <Button onClick={this.sendAlert}>Send alert</Button>
            </div>
        )
    }

}


export default SendAlert;