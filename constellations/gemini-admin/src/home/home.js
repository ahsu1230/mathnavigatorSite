"use strict";
require("./home.sass");
import React from "react";
import API from "../api.js";
import { Link } from "react-router-dom";

export class HomePage extends React.Component {
    state = {
        classes: [],
    };

    componentDidMount() {
        this.fetchData();
    }

    fetchData = () => {
        API.get("api/unpublished").then((res) => {
            const unpublishedList = res.data;
            this.setState({
                classes: unpublishedList.classes,
            });
        });
    };

    render() {
        let unpublishedClasses = this.state.classes.map((row, index) => {
            return <li key={index}> {row.classId} </li>;
        });

        return (
            <div id="view-home">
                <h1>Administrator Dashboard</h1>

                <div className="section">
                    <div className="container-class">
                        <h3 className="section-header">Unpublished Classes</h3>{" "}
                        <button id="publish">
                            <Link to={"/classes"}>View Details to Publish</Link>
                        </button>
                    </div>

                    <div className="class-section">
                        <div className="list-header">Class ID</div>
                        <ul>{unpublishedClasses}</ul>
                    </div>
                </div>

                <div className="section">
                    <h3 className="section-header">New Users</h3>
                    <div className="container-user">
                        <p>
                            Lorem ipsum dolor sit amet, vim ei tota dicant
                            interpretaris, sea consulatu scripserit ei, pri eu
                            accumsan contentiones. Usu ei iriure deleniti. Cum
                            ut periculis laboramus referrentur. Vis ea laoreet
                            imperdiet deterruisset, probo pertinax iudicabit qui
                            ea. Duo liber quodsi contentiones ne, vel ei movet
                            conceptam, id his quod iriure feugiat. Qui ex vide
                            labitur volumus. His cu fugit adolescens
                            voluptatibus, et per torquatos interesset. Iusto
                            deleniti invenire id cum. Te saepe alterum
                            appellantur pro, in est errem dicant suscipit,
                            oblique argumentum sed et. Quo cu partem inermis,
                            vix et minimum vivendum, ut illud delectus eos. Mei
                            te alia justo, per amet nemore quodsi cu, mel
                            eruditi copiosae contentiones cu. Sed ea quando
                            mediocrem, soleat dolorum no ius. Eu electram
                            iracundia mnesarchum his, eum ut hinc latine. Est
                            simul definiebas ut, debitis invenire eu usu. Eu
                            eros appetere mel, nullam delenit tincidunt duo at.
                            Eos option appetere torquatos in, cu debitis
                            singulis principes usu.
                        </p>
                    </div>
                </div>

                <div className="section">
                    <h3 className="section-header">New Registrations</h3>
                    <div className="container-registration">
                        <p>
                            Lorem ipsum dolor sit amet, vim ei tota dicant
                            interpretaris, sea consulatu scripserit ei, pri eu
                            accumsan contentiones. Usu ei iriure deleniti. Cum
                            ut periculis laboramus referrentur. Vis ea laoreet
                            imperdiet deterruisset, probo pertinax iudicabit qui
                            ea. Duo liber quodsi contentiones ne, vel ei movet
                            conceptam, id his quod iriure feugiat. Qui ex vide
                            labitur volumus. His cu fugit adolescens
                            voluptatibus, et per torquatos interesset. Iusto
                            deleniti invenire id cum. Te saepe alterum
                            appellantur pro, in est errem dicant suscipit,
                            oblique argumentum sed et. Quo cu partem inermis,
                            vix et minimum vivendum, ut illud delectus eos. Mei
                            te alia justo, per amet nemore quodsi cu, mel
                            eruditi copiosae contentiones cu. Sed ea quando
                            mediocrem, soleat dolorum no ius. Eu electram
                            iracundia mnesarchum his, eum ut hinc latine. Est
                            simul definiebas ut, debitis invenire eu usu. Eu
                            eros appetere mel, nullam delenit tincidunt duo at.
                            Eos option appetere torquatos in, cu debitis
                            singulis principes usu.
                        </p>
                    </div>
                </div>
            </div>
        );
    }
}
