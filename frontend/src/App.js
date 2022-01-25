import React, { Component } from "react";

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
    this.sendAlert = this.sendAlert.bind(this);


  }




  componentDidMount() {
    client.onopen = () => {
      console.log('WebSocket Client Connected');
    };
    client.onmessage = (message) => {
      console.log(message);
    };
  }


  sendAlert() {
    client.send(this.state.value);
    this.setState({ value: '' });
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



        <Container maxWidth="sm">
        <CssBaseline />
          <SendAlert websocket={client} />
          <AlertList />
        </Container>




      </div>

    );
  }
}

export default App;
