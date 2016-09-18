import React, { Component } from 'react'

class Client extends Component {
  constructor(props) {
    super(props)
    this.state = {
      edit: false || this.props.edit,
    }
    this.inputName = {}
  }

  save() {
    if (this.props.client) {
      this.props.update({
        name: this.inputName.value.trim(),
        token: this.props.client.token
      })
      this.setState({ edit: false })
    } else {
      this.props.create({
        name: this.inputName.value.trim(),
      })
    }
  }

  edit() {
    this.setState({ edit: true })
  }

  delete() {
    this.setState({ edit: false })
    this.props.del(this.props.client)
  }

  render() {
    const { edit } = this.state
    const { client } = this.props
    return edit ? (
        <div>
        <input ref={node => this.inputName = node}></input>
        <button onClick={::this.save}> Save </button>
        <button onClick={::this.delete}> Delete </button>
        </div>
      ) : (
        <div>
        Client {client.name}
        <button onClick={::this.edit}> Edit </button>
        </div>
    )
  }
}


export default Client
