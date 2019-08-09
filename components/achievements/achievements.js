'use strict';
require('./achievements.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import {
  getAchievementKeys,
  getAchievements
} from '../repos/mainRepo.js';
const classnames = require('classnames');

export class AchievementPage extends React.Component {
  constructor(props) {
    super(props);

    // Guarantee from most recent year to latest
    this.achievementYears = getAchievementKeys().sort().reverse();
  }

	render() {
    const cards = this.achievementYears.map(function(year, index) {
      const achievements = getAchievements(year);
      return (
        <AchievementCard key={index} year={year} achievements={achievements}/>
      );
    });

		return (
      <div id="view-achievements">
        <div id="view-achievements-container">
          <h1>Our Student Achievements</h1>
          <div className="subheaders">
            <h3>Congratulations to our students!</h3>
            <h3>With their hard work, we all succeed!</h3>
          </div>
          <div className="achievement-cards">
            {cards}
          </div>
        </div>
      </div>
		);
	}
}

class AchievementCard extends React.Component {
  render() {
    const year = this.props.year;
    const achievements = this.props.achievements;
    const lines = achievements.map((a, index) =>
      <li key={index}>
        <AchievementLine achievement={a}/>
      </li>
    );
    return (
      <div className="achieve-card">
        <h2>{year}</h2>
        <ul>
          {lines}
        </ul>
        <Link to="/programs">View our Programs &#62;</Link>
      </div>
    );
  }
}

class AchievementLine extends React.Component {
  render() {
    const achievement = this.props.achievement;
    const lineClasses = classnames({
      highlight: achievement.highlight
    });
    return (
      <div className={lineClasses}>
        <div className="bullet"></div>
        <p>{achievement.message}</p>
      </div>
    );
  }
}
