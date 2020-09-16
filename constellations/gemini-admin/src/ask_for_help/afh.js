"use strict";
require("./afh.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import AllPageHeader from "../utils/allPageHeader.js";
import RowCardBasic from "../utils/rowCardBasic.js";

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
                    editUrl={"/afhs/" + afh.afhId + "/edit"}
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
                    description={
                        "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."
                    }
                />

                <div className="cards-wrapper">{cards}</div>
            </div>
        );
    }
}

function generateFields(afh) {
    const date = moment(afh.startsAt).format("dddd, MMMM Do YYYY");
    const startTime = moment(afh.startsAt).format("h:mm:ss a");
    const endTime = moment(afh.endsAt).format("h:mm:ss a");

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
