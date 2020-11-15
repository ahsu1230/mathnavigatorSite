"use strict";
import React from "react";
import API from "../../api.js";
import { getFullName } from "../../common/userUtils.js";
import UserSelector from "./userSelector.js";

export default class UserAfhs extends React.Component {
    render() {
        return (
            <section>
                <h2>User AFH Registrations</h2>
            </section>
        );
    }
}
