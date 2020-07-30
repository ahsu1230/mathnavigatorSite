"use strict";
require("./location.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";

export class LocationPage extends React.Component {
    state = {
        list: [],
    };
    componentDidMount() {
        API.get("api/locations/all").then((res) => {
            const locations = res.data;
            this.setState({ list: locations });
        });
    }
    render() {
        const locations = this.state.list.map((location, index) => {
            return <LocationRow key={index} location={location} />;
        });
        const numLocations = locations.length;
        return (
            <div id="view-location">
                <div>
                    <h1>All Locations ({numLocations})</h1>
                </div>
                <ul id="list-heading">
                    <li className="li-med">Location ID</li>
                    <li className="li-large">Address</li>
                    <li className="li-large">Room</li>
                </ul>
                <ul>{locations}</ul>
                <div id="list-buttons">
                    <button>
                        <Link to={"/locations/add"} id="add-location">
                            Add Location
                        </Link>
                    </button>
                </div>
            </div>
        );
    }
}

class LocationRow extends React.Component {
    render() {
        const locationId = this.props.location.locationId;
        const address1 = this.props.location.street;
        const address2 =
            this.props.location.city +
            ", " +
            this.props.location.state +
            " " +
            this.props.location.zipcode;
        const room = this.props.location.room;
        const url = "/locations/" + locationId + "/edit";
        return (
            <ul id="location-row">
                <li className="li-med">{locationId}</li>
                <li className="li-large">
                    <div> {address1} </div>
                    <div> {address2} </div>
                </li>
                <li className="li-large">{room}</li>
                <Link to={url} className="editButton">
                    Edit
                </Link>
            </ul>
        );
    }
}
