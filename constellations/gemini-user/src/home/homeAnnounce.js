"use strict";
require("./homeAnnounce.sass");
import React from "react";
import { Link } from "react-router-dom";
import { isEmpty, filter } from "lodash";
import srcClose from "../../assets/close_white.svg";
import API from "../utils/api.js";

export class HomeAnnounce extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            show: false,
            targetAnnounce: {},
        };
    }

    componentDidMount() {
        console.log("api attempt ");
        API.get("api/announcements/all").then((res) => {
            const announceList = res.data;
            console.log("api success!");

            let valid = filter(announceList, (a) => a.onHomePage);
            let targetAnnounce = valid.length > 0 ? valid[0] : undefined;

            targetAnnounce.shortMessage = shortenMessage(
                targetAnnounce.message
            );

            this.showTimeout = setTimeout(() => {
                this.setState({ show: true });
            }, 1000);

            this.setState({
                targetAnnounce: targetAnnounce,
            });
        });
    }

    dismiss = () => {
        this.setState({ show: false });
    };

    render() {
        const announce = this.state.targetAnnounce;
        var showAnnounce = !isEmpty(announce);
        var component;
        if (showAnnounce) {
            component = (
                <Popup
                    announce={announce}
                    show={this.state.show}
                    announceHeight={this.state.announceHeight}
                    dismiss={this.dismiss}
                />
            );
        } else {
            component = <div></div>;
        }
        return component;
    }
}

class Popup extends React.Component {
    render() {
        const announce = this.props.announce;
        const show = this.props.show;
        const announceClass = show ? "show" : "";
        return (
            <div key="real" id="home-announce" className={announceClass}>
                <div className="header-container">
                    <h3>New Announcement!</h3>
                    <button className="close-x" onClick={this.props.dismiss}>
                        <div className="img-container">
                            <img src={srcClose} />
                        </div>
                    </button>
                </div>
                <div className="text-container">
                    <p>{announce.shortMessage}</p>
                </div>
                <Link to="/announcements" onClick={this.props.dismiss}>
                    Read more &#62;
                </Link>
            </div>
        );
    }
}

function shortenMessage(message) {
    var array = message.split(" ");
    var shortMessage = "";
    var needEllipsis = false;

    var i = 0;
    while (i < array.length) {
        var append = shortMessage + array[i];
        if (append.length > 120) {
            needEllipsis = true;
            break;
        }
        shortMessage += array[i] + " ";
        i++;
    }

    if (needEllipsis) {
        shortMessage = shortMessage.slice(0, -1); // remove last "space"
        shortMessage += "...";
    }
    return shortMessage;
}
