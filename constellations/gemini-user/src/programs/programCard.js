"use strict";
require("./programCard.sass");
import React from "react";
import { Modal } from "../modals/modal.js";
import { ProgramModal } from "./programModal.js";
import { capitalizeWord } from "../utils/displayUtils.js";

import srcPointGreen from "../../assets/point_right_green.svg";
import srcPointWhite from "../../assets/point_right_white.svg";
import srcStar from "../../assets/star_green.svg";

export class ProgramCard extends React.Component {
    state = {
        showModal: false,
        srcPoint: srcPointGreen,
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

    renderFeatured = () => {
        const program = this.props.program || {};
        const isFeatured = program.featured != "none";
        const featuredString = isFeatured
            ? capitalizeWord(program.featured)
            : "";

        if (isFeatured) {
            return (
                <div className="featured">
                    <div className="container">
                        <img src={srcStar} />
                        <span>{featuredString}</span>
                    </div>
                </div>
            );
        } else {
            return <div></div>;
        }
    };

    render = () => {
        const program = this.props.program || {};
        const grades = "Grades " + program.grade1 + " - " + program.grade2;
        const featured = this.renderFeatured();

        return (
            <div className="program-card-container">
                <div
                    className="program-card"
                    onClick={this.handleClick}
                    onMouseOver={() =>
                        this.setState({ srcPoint: srcPointWhite })
                    }
                    onMouseOut={() =>
                        this.setState({ srcPoint: srcPointGreen })
                    }>
                    {featured}

                    <div className="content">
                        <h1>{program.title}</h1>
                        <h3>{grades}</h3>
                        <button className="footer">
                            View Details
                            <img src={this.state.srcPoint} />
                        </button>
                    </div>
                </div>
                {this.renderModal()}
            </div>
        );
    };
}
