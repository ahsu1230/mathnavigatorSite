'use strict';
require('./class.styl');
import React from 'react';
import ReactDOM from 'react-dom';
import { Link } from 'react-router-dom';

export class ClassPage extends React.Component {
    render() {
        var numClasses = 5;
        var fakeLocation = {
            locId: 123,
            address: "123 Sesame Street",
            roomNum: 23
        }

        return (
            <div id="view-class">
                <h1>All Classes ({numClasses})</h1>
                <ul id="list-heading">
                    <li className="li-med">program_id</li>
                    <li className="li-small">class_key</li>
                    <li className="li-small">class_id</li>
                    <li className="li-med">semester_id</li>
                    <li className="li-small">location_id</li>
                    <li className="li-small">times</li>
                    <li className="li-small">start_date</li>
                    <li className="li-small">end_date</li>
                </ul>
                <ul>
                    <ClassRow locationObj = {fakeLocation}/>
                </ul>
                    <button id="add-class">
                        <Link to={"/classes/add"}>Add Class</Link>
                    </button>
            </div>
        );
    }
}

class ClassRow extends React.Component {
    render() {
        const locId = this.props.locationObj.locId;
        const address = this.props.locationObj.address;
        const roomNum = this.props.locationObj.roomNum;
        const url = "/classes/" + "/edit";
        return (
            <ul id="class-row">
                <li className="li-med">{locId}</li>
                <li className="li-med">{address}</li>
                <li className="li-small">{roomNum}</li>
                <Link to={url}>Edit</Link>
            </ul>
        );
    }
}
