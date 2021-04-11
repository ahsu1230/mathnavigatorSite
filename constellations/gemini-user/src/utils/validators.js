import { reduce } from "lodash";

const regexEmail = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

export const validateEmail = (email) => {
    return regexEmail.test(String(email).toLowerCase());
};

export const validatePhone = (phone) => {
    return (
        reduce(
            phone.split(""),
            (sum, c) => {
                return sum + (parseInt(c) >= 0 ? 1 : 0);
            },
            0
        ) >= 10
    );
};

export const validateGrade = (grade) => {
    return grade >= 1 && grade <= 12;
};

export const validateGradYear = (gradYear) => {
    return gradYear >= 2018 && gradYear <= 2050;
};
