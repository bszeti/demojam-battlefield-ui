import React from 'react';

import {
  HashRouter as Router,
  Route
  } from 'react-router-dom';
import './App.css';
import LiveScore from './components/livescore';
import Welcome from './components/welcome';
import "./assets/scss/black-dashboard-react.scss";
import "./assets/demo/demo.css";

function App() {
  return (
    <Router>
        <Route exact path="/" component={Welcome} />
        <Route path="/dashboard/:game" component={LiveScore} />
    </Router>
  );
}

export default App;


