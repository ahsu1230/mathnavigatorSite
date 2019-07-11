'use strict';
require('./../styl/class.styl');
import React from 'react';
import ReactDOM from 'react-dom';

export class ClassPage extends React.Component {
	render() {
    const slug = this.props.slug;
		return (
      <div id="view-class">
        <div id="view-class-container">
          Class {slug}
        </div>
      </div>
		);
	}
}
