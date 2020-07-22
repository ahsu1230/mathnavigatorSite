"use strict";
require("./app.styl");
import React from "react";
import ReactDOM from "react-dom";
import { withRouter } from "react-router";
import { HashRouter as Router, Route, Switch } from "react-router-dom";
import { HeaderSection } from "./header/header.js";
import { HomePage } from "./home/home.js";
import { ProgramPage } from "./programs/program.js";
import { ProgramEditPage } from "./programs/programEdit.js";
import { ClassPage } from "./classes/class.js";
import { ClassEditPage } from "./classes/classEdit.js";
import { SessionPage } from "./session/session.js";
import { SessionEditPage } from "./session/sessionEdit.js";
import { AnnouncePage } from "./announce/announce.js";
import { AnnounceEditPage } from "./announce/announceEdit.js";
import { AchievePage } from "./achieve/achieve.js";
import { AchieveEditPage } from "./achieve/achieveEdit.js";
import { LocationPage } from "./location/location.js";
import { LocationEditPage } from "./location/locationEdit.js";
import { SemesterPage } from "./semester/semester.js";
import { SemesterEditPage } from "./semester/semesterEdit.js";

import { UserPage } from "./user/user.js";
import { UserEditPage } from "./user/userEdit.js";
import { UserClassPage } from "./user/userClass.js";
import { UserAFHPage } from "./user/userAFH.js";
import { AccountPage } from "./account/account.js";
import { AccountEditPage } from "./account/accountEdit.js";
import { TransactionEditPage } from "./account/accountTransactionEdit.js";

import { HelpPage } from "./help/help.js";
import { AskForHelpPage } from "./ask_for_help/afh.js";
import { AskForHelpEditPage } from "./ask_for_help/afhEdit.js";

const Header = () => <HeaderSection />;
const Home = () => <HomePage />;

const Programs = () => <ProgramPage />;
const ProgramEdit = () => <ProgramEditPage />;
const ProgramEditMatch = ({ match }) => (
    <ProgramEditPage programId={match.params.programId} />
);
const Class = () => <ClassPage />;
const ClassEdit = () => <ClassEditPage />;
const ClassEditMatch = ({ match }) => (
    <ClassEditPage classId={match.params.classId} />
);
const Session = () => <SessionPage />;
const SessionEditMatch = ({ match }) => (
    <SessionEditPage classId={match.params.classId} id={match.params.id} />
);
const Announce = () => <AnnouncePage />;
const AnnounceEdit = () => <AnnounceEditPage />;
const AnnounceEditMatch = ({ match }) => (
    <AnnounceEditPage announceId={match.params.announceId} />
);
const Achieve = () => <AchievePage />;
const AchieveEdit = () => <AchieveEditPage />;
const AchieveEditMatch = ({ match }) => (
    <AchieveEditPage id={match.params.id} />
);
const Location = () => <LocationPage />;
const LocationEdit = () => <LocationEditPage />;
const LocationEditMatch = ({ match }) => (
    <LocationEditPage locationId={match.params.locationId} />
);
const Semester = () => <SemesterPage />;
const SemesterEdit = () => <SemesterEditPage />;
const SemesterEditMatch = ({ match }) => (
    <SemesterEditPage semesterId={match.params.semesterId} />
);

const User = () => <UserPage />;
const UserEditMatch = ({ match }) => <UserEditPage id={match.params.id} />;
const UserAddMatch = ({ match }) => (
    <UserEditPage accountId={match.params.accountId} />
);
const UserClassMatch = ({ match }) => <UserClassPage id={match.params.id} />;
const UserAFHMatch = ({ match }) => <UserAFHPage id={match.params.id} />;
const Account = () => <AccountPage />;
const AccountEdit = () => <AccountEditPage />;
const AccountTransactionEdit = () => <TransactionEditPage />;
const AccountTransactionEditMatch = ({ match }) => (
    <TransactionEditPage id={match.params.id} />
);

const Help = () => <HelpPage />;
const AFH = () => <AskForHelpPage />;
const AFHEdit = () => <AskForHelpEditPage />;
const AFHMatch = ({ match }) => (
    <AskForHelpEditPage afhId={match.params.afhId} />
);

class AppContainer extends React.Component {
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
                <Header />
                <Switch>
                    <Route path="/" exact component={Home} />
                    <Route
                        path="/programs/:programId/edit"
                        component={ProgramEditMatch}
                    />
                    <Route path="/programs/add" component={ProgramEdit} />
                    <Route path="/programs" component={Programs} />
                    <Route
                        path="/classes/:classId/edit"
                        component={ClassEditMatch}
                    />
                    <Route path="/classes/add" component={ClassEdit} />
                    <Route path="/classes" component={Class} />
                    <Route
                        path="/sessions/:classId/:id/edit"
                        component={SessionEditMatch}
                    />
                    <Route path="/sessions" component={Session} />
                    <Route
                        path="/announcements/:announceId/edit"
                        component={AnnounceEditMatch}
                    />
                    <Route path="/announcements/add" component={AnnounceEdit} />
                    <Route path="/announcements" component={Announce} />
                    <Route
                        path="/achievements/:Id/edit"
                        component={AchieveEditMatch}
                    />
                    <Route path="/achievements/add" component={AchieveEdit} />
                    <Route path="/achievements" component={Achieve} />
                    <Route
                        path="/locations/:locationId/edit"
                        component={LocationEditMatch}
                    />
                    <Route path="/locations/add" component={LocationEdit} />
                    <Route path="/locations" component={Location} />
                    <Route
                        path="/semesters/:semesterId/edit"
                        component={SemesterEditMatch}
                    />
                    <Route path="/semesters/add" component={SemesterEdit} />
                    <Route path="/semesters" component={Semester} />

                    <Route path="/users/:id/edit" component={UserEditMatch} />
                    <Route
                        path="/users/:accountId/add"
                        component={UserAddMatch}
                    />
                    <Route
                        path="/users/:id/class/edit"
                        component={UserClassMatch}
                    />
                    <Route
                        path="/users/:id/afh/edit"
                        component={UserAFHMatch}
                    />
                    <Route path="/users" component={User} />
                    <Route
                        path="/accounts/transaction/:id/edit"
                        component={AccountTransactionEditMatch}
                    />
                    <Route
                        path="/accounts/transaction/add"
                        component={AccountTransactionEdit}
                    />
                    <Route path="/accounts/add" component={AccountEdit} />
                    <Route path="/accounts" component={Account} />

                    <Route path="/help" component={Help} />
                    <Route path="/afh/:afhId/edit" component={AFHMatch} />
                    <Route path="/afh/add" component={AFHEdit} />
                    <Route path="/afh" component={AFH} />
                </Switch>
            </div>
        );
    }
}

const AppWithRouter = withRouter(App);

ReactDOM.render(<AppContainer />, document.getElementById("root"));
