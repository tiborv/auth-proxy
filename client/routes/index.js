import React from 'react'
import { Router, Route, browserHistory } from 'react-router'
import { syncHistoryWithStore } from 'react-router-redux'
import Root from '../containers/Root'
import LoginView from '../views/Login'


export default ({ store }) => {
  const history = syncHistoryWithStore(browserHistory, store)
  return (
  <Router history={history}>
    <Route path="/" component={Root}>
      <Route path="login" component={LoginView}/>
    </Route>
  </Router>
)}
