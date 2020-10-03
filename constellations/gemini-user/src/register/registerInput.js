"use strict";
require("./registerInput.sass");
import React from "react";

export default class RegisterInput extends React.Component {
    render() {
        const description = this.props.description ? 
            (<p>{this.props.description}</p>) : 
            <div></div>;
        return (
            <div className={"register-input " + this.props.className}>
                <h3>{this.props.title}</h3>
                {description}
                <input
                    type="text"
                    value={this.props.value}
                    placeholder={this.props.placeholder}
                    onChange={(e) => this.props.onChangeCallback(e)}
                />
            </div>
        );
    }
}