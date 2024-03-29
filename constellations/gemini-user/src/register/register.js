"use strict";
require("./register.sass");
import React from "react";
import axios from "axios";
import API from "../utils/api.js";
import { keyBy } from "lodash";
import { trackAnalytics } from "../utils/analyticsUtils.js";
import { parseQueryParams } from "../utils/urlUtils.js";
import {
    generateEmailMessageForClass,
    generateEmailMessageForAfh,
    sendEmail,
} from "./email.js";
import Loader from "./loader.js";
import RegisterSticky from "./registerSticky.js";
import RegisterSelectClass from "./registerClass.js";
import RegisterSelectAfh from "./registerAfh.js";
import { RegisterFormStudent, RegisterFormGuardian } from "./registerForms.js";
import {
    validateEmail,
    validatePhone,
    validateGrade,
    validateGradYear,
} from "../utils/validators.js";
import { isFullClass } from "../utils/classUtils.js";
import { changePage } from "../utils/historyUtils.js";

const CHOICE_NONE = "";
const CHOICE_CLASS = "class";
const CHOICE_AFH = "afh";

export default class RegisterPage extends React.Component {
    state = {
        location: {},
        choice: CHOICE_NONE,
        selectedAfhId: null,
        selectedClassId: null,
        selectedFromQuery: false,

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
        guardianAdditionalInfo: "",

        allAFHs: [],
        afhMap: {},
        allClasses: [],
        classMap: {},
        programMap: {},
        semesterMap: {},
        locationMap: {},
    };

