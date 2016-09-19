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

  delete(e) {
    e.preventDefault()
    this.setState({ edit: false })
    this.props.del(this.props.client)
  }


  render() {
    const { edit } = this.state
    const { client, addNew } = this.props
    return edit ? (
      <ClientForm
        formData={client}
        client={client}
        onSubmit={::this.save}
        new={addNew}>

        <div className="btn-group" role="group">
          <button type="submit" className="btn btn-info">Submit</button>
          <button className="btn btn-info">Cancel</button>
          { addNew ? (<div/>) : (
            <button className="btn btn-danger" onClick={::this.delete}> Delete </button>
          )}
        </div>
      </ClientForm>
    ) : (
      <button type="button" className="list-group-item" onClick={::this.toggleEdit}>
        {client.name}
      </button>
    )
  }
}


export default Client
