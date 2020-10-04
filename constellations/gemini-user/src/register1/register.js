"use strict";
require("./register.sass");
import React from "react";
import axios from "axios";
import API from "../utils/api.js";
import { keyBy } from "lodash";
import { parseQueryParams } from "../utils/utils.js";
import RegisterSectionSelect from "./registerSelect.js";
import RegisterSectionFormStudent from "./registerFormStudent.js";
import RegisterSectionFormGuardian from "./registerFormGuardian.js";
import RegisterSectionConfirm from "./registerConfirm.js";
import RegisterSectionSuccess from "./registerSuccess.js";
import {
    REGISTER_SECTION_SELECT,
    REGISTER_SECTION_FORM_STUDENT,
    REGISTER_SECTION_FORM_GUARDIAN,
    REGISTER_SECTION_CONFIRM,
    REGISTER_SECTION_SUCCESS,
} from "./registerBase.js";

export default class RegisterPage extends React.Component {
    state = {
        location: {},
        currentSection: REGISTER_SECTION_SELECT,
        selectedAfhId: null,
        selectedClassId: null,

        studentFirstName: "",
        studentLastName: "",
        studentSchool: "",
        studentEmail: "",
        studentGraduationYear: 2024,
        studentGrade: 9,

        guardianFirstName: "",
        guardianLastName: "",
        guardianEmail: "",
        guardianPhone: "",

        allAFHs: [],
        afhMap: {},
        allClasses: [],
        classMap: {},
        programMap: {},
        semesterMap: {},
        locationMap: {},
    };

    componentDidMount = () => {
        this.setState({ location: this.props.history.location });

        const allApiCalls = [
            API.get("api/classes/all"),
            API.get("api/askforhelp/all"),
            API.get("api/programs/all"),
            API.get("api/semesters/all"),
            API.get("api/locations/all"),
        ];

        axios
            .all(allApiCalls)
            .then(
                axios.spread((...responses) => {
                    const allClasses = responses[0].data;
                    const allAFHs = responses[1].data;
                    const allPrograms = responses[2].data;
                    const allSemesters = responses[3].data;
                    const allLocations = responses[4].data;

                    let queries = parseQueryParams(
                        this.props.history.location.search
                    );
                    console.log(JSON.stringify(queries));
                    console.log(JSON.stringify(queries["afhId"]));
                    console.log(JSON.stringify(queries["classId"]));
                    this.setState({
                        selectedAfhId: queries["afhId"] || null,
                        selectedClassId: queries["classId"] || null,
                        allClasses: allClasses,
                        classMap: keyBy(allClasses, "classId"),
                        allAFHs: allAFHs,
                        afhMap: keyBy(allAFHs, "id"),
                        programMap: keyBy(allPrograms, "programId"),
                        semesterMap: keyBy(allSemesters, "semesterId"),
                        locationMap: keyBy(allLocations, "locationId"),
                    });
                })
            )
            .catch((err) =>
                console.log("Error: could not fetch class: " + err)
            );
    };

    componentWillUnmount() {}

    changeSection = (section) => {
        this.setState({ currentSection: section });
    };

    changeStateValue = (stateName, value) => {
        this.setState({
            [stateName]: value,
        });
    };

    getSectionClass = () => {
        let sectionClass = "";
        switch (this.state.currentSection) {
            case REGISTER_SECTION_FORM_STUDENT:
                sectionClass = "shift1";
                break;
            case REGISTER_SECTION_FORM_GUARDIAN:
                sectionClass = "shift2";
                break;
            case REGISTER_SECTION_CONFIRM:
                sectionClass = "shift3";
                break;
            case REGISTER_SECTION_SUCCESS:
                sectionClass = "shift4";
                break;
            default:
                sectionClass = "";
        }
        return sectionClass;
    };

    render() {
        const sectionClass = this.getSectionClass();
        const studentInfo = {
            firstName: this.state.studentFirstName,
            lastName: this.state.studentLastName,
            email: this.state.studentEmail,
            school: this.state.studentSchool,
            grade: this.state.studentGrade,
            graduationYear: this.state.studentGraduationYear,
        };
        const guardianInfo = {
            firstName: this.state.guardianFirstName,
            lastName: this.state.guardianLastName,
            email: this.state.guardianEmail,
            phone: this.state.guardianPhone,
        };

        let selectedClass = {};
        if (this.state.selectedClassId) {
            selectedClass = this.state.classMap[this.state.selectedClassId];
        }
        let selectedAfh = undefined;
        if (this.state.selectedAfhId) {
            selectedAfh = this.state.afhMap[this.state.selectedAfhId];
        }
        return (
            <div id="view-register">
                <div className="window">
                    <div className={"scrollable " + sectionClass}>
                        <RegisterSectionSelect
                            onChangeSection={this.changeSection}
                            onChangeStateValue={this.changeStateValue}
                            afhId={this.state.selectedAfhId}
                            afhs={this.state.allAFHs}
                            afhMap={this.state.afhMap}
                            classId={this.state.selectedClassId}
                            classes={this.state.allClasses}
                            classMap={this.state.classMap}
                            programMap={this.state.programMap}
                            semesterMap={this.state.semesterMap}
                            locationMap={this.state.locationMap}
                        />
                        <RegisterSectionFormStudent
                            onChangeSection={this.changeSection}
                            onChangeStateValue={this.changeStateValue}
                            student={studentInfo}
                            isAfh={!!this.state.selectedAfhId}
                        />
                        <RegisterSectionFormGuardian
                            onChangeSection={this.changeSection}
                            onChangeStateValue={this.changeStateValue}
                            guardian={guardianInfo}
                            studentEmail={studentInfo.email}
                        />
                        <RegisterSectionConfirm
                            onChangeSection={this.changeSection}
                            student={studentInfo}
                            guardian={guardianInfo}
                            selectedAfh={selectedAfh}
                            selectedClass={selectedClass}
                            programMap={this.state.programMap}
                            semesterMap={this.state.semesterMap}
                            locationMap={this.state.locationMap}
                        />
                        <RegisterSectionSuccess
                            onChangeSection={this.changeSection}
                            selectedAfh={selectedAfh}
                            selectedClass={selectedClass}
                        />
                    </div>
                </div>
                <p className="register-footer">
                    If you have any questions, please email us at
                    <br />
                    <b>andymathnavigator@gmail.com</b>.
                </p>
            </div>
        );
    }
}
