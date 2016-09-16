import React from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router'


const LoginUser = ({ dispatch }) => {
  let username
  let password

  return (
    <div>
      <a href='/api/oauth/login'>Login</a>
    </div>
  )
}
export default connect()(LoginUser)
