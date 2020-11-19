"use strict";
require("./app.sass");
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
import { AskForHelpPage } from "./askForHelp/afh.js";
import { AskForHelpEditPage } from "./askForHelp/afhEdit.js";

// import { UserPage } from "./user/user.js";
// import { UserEditPage } from "./user/userEdit.js";
import { AccountSearchPage } from "./account/accountSearch.js";
import { AccountCreatePage } from "./account/accountCreate.js";
import { AccountViewPage } from "./account/accountView.js";
import { UserEditPage } from "./account/tabs/userEdit.js";
import { UserMovePage } from "./account/tabs/userMove.js";
import { TransactionEditPage } from "./account/tabs/transactionEdit.js";
import { UserClassesPage } from "./account/userClassesPage.js";
import { UserAFHPage } from "./account/userAfhPage.js";

import { HelpPage } from "./help/help.js";

// import { EmailPaymentsPage } from "./emailPayment/emailPayments.js";
// import { EmailProgramsPage } from "./emailProgram/emailPrograms.js";

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
const AFH = () => <AskForHelpPage />;
const AFHEdit = () => <AskForHelpEditPage />;
const AFHMatch = ({ match }) => (
    <AskForHelpEditPage afhId={match.params.afhId} />
);

// const User = () => <UserPage />;
// const UserEditMatch = ({ match }) => <UserEditPage id={match.params.id} />;
// const UserAddMatch = ({ match }) => (
//     <UserEditPage accountId={match.params.accountId} />
// );
const AccountCreate = () => <AccountCreatePage />;
const AccountSearch = () => <AccountSearchPage />;
const AccountView = ({ match }) => (
    <AccountViewPage accountId={match.params.id} />
);
const UserAdd = ({ match }) => (
    <UserEditPage accountId={match.params.accountId} />
);
const UserEdit = ({ match }) => (
    <UserEditPage accountId={match.params.accountId} userId={match.params.id} />
);
const UserMove = ({ match }) => (
    <UserMovePage accountId={match.params.accountId} userId={match.params.id} />
);
const TransactionAdd = ({ match }) => (
    <TransactionEditPage accountId={match.params.accountId} />
);
const TransactionEdit = ({ match }) => (
    <TransactionEditPage
        accountId={match.params.accountId}
        transactionId={match.params.id}
    />
);
const UserClasses = () => <UserClassesPage />;
const UserAfhs = () => <UserAFHPage />;

// const AccountTransactionEdit = () => <TransactionEditPage />;
// const AccountTransactionEditMatch = ({ match }) => (
//     <TransactionEditPage id={match.params.id} />
// );

const Help = () => <HelpPage />;

// const EmailPayments = () => <EmailPaymentsPage />;
// const EmailPrograms = () => <EmailProgramsPage />;

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
                        path="/achievements/:id/edit"
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

                    <Route path="/afh/:afhId/edit" component={AFHMatch} />
                    <Route path="/afh/add" component={AFHEdit} />
                    <Route path="/afh" component={AFH} />

                    {/* <Route path="/users/:id/edit" component={UserEditMatch} />
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
                    /> */}
                    {/* <Route path="/users" component={User} /> */}
                    {/* <Route
                        path="/accounts/transaction/:id/edit"
                        component={AccountTransactionEditMatch}
                    />
                    <Route
                        path="/accounts/transaction/add"
                        component={AccountTransactionEdit}
                    /> */}
                    <Route path="/accounts" component={AccountSearch} />

                    <Route
                        path="/account/:accountId/user/add"
                        component={UserAdd}
                    />
                    <Route
                        path="/account/:accountId/user/:id/edit"
                        component={UserEdit}
                    />
                    <Route
                        path="/account/:accountId/user/:id/move"
                        component={UserMove}
                    />
                    <Route
                        path="/account/:accountId/transaction/add"
                        component={TransactionAdd}
                    />
                    <Route
                        path="/account/:accountId/transaction/:id/edit"
                        component={TransactionEdit}
                    />
                    <Route path="/account/create" component={AccountCreate} />
                    <Route path="/account/:id" component={AccountView} />
                    <Route path="/users/classes" component={UserClasses} />
                    <Route path="/users/afhs" component={UserAfhs} />

                    <Route path="/help" component={Help} />

                    {/* <Route path="/emailPayments" component={EmailPayments} /> */}
                    {/* <Route path="/emailPrograms" component={EmailPrograms} /> */}
                </Switch>
            </div>
        );
    }
}

const AppWithRouter = withRouter(App);

ReactDOM.render(<AppContainer />, document.getElementById("root"));
