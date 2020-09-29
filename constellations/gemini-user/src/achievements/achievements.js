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
        console.log("api attempt ");
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

    renderFirstSection = () => {
        const years = this.state.years;
        const groupedByYear = this.state.groupedByYear;
        let section = <div></div>;
        if (years.length > 0) {
            section = (
                <section className="first">
                    <img src={srcSuccess}/>
                    <AchieveCard
                        year={years[0]} 
                        achievements={groupedByYear[years[0]]}/>
                </section>
            );
        }
        return section;
    }

    renderSecondSection = () => {
        const years = this.state.years;
        const groupedByYear = this.state.groupedByYear;
        let section = <div></div>;
        if (years.length > 1) {
            section = (
                <section className="second">
                    <div className="action-box">
                        <p>
                            Join our community today<br/>
                            and achieve success!
                        </p>
                        <Link to="/register">
                            <button>Enroll</button>
                        </Link>
                    </div>
                    <AchieveCard
                        year={years[1]} 
                        achievements={groupedByYear[years[1]]}/>
                </section>
            );
        }
        return section;
    }

    renderLastSection = () => {
        const years = this.state.years;
        const groupedByYear = this.state.groupedByYear;
        let section = <div></div>;
        if (this.state.years.length > 2) {
            let lastYears = years.slice(2);
            let items = lastYears.map((year, index) => {
                return (
                    <section key={index} className="last">
                        <GroupedCards
                            year={lastYears[index]} 
                            achievements={groupedByYear[lastYears[index]]}/>
                    </section>
                );
            });
            section = (<div>{items}</div>);
        }
        return section;
    }

    render() {
        const firstSection = this.renderFirstSection();
        const secondSection = this.renderSecondSection();
        const lastSection = this.renderLastSection();

        return (
            <div id="view-achieve">
                <h1>Our Student Achievements</h1>

                <div className="subheaders">
                    <h2>
                        Congratulations to our students!<br/>
                        With their hard work, we all succeed!
                    </h2>
                </div>

                {firstSection}
                {secondSection}

                <Link to="/programs" className="link-bar">
                    View our <b>Program Catalog</b> to find the right program for you.
                </Link>

                {lastSection}
            </div>
        );
    }
}

class AchieveCard extends React.Component {
    render() {
        let achievements = this.props.achievements || [];
        let year = this.props.year;
        let list = achievements.map((achievement, index) => (
            <li key={achievement.id}>
                <div className="bullet"></div>
                <p>{achievement.message}</p>
            </li>
        ));

        return (
            <div className="achieve-card">
                <h3>Achievements in {year}</h3>
                <ul>{list}</ul>
            </div>
        );
    }
}

class GroupedCards extends React.Component {
    render() {
        const year = this.props.year;
        const achievements = this.props.achievements || [];
        const items = achievements.map((achieve, index) => {
            return (
                <div className="card" key={index}>
                    <p>{achieve.message}</p>
                </div>
            );
        })
        return (
            <div className="grouped-cards">
                <h3>Achievements in {year}</h3>
                <div className="item-container">
                    {items}
                </div>
            </div>
        );
    }
}