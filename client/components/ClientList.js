import React from 'react'
import Client from './Client'

const ClientList = ({ clients, onClick }) => {
  console.log(clients);
  return (
    <div>
      <button onClick={onClick}>SWAG</button>
      {clients.map((client, i) =>
        <Client
          key={i}
          id={client.slug}
        />
      )}
    </div>
  )
}


export default ClientList
