const initialState = {
    // { alertName: "", alertPriority: "", alertState: "", resolved: bool}
    alertsList: []

}

function nextTodoId(alertsList) {
    const maxId = alertsList.reduce((maxId, alert) => Math.max(alert.id, maxId), -1)
    return maxId + 1
}

// Use the initialState as a default value
export default function alertsReducer(state = initialState, action) {
    // The reducer normally looks at the action type field to decide what happens
    switch (action.type) {
        // 1) React -> Pod Server
        // React send the "FIRE_ALERT" To web-socket pod
        // Add the alert to the react state
        case 'alert/fire_alert': {
            console.log("REDUCER: alert/fire_alert");
            return {
                ...state,
                alertsList: [
                    ...state.alertsList,
                    {
                        id: nextTodoId(state.alertsList),
                        alertName: action.payload.alertName,
                        alertPriority: action.payload.alertPriority,
                        resolved: false,
                        messageState: "FIRE_ALERT"


                    }
                ]
            }
        };

        // 2) Pod Server -> React 
        // 
        // Alert firing on Icinga
        // mod the "messageState" of the selected alert
        case 'alert/alert_received_in_ws_server': {
            console.log(action.payload)
            const index = state.alertsList.findIndex(alert => alert.alertName === action.payload.labels.alertname);
            console.log(index);
           
            let tmpAlertList = [...state.alertsList]
            tmpAlertList[index].messageState="ALERT_RECEIVED_IN_WS_SERVER"
            return {
                ...state,
                alertsList: [
                    ...tmpAlertList,
                   
                ]
            }
        };

        // 3) Pod Server -> React 
        // 
        // Alert sent  to Alert Manager
        case 'alert/ALERT_SENT_TO_ALERTMANAGER': {
            console.log(action.payload)
            const index = state.alertsList.findIndex(alert => alert.alertName === action.payload.labels.alertname);
            console.log(index);
           
            let tmpAlertList = [...state.alertsList]
            tmpAlertList[index].messageState="ALERT_SENT_TO_ALERTMANAGER"
            return {
                ...state,
                alertsList: [
                    ...tmpAlertList,
                   
                ]
            }
        };

        //WAITING_FOR_ICINGA_CONFIRMATION

        case 'alert/WAITING_FOR_ICINGA_CONFIRMATION': {
            console.log(action.payload)
            const index = state.alertsList.findIndex(alert => alert.alertName === action.payload.labels.alertname);
            console.log(index);
           
            let tmpAlertList = [...state.alertsList]
            tmpAlertList[index].messageState="WAITING_FOR_ICINGA_CONFIRMATION"
            return {
                ...state,
                alertsList: [
                    ...tmpAlertList,
                   
                ]
            }
        };

        // 4) Pod Server -> React 
        // 
        // Alert sent  to Icinga
        case 'alert/alert_firing_on_icinga': {

          



            return {
                ...state,
                alertsList: [
                    ...state.alertsList,
                    {
                        id: nextTodoId(state.alertsList),
                        alertName: action.payload.alertName,
                        alertPriority: action.payload.alertPriority,
                        resolved: false,
                        messageState: "alert_firing_on_icinga"


                    }
                ]
            }
        };





        case 'alert/resolved': {
            return {

                alertsList: [
                    ...state.alertsList.filter(alert => alert.id !== action.payload.alertId)
                ]
            }


        };



        // Do something here based on the different types of actions
        default:
            // If this reducer doesn't recognize the action type, or doesn't
            // care about this specific action, return the existing state unchanged
            return state
    }
}