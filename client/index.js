import React from 'react'
import { render } from 'react-dom'
import { Provider } from 'react-redux'

import Router from './routes'
import store from './reducers'

render(
  <Provider store={store}>
  <Router store={store}/>
  </Provider>,
  document.getElementById('root')
)
