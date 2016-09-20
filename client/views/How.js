import React from 'react'

const HowView = () => (
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
    <small>Execute a HTTP POST/GET/WHATEVER towards /api/proxy/{'<service-slug>'} and set header 'Authorization: Bearer {'<client-token>'}'</small>
  </div>
)

export default HowView
