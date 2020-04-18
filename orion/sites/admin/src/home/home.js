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
            rows: [],
            selection: "",
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
        let programs = this.state.programs.map((row, index) => {
            return <DashboardRow key={index} title={row.name} />;
        });
        this.setState({
            noUnpub: "",
            rows: programs,
            selection: "programs",
            switch: true,
        });
        if (programs.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
                switch: false,
            });
        }
    }

    onClickClasses() {
        let classes = this.state.classes.map((row, index) => {
            return <DashboardRow key={index} title={row.classId} />;
        });
        this.setState({
            noUnpub: "",
            rows: classes,
            selection: "classes",
            switch: true,
        });
        if (classes.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
                switch: false,
            });
        }
    }

    onClickLocations() {
        let locations = this.state.locations.map((row, index) => {
            return <DashboardRow key={index} title={row.locId} />;
        });
        this.setState({
            noUnpub: "",
            rows: locations,
            selection: "locations",
            switch: true,
        });
        if (locations.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
                switch: false,
            });
        }
    }

    onClickAchievements() {
        let achieves = this.state.achieves.map((row, index) => {
            return <DashboardRow key={index} title={row.message} />;
        });
        this.setState({
            noUnpub: "",
            rows: achieves,
            selection: "achievements",
            switch: true,
        });
        if (achieves.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
                switch: false,
            });
        }
    }

    onClickSemesters() {
        let semesters = this.state.semesters.map((row, index) => {
            return <DashboardRow key={index} title={row.semesterId} />;
        });
        this.setState({
            noUnpub: "",
            rows: semesters,
            selection: "semesters",
            switch: true,
        });
        if (semesters.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
                switch: false,
            });
        }
    }

    onClickSessions() {
        let sessions = this.state.sessions.map((row, index) => {
            return <DashboardRow key={index} title={row.sessionId} />;
        });
        this.setState({
            noUnpub: "",
            rows: sessions,
            selection: "sessions",
            switch: true,
        });
        if (sessions.length == 0) {
            this.setState({
                noUnpub: "All items are published! Your website is up to date!",
                switch: false,
            });
        }
    }

    onClickPage() {
        console.log("Go to Page clicked");
    }

    render() {
        let unpubMessage = <div> </div>;
        if (this.state.switch) {
            unpubMessage = (
                <h4>
                    {" "}
                    You have {this.state.rows.length} unpublished{" "}
                    {this.state.selection}{" "}
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
                        {this.state.rows}
                        {this.state.noUnpub}
                    </div>
                    <button id="go-to-page" onClick={this.onClickPage}>
                        Go to Page
                    </button>
                </div>
            </div>
        );
    }
}

class DashboardRow extends React.Component {
    render() {
        const rows = this.props.title;
        return <div className="dashboard-row">{rows}</div>;
    }
}
