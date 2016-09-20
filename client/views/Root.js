import React from 'react'
import { connect } from 'react-redux'

import How from './How'
import LoginView from './LoginView'
import NavBar from '../components/NavBar'

const mapStateToProps = (state, ownProps) => ({
  auth: state.auth,
  location: state.routing
})

const Root = ({ children, auth, location }) => (
  <div className="container">
  { auth ? (
    <div>
    <NavBar {...location}/>
    { children || (<How />) }
    </div>
  ) : (
    <LoginView />
  )}
  </div>
)


export default connect(
  mapStateToProps,
)(Root)
