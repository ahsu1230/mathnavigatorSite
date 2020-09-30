"use strict";
require("./program.sass");
import React from "react";
import API from "../api.js";
import AllPageHeader from "../common/allPages/allPageHeader.js";
import RowCardBasic from "../common/rowCards/rowCardBasic.js";

const PAGE_DESCRIPTION = `
    Programs represent a course topic. If a general area of study has subjects like Math, English, or Programming, 
    programs are a subtopic of that subject (i.e. "SAT 1 Math" or "Grammar" or "AP Java Programming").
    Each program has recommended grade levels (grade1 to grade2) and are displayed in the user website "Program Catalog" page. 
    Some programs may have "Featured" labels to differentiate them. For example, some programs may be popular programs which are
    highly recommended to students. Some programs may be new and should be more visible to users.`;
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
                    description={PAGE_DESCRIPTION}
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
            label: "Subject",
            value: program.subject,
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
