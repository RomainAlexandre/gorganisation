import React, { Component } from 'react';
import logo from '../images/logo.svg';
import '../css/App.css';

class App extends Component {
  render() {
    return (
      <div>
        <div className="App">
          <div className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
          </div>
          <p className="App-intro">
            To get started, edit <code>src/App.js</code> save and reload (in development mode).
          </p>
          <h1>Welcome to Buffalo [v0.8.0] and React [v15.4.2]!</h1>
          <div className="App-content">
            <h2>
              <b>Buffalo:</b>
              <a href="https://github.com/gobuffalo/buffalo"> Github </a>&amp;
              <a href="http://gobuffalo.io"> Documentation </a>
            </h2>
          </div>
          <h2>
            <b>React:&nbsp;&nbsp;&nbsp;</b>
            <a href="https://github.com/facebook/react"> Github </a>&amp;
            <a href="https://facebook.github.io/react/docs/hello-world.html"> Documentation</a>
          </h2>
        </div>
      </div>
    );
  }
}

export default App;