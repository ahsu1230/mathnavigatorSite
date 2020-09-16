"use strict";
require("./achieve.sass");
import React from "react";
import API from "../api.js";
import AllPageHeader from "../utils/allPageHeader.js";
import RowCardBasic from "../utils/rowCardBasic.js";

export class AchievePage extends React.Component {
    state = {
        achievements: [],
    };

    componentDidMount = () => {
        API.get("api/achievements/years").then((res) =>
            this.setState({ achievements: res.data })
        );
    };

    render = () => {
        let numAchievements = 0;
        let rows = [];
        this.state.achievements.forEach((group, indexI) =>
            (group.achievements || []).forEach((achieve, indexJ) => {
                numAchievements++;
                const fields = generateFields(achieve);
                const texts = generateTexts(achieve);
                rows.push(
                    <div className="card-wrapper" key={numAchievements}>
                        <RowCardBasic
                            title={"Achievement in " + achieve.year}
                            editUrl={"/achievements/" + achieve.id + "/edit"}
                            fields={fields}
                            texts={texts}
                        />
                    </div>
                );
            })
        );

        return (
            <div id="view-achieve">
                <AllPageHeader
                    title={"All Achievements (" + numAchievements + ")"}
                    addUrl={"/achievements/add"}
                    addButtonTitle={"Add Achievement"}
                    description={
                        "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."
                    }
                />
                <div className="cards">{rows}</div>
            </div>
        );
    };
}

function generateFields(achieve) {
    return [
        {
            label: "Position",
            value: achieve.position,
        },
    ];
}

function generateTexts(achieve) {
    return [
        {
            label: "Message",
            value: achieve.message,
        },
    ];
}
