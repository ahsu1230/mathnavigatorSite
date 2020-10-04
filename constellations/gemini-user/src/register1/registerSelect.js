"use strict";
require("./registerSelect.sass");
import React from "react";
import moment from "moment";
import { isFullClass } from "../utils/classUtils.js";
import { capitalizeWord } from "../utils/utils.js";
import { 
    RegisterSectionBase,
    REGISTER_SECTION_FORM_STUDENT,
} from "./registerBase.js";

export default class RegisterSectionSelect extends React.Component {
    state = {
        chosen: "",
    }

    renderContent = () => {
        let selectSection;
        if (this.state.chosen == "class") {
            selectSection=(<SelectClass 
                onChangeSection={this.props.onChangeSection}
                onChangeStateValue={this.props.onChangeStateValue}
                classId={this.props.classId}
                classes={this.props.classes}
                classMap={this.props.classMap}
                programMap={this.props.programMap}
                semesterMap={this.props.semesterMap}
                locationMap={this.props.locationMap}/>);
        } else if (this.state.chosen == "afh") {
            selectSection=(<SelectAfh 
                onChangeSection={this.props.onChangeSection}
                onChangeStateValue={this.props.onChangeStateValue}
                afhId={this.props.afhId}
                afhs={this.props.afhs}
                afhMap={this.props.afhMap}
                locationMap={this.props.locationMap}/>);
        } else {
            selectSection=(<div></div>);
        }

        return (
            <div className="content">
                <div className="intro-wrapper">
                    <p>
                        Choose either to enroll into a Math Navigator class or an ask-for-help session.
                        Ask For Help sessions are only for students who are already enrolled into one of our classes.
                    </p>
                    <div className="buttons">
                        <button className="class-btn" onClick={() => this.setState({chosen: "class"})}>
                            Enroll into a class
                        </button>
                        <button className="afh-btn" onClick={() => this.setState({chosen: "afh"})}>
                            Ask For Help
                        </button>
                    </div>
                </div>
                {selectSection}
            </div>
        );
    }

    render() {
        return (
            <RegisterSectionBase
                sectionName="select"
                title={"Registration"}
                content={this.renderContent()}
                onChangeSection={this.props.onChangeSection}
            />
        );
    }
}

export class SelectClass extends React.Component {
    onChangeClass = (e) => {
        this.props.onChangeStateValue("selectedClassId", e.target.value);
        this.props.onChangeStateValue("selectedAfhId", null);
    }

    renderCurrentClass = () => {
        if (this.props.classId) {
            const currentClass = this.props.classMap[this.props.classId];
            const program = this.props.programMap[currentClass.programId];
            const semester = this.props.semesterMap[currentClass.semesterId];
            const location = this.props.locationMap[currentClass.locationId];

            const fullTitle = program.title + " " + capitalizeWord(currentClass.classKey);
            const fullSection = isFullClass(currentClass) ? <p className="error">This class is full. Please select another class to enroll.</p> : <div></div>;
            let showNextButton = isFullClass(currentClass) ? <div></div> : (
                <div className="next-section">
                    <button onClick={() => this.props.onChangeSection(REGISTER_SECTION_FORM_STUDENT)}>
                        Next
                    </button>
                </div>
            );
            return (
                <section className="selected">
                    <div>
                        You have selected to enroll into:
                        <div className="info">
                            {fullSection}
                            <h3>{fullTitle}</h3>
                            <h4>{semester.title}</h4>
                            <p>Times: {currentClass.timesStr}</p>
                            <p>
                                Location: {location.title}<br/>
                                {location.street}<br/>
                                {location.city + ", " + location.state + " " + location.zipcode}<br/>
                                {location.room}
                            </p>
                        </div>
                    </div>
                    {showNextButton}
                </section>
            );
        } else {
            return (<div></div>);
        }
    }

    render() {
        const optionsClasses = this.props.classes.map((classObj, index) => {
            const program = this.props.programMap[classObj.programId];
            const semester = this.props.semesterMap[classObj.semesterId];
            const fullTitle = program.title + " " + capitalizeWord(classObj.classKey) + " " + semester.title;
            return (<option key={index} value={classObj.classId}>{fullTitle}</option>);
        });

        const currentClassSection = this.renderCurrentClass();
        return (
            <div className="selection class">
                <p>Select the class you would like to enroll into.</p>
                <select 
                    value={this.props.classId || "none"}
                    onChange={this.onChangeClass}>
                    <option disabled value={"none"}>
                        -- Select a class --
                    </option>
                    {optionsClasses}
                </select>
                {currentClassSection}
            </div>
        );
    }
}

export class SelectAfh extends React.Component {
    onChangeAfh = (e) => {
        this.props.onChangeStateValue("selectedClassId", null);
        this.props.onChangeStateValue("selectedAfhId", e.target.value);
    }

    renderCurrentAfh = () => {
        if (this.props.afhId) {
            const currentAfh = this.props.afhMap[this.props.afhId]
            const datetime = moment(currentAfh.startsAt).format("MM/DD/YY h:mm a") + 
                                " - " + 
                                moment(currentAfh.endsAt).format("h:mm a");
            const location = this.props.locationMap[currentAfh.locationId];
            return (
                <section className="selected">
                    <div>
                        You have selected to attend:
                        <div className="info">
                            <h3>{currentAfh.title}</h3>
                            <h4>{datetime}</h4>
                            <p>
                                Location: {location.title}<br/>
                                {location.street}<br/>
                                {location.city + ", " + location.state + " " + location.zipcode}<br/>
                                {location.room}
                            </p>
                        </div>
                    </div>
                    
                    <div className="next-section">
                        <button onClick={() => this.props.onChangeSection(REGISTER_SECTION_FORM_STUDENT)}>
                            Next
                        </button>
                    </div>
                </section>
            );
        } else {
            return (<div></div>);
        }
    }

    render() {
        const optionsAfh = this.props.afhs.map((afh, index) => {
            const afhTime = moment(afh.startsAt).format("MM/DD/YY h:mm a");
            const fullTitle = afh.title + " " + afhTime;
            return (<option key={index} value={afh.id}>{fullTitle}</option>);
        });

        const currentAfhSection = this.renderCurrentAfh();
        return (
            <div className="selection afh">
                <p>Select the ask-for-help session you would like to attend.</p>
                <select
                    value={this.props.afhId || "none"}
                    onChange={this.onChangeAfh}>
                    <option disabled value={"none"}>
                        -- Select an AskForHelp session --
                    </option>
                    {optionsAfh}
                </select>
                {currentAfhSection}
            </div>
        );
    }
}