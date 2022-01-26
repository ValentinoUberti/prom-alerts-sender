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

        case 'alert/firing': {
            return {
                ...state,
                alertsList: [
                    ...state.alertsList,
                    {
                        id: nextTodoId(state.alertsList),
                        alertName: action.payload.alertName,
                        alertPriority: action.payload.alertPriority,
                        resolved: false


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