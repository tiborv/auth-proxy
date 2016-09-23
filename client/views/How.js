import React from 'react'

const HowView = () => (
  <div className="page-header">
    <h1>
      How
    </h1>
    <h2>
      Step 1.
    </h2>
    <small>Create a client -> obtain a token </small>
    <h2>
      Step 2.
    </h2>
    <small>Add client to a service -> remember slug</small>
    <h2>
      Step 3.
    </h2>
    <small>Execute a HTTP POST/GET/WHATEVER towards /api/proxy/{'<service-slug>'} and set header 'Authorization: Bearer {'<client-token>'}'</small>
  </div>
)

export default HowView
