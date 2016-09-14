import { combineReducers, createStore, applyMiddleware } from 'redux'
import { routerReducer } from 'react-router-redux'
import thunkMiddleware from 'redux-thunk'

import auth from './auth'
import tokens from './token'


export default createStore(
  combineReducers({
    tokens,
    auth,
    routing: routerReducer
  }),
    applyMiddleware(thunkMiddleware)
)
