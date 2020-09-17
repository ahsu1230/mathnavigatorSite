"use strict";
require("./achieve.sass");
import React from "react";
import API from "../api.js";
import AllPageHeader from "../utils/allPageHeader.js";
import RowCardBasic from "../utils/rowCardBasic.js";

const PAGE_DESCRIPTION = `
    Achievements are student accomplishments that have occured throughout the years of Math Navigator's programs. 
    They are shown in the "Student Achievements" section of the user website and are grouped by year (starting from the most recent). 
    Achievements are sorted by year (descending) and then by position (ascending). 
    Use the 'Position' field value to determine which achievement comes first within each year group.`;

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
                    description={PAGE_DESCRIPTION}
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