    componentDidMount = () => {
        let queries = parseQueryParams(this.props.history.location.search);

        let selectedFromQuery = false;
        let selectedAfhId;
        let selectedClassId;
        let choice = CHOICE_NONE;
        if (queries["afhId"]) {
            selectedAfhId = queries["afhId"] || null;
            choice = CHOICE_AFH;
            selectedFromQuery = true;
        }
        if (queries["classId"]) {
            selectedClassId = queries["classId"] || null;
            choice = CHOICE_CLASS;
            selectedFromQuery = true;
        }

        trackAnalytics("register", { choice: choice });
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

                    this.setState({
                        choice: choice,
                        selectedFromQuery: selectedFromQuery,
                        selectedAfhId: selectedAfhId,
                        selectedClassId: selectedClassId,
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
                console.log("Error: could not fetch data from server: " + err)
            );
    };

    componentWillUnmount() {}

    changeStateValue = (stateName, value) => {
        this.setState({
            [stateName]: value,
        });
    };

    changeChoice = (newChoice) => {
        trackAnalytics("register", { choice: newChoice });
        this.setState({
            choice: newChoice,
        });
    };

    invokeEmail = () => {
        trackAnalytics("register-submit", { choice: this.state.choice });
        const studentInfo = {
            firstName: this.state.studentFirstName,
            lastName: this.state.studentLastName,
            email: this.state.studentEmail,
            isGuardian: false,
            school: this.state.studentSchool,
            grade: this.state.studentGrade,
            graduationYear: this.state.studentGraduationYear,
        };
        const guardianInfo = {
            firstName: this.state.guardianFirstName,
            lastName: this.state.guardianLastName,
            email: this.state.guardianEmail,
            isGuardian: true,
            phone: this.state.guardianPhone,
            notes: this.state.guardianAdditionalInfo,
        };

        let registerApiUrl = "";
        let emailMessage = "";
        if (this.state.choice == CHOICE_CLASS) {
            const classId = this.state.selectedClassId;
            emailMessage = generateEmailMessageForClass(
                classId,
                studentInfo,
                guardianInfo
            );
            registerApiUrl = "/api/register/class/" + classId;
        } else if (this.state.choice == CHOICE_AFH) {
            const afhId = this.state.selectedAfhId;
            emailMessage = generateEmailMessageForAfh(
                afhId,
                this.state.afhMap[afhId],
                studentInfo
            );
            registerApiUrl = "/api/register/afh/" + afhId;
        } else {
            return;
        }
        console.log("Sending email... " + emailMessage);
        sendEmail(
            emailMessage,
            () => {
                console.log("EMAIL SUCCESS!");
            },
            () => {
                console.log("EMAIL FAILED!");
            }
        );

        console.log("Invoke register API...");
        const registerBody = {
            student: studentInfo,
            guardian: guardianInfo,
        };
        API.post(registerApiUrl, registerBody)
            .then((res) => console.log("Register succeeded!"))
            .catch((err) => console.log("Register failed! " + err));
    };

    isClassValid = () => {
        const classId = this.state.selectedClassId;
        return classId && !isFullClass(this.state.classMap[classId]);
    };

    isAfhValid = () => {
        return !!this.state.selectedAfhId;
    };

    isStudentInfoValid = () => {
        return (
            this.state.studentFirstName &&
            this.state.studentLastName &&
            validateEmail(this.state.studentEmail) &&
            this.state.studentSchool &&
            validateGradYear(this.state.studentGraduationYear) &&
            validateGrade(this.state.studentGrade)
        );
    };

    isGuardianInfoValid = () => {
        return (
            this.state.guardianFirstName &&
            this.state.guardianLastName &&
            !(
                this.state.guardianFirstName == this.state.studentFirstName &&
                this.state.guardianLastName == this.state.studentLastName
            ) &&
            validateEmail(this.state.guardianEmail) &&
            this.state.studentEmail != this.state.guardianEmail &&
            validatePhone(this.state.guardianPhone)
        );
    };

    render() {
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
            additionalInfo: this.state.guardianAdditionalInfo,
        };
        return (
            <div id="view-register">
                <h1>Registration</h1>
                <div className="intro-wrapper">
                    <div>
                        {this.state.selectedFromQuery &&
                            this.state.choice == CHOICE_AFH && (
                                <p>
                                    You have selected to register for an
                                    Ask-for-Help session.
                                </p>
                            )}
                        {this.state.selectedFromQuery &&
                            this.state.choice == CHOICE_CLASS && (
                                <p>
                                    You have selected to enroll into a Math
                                    Navigator class.
                                </p>
                            )}
                        {!this.state.selectedFromQuery && (
                            <p>
                                Are you enrolling into a class or attending an
                                ask-for-help session? Ask-for-Help sessions are
                                only for students who are already enrolled in
                                one of our classes.
                            </p>
                        )}
                        <p className="instruction">
                            Please fill out the following information about
                            yourself. We use this information to contact you
                            about important class updates, so please use a valid
                            email you frequently use. This information is used
                            for <i>contacting purposes only</i> and will NOT be
                            shared with anyone.
                        </p>
                    </div>
                    {!this.state.selectedFromQuery && (
                        <div className="choose-btns">
                            <button
                                className={
                                    "class-btn" +
                                    (this.state.choice == CHOICE_CLASS
                                        ? " active"
                                        : "")
                                }
                                onClick={() => this.changeChoice(CHOICE_CLASS)}>
                                Enroll into a class
                            </button>
                            <button
                                className={
                                    "afh-btn" +
                                    (this.state.choice == CHOICE_AFH
                                        ? " active"
                                        : "")
                                }
                                onClick={() => this.changeChoice(CHOICE_AFH)}>
                                Ask For Help
                            </button>
                        </div>
                    )}
                </div>
                {this.state.choice != CHOICE_NONE && (
                    <div id="form-container">
                        <div className="register-form">
                            {this.state.choice == CHOICE_CLASS && (
                                <RegisterSelectClass
                                    onChangeStateValue={this.changeStateValue}
                                    classId={this.state.selectedClassId}
                                    classes={this.state.allClasses}
                                    classMap={this.state.classMap}
                                    programMap={this.state.programMap}
                                    semesterMap={this.state.semesterMap}
                                    locationMap={this.state.locationMap}
                                />
                            )}

                            {this.state.choice == CHOICE_AFH && (
                                <RegisterSelectAfh
                                    onChangeStateValue={this.changeStateValue}
                                    afhId={this.state.selectedAfhId}
                                    afhs={this.state.allAFHs}
                                    afhMap={this.state.afhMap}
                                    locationMap={this.state.locationMap}
                                />
                            )}

                            <RegisterFormStudent
                                onChangeStateValue={this.changeStateValue}
                                student={studentInfo}
                                valid={this.isStudentInfoValid()}
                            />

                            {/* No Guardian section if selecting AFH */}
                            {this.state.choice == CHOICE_CLASS && (
                                <RegisterFormGuardian
                                    onChangeStateValue={this.changeStateValue}
                                    student={studentInfo}
                                    guardian={guardianInfo}
                                    valid={this.isGuardianInfoValid()}
                                />
                            )}

                            {((this.isClassValid() &&
                                this.isStudentInfoValid() &&
                                this.isGuardianInfoValid()) ||
                                (this.isAfhValid() &&
                                    this.isStudentInfoValid())) && (
                                <SubmitButton
                                    choice={this.state.choice}
                                    invokeEmail={this.invokeEmail}
                                />
                            )}
                        </div>
                        <RegisterSticky
                            choice={this.state.choice}
                            invokeEmail={this.invokeEmail}
                            validClass={this.isClassValid()}
                            validAfh={this.isAfhValid()}
                            validStudent={this.isStudentInfoValid()}
                            validGuardian={this.isGuardianInfoValid()}
                        />
                    </div>
                )}
            </div>
        );
    }
}

class SubmitButton extends React.Component {
    state = {
        submitted: false,
    };

    onSubmit = () => {
        // invoke email
        this.props.invokeEmail();

        // load animation
        this.setState({ submitted: true });

        // go to page (after timeout)
        setTimeout(() => {
            changePage("/register-success/" + this.props.choice);
        }, 1800);
    };

    render() {
        return (
            <div className="submit-bar">
                <button className="submit" onClick={this.onSubmit}>
                    Confirm & Submit Registration
                </button>
                <Loader in={this.state.submitted} />
            </div>
        );
    }
}
