"use strict";
require("./home.sass");
import React from "react";
import API from "../api.js";
import { Tab, Tabs, TabList, TabPanel } from "react-tabs";
import "react-tabs/style/react-tabs.css";
import { HomeTabSectionClasses } from "./homeClasses.js";
import { HomeTabSectionUsers } from "./homeUsers.js";
import { HomeTabSectionRegistrations } from "./homeRegistrations.js";
import { HomeTabSectionAccounts } from "./homeAccounts.js";

export class HomePage extends React.Component {
    state = {
        unpublishedClasses: [],
        newUsers: [],
        newUserClasses: [],
        newUserAfh: [],
        unpaidAccounts: [],
    };

    componentDidMount = () => {
        API.get("api/classes/unpublished").then((res) => {
            const unpublishedList = res.data || [];
            this.setState({
                unpublishedClasses: unpublishedList,
            });
        });

        API.get("api/users/new").then((res) => {
            const users = res.data;
            this.setState({
                newUsers: users,
            });
        });

        API.get("api/user-classes/new").then((res) => {
            const userClass = res.data;
            this.setState({
                newUserClasses: userClass,
            });
        });

        API.get("api/user-afhs/new").then((res) => {
            const userAfh = res.data;
            this.setState({
                newUserAfh: userAfh,
            });
        });

        API.get("api/accounts/unpaid").then((res) => {
            const accounts = res.data;
            this.setState({
                unpaidAccounts: accounts,
            });
        });
    };

    render() {
        return (
            <div id="view-home">
                <h1>Administrator Dashboard</h1>
                <Tabs>
                    <TabList>
                        <Tab>New Users ({this.state.newUsers.length})</Tab>
                        <Tab>
                            New Registrations (
                            {this.state.newUserClasses.length +
                                this.state.newUserAfh.length}
                            )
                        </Tab>
                        <Tab>
                            Unpublished Classes (
                            {this.state.unpublishedClasses.length})
                        </Tab>
                        <Tab>
                            Unpaid Accounts ({this.state.unpaidAccounts.length})
                        </Tab>
                    </TabList>

                    <TabPanel>
                        <HomeTabSectionUsers users={this.state.newUsers} />
                    </TabPanel>
                    <TabPanel>
                        <HomeTabSectionRegistrations
                            newUserClasses={this.state.newUserClasses}
                            newUserAfh={this.state.newUserAfh}
                        />
                    </TabPanel>
                    <TabPanel>
                        <HomeTabSectionClasses
                            unpublishedClasses={this.state.unpublishedClasses}
                        />
                    </TabPanel>
                    <TabPanel>
                        <HomeTabSectionAccounts
                            unpaidAccounts={this.state.unpaidAccounts}
                        />
                    </TabPanel>
                </Tabs>
            </div>
        );
    }
}
