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
import { AchievePage } from './achieve/achieve.js';
import { AnnouncePage } from './announce/announce.js';
import { AnnounceEditPage } from './announce/announceEdit.js';

const Achieve = () => <AchievePage/>;
const Announce = () => <AnnouncePage/>;
const AnnounceEdit = () => <AnnounceEditPage/>;
const AnnounceEditMatch = ({match}) => <AnnounceEditPage programId={match.params.AnnounceId}/>;
const Header = () => <HeaderSection/>;
const Home = () => <HomePage/>;
const Programs = () => <ProgramPage/>;
const ProgramEdit = () => <ProgramEditPage/>;
const ProgramEditMatch = ({match}) => <ProgramEditPage programId={match.params.programId}/>;

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
          <Route path="/announce/:announceId/edit" component={AnnounceEditMatch}/>
          <Route path="/announce/add" component={AnnounceEdit}/>
          <Route path="/announcements" component={Announce}/>
          <Route path="/achievements" component={Achieve}/>
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
