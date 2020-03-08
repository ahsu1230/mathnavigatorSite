"use strict";
require("./location.styl");
import React from "react";
import ReactDOM from "react-dom";
import { Link } from "react-router-dom";

export class LocationPage extends React.Component {
    render() {
        var numLocations = 5;
        var fakeLocation = {
            locId: 123,
            address: "123 Sesame Street",
            roomNum: 23
        };
        var fakeLocation1 = {
            locId: 4567,
            address: "456 Sesame St",
            roomNum: 892
        };
        var fakeLocation2 = {
            locId: 123456,
            address: "12345 Park Place",
            roomNum: 212
        };
        return (
            <div id="view-location">
                <h1>All Locations ({numLocations})</h1>
                <ul id="list-heading">
                    <li className="li-med">Location ID</li>
                    <li className="li-med">Address</li>
                    <li className="li-large">Room Number</li>
                </ul>
                <ul>
                    <LocationRow locationObj={fakeLocation} />
                    <LocationRow locationObj={fakeLocation1} />
                    <LocationRow locationObj={fakeLocation2} />
                </ul>
                <button id="add-location">
                    <Link to={"/locations/add"}>Add Location</Link>
                </button>
            </div>
        );
    }
}

class LocationRow extends React.Component {
    render() {
        const locId = this.props.locationObj.locId;
        const address = this.props.locationObj.address;
        const roomNum = this.props.locationObj.roomNum;
        const url = "/location/" + "/edit";
        return (
            <ul id="location-row">
                <li className="li-med">{locId}</li>
                <li className="li-med">{address}</li>
                <li className="li-small">{roomNum}</li>
                <Link to={url}>Edit</Link>
            </ul>
        );
    }
}
