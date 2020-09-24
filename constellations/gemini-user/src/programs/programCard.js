"use strict";
require("./programCard.sass");
import React from "react";
import { Modal } from "../modals/modal.js";
import { ProgramModal } from "./programModal.js";
import { capitalizeWord } from "../utils/utils.js";

export class ProgramCard extends React.Component {
    state = {
        showModal: false,
    };

    handleClick = () => {
        const classes = this.props.classes;
        if (classes.length == 1) {
            window.location.hash = "/class/" + classes[0].classId;
        } else if (classes.length > 1) {
            this.setState({ showModal: true });
        }
    };

    renderModal = () => {
        const classes = this.props.classes || [];
        let modalDiv = <div></div>;

        if (classes.length > 1) {
            const modalContent = (
                <ProgramModal
                    semester={this.props.semester}
                    program={this.props.program}
                    classes={classes}
                    fullStates={this.props.fullStates}
                />
            );
            modalDiv = (
                <Modal
                    content={modalContent}
                    show={this.state.showModal}
                    withClose={true}
                    onDismiss={() => this.setState({ showModal: false })}
                />
            );
        }
        return modalDiv;
    };

    render = () => {
        const program = this.props.program || {};
        const classes = this.props.classes || [];
        const grades = "Grades " + program.grade1 + " - " + program.grade2;
        const classesString =
            classes.length == 1 ? "1 class" : classes.length + " classes";
        const buttonString =
            program.featured != "none"
                ? capitalizeWord(program.featured)
                : "View Details";

        return (
            <div className="program-card-container">
                <div className="program-card" onClick={this.handleClick}>
                    <div className="content">
                        <h2>{program.title}</h2>
                        <h3>{grades}</h3>
                        <div className="classes-avail">
                            {classesString} available
                        </div>
                        <button>{buttonString}</button>
                    </div>
                </div>
                {this.renderModal()}
            </div>
        );
    };
}
