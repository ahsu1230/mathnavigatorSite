'use strict';
require('./app.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { withRouter, hashHistory } from 'react-router';
import {
  HashRouter as Router,
  // By switching back to HashRouter, we lose functionality (scrollMemory)
  Route,
  Switch
} from 'react-router-dom';
import { history } from './history.js';
import { createPageTitle, getNavByUrl } from '../constants.js';
import ScrollMemory from 'react-router-scroll-memory'; // Requires BrowserRouter

import { Header as HeaderComponent} from '../header/header.js';
import { HomePage } from '../home/home.js';
import { AchievementPage } from '../achievements/achievements.js';
import { AFHPage } from '../afh/afh.js';
import { AnnouncePage } from '../announcements/announce.js';
import { ProgramsPage } from '../programs/programs.js';
import { ContactPage } from '../contact/contact.js';
import { Footer } from '../footer/footer.js';
import { ClassPage } from '../class/class.js';
import { ErrorPage } from '../errorPage/error.js';

const Header = withRouter(HeaderComponent);
const Home = () => <HomePage/>;
const Announce = () => <AnnouncePage/>;
const Programs = () => <ProgramsPage/>;
const ContactPageRouter = withRouter(ContactPage);
const Contact = () => <ContactPageRouter/>;
const ClassPageWithSlug = ({match}) => <ClassPage slug={match.params.slug}/>;
const Achievements = () => <AchievementPage/>;
const AFH = () => <AFHPage/>;
const Error = () => <ErrorPage/>;

class AppContainer extends React.Component {
	render() {
		return (
      <Router>
        <ScrollMemory/>
        <AppWithRouter/>
      </Router>
		);
	}
}

class App extends React.Component {
  componentDidMount() {
      this.props.history.listen((location, action) => {
        var nav = getNavByUrl(location.pathname);
        if (nav) {
          document.title = createPageTitle(nav.name);
        }
        // if not in Nav, component must set it's own title!
      });
  }

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

const AppWithRouter = withRouter(App);

ReactDOM.render(
  <AppContainer/>,
  document.getElementById('root')
);
