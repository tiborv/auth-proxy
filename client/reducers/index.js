import { combineReducers, createStore, applyMiddleware } from 'redux'
import { routerReducer } from 'react-router-redux'
import thunkMiddleware from 'redux-thunk'

import auth from './auth'
import clients from './client'


export default createStore(
  combineReducers({
    clients,
    auth,
    routing: routerReducer
  }),
    applyMiddleware(thunkMiddleware)
)
