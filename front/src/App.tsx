// import React from 'react';
import { NotFound } from "./components/NotFound";
import "./App.css";
import { BrowserRouter, Link } from "react-router-dom";

function App() {
  return (
    <BrowserRouter>
      <h1>Hello React Router</h1>
      <Link to="/404" component={NotFound}>
        NotFoud Page
      </Link>
    </BrowserRouter>
  );
}

export default App;
