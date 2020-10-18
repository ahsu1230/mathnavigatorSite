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
import ScrollMemory from "react-router-scroll-memory"; // Requires BrowserRouter
import { getLinkByUrl } from "../utils/links.js";

import { AchievementPage } from "../achievements/achievements.js";
import { AFHPage } from "../afh/afh.js";
import { AnnouncePage } from "../announcements/announce.js";
import { ClassPage } from "../class/class.js";
import { Header as HeaderComponent } from "../header/header.js";
import { HomePage } from "../home/home.js";
import { InternshipPage } from "../internship/internship.js";
import { ProgramsPage } from "../programs/programs.js";
import RegisterPage from "../register/register.js";
import RegisterSuccessPage from "../register/registerSuccess.js";
import Footer from "../footer/footer.js";

const Achievements = () => <AchievementPage />;
const AFH = () => <AFHPage />;
const Announce = () => <AnnouncePage />;
const Class = ({ match }) => <ClassPage classId={match.params.classId} />;
const Header = withRouter(HeaderComponent);
const Home = () => <HomePage />;
const Internship = () => <InternshipPage />;
const Programs = () => <ProgramsPage />;
const Register = withRouter(RegisterPage);
const RegisterSuccessAfh = () => <RegisterSuccessPage registered="afh" />;
const RegisterSuccessClass = () => <RegisterSuccessPage registered="class" />;

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
    componentDidMount() {
        this.props.history.listen((location, action) => {
            let nav = getLinkByUrl(location.pathname);
            if (nav) {
                document.title = nav.name;
            } else {
                document.title = "Math Navigator";
            }
        });
    }

    render() {
        return (
            <div>
                <Header />
                <Switch>
                    <Route path="/" exact component={Home} />
                    <Route path="/announcements" component={Announce} />
                    <Route path="/ask-for-help" component={AFH} />
                    <Route path="/class/:classId" component={Class} />
                    <Route path="/internship" component={Internship} />
                    <Route path="/programs" component={Programs} />
                    <Route path="/register" component={Register} />
                    <Route
                        path="/register-success/afh"
                        component={RegisterSuccessAfh}
                    />
                    <Route
                        path="/register-success/class"
                        component={RegisterSuccessClass}
                    />
                    <Route
                        path="/student-achievements"
                        component={Achievements}
                    />
                </Switch>
                <Footer />
            </div>
        );
    }
}

const AppWithRouter = withRouter(App);

ReactDOM.render(<AppContainer />, document.getElementById("root"));
