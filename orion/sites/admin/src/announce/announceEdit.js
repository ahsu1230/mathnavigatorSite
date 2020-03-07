'use strict';
require('./announceEdit.styl');
import React from 'react';
import ReactDOM from 'react-dom';

export class AnnounceEditPage extends React.Component {
  render() {

/*I have no idea how this looks because this page has "Cannot GET /announceEdit.js" error. fix later*/
    return (
      <div id ="view-announce-edit">
        <h2>Add Announcement</h2>

        <h4>Date</h4>
        <input value={this.state.inputDate}
                onChange={(e) => this.handleChange(e, "inputDate")}/>

        <h4>Author</h4>
        <input value={this.state.inputAuthor}
                onChange={(e) => this.handleChange(e, "inputAuthor")}/>

        <h4>Message</h4>
        <input value={this.state.inputMessage}
                onChange={(e) => this.handleChange(e, "inputMessage")}/>

        <div className="buttons">
          <button className="btn-save" onClick={this.onClickSave}>Save</button>
          <button className="btn-cancel" onClick={this.onClickCancel}>Cancel</button>
          {deleteButton}
        </div>
      </div>
    );
  }
}
