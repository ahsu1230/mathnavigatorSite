'use strict';
require('./app.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { withRouter } from 'react-router';
import {
    HashRouter as Router,
    Route,
    Switch
} from 'react-router-dom';
import { HeaderSection } from './header/header.js';
import { HomePage } from './home/home.js';
import { ProgramPage } from './programs/program.js';
import { ProgramEditPage } from './programs/programEdit.js';
import { AnnouncePage } from './announce/announce.js';
import { LocationPage } from './location/location.js';
import { LocationEditPage } from './location/locationEdit.js';

const Announce = () => <AnnouncePage/>;
const Header = () => <HeaderSection/>;
const Home = () => <HomePage/>;
const Programs = () => <ProgramPage/>;
const ProgramEdit = () => <ProgramEditPage/>;
const ProgramEditMatch = ({match}) => <ProgramEditPage programId={match.params.programId}/>;
const Location = () => <LocationPage/>;
const LocationEdit = () => <LocationEditPage/>;
const LocationEditMatch = ({match}) => <LocationEditPage locId={match.params.locId}/>;

class AppContainer extends React.Component {
	render() {
		return (
      <Router>
        <AppWithRouter/>
      </Router>
		);
	}
}

class App extends React.Component {
  render() {
    return (
      <div>
        <Header/>
        <Switch>
          <Route path="/" exact component={Home}/>
          <Route path="/program/:programId/edit" component={ProgramEditMatch}/>
          <Route path="/programs/add" component={ProgramEdit}/>
          <Route path="/programs" component={Programs}/>
          <Route path="/announcements" component={Announce}/>
          <Route path="/locations/:locId/edit" component={LocationEditMatch}/>
          <Route path="/locations/add" component={LocationEdit}/>
          <Route path="/locations" component={Location}/>
        </Switch>
      </div>
    );
  }
}

const AppWithRouter = withRouter(App);

ReactDOM.render(
  <AppContainer/>,
  document.getElementById('root')
);
