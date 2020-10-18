"use strict";
import moment from "moment";

export const generateEmailMessageForClass = (
    classId,
    studentInfo,
    guardianInfo
) => {
    return [
        "<html>",
        "<body>",
        "<h1>New Class Enrollment</h>",
        "<h2>ClassId: " + classId + "</h2>",
        "<h2>Student: " +
            studentInfo.firstName +
            "	&nbsp; " +
            studentInfo.lastName +
            "</h2>",
        "<h3>Email: " + studentInfo.email + "</h3>",
        "<h3>Grade: " + studentInfo.grade + "</h3>",
        "<h3>School: " + studentInfo.school + "</h3>",
        "<h3>GraduationYear: " + studentInfo.graduationYear + "</h3>",
        "<br/>",
        "<h2>Guardian: " +
            guardianInfo.firstName +
            "	&nbsp; " +
            guardianInfo.lastName +
            "</h2>",
        "<h3>Phone: " + guardianInfo.phone + "</h3>",
        "<h3>Email: " + guardianInfo.email + "</h3>",
        "<br/>",
        "<p>Additional Info: " + guardianInfo.additionalText + "</p>",
        "</body>",
        "</html>",
    ].join("\n");
};

export const generateEmailMessageForAfh = (afhId, afh, studentInfo) => {
    return [
        "<html>",
        "<body>",
        "<h1>New AFH Registration</h>",
        "<h2>AskForHelp Id: " + afhId + "</h2>",
        "<h3>AskForHelp session: " + afh.title + "</h3>",
        "<h3>DateTime: " + moment(afh.startsAt).format("llll") + "</h3>",
        "<br/>",
        "<h2>Student: " +
            studentInfo.firstName +
            "	&nbsp; " +
            studentInfo.lastName +
            "</h2>",
        "<h3>Email: " + studentInfo.email + "</h3>",
        "<h3>Grade: " + studentInfo.grade + "</h3>",
        "<h3>School: " + studentInfo.school + "</h3>",
        "<h3>GraduationYear: " + studentInfo.graduationYear + "</h3>",
        "<br/>",
        "</body>",
        "</html>",
    ].join("\n");
};

export const sendEmail = (message, onSuccess, onFail) => {
    if (process.env.NODE_ENV === "development") {
        sendTestEmail(true, onSuccess, onFail);
    } else {
        sendProdEmail(message, onSuccess, onFail);
    }
};

const sendProdEmail = (message, onSuccess, onFail) => {
    console.log("Sending prod email...");
    const templateId = "mathnavigatorwebsitecontact";
    const receiverEmail = "andymathnavigator@gmail.com";
    const senderEmail = "anonymous@andymathnavigator.com";
    var templateParams = {
        emailMessage: message,
    };

    window.emailjs
        .send("mailgun", templateId, templateParams)
        .then((res) => {
            console.log("Email successfully sent!");
            if (onSuccess) {
                onSuccess();
            }
        })
        .catch((err) => {
            console.error("Failed to send email. Error: ", err);
            if (onFail) {
                onFail();
            }
        });
};

const sendTestEmail = (success, onSuccess, onFail) => {
    console.log("Sending test email...");
    if (success && onSuccess) {
        onSuccess();
    } else if (onFail) {
        onFail();
    }
};
