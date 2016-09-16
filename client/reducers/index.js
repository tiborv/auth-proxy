import { combineReducers, createStore, applyMiddleware } from 'redux'
import { routerReducer } from 'react-router-redux'
import thunkMiddleware from 'redux-thunk'

import clients from './client'


export default createStore(
  combineReducers({
    clients,
    routing: routerReducer
  }),
    applyMiddleware(thunkMiddleware)
)
