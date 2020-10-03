"use strict";
require("./registerConfirm.sass");
import React from "react";
import { isEmpty } from "lodash";
import { 
    RegisterSectionBase,
    REGISTER_SECTION_SELECT,
    REGISTER_SECTION_FORM_STUDENT,
    REGISTER_SECTION_FORM_GUARDIAN,
    REGISTER_SECTION_SUCCESS
} from "./registerBase.js";
import { capitalizeWord } from "../utils/utils.js";

export default class RegisterSectionConfirm extends React.Component {
    renderContent = () => {
        return (
            <div className="content">
                <p className="review">
                    Please review the information below. 
                    Double check your contact information because that is what will be used to contact you.
                    If everything is correct, please click on <b>Confirm</b> at the bottom of the page to submit your request.
                </p>
                <SectionClass
                    onEdit={() => this.props.onChangeSection(REGISTER_SECTION_SELECT)}
                    selectedClass={this.props.selectedClass}
                    programMap={this.props.programMap}
                    semesterMap={this.props.semesterMap}
                    locationMap={this.props.locationMap}/>
                <SectionAfh
                    onEdit={() => this.props.onChangeSection(REGISTER_SECTION_SELECT)}
                    selectedAfh={this.props.selectedAfh}
                    locationMap={this.props.locationMap}/>
                <SectionStudent 
                    onEdit={() => this.props.onChangeSection(REGISTER_SECTION_FORM_STUDENT)}
                    student={this.props.student}/>
                <SectionGuardian 
                    onEdit={() => this.props.onChangeSection(REGISTER_SECTION_FORM_GUARDIAN)}
                    guardian={this.props.guardian}/>
            </div>
        );
    }

    render() {
        return (
            <RegisterSectionBase
                sectionName="confirm"
                title={"Confirm Registration"}
                nextAllowed={true}
                next={REGISTER_SECTION_SUCCESS}
                prev={REGISTER_SECTION_FORM_GUARDIAN}
                content={this.renderContent()}
                onChangeSection={this.props.onChangeSection}
            />
        );
    }
}

class SectionClass extends React.Component {
    render() {
        const classObj = this.props.selectedClass;
        if (isEmpty(classObj)) {
            return (<div></div>);
        } else {
            const program = this.props.programMap[classObj.programId];
            const semester = this.props.semesterMap[classObj.semesterId];
            const location = this.props.locationMap[classObj.locationId];
            const fullTitle = program.title + " " + capitalizeWord(classObj.classKey);
            return(
                <section className="confirm class">
                    <div className="container">
                        <h3>Class Enrollment Request</h3>
                        <a onClick={this.props.onEdit}>Edit</a>
                    </div>
                    <p>
                        {fullTitle}<br/>
                        {semester.title}<br/>
                        {classObj.timesStr}<br/>
                        {/* class price */}
                    </p>
                    <p>
                        Location: {location.title}<br/>
                        {location.street}<br/>
                        {location.city + ", " + location.state + " " + location.zipcode}<br/>
                        {location.room}
                    </p>
                </section>
            );
        }
    }
}

class SectionAfh extends React.Component {
    render() {
        const afh = this.props.selectedAfh;
        if (isEmpty(afh)) {
            return (<div></div>);
        } else {
            const location = this.props.locationMap[afh.locationId];
            return(
                <section className="confirm afh">
                    <div className="container">
                        <h3>Ask-for-Help Session Request</h3>
                        <a onClick={this.props.onEdit}>Edit</a>
                    </div>
                    <p>
                        {afh.title}<br/>
                        {/* afh times (Display Utils) */}
                    </p>
                    <p>
                        Location: {location.title}<br/>
                        {location.street}<br/>
                        {location.city + ", " + location.state + " " + location.zipcode}<br/>
                        {location.room}
                    </p>
                </section>
            );
        }
    }
}

class SectionStudent extends React.Component {
    render() {
        const student = this.props.student;
        return(
            <section className="confirm student">
                <div className="container">
                    <h3>Student Information</h3>
                    <a onClick={this.props.onEdit}>Edit</a>
                </div>
                <p>
                    {student.firstName}{" "}{student.lastName}<br/>
                    {student.email}<br/>
                    School: {student.school}<br/>
                    Grade: {student.grade}<br/>
                    Graduation Year: {student.graduationYear}
                </p>
                
            </section>
        );
    }
}

class SectionGuardian extends React.Component {
    render() {
        const guardian = this.props.guardian;
        if (isEmpty(guardian)) {
            return (<div></div>);
        } else {
            return(
                <section className="confirm guardian">
                    <div className="container">
                        <h3>Guardian Information</h3>
                        <a onClick={this.props.onEdit}>Edit</a>
                    </div>
                    <p>
                        {guardian.firstName}{" "}{guardian.lastName}<br/>
                        {guardian.email}<br/>
                        {guardian.phone}
                    </p>
                </section>
            );
        }
    }
}