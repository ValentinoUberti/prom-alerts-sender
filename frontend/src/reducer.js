
import { combineReducers } from 'redux'

import alertsReducer from "./features/alertList/alertsReducer";


const rootReducer = combineReducers({
    // Define a top-level state field named `todos`, handled by `todosReducer`
    alertStore: alertsReducer

})

export default rootReducer