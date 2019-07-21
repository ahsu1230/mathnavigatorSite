'use strict';
require('./../styl/modal.styl');
import React from 'react';
import ReactDOM from 'react-dom';
const classnames = require('classnames');
const srcClose = require('../assets/close_black.svg');

export class Modal extends React.Component {
  constructor(props) {
    super(props);
    this.handleDismiss = this.handleDismiss.bind(this);
  }

  handleDismiss() {
    if (this.props.onDismiss) {
      this.props.onDismiss();
    }
  }

  render() {
    const modalContent = this.props.content;
    const modalViewClasses = classnames("modal-view", {
      show: this.props.show
    });
    var modalClasses = classnames("modal", this.props.modalClassName);
    return (
      <div className={modalViewClasses}>
        <div className="modal-overlay" onClick={this.handleDismiss}></div>
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
