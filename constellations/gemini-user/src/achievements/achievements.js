"use strict";
require("./achievements.sass");
import React from "react";
import API from "../utils/api.js";

export class YearList extends React.Component {
    render() {
		let a = this.props.achievements
		let year = this.props.year
		let achievements = a.map((achievement, index) => (
			<li key={achievement.id}>
				<div className="bullet"></div>
				<p>{achievement.message}</p>
			</li>
		))
		
        return (
			<div className="achieve-card">
				<h2>{year}</h2>
				
				<ul>
					{achievements}
				</ul>
				
				<a href="#/programs">View our Programs &gt;</a>
			</div>
        );
    }
}

export class AchievementPage extends React.Component {
    state = {
        achieveList: [],
    };

    componentDidMount() {
        console.log("api attempt ");
        API.get("api/achievements/all").then((res) => {
            const achieveList = res.data;
            console.log("api success!");
            this.setState({ achieveList });
        });
    }

	// {"id":1,"publishedAt":null,"year":2020,"message":"message"}
    render() {
		const achievements = this.state.achieveList
		
		let sorted = {}
		achievements.forEach( achievement => {			
			if (!(achievement.year in sorted)) {
				sorted[achievement.year] = []
			}
			sorted[achievement.year].push(achievement)
		} )
		
		let years = Object.keys(sorted)
		years.sort().reverse()
		
		let items = []
		years.forEach(year => {
			items.push((
				<YearList key={year} year={year} achievements={sorted[year]} />
			))
		})
		
        return (
            <div id="view-achieve">
				<div id="view-achievements-container">
					<h1>Math Navigator Achievements</h1>
					
					<div className="subheaders">
						<h3>Congratulations to our students!</h3>
						<h3>With their hard work, we all succeed!</h3>
					</div>
					
					{items}
				</div>
            </div>
        );
    }
}
