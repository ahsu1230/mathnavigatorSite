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
            list: [],
            numUnpublished: 0,
            numSelected: 0,
        };
        this.onClickSelectAll = this.onClickSelectAll.bind(this);
        this.onCheckRow = this.onCheckRow.bind(this);
    }

    componentDidMount() {
        API.get("api/locations/v1/all").then((res) => {
            const locations = res.data;
            this.setState({ list: locations });
        });
    }

    onCheckRow() {
        const elements = document.getElementsByName("unpublished") || [];
        let numChecked = 0;
        for (let i = 0; i < elements.length; i++) {
            if (elements[i].checked) {
                numChecked++;
            }
        }
        this.setState({ numSelected: numChecked });
    }

    onClickSelectAll() {
        var items = document.getElementsByName("unpublished");
        for (var i = 0; i < items.length; i++) {
            if (items[i].type == "checkbox") {
                items[i].checked = true;
            }
        }
        this.setState({ numSelected: items.length });
    }

    onClickPublish() {
        console.log("clicked publish");
    }

    render() {
        const location = this.state.list.map((location, index) => {
            return (
                <LocationRow
                    key={index}
                    location={location}
                    onCheckRow={this.onCheckRow}
                />
            );
        });
        const numLocations = location.length;
        return (
            <div id="view-location">
                <div>
                    <h1>All Locations ({numLocations})</h1>
                    <p>
                        You have {this.state.numUnpublished} unpublished items.{" "}
                        <br />
                        You have selected {this.state.numSelected} items to
                        publish.
                    </p>
                </div>
                <ul id="list-heading">
                    <button
                        className="li-small"
                        onClick={this.onClickSelectAll}>
                        Select All
                    </button>
                    <li className="li-med">Location ID</li>
                    <li className="li-large">Address</li>
                    <li className="li-large">Room</li>
                </ul>
                <ul>{location}</ul>
                <div id="list-buttons">
                    <button>
                        <Link to={"/locations/add"} id="add-location">
                            Add Location
                        </Link>
                    </button>
                    <button id="publish" onClick={this.onClickPublish}>
                        Publish
                    </button>
                </div>
            </div>
        );
    }
}

class LocationRow extends React.Component {
    renderCheckbox(isUnpublished) {
        let checkbox = <div> </div>;
        if (isUnpublished) {
            return (checkbox = (
                <input
                    className="li-small"
                    type="checkbox"
                    name="unpublished"
                    onClick={this.props.onCheckRow}
                />
            ));
        } else {
            return (checkbox = <div className="li-small"></div>);
        }
    }

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
        let checkbox = this.renderCheckbox(true);
        return (
            <ul id="location-row">
                {checkbox}
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
