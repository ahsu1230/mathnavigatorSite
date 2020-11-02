"use strict";
require("./achievements.sass");
import React from "react";
import { Link } from "react-router-dom";
import API from "../utils/api.js";
import { keys, sortBy, groupBy } from "lodash";
import srcSuccess from "../../assets/map_meet.jpg";

export class AchievementPage extends React.Component {
    state = {
        achieveList: [],
        groupedByYear: {},
        years: [],
    };

    componentDidMount() {
        API.get("api/achievements/all").then((res) => {
            const achieveList = res.data;
            let groupedByYear = groupBy(achieveList, (a) => a.year);
            let years = sortBy(keys(groupedByYear)).reverse();
            this.setState({
                achieveList: achieveList,
                groupedByYear: groupedByYear,
                years: years,
            });
        });
    }

    render() {
        const grouped = this.state.years.map((year, index) => (
            <GroupedItems
                key={index}
                year={year}
                achievements={this.state.groupedByYear[year]}
            />
        ));

        return (
            <div id="view-achieve">
                <div id="banner-container">
                    <div id="banner-bg-img"></div>
                    <div id="banner-bg-overlay"></div>
                    <div id="banner-content">
                        <h1>Our Student Achievements</h1>
                        <h2>
                            Congratulations to our students!
                            <br />
                            With their hard work, we all succeed!
                        </h2>
                    </div>
                </div>
                <div className="content">
                    {grouped}
                    <div className="timeline"></div>

                    <div className="last-item">
                        <div className="last-dot"></div>
                        <div className="message-container">
                            <p>
                                And many more! Math Navigator has been proud to
                                serve thousands of Montgomery County students
                                since 2004.
                            </p>
                        </div>
                    </div>
                </div>

                <div className="actions">
                    <div className=""></div>
                    <p>
                        Join our community today!
                        <br />
                        Success is just around the corner!
                    </p>
                    <Link to="/register">
                        <button>Enroll Now</button>
                    </Link>
                    <Link to="/programs">
                        <button>View Programs</button>
                    </Link>
                </div>
            </div>
        );
    }
}

class GroupedItems extends React.Component {
    render() {
        const year = this.props.year;
        const achievements = this.props.achievements || [];
        const items = achievements.map((achieve, index) => {
            return (
                <div className="item" key={index}>
                    <div className="dot"></div>
                    <p>{achieve.message}</p>
                </div>
            );
        });
        return (
            <div className="group">
                <div className="year-container">
                    <div>{year}</div>
                </div>
                <div className="items">{items}</div>
            </div>
        );
    }
}
