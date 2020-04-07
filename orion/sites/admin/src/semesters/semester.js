require("./semester.styl");
import React from "react";
import ReactDOM from "react-dom";
import { Link } from "react-router-dom";

export class SemesterPage extends React.Component {
<<<<<<< HEAD
<<<<<<< HEAD
  constructor(props) {
    super(props);
    this.state = {
      list: [],
    };
  }

  render() {
    const rows = this.state.list.map((row, index) => {
      return <SemesterRow key={index} row={row} />;
    });
    const numRows = rows.length;
    return (
      <div id="view-semester">
        <h1>All Semesters ({numRows}) </h1>
        <ul className="semester-lists">
          <li className="li-large">SemesterID</li>
          <li className="li-large">Title</li>
          <li className="li-small"> </li>
        </ul>
        <ul id="view-semester">{rows}</ul>
        <Link to={"/semesters/add"}>
          {" "}
          <button className="semester-button"> Add Semester</button>{" "}
        </Link>
      </div>
    );
  }
}

class SemesterRow extends React.Component {
  render() {
    const semesterId = this.props.semesterObj.semesterId;
    const title = this.props.semesterObj.title;
    const url = "/semester/" + "/edit";
    return (
      <ul className="semester-lists">
        <li className="li-large"> {semesterId} </li>
        <li className="li-large"> {title} </li>
        <Link to={url}> Edit </Link>
      </ul>
    );
  }
=======
=======
>>>>>>> c15f24dc4318ffae807d39aef3ef62f1b6948b26
    constructor(props) {
        super(props);
        this.state = {
            list: [],
        };
    }

    render() {
        const rows = this.state.list.map((row, index) => {
            return <SemesterRow key={index} row={row} />;
        });
        const numRows = rows.length;
        return (
            <div id="view-semester">
                <h1>All Semesters ({numRows}) </h1>
                <ul className="semester-lists">
                    <li className="li-large">SemesterID</li>
                    <li className="li-large">Title</li>
                    <li className="li-small"> </li>
                </ul>
                <ul id="view-semester">{rows}</ul>
                <Link to={"/semesters/add"}>
                    {" "}
                    <button className="semester-button">
                        {" "}
                        Add Semester
                    </button>{" "}
                </Link>
            </div>
        );
    }
}

class SemesterRow extends React.Component {
    render() {
        const semesterId = this.props.semesterObj.semesterId;
        const title = this.props.semesterObj.title;
        const url = "/semester/" + "/edit";
        return (
            <ul className="semester-lists">
                <li className="li-large"> {semesterId} </li>
                <li className="li-large"> {title} </li>
                <Link to={url}> Edit </Link>
            </ul>
        );
    }
<<<<<<< HEAD
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
=======
>>>>>>> c15f24dc4318ffae807d39aef3ef62f1b6948b26
}
