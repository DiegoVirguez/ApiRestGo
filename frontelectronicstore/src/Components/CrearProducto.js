import React from 'react';


class CrearProductoComponent extends React.Component{
    constructor(props) {
    super(props);
    this.state = {
      nombre: "",
      precio: 0,
      caracteristicas: ""
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    const target = event.target;
    const name = target.name;
    const value = name != "precio" ? target.value : parseInt(target.value);
    this.setState({
      [name]: value
    });
  }

  handleSubmit(event) {
    event.preventDefault();
    fetch('http://localhost:8081/productos',{
            method: "POST",
            body: JSON.stringify(this.state)
        })
        .then(result=>result.json())
        .then(items=>this.setState({
            done: true,
            items
         }))
        .catch(() => {
            this.setState({
                done: true,
                success: false
            })
        })
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit} className="col-lg-8" id="form-producto">
        <label>
          Nombre
    </label>
    <input type="text" name="nombre"  onChange={this.handleChange} className="form-control" required />
    <label>
          Precio
    </label>
    <input type="text" name="precio"  onChange={this.handleChange} className="form-control" required />
    <label>
          caracteristicas:
          
        </label>
        <textarea onChange={this.handleChange} name="caracteristicas" />
        <input type="submit" value="Guardar Producto" className="btn btn-success"/>
      </form>
    );
  }
}

export default CrearProductoComponent