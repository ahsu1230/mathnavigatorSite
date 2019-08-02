'use strict';
import 'core-js/es/map';
import 'core-js/es/set';
import 'babel-polyfill';
import 'react-app-polyfill/ie9';
import 'react-app-polyfill/ie11';
import 'react-app-polyfill/stable';

require('./app.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import {
  HashRouter as Router,
  Route,
  Switch
} from 'react-router-dom';
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

class MainContainer extends React.Component {
	render() {
		return (
      <Router onUpdate={() => window.scrollTo(0, 0)}>
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
      </Router>
		);
	}
}

ReactDOM.render(
  <MainContainer/>,
  document.getElementById('root')
);
