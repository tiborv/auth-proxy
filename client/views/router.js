import React from 'react'
import { Router, Route, browserHistory, Redirect } from 'react-router'
import { syncHistoryWithStore } from 'react-router-redux'
import Index from './Index'
import LoginView from './Login'


export default ({ store }) => {
  const history = syncHistoryWithStore(browserHistory, store)
  return (
  <Router history={history}>
    <Redirect from="/" to="/login" />
    <Route path="/">
      <Route path="login" component={LoginView}/>
    </Route>
  </Router>
)}
