"use strict";
require("./class.sass");
import React from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import { isEmpty } from "lodash";
import API from "../utils/api.js";
import { getFullTitle } from "../utils/classUtils.js";

import { ClassBreadcrumbs } from "./classBreadcrumbs.js";
import { ClassInfo } from "./classInfo.js";
import { ClassProgramInfo } from "./classProgramInfo.js";
import { ClassRegister } from "./classRegister.js";
import { ClassSchedule } from "./classSchedule.js";
import { ClassErrorPage } from "./classError.js";

export class ClassPage extends React.Component {
    state = {
        classObj: {},
        sessions: [],
        program: {},
        semester: {},
        location: {},
        otherClasses: [],
        allSemesters: [],
        fetchedData: false,
    };

    componentDidMount = () => {
        this.fetchData(this.props.classId);
    };

    fetchData = (classId) => {
        const apiCallsPerClass = [
            API.get("api/classes/class/" + classId),
            API.get("api/sessions/class/" + classId),
        ];

        axios
            .all(apiCallsPerClass)
            .then(
                axios.spread((...responses) => this.fetchOtherData(responses))
            )
            .catch((err) =>
                console.log("Error: could not fetch class: " + err)
            );
    };

    fetchOtherData = (responses) => {
        const classObj = responses[0].data;
        const sessions = responses[1].data;
        const apiCalls = [
            API.get("api/programs/program/" + classObj.programId),
            API.get("api/semesters/all"),
            API.get("api/locations/location/" + classObj.locationId),
            API.get("api/classes/program/" + classObj.programId),
        ];
        axios
            .all(apiCalls)
            .then(
                axios.spread((...responses) => {
                    const program = responses[0].data;
                    const allSemesters = responses[1].data;
                    const semester = allSemesters.find(
                        (semester) => semester.semesterId == classObj.semesterId
                    );
                    this.setState({
                        classObj: classObj || {},
                        sessions: sessions || [],
                        program: program,
                        allSemesters: allSemesters,
                        semester: semester,
                        location: responses[2].data,
                        otherClasses: responses[3].data,
                        fetchedData: true,
                    });
                    document.title = getFullTitle(program, classObj, semester);
                })
            )
            .catch((err) =>
                console.log("Error: could not fetch other data: " + err)
            );
    };

    render = () => {
        const classObj = this.state.classObj;
        const sessions = this.state.sessions;

        if (this.state.fetchedData && isEmpty(classObj))
            return <ClassErrorPage classId={this.props.classId} />;
        else
            return (
                <div id="view-class">
                    <ClassBreadcrumbs
                        program={this.state.program}
                        classObj={classObj}
                        semester={this.state.semester}
                    />
                    <article className="intro">
                        <ClassProgramInfo program={this.state.program} />
                        <ClassRegister classObj={classObj} />
                    </article>
                    <ClassInfo
                        classObj={classObj}
                        location={this.state.location}
                        sessions={this.state.sessions}
                    />
                    <ClassSchedule sessions={sessions} />

                    <div id="class-footer">
                        <Link to="/programs">{"< More Programs"}</Link>
                        <p>
                            Still have questions about our programs and classes?
                            <br />
                            Email us at <u>andymathnavigator@gmail.com</u>.
                        </p>
                    </div>
                </div>
            );
    };
}
