"use strict";
require("./homeStrategy.sass");
import React from "react";
import { strategyText } from "./homeText.js";

export default class HomeSectionStrategy extends React.Component {
    render() {
        return (
            <div className="section strategy">
                <h2>Navigate to Success</h2>
                <div className="card-container">
                    <Card card={strategyText[0]} />
                    <Card card={strategyText[1]} />
                    <Card card={strategyText[2]} />
                </div>
            </div>
        );
    }
}

class Card extends React.Component {
    render() {
        var card = this.props.card;
        return (
            <div className="card">
                <div className="img-container">
                    <img src={card.imgSrc} />
                </div>
                <h3>{card.title}</h3>
                <p>{card.description}</p>
            </div>
        );
    }
}
