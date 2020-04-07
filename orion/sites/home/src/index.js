"use strict";
// require('./app.styl');
import React from "react";
import ReactDOM from "react-dom";
import { withRouter } from "react-router";
import { HashRouter as Router, Route, Switch } from "react-router-dom";
import { HomePage } from "./home/home.js";

const Home = () => <HomePage />;

class AppContainer extends React.Component {
<<<<<<< HEAD
<<<<<<< HEAD
  render() {
    return (
      <Router>
        <AppWithRouter />
      </Router>
    );
  }
}

class App extends React.Component {
  render() {
    return (
      <div>
        <Switch>
          <Route path="/" exact component={Home} />
        </Switch>
      </div>
    );
  }
=======
=======
>>>>>>> c15f24dc4318ffae807d39aef3ef62f1b6948b26
    render() {
        return (
            <Router>
                <AppWithRouter />
            </Router>
        );
    }
}

class App extends React.Component {
    render() {
        return (
            <div>
                <Switch>
                    <Route path="/" exact component={Home} />
                </Switch>
            </div>
        );
    }
<<<<<<< HEAD
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
=======
>>>>>>> c15f24dc4318ffae807d39aef3ef62f1b6948b26
}

const AppWithRouter = withRouter(App);

ReactDOM.render(<AppContainer />, document.getElementById("root"));
