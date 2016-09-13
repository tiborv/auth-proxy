import React from 'react'

const Index = ({ children }) => children ? (
  <div>
  { children }
  </div>
) : (
  <div>
    Index
  </div>
)

export default Index
