"use strict";
require("./home.sass");
import React from "react";
import ReactDOM from "react-dom";
import { successText } from "./homeText.js";

export class HomeSectionSuccess extends React.Component {
    render() {
        return (
            <div className="section success">
                <h2>Navigate to Success</h2>
                <div className="card-container">
                    <Card card={successText[0]} />
                    <Card card={successText[1]} />
                    <Card card={successText[2]} />
                </div>
            </div>
        );
    }
}

class Card extends React.Component {
    render() {
        var card = this.props.card;
        return (
            <div className="home-tile-card success">
                <div className="img-container">
                    <img src={card.imgSrc} />
                </div>
                <h3>{card.title}</h3>
                <p>{card.description}</p>
            </div>
        );
    }
}
