import React from 'react'
import { Router, Route, browserHistory, Redirect } from 'react-router'
import { syncHistoryWithStore } from 'react-router-redux'

import Root from './Root'
import ClientView from './Clients'
import ServiceView from './Services'

import { authorize } from '../actions/auth'
import { list as listClients } from '../actions/client'
import { list as listServices } from '../actions/service'

export default ({ store }) => {
  const history = syncHistoryWithStore(browserHistory, store)
  return (
  <Router history={history}>
      <Route path="/" component={Root} onEnter={() => authorize(store)}>
        <Route path="/clients" component={ClientView} onEnter={() => listClients(store)}/>
        <Route path="/services" component={ServiceView} onEnter={() => {listServices(store); listClients(store)}}/>
    </Route>
  </Router>
)}
