'use strict';
require('./../styl/announce.styl');
import React from 'react';
import ReactDOM from 'react-dom';

export class AnnouncePage extends React.Component {
	render() {
		return (
      <div id="view-announce">
        <div id="view-announce-container">
          Announcements
        </div>
      </div>
		);
	}
}
