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
				  <div>
				      <h2> Unpublished Content </h2>
					  	  <ul>
						      {unpubContent}
						  </ul>
					  <h2> Registrations </h2>
					      <ul> New Users </ul>
						  <ul> Questions </ul>
						  <ul> Complaints </ul>
				 </div>
			  	 <div>
				  	<div className="boxed">
				 	text
					</div>
			  	 </div>
			 </div>
			 <div id="view-button">
				 <button id="go-to-page">
					Go to Page
				 </button>
			 </div>
		  </div>


		);
	}
}
