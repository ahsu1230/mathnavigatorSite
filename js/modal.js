'use strict';
require('./../styl/modal.styl');
import React from 'react';
import ReactDOM from 'react-dom';
const classnames = require('classnames');
const srcClose = require('../assets/close_black.svg');

export class Modal extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      show: true
    }
    this.handleDismiss = this.handleDismiss.bind(this);
  }

  handleDismiss() {
    console.log("Dismissing modal...");
    this.setState({
      show: false
    });
    if (this.props.onDismiss) {
      this.props.onDismiss();
    }
  }

  render() {
    const modalContent = this.props.content;
    var modalClasses = classnames("modal", this.props.modalClassName);
    return (
      <div id="modal-view">
        <div id="modal-overlay" onClick={this.handleDismiss}></div>
        <div className={modalClasses}>
          <button className="close-x" onClick={this.handleDismiss}>
            <img src={srcClose}/>
          </button>
          {modalContent}
        </div>
      </div>
    );
  }
}
