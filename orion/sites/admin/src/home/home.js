"use strict";
require("./home.styl");
import React from "react";
import ReactDOM from "react-dom";
import API from "../api.js";

export class HomePage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            programs: [],
            classes: [],
            locations: [],
            achieves: [],
            semesters: [],
            sessions: [],
            row: [],
            selection: "",
            num: 0,
            noUnpub: "",
            switch: false,
        };
        this.onClickPrograms = this.onClickPrograms.bind(this);
        this.onClickClasses = this.onClickClasses.bind(this);
        this.onClickLocations = this.onClickLocations.bind(this);
        this.onClickAchievements = this.onClickAchievements.bind(this);
        this.onClickSemesters = this.onClickSemesters.bind(this);
        this.onClickSessions = this.onClickSessions.bind(this);
    }

    componentDidMount() {
        API.get("api/v1/unpublished").then((res) => {
            const unpublishedList = res.data;
            this.setState({
                programs: unpublishedList.programs,
                classes: unpublishedList.classes,
                locations: unpublishedList.locations,
                achieves: unpublishedList.achieves,
                semesters: unpublishedList.semesters,
                sessions: unpublishedList.sessions,
            });
        });
    }

    onClickPrograms() {
        this.setState({ noUnpub: "" });
        let programs = this.state.programs.map((row, index) => {
            return <DashboardRow key={index} row={row} />;
        });
        this.setState({ row: programs });
        this.setState({ selection: "programs" });
        this.setState({ num: programs.length });
        this.state.switch = true;
        if (programs.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
            });
            this.state.switch = false;
        }
    }

    onClickClasses() {
        this.setState({ noUnpub: "" });
        let classes = this.state.classes.map((row, index) => {
            return <DashboardRow key={index} row={row} />;
        });
        this.setState({ row: classes });
        this.setState({ selection: "classes" });
        this.setState({ num: classes.length });
        this.state.switch = true;
        if (classes.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
            });
            this.state.switch = false;
        }
    }

    onClickLocations() {
        this.setState({ noUnpub: "" });
        let locations = this.state.locations.map((row, index) => {
            return <DashboardRow key={index} row={row} />;
        });
        this.setState({ row: locations });
        this.setState({ selection: "locations" });
        this.setState({ num: locations.length });
        this.state.switch = true;
        if (locations.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
            });
            this.state.switch = false;
        }
    }

    onClickAchievements() {
        this.setState({ noUnpub: "" });
        let achieves = this.state.achieves.map((row, index) => {
            return <DashboardRow key={index} row={row} />;
        });
        this.setState({ row: achieves });
        this.setState({ selection: "achievements" });
        this.setState({ num: achieves.length });
        this.state.switch = true;
        if (achieves.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
            });
            this.state.switch = false;
        }
    }

    onClickSemesters() {
        this.setState({ noUnpub: "" });
        let semesters = this.state.semesters.map((row, index) => {
            return <DashboardRow key={index} row={row} />;
        });
        this.setState({ row: semesters });
        this.setState({ selection: "semesters" });
        this.setState({ num: semesters.length });
        this.state.switch = true;
        if (semesters.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
            });
            this.state.switch = false;
        }
    }

    onClickSessions() {
        this.setState({ noUnpub: "" });
        let sessions = this.state.sessions.map((row, index) => {
            return <DashboardRow key={index} row={row} />;
        });
        this.setState({ row: sessions });
        this.setState({ selection: "sessions" });
        this.setState({ num: sessions.length });
        this.state.switch = true;
        if (sessions.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
            });
            this.state.switch = false;
        }
    }

    render() {
        let unpubMessage = <div> </div>;
        if (this.state.switch) {
            unpubMessage = (
                <h4>
                    {" "}
                    You have {this.state.num} unpublished {this.state.selection}{" "}
                </h4>
            );
        }
        return (
            <div id="view-home">
                <div id="view-content">
                    <h2 id="unpublished-heading"> Unpublished Content </h2>
                    <ul>
                        <button onClick={this.onClickPrograms}>Programs</button>
                        <button onClick={this.onClickClasses}>Classes</button>
                        <button onClick={this.onClickLocations}>
                            Locations
                        </button>
                        <button onClick={this.onClickAchievements}>
                            Achievements
                        </button>
                        <button onClick={this.onClickSemesters}>
                            Semesters
                        </button>
                        <button onClick={this.onClickSessions}>Sessions</button>
                    </ul>
                    <h2> Registrations </h2>
                    <ul>
                        <button>New Users</button>
                        <button>Questions</button>
                        <button>Complaints</button>
                    </ul>
                </div>
                <div id="box-and-button">
                    <div className="boxed">
                        {unpubMessage}
                        {this.state.row}
                        {this.state.noUnpub}
                    </div>
                    <button id="go-to-page">Go to Page</button>
                </div>
            </div>
        );
    }
}

class DashboardRow extends React.Component {
    render() {
        const row = this.props.row;
        return (
            <div id="dashboard-row">
                <div>
                    {row.name}
                    {row.classId}
                    {row.locId}
                    {row.message}
                    {row.semesterId}
                    {row.sessionId}
                </div>
            </div>
        );
    }
}
