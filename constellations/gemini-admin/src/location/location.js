"use strict";
require("./location.sass");
import React from "react";
import API from "../api.js";
import AllPageHeader from "../utils/allPageHeader.js";
import RowCardBasic from "../utils/rowCardBasic.js";

const PAGE_DESCRIPTION = `
    Locations are typically physical addresses that are often used to host Math Navigator class sessions.
    Address1 consists of the Street # and Street name.
    Address2 consists of the City, State and Zipcode.
    Address3 is usually for a room number (i.e. Room #110, Basement, Gymnasium).
    All locations should have a title to display to the user.
    A location can also be online (i.e. Zoom meeting or Google Meets video conference). If this is the case,
    the location will not have Addresses 1 & 2.
    The full location is typically displayed in the user website's Class page.
`;
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
                    description={PAGE_DESCRIPTION}
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
