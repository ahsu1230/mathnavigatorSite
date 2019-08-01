'use strict';
require('./home.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';
import { getAllProfiles } from '../repos/mainRepo.js';

export class HomeSectionStories extends React.Component {
	render() {
    var profiles = getAllProfiles();
    // For now, only grab one
    var profile = profiles[0];

		return (
			<div className="section stories">
				<h2>Stories</h2>
        <StoryProfile profile={profile}/>
        <Link to='programs'>
          <button className="action-join">Join our community</button>
        </Link>
			</div>
		);
	}
}

class StoryProfile extends React.Component {
  render() {
    var profile = this.props.profile

    return (
      <div className="profile-container">
        <div className="profile-img-container">
          <img src={profile.imgSrc}/>
        </div>
        <div className="profile-text-container">
          <h2 className="title">{profile.name}</h2>
          <h3>{profile.subtitle1}</h3>
          <h3>{profile.subtitle2}</h3>
          <p>"{profile.quote}"</p>
        </div>
      </div>
    );
  }
}
