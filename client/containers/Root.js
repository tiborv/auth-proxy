import React from 'react'
import App from '../components/App'

const Root = ({ children }) => children ? (
  <div>
  { children }
  </div>
) : (
  <div>
    <App />
  </div>
)

export default Root
