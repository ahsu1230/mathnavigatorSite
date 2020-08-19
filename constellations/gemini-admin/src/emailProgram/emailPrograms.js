"use strict";
require("./emailPrograms.sass");
import axios from "axios";
import React from "react";
import API, { executeApiCalls } from "../api.js";

export class EmailPrograms extends React.Component {
    state = {
        selectProgramId: "",
        programs: [],
        classes: [],
        classesForProgram: [],
    };

    componentDidMount = () => {
        const apiCalls = [
            API.get("api/programs/all"),
            API.get("api/classes/all"),
        ];
        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) => {
                    const programs = responses[0].data;
                    const classes = responses[1].data;
                    const hasClassId = responses.length > 3;
                    let classObj = hasClassId ? responses[3].data : {};
                    let selectedProgramId = hasClassId
                        ? classObj.programId
                        : programs[0].programId;

                    this.setState({
                        selectProgramId: selectedProgramId,
                        programs: programs,
                        classes: classes,
                    });
                })
            )
            .catch((err) => {
                console.log("Error: api call failed. " + err.message);
            });
    };

    handleProgramChange = (event, value) => {
        const length = event.target.value.length;
        const classes = this.state.classes.map((classes) => classes.classId)
        var classesForProgram = [];
        for (var i = 0; i < classes.length; i++) {
            if  (classes[i].substring(0,length) == event.target.value) {
                classesForProgram.push(classes[i]);
            }
        }
        this.setState({ 
            [value]: event.target.value,
            classesForProgram: classesForProgram,
        });
    };

    render = () => {
        const programOptions = this.state.programs.map((program, index) => (
            <option key={index}>{program.programId}</option>
        ));
        const classOptions = this.state.classesForProgram.map((classes,index) => (
            <div>
                <input
                        type="checkbox"
                        onChange={(e) => this.onCheckUser(e, user.id)}
                />
                <span>{classes}</span>
            </div>
        ));
        return (
            <div id="view-program-emails">
                <section id="Select-program">
                    <h1>Generate Email to Program</h1>
                    <h4>Select a Program</h4>
                    <select
                        value={this.state.selectProgramId}
                        onChange={(e) =>
                            this.handleProgramChange(e, "selectProgramId")
                        }>
                        {programOptions}
                    </select>
                    <h4>Select a class for {this.state.selectProgramId}</h4>
                    {classOptions}
                </section>
            </div>
        );
    };
}
