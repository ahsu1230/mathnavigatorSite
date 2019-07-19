'use strict';
require('../styl/app.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import {
  HashRouter as Router,
  Route,
  Switch
} from 'react-router-dom';
import { Header } from './header.js';
import { HomePage } from './home.js';
import { AnnouncePage } from './announce.js';
import { ProgramsPage } from './programs.js';
import { ContactPage } from './contact.js';
import { Footer } from './footer.js';
import { ClassPage } from './class.js';
import { ErrorPage } from './error.js';

const Home = () => <HomePage/>;
const Announce = () => <AnnouncePage/>;
const Programs = () => <ProgramsPage/>;
const Contact = () => <ContactPage/>;
const ClassPageWithSlug = ({match}) => <ClassPage slug={match.params.slug}/>;
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
