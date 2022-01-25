import React from 'react'
import { connect } from 'react-redux'

class AlertList extends React.Component {


  constructor(props) {

    super(props);
    this.renderedListItems = this.renderedListItems.bind(this);
  }

  renderedListItems() {
    //console.log( this.props.mystate.alertStore.alertsList);

    return this.props.mystate.alertStore.alertsList.map(
      (alert) => <li key={alert.id}> {alert.alertName} - {alert.alertPriority}</li>
    );
  }

  render() {
    console.log(this.props);
    return (
      <div>
        <ul className="todo-list">{this.renderedListItems()}</ul>

      </div>
    )
  }

}
const mapStateToProps = state => ({
  mystate: state
})

export default connect(mapStateToProps)(AlertList)

