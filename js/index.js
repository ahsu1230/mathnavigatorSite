'use strict';
require('../styl/app.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import {
  HashRouter as Router,
  Route,
} from 'react-router-dom';
import { Header } from './header.js';
import { HomePage } from './home.js';
import { AnnouncePage } from './announce.js';
import { ProgramsPage } from './programs.js';
import { ContactPage } from './contact.js';
import { Footer } from './footer.js';

const Home = () => <HomePage/>;
const Announce = () => <AnnouncePage/>;
const Programs = () => <ProgramsPage/>;
const Contact = () => <ContactPage/>;


class MainContainer extends React.Component {
	render() {
		return (
      <Router onUpdate={() => window.scrollTo(0, 0)}>
        <div>
          <Header/>
          <Route exact path="/" component={Home}/>
          <Route path="/announcements" component={Announce}/>
          <Route path="/programs" component={Programs}/>
          <Route path="/contact" component={Contact}/>
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
