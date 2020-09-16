"use strict";
require("./announce.sass");
import React from "react";
import moment from "moment";
import API from "../api.js";
import AllPageHeader from "../utils/allPageHeader.js";
import RowCardBasic from "../utils/rowCardBasic.js";

export class AnnouncePage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            list: [],
        };
    }

    componentDidMount() {
        API.get("api/announcements/all").then((res) => {
            const announcements = res.data;
            this.setState({ list: announcements });
        });
    }

    onChangeCheckbox = (e) => {
        let successCallback = () => console.log("api success");
        let failCallback = (err) => {
            alert("Could not save announcement: " + err.response.data);
            console.log(err.response);
        };

        let previous = this.state.list.findIndex(
            (announcement) => announcement.onHomePage
        );
        let current = this.state.list.findIndex(
            (announcement) => announcement.id == e.target.id
        );

        let newList = this.state.list;

        newList[current].onHomePage = true;
        if (previous >= 0) {
            newList[previous].onHomePage = false;
        }

        this.setState({ list: newList });

        let indexes = previous == current ? [previous] : [previous, current];

        indexes.forEach((index) => {
            if (index >= 0) {
                API.post(
                    "api/announcements/announcement/" +
                        this.state.list[index].id,
                    this.state.list[index]
                )
                    .then((res) => successCallback())
                    .catch((err) => failCallback(err));
            }
        });
    };

    render() {
        const cards = this.state.list.map((announce, index) => {
            const postedAt = moment(announce.postedAt);
            const fields = generateFields(announce);
            const texts = generateTexts(announce);
            return (
                <div className="card-wrapper" key={index}>
                    <RowCardBasic
                        key={index}
                        title={"Announcement on " + postedAt.format("M/D/YYYY")}
                        subtitle={postedAt.format("hh:mm a")}
                        editUrl={"/announcements/" + announce.id + "/edit"}
                        fields={fields}
                        texts={texts}
                    />
                </div>
            );
        });

        return (
            <div id="view-announce">
                <AllPageHeader
                    title={"All Announcements (" + this.state.list.length + ")"}
                    addUrl={"/announcements/add"}
                    addButtonTitle={"Add Announcement"}
                    description={
                        "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."
                    }
                />
                <div className="cards">{cards}</div>
            </div>
        );
    }
}

function generateFields(announce) {
    const now = moment();
    const postedAt = moment(announce.postedAt);
    const isPublic = postedAt.isAfter(now);
    return [
        {
            label: "Author",
            value: announce.author,
        },
        {
            label: "OnHomePage",
            value: announce.onHomePage ? "true" : "false",
            highlightFn: () => announce.onHomePage,
        },
        {
            label: "Posted",
            value: isPublic ? "Published" : "Scheduled",
            highlightFn: () => !isPublic,
        },
    ];
}

function generateTexts(announce) {
    return [
        {
            label: "Message",
            value: announce.message,
        },
    ];
}

// class AnnounceRow extends React.Component {
//     render() {
//         const announceId = this.props.row.id;
//         const postedAt = moment(this.props.row.postedAt);

//         const checked = this.props.row.onHomePage || false;

//         const now = moment();
//         const isScheduled = postedAt.isAfter(now);

//         const author = this.props.row.author;
//         const message = this.props.row.message;

//         const url = "/announcements/" + announceId + "/edit";
//         return (
//             <ul className="announce-list-row">
//                 <li className="li-small">
//                     <input
//                         type="checkbox"
//                         onChange={this.props.onChangeCheckbox}
//                         id={announceId}
//                         checked={checked}
//                     />
//                 </li>
//                 <li
//                     className={
//                         "li-small " +
//                         (isScheduled ? " scheduled" : " published")
//                     }>
//                     <div>{isScheduled ? "Scheduled" : "Published"}</div>
//                     <div>{postedAt.fromNow()}</div>
//                 </li>
//                 <li className="li-small">
//                     <div>{postedAt.format("M/D/YYYY")}</div>
//                     <div>{postedAt.format("hh:mm a")}</div>
//                 </li>
//                 <li className="li-small"> {author} </li>
//                 <li className="li-large"> {message} </li>
//                 <Link to={url}>Edit</Link>
//             </ul>
//         );
//     }
// }
