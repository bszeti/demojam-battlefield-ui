import React from 'react';

import {
  HashRouter as Router,
  Route, Link
  } from 'react-router-dom';
import './App.css';
import LiveScore from './components/livescore';
import "./assets/scss/black-dashboard-react.scss";
import "./assets/demo/demo.css";

function App() {
  return (
    <Router>
        <Route exact path="/" component={Welcome} />
        <Route path="/dashboard" component={LiveScore} />
        <Route path="/finish" component={Topics} />
    </Router>
  );
}

export default App;

class Welcome extends React.Component {
  render () {
    return (
      <div className="App">
        <header className="App-header">
          <img src={require("./assets/img/battle-icon-10.jpg")} className="App-logo" alt="logo" />
          <p>
            GREAT BATTLE OF KILLER PODS
          </p>
          <Link className="App-link" to={`dashboard`}>Start</Link>
        </header>
      </div>
    );
  }
}


function Topics({ match }) {
  return (
    <div>
      <h2>Topics</h2>
      <ul>
        <li>
          <Link to={`${match.url}/rendering`}>Rendering with React</Link>
        </li>
        <li>
          <Link to={`${match.url}/components`}>Components</Link>
        </li>
        <li>
          <Link to={`${match.url}/props-v-state`}>Props v. State</Link>
        </li>
      </ul>

      <Route path={`${match.path}/:topicId`} component={Topic} />
      <Route
        exact
        path={match.path}
        render={() => <h3>Please select a topic.</h3>}
      />
    </div>
  );
}

function Topic({ match }) {
  return (
    <div>
      <h3>{match.params.topicId}</h3>
    </div>
  );
}