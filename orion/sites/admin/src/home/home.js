'use strict';
require('./home.styl');
import React from 'react';
import ReactDOM from 'react-dom';

export class HomePage extends React.Component {
	render() {
		const unpubContent = 5
		return (
	      <div id="view-home">
		      <h2> Unpublished Content </h2>
			  	  <ul>
				      {unpubContent}
				  </ul>
			  <h2> Registrations </h2>
			      <ul> New Users </ul>
				  <ul> Questions </ul>
				  <ul> Complaints </ul>
			  <button id="go-to-page">
			  	  Add Location
			  </button>
	      </div>
		);
	}
}
