import React, { Component } from "react";
import { Button, TextField } from '@material-ui/core';
import store from '../../store'
import createAlertJson from "../../alertHelper.js"

class SendResolved extends Component {

    constructor(props) {
        super(props);
        this.state = {
            client: props.websocket,
            alertId: props.alertId,
            alertName: props.alertName

        }

        this.sendResolved = this.sendResolved.bind(this);

    }



    sendResolved() {
        //this.state.client.send(this.state.alertDetails.alertName);


        var msg = createAlertJson("resolved", this.state.alertName,
            "warning","summary","message","description");

        console.log(msg)
        this.state.client.send(msg);

        console.log("sendResolved called")
        console.log('Store: ', store.getState())
        store.dispatch({ type: 'alert/resolved', payload: JSON.parse('{ "alertId": ' + this.state.alertId + '}') })
        console.log('Store after: ', store.getState())
        //store.dispatch({ type: 'alert/firing', payload: JSON.parse('{ "alertName": "' + this.state.alertDetails.alertName + '", "alertPriority": "warning", "alertState": "firing", "resolved": false}') })
        //var state = { ...this.state }
        //state.alertDetails.alertName = "";
        //this.setState({ state });
    }



    render() {
        return (



            <Button onClick={this.sendResolved}>Resolve alert</Button>



        )
    }

}


export default SendResolved;