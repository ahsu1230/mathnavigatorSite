"use strict";
require("./home.sass");
import React from "react";
import API from "../api.js";

export class HomePage extends React.Component {
    render() {
        return (
            <div id="view-home">
                <h1>Adminstrator Dashboard</h1>

                <div class="container-class">
                    <span>
                        <h3>Unpublished Classes</h3>{" "}
                    </span>
                    <button id="publish">View Details to Publish</button>
                </div>

                <div class="class-section">
                    <p class="classID">Class ID</p>
                    <p class="edit"> </p>
                </div>
                <div class="unpublished-classes">
                    <p class="classID">ap_calculus_spring_2020</p>
                    <p class="edit">
                        {" "}
                        <button class="editB">Edit</button>{" "}
                    </p>
                </div>
                <div class="unpublished-classes">
                    <p class="classID">ap_java_summer_2020_class_a</p>
                    <p class="edit">
                        {" "}
                        <button class="editB">Edit</button>{" "}
                    </p>
                </div>

                <div class="divider">
                    <br></br> <br></br> <br></br>
                </div>

                <h3>New Users</h3>
                <div class="container-user">
                    <p>
                        Lorem ipsum dolor sit amet, vim ei tota dicant
                        interpretaris, sea consulatu scripserit ei, pri eu
                        accumsan contentiones. Usu ei iriure deleniti. Cum ut
                        periculis laboramus referrentur. Vis ea laoreet
                        imperdiet deterruisset, probo pertinax iudicabit qui ea.
                        Duo liber quodsi contentiones ne, vel ei movet
                        conceptam, id his quod iriure feugiat. Qui ex vide
                        labitur volumus. His cu fugit adolescens voluptatibus,
                        et per torquatos interesset. Iusto deleniti invenire id
                        cum. Te saepe alterum appellantur pro, in est errem
                        dicant suscipit, oblique argumentum sed et. Quo cu
                        partem inermis, vix et minimum vivendum, ut illud
                        delectus eos. Mei te alia justo, per amet nemore quodsi
                        cu, mel eruditi copiosae contentiones cu. Sed ea quando
                        mediocrem, soleat dolorum no ius. Eu electram iracundia
                        mnesarchum his, eum ut hinc latine. Est simul definiebas
                        ut, debitis invenire eu usu. Eu eros appetere mel,
                        nullam delenit tincidunt duo at. Eos option appetere
                        torquatos in, cu debitis singulis principes usu.
                    </p>
                </div>

                <div class="divider">
                    <br></br> <br></br> <br></br>
                </div>

                <h3>New Registrations</h3>
                <div class="container-registration">
                    <p>
                        Lorem ipsum dolor sit amet, vim ei tota dicant
                        interpretaris, sea consulatu scripserit ei, pri eu
                        accumsan contentiones. Usu ei iriure deleniti. Cum ut
                        periculis laboramus referrentur. Vis ea laoreet
                        imperdiet deterruisset, probo pertinax iudicabit qui ea.
                        Duo liber quodsi contentiones ne, vel ei movet
                        conceptam, id his quod iriure feugiat. Qui ex vide
                        labitur volumus. His cu fugit adolescens voluptatibus,
                        et per torquatos interesset. Iusto deleniti invenire id
                        cum. Te saepe alterum appellantur pro, in est errem
                        dicant suscipit, oblique argumentum sed et. Quo cu
                        partem inermis, vix et minimum vivendum, ut illud
                        delectus eos. Mei te alia justo, per amet nemore quodsi
                        cu, mel eruditi copiosae contentiones cu. Sed ea quando
                        mediocrem, soleat dolorum no ius. Eu electram iracundia
                        mnesarchum his, eum ut hinc latine. Est simul definiebas
                        ut, debitis invenire eu usu. Eu eros appetere mel,
                        nullam delenit tincidunt duo at. Eos option appetere
                        torquatos in, cu debitis singulis principes usu.
                    </p>
                </div>
            </div>
        );
    }
}
