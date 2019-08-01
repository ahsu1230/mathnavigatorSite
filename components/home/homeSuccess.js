'use strict';
require('./home.styl');
import React from 'react';
import ReactDOM from 'react-dom';

const srcUserCheck = require('../../assets/user_check_white.svg');
const srcLightBulb = require('../../assets/lightbulb_white.svg');
const srcLibrary = require('../../assets/library_white.svg');

const cards = [
  {
    title: "Dedicated Professionalism",
    description: "Professional and responsible. Andy cares about each of his students and will not fail to point out their strengths and weaknesses.",
    imgSrc: srcUserCheck
  },
  {
    title: "Strategical Thinking",
    description: "Professional and responsible. Andy cares about each of his students and will not fail to point out their strengths and weaknesses.",
    imgSrc: srcLightBulb
  },
  {
    title: "Abundant Resources",
    description: "Professional and responsible. Andy cares about each of his students and will not fail to point out their strengths and weaknesses.",
    imgSrc: srcLibrary
  }
];

export class HomeSectionSuccess extends React.Component {
	render() {
		return (
			<div className="section success">
				<h2>Navigate to Success</h2>
        <div className="card-container">
          <Card card={cards[0]}/>
          <Card card={cards[1]}/>
          <Card card={cards[2]}/>
        </div>
			</div>
		)
	}
}

class Card extends React.Component {
  render() {
    var card = this.props.card;
    return (
      <div className="home-tile-card success">
        <div className="img-container">
          <img src={card.imgSrc}/>
        </div>
        <h3>{card.title}</h3>
        <p>{card.description}</p>
      </div>
    );
  }
}
