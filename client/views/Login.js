import React from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router'

const style = {
  display: '-webkit-flex',
  height: '100%',
  WebkitAlignItems: 'center',
  WebkitBoxAlign: 'center',
  alignItems: 'center',
}

const LoginView = ({ dispatch }) => (
  <div style={style} className='jumbotron vertical-center'>
    <div className='col-md-4 col-md-offset-4'>
      <a className="btn btn-block btn-social btn-github" href="/api/oauth/login">
      <span className="fa fa-github"></span> Sign in with Github
      </a>
    </div>
  </div>
)

export default connect()(LoginView)
