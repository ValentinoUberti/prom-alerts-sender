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
    //console.log( this.props.mystate.alertStore.alertsList);

    return this.props.mystate.alertStore.alertsList.map(
      (alert) => <li key={alert.id}> {alert.alertName} - ({alert.alertPriority}) - <SendResolved alertName={alert.alertName} alertId={alert.id} websocket={this.state.client}/></li>
    );
  }

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

