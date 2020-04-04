"use strict";
require("./location.styl");
import React from "react";
import ReactDOM from "react-dom";
import API from "../api.js";
import { Link } from "react-router-dom";

export class LocationPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            list: []
        };
    }

    componentDidMount() {
        API.get("api/locations/v1/all").then(res => {
            const locations = res.data;
            this.setState({ list: locations });
        });
    }

    render() {
        const location = this.state.list.map((location, index) => {
            return <LocationRow key={index} location={location} />;
        });
        const numLocations = location.length;
        return (
            <div id="view-location">
                <h1>All Locations ({numLocations})</h1>
                <ul id="list-heading">
                    <li className="li-med">Location ID</li>
                    <li className="li-large">Address</li>
                    <li className="li-large">Room</li>
                </ul>
                <ul>{location}</ul>
                <button id="add-location">
                    <Link to={"/locations/add"}>Add Location</Link>
                </button>
            </div>
        );
    }
}

class LocationRow extends React.Component {
    render() {
        const locId = this.props.location.locId;
        const address1 = this.props.location.street;
        const address2 =
            this.props.location.city +
            ", " +
            this.props.location.state +
            " " +
            this.props.location.zipcode;
        const room = this.props.location.room;
        const url = "/locations/" + locId + "/edit";
        return (
            <ul id="location-row">
                <li className="li-med">{locId}</li>
                <li className="li-large">
                    <div> {address1} </div>
                    <div> {address2} </div>
                </li>
                <li className="li-small">{room}</li>
                <Link to={url}>Edit</Link>
            </ul>
        );
    }
}
