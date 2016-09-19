import React from 'react'


const IndexView = () => (
  <div className="page-header">
    <h1>
      Step 1.
    </h1>
    <small>Create a client -> obtain a token </small>
    <h1>
      Step 2.
    </h1>
    <small>Add client to a service -> remember slug</small>
    <h1>
      Step 3.
    </h1>
    <small>Execute HTTP POST/GET towrads /api/porxy/{'<service-slug>'} and set header 'Auth-Token' to the client token</small>
  </div>
)

export default IndexView
