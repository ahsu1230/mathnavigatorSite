"use strict";
require("./account.sass");
import React from "react";
import API from "../utils/api.js";
import { union } from "lodash";

import { SettingsTab } from "./settings.js";
import {
    RegistrationsTab,
    RegistrationsTabMain,
    RegistrationsTabAllClasses,
} from "./registrations.js";
import { PaymentTab } from "./payment.js";

export class AccountPage extends React.Component {
    state = {
        id: 0,
        selectedTab: "settings",

        users: [],
    };

    componentDidMount = () => {
        const id = 1; // TODO: Replace with signin later
        this.setState({ id: id });
        if (id) {
            API.get("api/users/account/" + id)
                .then((res) => {
                    const users = res.data;
                    this.setState({ users: users });
                })
                .catch((err) => alert("Could not fetch data: " + err));
        }
    };

    onTabSelect = (tab) => {
        this.setState({ selectedTab: tab.toLowerCase() });
    };

    render = () => {
        const tabButtons = ["Settings", "Registrations", "Payment"].map(
            (item, index) => {
                return (
                    <div
                        className={
                            this.state.selectedTab == item.toLowerCase()
                                ? "selected"
                                : ""
                        }
                        onClick={(e) => this.onTabSelect(item)}
                        key={index}>
                        {item}
                    </div>
                );
            }
        );

        let tabContent;
        switch (this.state.selectedTab) {
            case "settings":
                tabContent = (
                    <SettingsTab
                        accountId={this.state.id}
                        users={this.state.users}
                        key={this.state.id}
                    />
                );
                break;
            case "registrations":
                tabContent = (
                    <RegistrationsTab
                        accountId={this.state.id}
                        users={this.state.users}
                        key={this.state.id}
                    />
                );
                break;
            case "payment":
                tabContent = (
                    <PaymentTab accountId={this.state.id} key={this.state.id} />
                );
                break;
        }

        return (
            <div id="view-account">
                <h1>Your Account</h1>

                <div id="tab-container">{tabButtons}</div>

                {tabContent}
            </div>
        );
    };
}
