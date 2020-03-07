'use strict';
require('./achieve.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';

export class AchievePage extends React.Component {
	constructor(props) {
	    super(props);
	    this.state = {
	      list: []
	    };
 	}


	render() {
		var numAchievements = 5;
		var fakeAchieve = {
			year: 2020,
			message: "Hi"
		};
		var fakeAchieve2 = {
			year: 2019,
			message: "Hello"
		};
		return (
      	<div id="view-achieve">
	      	<h1>All Achievements ({numAchievements})</h1>
			<ul id="list-heading">
	          	<li className="li-med">Year</li>
	          	<li className="li-med"> Message</li>
	        </ul>
			<ul>
				<AchieveRow achieveObj = {fakeAchieve}/>
				<AchieveRow achieveObj = {fakeAchieve2}/>
			</ul>
				<button>
					<Link className="add-achievement" to={"/achievements/add"}>Add Achievement</Link>
				</button>
      	</div>
		);
	}
}

class AchieveRow extends React.Component {
  	render() {
  		const year = this.props.achieveObj.year;
		const message = this.props.achieveObj.message;
		const url = "/achievements/"  + "/edit";
	    return (
	    	<ul id="achieve-row">
	        	<li className="li-med">{year}</li>
	        	<li className="li-med">{message}</li>
	        	<Link to={url}>Edit</Link>
	      	</ul>
	    )
  	}
}
