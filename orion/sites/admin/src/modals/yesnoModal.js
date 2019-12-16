'use strict';
require('./yesnoModal.styl');
import React from 'react';
import ReactDOM from 'react-dom';

export class YesNoModal extends React.Component {
  render() {
    const text = this.props.text;
    const rejectText = this.props.rejectText || "No";
    const acceptText = this.props.acceptText || "Yes";
    const onReject = this.props.onReject;
    const onAccept = this.props.onAccept;
    return (
      <div id="modal-view-yesno">
        <p>{text}</p>
        <button className="reject" onClick={onReject}>{rejectText}</button>
        <button className="accept" onClick={onAccept}>{acceptText}</button>
      </div>
    );
  }
}
