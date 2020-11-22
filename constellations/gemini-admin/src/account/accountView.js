"use strict";
require("./accountView.sass");
import React from "react";
import { keyBy } from "lodash";
import queryString from "query-string";
import { Tab, Tabs, TabList, TabPanel } from "react-tabs";
import "react-tabs/style/react-tabs.css";
import API from "../api.js";
import AccountInfo from "./tabs/accountInfo.js";
import UserInfos from "./tabs/accountUsers.js";
import UserClasses from "./tabs/accountUserClasses.js";
import UserAfhs from "./tabs/accountUserAfhs.js";
import AccountTransactions from "./tabs/accountTransactions.js";

const TAB_VAL_ACCOUNT = "account";
const TAB_VAL_EDIT_USERS = "edit-users";
const TAB_VAL_USER_CLASSES = "user-classes";
const TAB_VAL_USER_AFH = "user-afh";
const TAB_VAL_TRANSACTIONS = "transactions";
const TabIndexValues = {
    [TAB_VAL_ACCOUNT]: 0,
    [TAB_VAL_EDIT_USERS]: 1,
    [TAB_VAL_USER_CLASSES]: 2,
    [TAB_VAL_USER_AFH]: 3,
    [TAB_VAL_TRANSACTIONS]: 4,
};

export class AccountViewPage extends React.Component {
    state = {
        users: [],
        userMap: {},
        selectedUserId: 0,
        tabIndex: 0,
    };

    componentDidMount() {
        const accountId = this.props.accountId;
        API.get("api/users/account/" + accountId)
            .then((res) => {
                const users = res.data || [];

                this.setState({
                    users: users,
                    userMap: keyBy(users, (user) => user.id),
                    selectedUserId: users[0].id,
                });
            })
            .catch((err) => {
                console.log("Error searching account " + err);
            })
            .finally(() => {
                // Identify correct tab to be on
                let indexOfQuery = location.hash.indexOf("?");
                let tabView = "";
                if (indexOfQuery >= 0) {
                    tabView = queryString.parse(
                        location.hash.substring(indexOfQuery)
                    ).view;
                    this.setState({
                        tabIndex: TabIndexValues[tabView] || 0,
                    });
                }
            });
    }

    onSelectTab = (index, lastIndex, e) => {
        if (index != lastIndex) {
            this.setState({
                tabIndex: index,
            });
            return true;
        }
        return false;
    };

    onSwitchUser = (newUserId) => {
        this.setState({
            selectedUserId: newUserId,
        });
    };

    render() {
        const accountId = this.props.accountId;
        const users = this.state.users;
        const userMap = this.state.userMap;
        const selectedUserId = this.state.selectedUserId;
        return (
            <main id="view-account">
                <h1>Account No. {accountId}</h1>
                <Tabs
                    onSelect={this.onSelectTab}
                    selectedIndex={this.state.tabIndex}>
                    <TabList>
                        <Tab>Account</Tab>
                        <Tab>Edit Users</Tab>
                        <Tab>User Classes</Tab>
                        <Tab>User AskForHelp</Tab>
                        <Tab>Transactions</Tab>
                    </TabList>

                    <TabPanel>
                        <AccountInfo accountId={accountId} />
                    </TabPanel>
                    <TabPanel>
                        <UserInfos
                            accountId={accountId}
                            users={users}
                            selectedUser={userMap[selectedUserId]}
                            onSwitchUser={this.onSwitchUser}
                        />
                    </TabPanel>
                    <TabPanel>
                        <UserClasses
                            accountId={accountId}
                            users={users}
                            selectedUser={userMap[selectedUserId]}
                            onSwitchUser={this.onSwitchUser}
                        />
                    </TabPanel>
                    <TabPanel>
                        <UserAfhs
                            accountId={accountId}
                            users={users}
                            selectedUser={userMap[selectedUserId]}
                            onSwitchUser={this.onSwitchUser}
                        />
                    </TabPanel>
                    <TabPanel>
                        <AccountTransactions accountId={accountId} />
                    </TabPanel>
                </Tabs>
            </main>
        );
    }
}
