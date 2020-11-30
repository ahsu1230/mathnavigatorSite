import { sortBy } from "lodash";

export const getAccountBalance = (transactions) => {
    return transactions.reduce((accum, curr) => accum + curr.amount, 0);
};

export const sortTransactionsLatestFirst = (transactions) => {
    return sortBy(transactions, "createdAt").reverse();
};
