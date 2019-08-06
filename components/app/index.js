'use strict';
require('./app.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import {
  Router,
  Route,
  Switch
} from 'react-router-dom';
import { history } from './history.js';
import { createPageTitle, getNavByUrl } from '../constants.js';

import { Header } from '../header/header.js';
import { HomePage } from '../home/home.js';
import { AchievementPage } from '../achievements/achievements.js';
import { AFHPage } from '../afh/afh.js';
import { AnnouncePage } from '../announcements/announce.js';
import { ProgramsPage } from '../programs/programs.js';
import { ContactPage } from '../contact/contact.js';
import { Footer } from '../footer/footer.js';
import { ClassPage } from '../class/class.js';
import { ErrorPage } from '../errorPage/error.js';

const Home = () => <HomePage/>;
const Announce = () => <AnnouncePage/>;
const Programs = () => <ProgramsPage/>;
const Contact = () => <ContactPage/>;
const ClassPageWithSlug = ({match}) => <ClassPage slug={match.params.slug}/>;
const Achievements = () => <AchievementPage/>;
const AFH = () => <AFHPage/>;
const Error = () => <ErrorPage/>;


history.listen((location, action) => {
  var nav = getNavByUrl(location.pathname);
  if (nav) {
    document.title = createPageTitle(nav.name);
  } // if not in Nav, component must set it's own title!
});

class AppContainer extends React.Component {
	render() {
		return (
      <Router history={history}>
        <App/>
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
          <Route path="/announcements" component={Announce}/>
          <Route path="/programs" component={Programs}/>
          <Route path="/contact" component={Contact}/>
          <Route path="/class/:slug" component={ClassPageWithSlug}/>
          <Route path="/askforhelp" component={AFH}/>
          <Route path="/student-achievements" component={Achievements}/>
          <Route path="/" component={Error}/>
        </Switch>
        <Footer/>
      </div>
    );
  }
}

ReactDOM.render(
  <AppContainer/>,
  document.getElementById('root')
);
