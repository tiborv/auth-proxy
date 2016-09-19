import { combineReducers, createStore, applyMiddleware } from 'redux'
import { routerReducer } from 'react-router-redux'
import createLogger from 'redux-logger'
import thunkMiddleware from 'redux-thunk'
import { reducer as form } from 'redux-form'

import clients from './clients'
import auth from './auth'
import services from './services'

const middlewares = [thunkMiddleware]

if (process.env.NODE_ENV === 'development') {
  middlewares.push(createLogger())
}

export default createStore(
  combineReducers({
    services,
    auth,
    clients,
    form,
    routing: routerReducer
  }),
  applyMiddleware(...middlewares)
)
