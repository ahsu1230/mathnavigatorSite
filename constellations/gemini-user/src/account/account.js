"use strict";
require("./account_base.sass");
import React from "react";
import API from "../utils/api.js";
import { union } from "lodash";

import { SettingsTab } from "./settings.js";
import { RegistrationsTab } from "./registrations.js";
import { PaymentTab } from "./payment.js";

const TAB_SETTINGS = "settings";
const TAB_REGISTRATIONS = "registrations";
const TAB_PAYMENT = "payment";

export class AccountPage extends React.Component {
    state = {
        id: 0,
        selectedTab: TAB_SETTINGS,

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
            case TAB_SETTINGS:
                tabContent = (
                    <SettingsTab
                        accountId={this.state.id}
                        users={this.state.users}
                        key={this.state.id}
                    />
                );
                break;
            case TAB_REGISTRATIONS:
                tabContent = (
                    <RegistrationsTab
                        accountId={this.state.id}
                        users={this.state.users}
                        key={this.state.id}
                    />
                );
                break;
            case TAB_PAYMENT:
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
