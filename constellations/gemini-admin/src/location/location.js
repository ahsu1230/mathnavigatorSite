"use strict";
require("./location.sass");
import React from "react";
import API from "../api.js";
import AllPageHeader from "../utils/allPageHeader.js";
import RowCardBasic from "../utils/rowCardBasic.js";

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
        const cards = this.state.list.map((location, index) => {
            const fields = generateFields(location);
            return (
                <RowCardBasic
                    key={index}
                    title={location.title}
                    subtitle={location.locationId}
                    editUrl={"/locations/" + location.locationId + "/edit"}
                    fields={fields}
                />
            );
        });

        const numLocations = cards.length;
        return (
            <div id="view-location">
                <AllPageHeader
                    title={"All Locations (" + numLocations + ")"}
                    addUrl={"/locations/add"}
                    addButtonTitle={"Add Location"}
                    description={
                        "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."
                    }
                />

                <div className="cards-wrapper">{cards}</div>
            </div>
        );
    }
}

function generateFields(location) {
    return [
        {
            label: "IsOnline",
            value: "" + (location.isOnline || false),
            highlightFn: () => location.isOnline,
        },
        {
            label: "Address1",
            value: location.street,
        },
        {
            label: "Address2",
            value:
                location.city + ", " + location.state + " " + location.zipcode,
        },
        {
            label: "Address3",
            value: location.room,
        },
    ];
}
