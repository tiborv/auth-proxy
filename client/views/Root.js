import React from 'react'
import { connect } from 'react-redux'

import LoginBox from '../containers/LoginBox'


const mapStateToProps = (state, ownProps) => ({
  auth: state.auth
})

const Root = ({ children, auth }) => (
  <div>
  { auth ? children : (<LoginBox />) }
  </div>
)


export default connect(
  mapStateToProps,
)(Root)
