"use strict";
require("./achieve.styl");
import React from "react";
import ReactDOM from "react-dom";
import API from "../api.js";
import { Link } from "react-router-dom";

export class AchievePage extends React.Component {
<<<<<<< HEAD
  constructor(props) {
    super(props);
    this.state = {
      list: [],
    };
  }

  componentDidMount() {
    API.get("api/achievements/v1/all").then((res) => {
      const achievements = res.data;
      this.setState({ list: achievements });
    });
  }

  render() {
    const achievements = this.state.list.map((achieve, index) => {
      return <AchieveRow key={index} achieve={achieve} />;
    });
    const numAchievements = achievements.length;
    return (
      <div id="view-achieve">
        <h1>All Achievements ({numAchievements})</h1>
        <ul id="list-heading">
          <li className="li-med">Year</li>
          <li className="li-med"> Message</li>
        </ul>
        <ul>{achievements}</ul>
        <button>
          <Link className="add-achievement" to={"/achievements/add"}>
            Add Achievement
          </Link>
        </button>
      </div>
    );
  }
}

class AchieveRow extends React.Component {
  render() {
    const achieve = this.props.achieve;
    const url = "/achievements/" + achieve.Id + "/edit";
    return (
      <ul id="achieve-row">
        <li className="li-med">{achieve.year}</li>
        <li className="li-med">{achieve.message}</li>
        <Link to={url}>Edit</Link>
      </ul>
    );
  }
=======
    constructor(props) {
        super(props);
        this.state = {
            list: [],
        };
    }

    componentDidMount() {
        API.get("api/achievements/v1/all").then((res) => {
            const achievements = res.data;
            this.setState({ list: achievements });
        });
    }

    render() {
        const achievements = this.state.list.map((achieve, index) => {
            return <AchieveRow key={index} achieve={achieve} />;
        });
        const numAchievements = achievements.length;
        return (
            <div id="view-achieve">
                <h1>All Achievements ({numAchievements})</h1>
                <ul id="list-heading">
                    <li className="li-med">Year</li>
                    <li className="li-med"> Message</li>
                </ul>
                <ul>{achievements}</ul>
                <button>
                    <Link className="add-achievement" to={"/achievements/add"}>
                        Add Achievement
                    </Link>
                </button>
            </div>
        );
    }
}

class AchieveRow extends React.Component {
    render() {
        const achieve = this.props.achieve;
        const url = "/achievements/" + achieve.Id + "/edit";
        return (
            <ul id="achieve-row">
                <li className="li-med">{achieve.year}</li>
                <li className="li-med">{achieve.message}</li>
                <Link to={url}>Edit</Link>
            </ul>
        );
    }
>>>>>>> a27fb3b5070f8e1928daed628fb9a9038d1e89b9
}
