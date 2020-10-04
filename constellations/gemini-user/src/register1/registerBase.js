"use strict";
require("./register.sass");
import React from "react";

export const REGISTER_SECTION_SELECT = "select";
export const REGISTER_SECTION_FORM_STUDENT = "form-student";
export const REGISTER_SECTION_FORM_GUARDIAN = "form-guardian";
export const REGISTER_SECTION_CONFIRM = "confirm";
export const REGISTER_SECTION_SUCCESS = "success";

export class RegisterSectionBase extends React.Component {
    onClickNext = () => {
        if (this.props.nextAllowed) {
            this.props.onChangeSection(this.props.next);
        }
    };

    renderButton = () => {
        const next = this.props.next;
        const nextAllowed = this.props.nextAllowed;

        if (this.props.sectionName == "confirm") {
            return (
                <button className="next confirm" onClick={this.onClickNext}>
                    Confirm
                </button>
            );
        } else if (next && nextAllowed) {
            return (
                <button className="next allowed" onClick={this.onClickNext}>
                    Next
                </button>
            );
        } else {
            return <div></div>;
        }
    };

    render() {
        const sectionName = this.props.sectionName;
        const active = this.props.index == this.props.currentIndex;
        return (
            <section
                className={
                    "register-section " +
                    sectionName +
                    (active ? " active" : "")
                }>
                <h1>{this.props.title}</h1>
                <div className="content-wrapper">{this.props.content}</div>
                <div className="buttons-footer">
                    {this.props.prev ? (
                        <button
                            className="back"
                            onClick={() =>
                                this.props.onChangeSection(this.props.prev)
                            }>
                            Back
                        </button>
                    ) : (
                        <div></div>
                    )}
                    {this.renderButton()}
                </div>
            </section>
        );
    }
}
