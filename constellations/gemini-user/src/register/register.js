"use strict";
require("./register.sass");
import React from "react";
import axios from "axios";
import API from "../utils/api.js";
import { keyBy } from "lodash";
import { parseQueryParams } from "../utils/utils.js";
import RegisterSticky from "./registerSticky.js";
import RegisterSelectClass from "./registerClass.js";
import RegisterSelectAfh from "./registerAfh.js";
import { RegisterFormStudent, RegisterFormGuardian} from "./registerForms.js";
import { validateEmail, validatePhone } from "../utils/validators.js";
import { isFullClass } from "../utils/classUtils.js";

const CHOICE_NONE = "";
const CHOICE_CLASS = "class";
const CHOICE_AFH = "afh";

export default class RegisterPage extends React.Component {
    state = {
        location: {},
        choice: CHOICE_NONE,
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
        guardianAdditionalInfo: "",

        allAFHs: [],
        afhMap: {},
        allClasses: [],
        classMap: {},
        programMap: {},
        semesterMap: {},
        locationMap: {},
    }

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

                    let queries = parseQueryParams(this.props.history.location.search);
                    this.setState({
                        selectedAfhId: queries["afhId"] || null,
                        selectedClassId: queries["classId"] || null,
                        allClasses: allClasses,
                        classMap: keyBy(allClasses, 'classId'),
                        allAFHs: allAFHs,
                        afhMap: keyBy(allAFHs, "id"),
                        programMap: keyBy(allPrograms, "programId"),
                        semesterMap: keyBy(allSemesters, "semesterId"),
                        locationMap: keyBy(allLocations, "locationId")
                    });
                }
            )).catch(err =>
                console.log("Error: could not fetch class: " + err)
            );
    }

    componentWillUnmount() {
    }

    changeStateValue = (stateName, value) => {
        this.setState({
            [stateName]: value
        });
    }

    isClassValid = () => {
        const classId = this.state.selectedClassId;
        return classId && !isFullClass(this.state.classMap[classId]);
    }

    isAfhValid = () => {
        return !!this.state.selectedAfhId;
    }

    isStudentInfoValid = () => {
        return this.state.studentFirstName 
            && this.state.studentLastName
            && validateEmail(this.state.studentEmail)
            && this.state.studentSchool
            && this.state.studentGraduationYear > 2019
            && this.state.studentGraduationYear < 2030
            && this.state.studentGrade > 5
            && this.state.studentGrade < 13;
    }

    isGuardianInfoValid = () => {
        return this.state.guardianFirstName 
            && this.state.guardianLastName
            && !(this.state.guardianFirstName == this.state.studentFirstName && this.state.guardianLastName == this.state.studentLastName)
            && validateEmail(this.state.guardianEmail)
            && this.state.studentEmail != this.state.guardianEmail
            && validatePhone(this.state.guardianPhone);
    }

    render() {
        const studentInfo = {
            firstName: this.state.studentFirstName,
            lastName: this.state.studentLastName,
            email: this.state.studentEmail,
            school: this.state.studentSchool,
            grade: this.state.studentGrade,
            graduationYear: this.state.studentGraduationYear
        };
        const guardianInfo = {
            firstName: this.state.guardianFirstName,
            lastName: this.state.guardianLastName,
            email: this.state.guardianEmail,
            phone: this.state.guardianPhone,
            additionalInfo: this.state.guardianAdditionalInfo
        }
        return (
            <div id="view-register">
                <h1>Registration</h1>
                <div className="intro-wrapper">
                    <div>
                        <p>
                            Are you enrolling into a class or attending an ask-for-help session? 
                            Ask-for-Help sessions are only for students who are already enrolled into 
                            one of our classes.<br/><br/>
                            You will be asked to fill out some information about yourself.
                            We use this information to contact you about important class updates,
                            so please use a valid email you frequently use.
                            This information is used for <i>contacting purposes only</i> and will NOT be shared with anyone.
                        </p>
                    </div>
                    <div className="choose-btns">
                        <button className={"class-btn" + (this.state.choice == CHOICE_CLASS ? " active" : "")} 
                                onClick={() => this.setState({choice: CHOICE_CLASS})}>
                            Enroll into a class
                        </button>
                        <button className={"afh-btn" + (this.state.choice == CHOICE_AFH ? " active" : "")} 
                                onClick={() => this.setState({choice: CHOICE_AFH})}>
                            Ask For Help
                        </button>
                    </div>
                </div>
                {
                    this.state.choice != CHOICE_NONE && 
                    <div id="form-container">
                        <div className="register-form">
                            
                            {this.state.choice == CHOICE_CLASS && 
                                <RegisterSelectClass
                                    onChangeStateValue={this.changeStateValue}
                                    classId={this.state.selectedClassId}
                                    classes={this.state.allClasses}
                                    classMap={this.state.classMap}
                                    programMap={this.state.programMap}
                                    semesterMap={this.state.semesterMap}
                                    locationMap={this.state.locationMap}/>}
                            
                            {this.state.choice == CHOICE_AFH && 
                                <RegisterSelectAfh
                                    onChangeStateValue={this.changeStateValue}
                                    afhId={this.state.selectedAfhId}
                                    afhs={this.state.allAFHs}
                                    afhMap={this.state.afhMap}
                                    locationMap={this.state.locationMap}/>}
                            
                            <RegisterFormStudent
                                onChangeStateValue={this.changeStateValue}
                                student={studentInfo}
                                valid={this.isStudentInfoValid()}
                                />
                            
                            {/* No Guardian section if selecting AFH */}
                            {this.state.choice == CHOICE_CLASS && 
                                <RegisterFormGuardian
                                    onChangeStateValue={this.changeStateValue}
                                    student={studentInfo}
                                    guardian={guardianInfo}
                                    valid={this.isGuardianInfoValid()}
                                    />}
                            
                        </div>
                        <RegisterSticky
                            choice={this.state.choice}
                            validClass={this.isClassValid()}
                            validAfh={this.isAfhValid()}
                            validStudent={this.isStudentInfoValid()}
                            validGuardian={this.isGuardianInfoValid()}/>
                    </div>
                }
            </div>
        );
    }
}