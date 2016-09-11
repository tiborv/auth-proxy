import { combineReducers, createStore, applyMiddleware } from 'redux'
import { routerReducer } from 'react-router-redux'
import thunkMiddleware from 'redux-thunk'

import auth from './auth'


export default createStore(
  combineReducers({
    auth,
    routing: routerReducer
  }),
    applyMiddleware(thunkMiddleware)
)
