"use strict";
require("./afh.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import AllPageHeader from "../utils/allPageHeader.js";
import RowCardBasic from "../utils/rowCardBasic.js";

const PAGE_DESCRIPTION = `
    AskForHelp (AFH) sessions can be scheduled and presented in the "Ask For Help" page in the user website. 
    These AFH sessions are grouped by subject (math, english, programming) and are held over a time interval at a certain location. 
    You must create a Location before creating an AFH session. 
    A "notes" field is available per AFH session to give a brief description of what they AFH session is dedicated for 
    (i.e. "Reviewing Problem Set 4A" or "Final Practice Review" or "Essay Writing AMA").`;
export class AskForHelpPage extends React.Component {
    state = {
        list: [],
    };

    componentDidMount() {
        API.get("api/askforhelp/all").then((res) => {
            const afh = res.data;
            this.setState({
                list: afh,
            });
        });
    }

    render() {
        const cards = this.state.list.map((afh, index) => {
            const fields = generateFields(afh);
            const texts = generateTexts(afh);
            return (
                <RowCardBasic
                    key={index}
                    title={afh.title}
                    subtitle={afh.subject}
                    editUrl={"/afh/" + afh.id + "/edit"}
                    fields={fields}
                    texts={texts}
                />
            );
        });
        const numAfhs = cards.length;

        return (
            <div id="view-afh">
                <AllPageHeader
                    title={"All AskForHelp sessions (" + numAfhs + ")"}
                    addUrl={"/afh/add"}
                    addButtonTitle={"Add AskForHelp"}
                    description={PAGE_DESCRIPTION}
                />

                <div className="cards-wrapper">{cards}</div>
            </div>
        );
    }
}

function generateFields(afh) {
    const date = moment(afh.startsAt).format("dddd, MMMM Do YYYY");
    const startTime = moment(afh.startsAt).format("h:mm a");
    const endTime = moment(afh.endsAt).format("h:mm a");

    return [
        {
            label: "Date",
            value: date,
        },
        {
            label: "Times",
            value: startTime + " - " + endTime,
        },
        {
            label: "LocationId",
            value: afh.locationId,
        },
    ];
}

function generateTexts(afh) {
    return [
        {
            label: "Notes",
            value: afh.notes,
        },
    ];
}
