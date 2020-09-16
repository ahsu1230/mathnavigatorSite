"use strict";
require("./program.sass");
import React from "react";
import API from "../api.js";
import AllPageHeader from "../utils/allPageHeader.js";
import RowCardBasic from "../utils/rowCardBasic.js";

export class ProgramPage extends React.Component {
    state = {
        programs: [],
    };

    componentDidMount = () => {
        this.fetchData();
    };

    fetchData = () => {
        API.get("api/programs/all").then((res) => {
            const programs = res.data;
            this.setState({
                programs: programs,
            });
        });
    };

    render = () => {
        const cards = this.state.programs.map((program, index) => {
            const fields = generateFields(program);
            const texts = generateTexts(program);
            return (
                <RowCardBasic
                    key={index}
                    title={program.title}
                    subtitle={program.programId}
                    editUrl={"/programs/" + program.programId + "/edit"}
                    fields={fields}
                    texts={texts}
                />
            );
        });
        const numPrograms = this.state.programs.length;

        return (
            <div id="view-program">
                <AllPageHeader
                    title={"All Programs (" + numPrograms + ")"}
                    addUrl={"/programs/add"}
                    addButtonTitle={"Add Program"}
                    description={
                        "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."
                    }
                />

                <div className="cards-wrapper">{cards}</div>
            </div>
        );
    };
}

function generateFields(program) {
    return [
        {
            label: "Grade1",
            value: program.grade1,
        },
        {
            label: "Grade2",
            value: program.grade2,
        },
        {
            label: "Featured",
            value: program.featured,
            highlightFn: () => program.featured != "none",
        },
    ];
}

function generateTexts(program) {
    return [
        {
            label: "Description",
            value: program.description,
        },
    ];
}
