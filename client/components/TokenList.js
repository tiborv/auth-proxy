import React from 'react'
import Token from './Token'

const TokenList = ({ tokens, onClick }) => {
  return (
    <div>
      <button onClick={onClick}>SWAG</button>
      {tokens.map((token, i) =>
        <Token
          key={i}
          id={token.id}
        />
      )}
    </div>
  )
}


export default TokenList
