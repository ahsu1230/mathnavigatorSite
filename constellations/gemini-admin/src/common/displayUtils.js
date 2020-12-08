import moment from "moment";

export const formatCurrency = (amount) => {
    return new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "USD",
    }).format(amount);
};

export const getAfhTitle = (afh) => {
    afh = afh || {};
    const time =
        moment(afh.startsAt).format("MM/DD/yy hh:mm") +
        "-" +
        moment(afh.endsAt).format("hh:mm a");
    return afh.id + " " + afh.title + " (" + afh.subject + ") " + time;
};
