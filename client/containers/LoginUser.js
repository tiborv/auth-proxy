import React from 'react'
import { connect } from 'react-redux'
import { login } from '../actions/login'


const LoginUser = ({ dispatch }) => {
  let username
  let password

  return (
    <div>
      <input ref={node => {
        username = node
      }} />
      <input ref={node => {
        password = node
      }} />
      <button onClick={() => dispatch(login(username.value, password.value))}>
        Log in
      </button>
    </div>
  )
}
export default connect()(LoginUser)
