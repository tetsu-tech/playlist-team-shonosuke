// import React from 'react';
import { NotFound } from "./components/NotFound";
import "./App.css";
import { BrowserRouter, Link } from "react-router-dom";
import { login } from "components/login";

function App() {
  return (
    <BrowserRouter>
      <h1>Hello React Router</h1>
      <Link to="/404" component={NotFound}></Link>
      <Link to="/login" component={login} />
    </BrowserRouter>
  );
}

export default App;
