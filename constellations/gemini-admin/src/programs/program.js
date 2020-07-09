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
        const rows = this.state.programs.map((row, index) => {
            return (
                <li key={index}>
                    <ProgramRow row={row} />
                </li>
            );
        });

        const count = rows.length;

        return (
            <div id="view-program">
                <h1>All Programs ({count}) </h1>

                <div id="header">
                    <span className="medium">ProgramId</span>
                    <span className="medium">Name</span>
                    <span className="small">Grade1</span>
                    <span className="small">Grade2</span>
                </div>
                <ul id="rows">{rows}</ul>

                <button>
                    <Link id="add-program" to={"/programs/add"}>
                        Add Program
                    </Link>
                </button>
            </div>
        );
    };
}
