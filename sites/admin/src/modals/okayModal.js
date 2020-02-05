'use strict';
require('./okayModal.styl');
import React from 'react';
import ReactDOM from 'react-dom';

export class OkayModal extends React.Component {
  render() {
    const text = this.props.text;
    const onOkay = this.props.onOkay;
    return (
      <div id="modal-view-okay">
        <p>{text}</p>
        <button onClick={onOkay}>OK</button>
      </div>
    );
  }
}
