'use strict';
require('./../styl/formInput.styl');
import React from 'react';
import ReactDOM from 'react-dom';
const classnames = require('classnames');

export class FormInput extends React.Component {
	constructor(props){
    super(props);
		this.state = {
			inputValue: ""
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
		const inputValue = this.state.inputValue;
		const validator = this.props.validator;

		const isEmpty = !inputValue;
		const isValid = validator ? validator.validate(inputValue) : true;
		var errorMsg = isEmpty || isValid ? "" : validator.errorMsg;

		const classNames = classnames("form-input", this.props.addClasses);
		const title = this.props.title;
		const placeholder = this.props.placeholder;
		const onErrorFn = this.props.onError;
		return (
			<div className={classNames}>
				<label>{title}</label>
				<input placeholder={placeholder}
							value={inputValue}
							onChange={this.updateInputValue}/>
				<label>{errorMsg}</label>
			</div>
		);
	}
}
