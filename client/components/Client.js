import React, { Component } from 'react'
import ClientForm from '../forms/client'

class Client extends Component {
  constructor(props) {
    super(props)
    this.state = {
      edit: false || this.props.edit,
    }
    this.inputName = {}
  }

  save(form) {
    if (this.props.client) {
      this.props.update(form.formData)
      this.setState({ edit: false })
      return
    }
    this.props.create(form.formData)
    this.props.addNew()
  }

  toggleEdit() {
    this.setState({ edit: !this.state.edit })
    if (this.props.addNew) this.props.addNew()
  }

  delete() {
    this.setState({ edit: false })
    this.props.del(this.props.client)
  }


  render() {
    const { edit } = this.state
    const { client, addNew } = this.props
    return edit ? (
      <div>
        <ClientForm
          formData={client}
          client={client}
          onSubmit={::this.save}
          new={addNew}/>
        <button className="btn btn-info" onClick={::this.toggleEdit}>Cancel</button>
      { addNew ? (<div/>) : (
        <button className="btn btn-danger" onClick={::this.delete}> Delete </button>
      )}
      </div>
      ) : (
        <div>
        Client {client.name}
        <button className="btn btn-info" onClick={::this.toggleEdit}> Edit </button>
        </div>
    )
  }
}


export default Client
