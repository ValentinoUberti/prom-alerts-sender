import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import store from './store'
import { Provider } from 'react-redux'


console.log('Initial state: ', store.getState())

//store.dispatch({ type: 'alert/firing', payload: JSON.parse('{ "alertName": "a", "alertPriority": "b", "alertState": "c", "resolved": true}') })

//console.log('Second state: ', store.getState())

ReactDOM.render(
  
    <Provider store={store}>
      <App />
    </Provider>
  ,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
