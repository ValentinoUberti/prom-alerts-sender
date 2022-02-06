import React from 'react'
import { connect } from 'react-redux'
import SendResolved from "../SendResolved";
import Divider from '@mui/material/Divider';

class AlertList extends React.Component {

  constructor(props) {

    super(props);
    this.renderedListItems = this.renderedListItems.bind(this);
    this.state = {
      client: props.websocket

    }
  }

  renderedListItems() {
    console.log( this.props.mystate.alertStore.alertsList);

    return this.props.mystate.alertStore.alertsList.map(
      (alert) => 
      (alert.messageState == "ALERT_RECEIVED_IN_WS_SERVER") ?
      <li key={alert.id}> {alert.alertName} - ({alert.alertPriority}) - Message received from WS SERVER</li>:
      (alert.messageState == "ALERT_SENT_TO_ALERTMANAGER") ?
      <li key={alert.id}> {alert.alertName} - ({alert.alertPriority}) - Alert sent to Alert Manager</li>:
      (alert.messageState == "WAITING_FOR_ICINGA_CONFIRMATION") ?
      <li key={alert.id}> {alert.alertName} - ({alert.alertPriority}) - Waiting for Icinga confirmation</li>:
      (alert.messageState == "ALERT_FIRING_ON_ICINGA") ?
      <li key={alert.id}> {alert.alertName} - ({alert.alertPriority}) - <SendResolved alertName={alert.alertName} alertPriority={alert.alertPriority} alertId={alert.id} websocket={this.state.client}/></li>:
      (alert.messageState == "WAITING_FOR_ICINGA_RESOLVED_CONFIRMATION") ?
      <li key={alert.id}> {alert.alertName} - ({alert.alertPriority}) - WAITING_FOR_ICINGA_RESOLVED_CONFIRMATION </li>:
      (alert.messageState == "ALERT_RESOLVED_IN_ICINGA") ?
      <li key={alert.id}> {alert.alertName} - ({alert.alertPriority}) - ALERT_RESOLVED_IN_ICINGA </li>:null
    )
  }

  //WAITING_FOR_ICINGA_RESOLVED_CONFIRMATION

  render() {

    return (
      <div>
         <Divider light />
        <ul className="todo-list">{this.renderedListItems()}</ul>

      </div>
    )
  }

}
const mapStateToProps = state => ({
  mystate: state
})

export default connect(mapStateToProps)(AlertList)

