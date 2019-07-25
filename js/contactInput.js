'use strict';
require('./../styl/contact.styl');
import React from 'react';
import ReactDOM from 'react-dom';
const classnames = require('classnames');

export class ContactInput extends React.Component {
	constructor(props){
    super(props);
		this.state = {
			inputValue: "",
			errorMsg: ""
		};
		this.updateInputValue = this.updateInputValue.bind(this);
  }

	updateInputValue(event) {
		const newValue = event.target.value;
		const propertyName = this.props.propertyName;
		this.setState({ inputValue: newValue });
		this.props.onUpdate(propertyName, newValue);
	}

	render() {
		const classNames = classnames("contact-input", this.props.addClasses);
		const title = this.props.title;
		const placeholder = this.props.placeholder;
		const onErrorFn = this.props.onError;
		return (
			<div className={classNames}>
				<label>{title}</label>
				<input placeholder={placeholder}
							value={this.state.inputValue}
							onChange={this.updateInputValue}/>
				<label>{this.state.errorMsg}</label>
			</div>
		);
	}
}
