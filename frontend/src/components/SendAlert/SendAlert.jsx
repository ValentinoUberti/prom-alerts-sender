import React, { Component } from "react";
import { Button, TextField } from '@material-ui/core';
import store from '../../store'
import createJSONMsg from "../../alertHelper.js"
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import Box from '@mui/material/Box';

//createAlertJosn(status,alertname,severity,summary,message,description)


class SendAlert extends Component {

    constructor(props) {
        super(props);
        this.state = {
            client: props.websocket,
            alertDetails: {
                alertName: "TestAlert",
                alertPriority: "warning",

            }
        }
        this.handleChangeAlertName = this.handleChangeAlertName.bind(this);
        this.handleChangeAlertPriority = this.handleChangeAlertPriority.bind(this);
        this.sendAlert = this.sendAlert.bind(this);
        this.handleChange = this.handleChange.bind(this);
        this.createAlertJson = this.createAlertJson.bind(this);

    }


    handleChange(evt) {
        const value = evt.target.value;
        this.setState({
            ...this.state,

        });
    }


    createAlertJson(status, alertname, severity, summary, message, description) {
        return createJSONMsg("FIRE_ALERT", false, status, alertname, severity, summary, message, description)
    }

    
    sendAlert() {

        var msg = this.createAlertJson("firing", this.state.alertDetails.alertName,
            this.state.alertDetails.alertPriority, "summary", "message", "description");

        console.log(msg)
        this.state.client.send(msg);

        store.dispatch({ type: 'alert/fire_alert', payload: JSON.parse('{"alertName": "' + this.state.alertDetails.alertName + '", "alertPriority": "' + this.state.alertDetails.alertPriority + '", "alertState": "firing", "resolved": false}') })
        var state = { ...this.state }
        state.alertDetails.alertName = "";
        this.setState({ state });
    }

    handleChangeAlertName(event) {
        var state = { ...this.state }
        state.alertDetails.alertName = event.target.value;
        this.setState({ state });
    }
    handleChangeAlertPriority(event) {
        var state = { ...this.state }
        state.alertDetails.alertPriority = event.target.value;
        this.setState({ state });
    }

    render() {
        return (

            <div>
                <Box
                    component="form"
                    sx={{
                        '& > :not(style)': { m: 2, width: '25ch' },
                    }}
                    noValidate
                    autoComplete="off"
                >



                    <TextField variant="outlined" label="Alert name" value={this.state.alertDetails.alertName} onChange={this.handleChangeAlertName} />

                    <TextField

                        id="severity"
                        value={this.state.alertDetails.alertPriority}
                        label="Severity"
                        onChange={this.handleChangeAlertPriority}
                        select
                    >
                       
                        <MenuItem value="warning">Warning</MenuItem>
                        <MenuItem value="critical">Critical</MenuItem>
                        <MenuItem value="none">Unknow</MenuItem>
                    </TextField>

                    <textarea
                        value={this.state.textAreaValue}
                        onChange={this.handleChange}
                        rows={5}
                        cols={5}
                    />

                    <textarea
                        value={this.state.textAreaValue}
                        onChange={this.handleChange}
                        rows={5}
                        cols={5}
                    />

                </Box>

                <Button onClick={this.sendAlert}>Fire alert</Button>
            </div>



        )
    }

}


export default SendAlert;