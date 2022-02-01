import React, { Component } from "react";
import store from './store'
import { w3cwebsocket as W3CWebSocket } from "websocket";
import TypoGraphy from '@material-ui/core/Typography'
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar'
//import NavBar from "./navbar.js"
import SendAlert from "./components/SendAlert"
import AlertList from "./components/AlertList"
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import CssBaseline from '@mui/material/CssBaseline';

const client = new W3CWebSocket('ws://127.0.0.1:8080/ws');

class App extends Component {
  constructor(props) {
    super(props);

  }




  componentDidMount() {
    client.onopen = () => {
      console.log('WebSocket Client Connected');
    };


    client.onmessage = (e) => {
      console.log("REACT");
      console.log(e);

      try {
        const event = JSON.parse(e.data);

        switch (event.command) {

          case 'ALERT_RECEIVED_IN_WS_SERVER':
            console.log("ALERT_RECEIVED_IN_WS_SERVER")
            store.dispatch({ type: 'alert/alert_received_in_ws_server', payload: event.data })
            break;

          case 'ALERT_SENT_TO_ALERTMANAGER':
            console.log("ALERT_SENT_TO_ALERTMANAGER")
            store.dispatch({ type: 'alert/ALERT_SENT_TO_ALERTMANAGER', payload: event.data })
            break;

          //WAITING_FOR_ICINGA_CONFIRMATION
          case 'WAITING_FOR_ICINGA_CONFIRMATION':
            console.log("WAITING_FOR_ICINGA_CONFIRMATION")
            store.dispatch({ type: 'alert/WAITING_FOR_ICINGA_CONFIRMATION', payload: event.data })
            break;
          default:
            console.log("MESSAGE NOT UNDERSTOOD")

        } // switch


      } // try

      catch (err)
      //TODO
      {
        console.log(err);
      }








    };
  }





  render() {
    return (
      <div className="App">

        <AppBar color="primary" position="static">
          <Toolbar>
            <TypoGraphy color="inherit">
              ISP Alert sender
            </TypoGraphy>

          </Toolbar>
        </AppBar>





        <SendAlert websocket={client} />
        <AlertList websocket={client} />





      </div>

    );
  }
}

export default App;
