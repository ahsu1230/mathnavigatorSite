'use strict';
require('./home.styl');
import React from 'react';
import ReactDOM from 'react-dom';

export class HomePage extends React.Component {
	render() {
		const unpubContent = 5
		return (
	      <div id="view-home">
		  	  <div id="view-content">
			      <h2> Unpublished Content </h2>
				  	  <ul>
					      <li>{unpubContent}</li>
					  </ul>
				  <h2> Registrations </h2>
				      <ul>
					  	  <li> New Users </li>
						  <li> Questions </li>
						  <li> Complaints </li>
					  </ul>
			  </div>
			  <div id="box-and-button">
				  <div className="boxed">
					text
				 </div>
				 <button id="go-to-page">
					Go to Page
				 </button>
			 </div>
		  </div>
		);
	}
}
