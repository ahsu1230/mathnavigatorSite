"use strict";
import React from "react";
import { Transition } from "react-transition-group";
import srcCheckmarkGreen from "../../assets/checkmark_light_blue.svg";

const defaultBarStyle = {
    position: "absolute",
    bottom: 0,
    left: 0,
    height: "4px",
    width: "0px",
    transition: `width 1200ms cubic-bezier(0.4, 0.6, 0.9, 0.1)`,
    backgroundColor: "#056571",
};

const transitionBarStyles = {
    entering: { width: "100%" },
    entered: { width: "100%" },
    exiting: { width: 0 },
    exited: { width: 0 },
};

const defaultCheckStyle = {
    position: "absolute",
    bottom: 0,
    right: "4px",
    width: "20px",
    opacity: 0.0,
    transition: `all 100ms ease-in 1250ms`,
};

const transitionCheckStyles = {
    entering: { opacity: 1.0, bottom: "8px" },
    entered: { opacity: 1.0, bottom: "8px" },
    exiting: { opacity: 0.0, bottom: "0px" },
    exited: { opacity: 0.0, bottom: "0px" },
};

export default class Loader extends React.Component {
    render() {
        return (
            <div className="loader">
                <Transition in={this.props.in} timeout={1200}>
                    {(state) => (
                        <div
                            className="highlight"
                            style={{
                                ...defaultBarStyle,
                                ...transitionBarStyles[state],
                            }}></div>
                    )}
                </Transition>
                <Transition in={this.props.in} timeout={300}>
                    {(state) => (
                        <img
                            src={srcCheckmarkGreen}
                            style={{
                                ...defaultCheckStyle,
                                ...transitionCheckStyles[state],
                            }}
                        />
                    )}
                </Transition>
            </div>
        );
    }
}
