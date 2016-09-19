import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootswatch/darkly/bootstrap.css'
import 'font-awesome/css/font-awesome.css'
import 'bootstrap-social/bootstrap-social.css'

import React from 'react'
import { render } from 'react-dom'
import { Provider } from 'react-redux'

import Router from './views/router'
import store from './reducers'

render(
  <Provider store={store}>
    <Router store={store}/>
  </Provider>,
  document.getElementById('root')
)
