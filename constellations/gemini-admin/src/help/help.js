"use strict";
require("./help.sass");
import React from "react";
import { Link } from "react-router-dom";

export class HelpPage extends React.Component {
    render() {
        return (
            <div id="view-help">
                <h1>Administrative Support</h1>
                <section>
                    <h2>How to Create a Class</h2>
                    <p>
                        A <i>class</i> represents a course that students attend.
                        Every class is under a certain program (i.e. AP
                        Calculus) and is held over a certain semester (Fall
                        2020). In addition, every class is held at a particular
                        physical location or online.
                    </p>
                    <p>
                        To create a class, you must first create at least one
                        program, one semester and one location. Then, visit this
                        page to add a class and use the respective ids from the
                        program, semester, and location to create a class. There
                        are a few other fields that are required as well, such
                        as payment options. Create a class{" "}
                        <Link to={"classes"}>here</Link>.
                    </p>
                </section>

                <section>
                    <h2>How to add sessions to a class</h2>
                    <p>
                        A class <i>session</i> represents a meeting time for
                        students to congregate for a class. Because every
                        session is associated with a class, so to create a
                        session, you must first create a class and link the
                        classId to the session. Every session also has required
                        fields like a date, starting time and end time. If a
                        session is canceled, it's advised to NOT delete the
                        session and instead mark it as "Canceled" to notify
                        students and parents. Create class sessions{" "}
                        <Link to={"sessions"}>here</Link>.
                    </p>
                </section>

                <section>
                    <h2>Why are some class published and some unpublished?</h2>
                    <p>
                        Published classes are classes with information available
                        to the public. In other words, the class is available on
                        the user site and can be seen by students and parents.
                        View all classes (published or unpublished){" "}
                        <Link to={"classes"}>here</Link>.
                    </p>
                    <p>
                        Unpublished classes are classes that have NOT been
                        released to the public yet. This feature was created to
                        allow an administrator to create class information
                        before releasing it to the public. You can think of this
                        as "scheduling" a new class to be released. View only
                        unpublished classes on the home dashboard{" "}
                        <Link to={"/"}>here</Link>.
                    </p>
                </section>
            </div>
        );
    }
}
