import React from 'react'
import { Link } from 'react-router'

const NavBar = ({ locationBeforeTransitions }) => {
  const path = (pathname, url) => (
    <li role="presentation" className={locationBeforeTransitions === url ? 'active' : ''}>
      <Link to={url}>
        {pathname}
      </Link>
    </li>
  )
  return (
    <ul className="nav nav-tabs">
    {path('How', '/')}
    {path('Services', '/services')}
    {path('Clients', '/clients')}
    </ul>
  )
}

export default NavBar
