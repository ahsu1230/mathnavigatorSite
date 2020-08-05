"use strict";
require("./base.sass");
import React from "react";
import ReactDOM from "react-dom";
import { withRouter, hashHistory } from "react-router";
import {
    HashRouter as Router,
    // By switching back to HashRouter, we lose functionality (scrollMemory)
    Route,
    Switch,
} from "react-router-dom";
import { history } from "./history.js";
// import { createPageTitle, getNavByUrl } from '../constants.js';
import ScrollMemory from "react-router-scroll-memory"; // Requires BrowserRouter

import { AchievementPage } from "../achievements/achievements.js";
import { InternshipPage } from "../internship/internship.js";
import { AFHPage } from "../afh/afh.js";
import { AnnouncePage } from "../announcements/announce.js";
// import { ClassPage } from '../class/class.js';
// import { ContactPage } from '../contact/contact.js';
// import { ErrorPage } from '../errorPage/error.js';
import Footer from "../footer/footer.js";
import { Header as HeaderComponent } from "../header/header.js";
import { HomePage } from "../home/home.js";
// import { ProgramsPage } from '../programs/programs.js';
// import { StudentProjectsPage } from '../student/studentProjects.js';
// import { StudentWebDevPage } from '../student/studentWebDev.js';
import { AccountPage } from "../account/account.js";

const Achievements = () => <AchievementPage />;
const AFH = () => <AFHPage />;
const Internship = () => <InternshipPage />;
const Announce = () => <AnnouncePage />;
// const ClassPageWithSlug = ({match}) => <ClassPage slug={match.params.slug}/>;
// const Contact = () => <ContactPageRouter/>;
// const ContactPageRouter = withRouter(ContactPage);
const Header = withRouter(HeaderComponent);
const Home = () => <HomePage />;
// const Programs = () => <ProgramsPage/>;
// const StudentWebDev = () => <StudentWebDevPage/>;
// const StudentProjects = () => <StudentProjectsPage/>;
// const AFH = () => <AFHPage/>;
// const Error = () => <ErrorPage/>;
const Account = () => <AccountPage />;

class AppContainer extends React.Component {
    render() {
        return (
            <Router>
                <ScrollMemory />
                <AppWithRouter />
            </Router>
        );
    }
}

class App extends React.Component {
    // componentDidMount() {
    //     this.props.history.listen((location, action) => {
    //         var nav = getNavByUrl(location.pathname);
    //         if (nav) {
    //             document.title = createPageTitle(nav.name);
    //         }
    //         // if not in Nav, component must set it's own title!
    //     });
    // }

    render() {
        return (
            <div>
                <Header />
                <Switch>
                    <Route path="/" exact component={Home} />

                    <Route path="/announcements" component={Announce} />
                    <Route path="/ask-for-help" component={AFH} />
                    {/* <Route path="/programs" component={Programs}/>
          <Route path="/contact" component={Contact}/>
          <Route path="/class/:slug" component={ClassPageWithSlug}/>
         */}
                    <Route
                        path="/student-achievements"
                        component={Achievements}
                    />

                    <Route path="/internship" component={Internship} />
                    {/* <Route path="/student-webdev" component={StudentWebDev}/>
          <Route path="/student-projects" component={StudentProjects}/>
          <Route path="/" component={Error}/> */}

                    <Route path="/account" component={Account} />
                </Switch>
                <Footer />
            </div>
        );
    }
}

const AppWithRouter = withRouter(App);

ReactDOM.render(<AppContainer />, document.getElementById("root"));
