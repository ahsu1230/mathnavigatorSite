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
    const persistent = this.props.persistent || false;
    const withClose = this.props.withClose || false;
    const modalViewClasses = classnames("modal-view", {
      show: this.props.show
    });

    const onClickOverlay = persistent ? undefined : this.handleDismiss;
    var closeButton;
    if (withClose) {
      closeButton = (
        <button className="close-x" onClick={this.handleDismiss}>
          <img src={srcClose}/>
        </button>
      );
    } else {
      <div></div>
    }

    var modalClasses = classnames("modal", this.props.modalClassName);
    return (
      <div className={modalViewClasses}>
        <div className="modal-overlay" onClick={onClickOverlay}></div>
        <div className={modalClasses}>
          {closeButton}
          {modalContent}
        </div>
      </div>
    );
  }
}
