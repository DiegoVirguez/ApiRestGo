import React from 'react';
import './App.css';

import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  useRouteMatch,
  useParams
} from "react-router-dom";
import "./css/bootstrap/css/bootstrap.min.css";
import "./css/styles.css";

//components
import ProductoComponent from './Components/ProductoComponent';
import CrearProductoComponent from './Components/CrearProducto';
import VerProductoComponent from './Components/VerProducto';
import EditarProductoComponent from './Components/EditarProducto';


function App() {
  return (
   <Router>
    <div id="content" className="col-lg-10 col-lg-offset-1" >
  <header id="header" className="col-lg-12"  >
    <h1>
      MarketElectronics
    </h1>
  </header>

  <nav id="nav" className="col-lg-12">
    <ul>
      <li>
      <Link to="/">Productos</Link>
      </li>
      <li>
      <Link to="/crearProducto">Crear producto</Link>
      </li>
    </ul>
  </nav>
   
   <section id="main" className="col-lg-12">
   <Switch>
          <Route exact path="/">
          <ProductoComponent/>
          </Route>
          <Route path="/crearProducto">
          <CrearProductoComponent/>
          </Route>
          <Route path="/verProducto/:codigo">
          <VerProductoComponent/>
          </Route>
            <Route path="/editarProducto/:codigo">
          <EditarProductoComponent/>
          </Route>
   </Switch>
  </section>
 

  <footer id="footer" className="col-lg-12">
    <p>Market</p>
  </footer>
</div>
</Router>
  );
}
export default App;
