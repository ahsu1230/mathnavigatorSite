"use strict";
require("./program.sass");
import React from "react";
import { Link } from "react-router-dom";
import API from "../api.js";
import { ProgramRow } from "./programRow.js";

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
        const programs = this.state.programs.map((program, index) => (
            <ProgramRow key={index} program={program} />
        ));

        return (
            <div id="view-program">
                <h1>All Programs ({programs.length}) </h1>

                <div className="header row">
                    <span className="medium-column">ProgramId</span>
                    <span className="medium-column">Name</span>
                    <span className="small-column">Grade1</span>
                    <span className="small-column">Grade2</span>
                    <span className="large-column">Description</span>
                    <span className="edit"></span>
                </div>
                {programs}

                <button>
                    <Link id="add-program" to={"/programs/add"}>
                        Add Program
                    </Link>
                </button>
            </div>
        );
    };
}
