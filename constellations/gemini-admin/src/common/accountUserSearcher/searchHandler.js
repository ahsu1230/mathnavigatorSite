"use strict";

import API from "../../api.js";

const API_SEARCH_ACCOUNT_BY_ID = "api/accounts/account/";
const API_SEARCH_ACCOUNT_BY_EMAIL = "api/accounts/search";
const API_SEARCH_USER_BY_ID = "api/users/user/";
const API_SEARCH_USER_BY_EMAIL = "api/users/search";
const API_SEARCH_USER_BY_ACCOUNT_ID = "api/users/account/";

// Given an account's primary email, search for the account
function searchAccountByEmail(email, onFoundAccount, onFoundUsers, onError) {
    API.post(API_SEARCH_ACCOUNT_BY_EMAIL, { primaryEmail: email })
        .then((res) => {
            const account = res.data;
            onFoundAccount(account);
            handleUsersForFoundAccount(account, onFoundUsers, onError);
        })
        .catch((err) => {
            console.log("Error searching account " + err);
            onError();
        });
}

// Given an accountId, search for the account
function searchAccountById(accountId, onFoundAccount, onFoundUsers, onError) {
    API.get(API_SEARCH_ACCOUNT_BY_ID + accountId)
        .then((res) => {
            const account = res.data;
            onFoundAccount(account);
            handleUsersForFoundAccount(account, onFoundUsers, onError);
        })
        .catch((err) => {
            console.log("Error searching account " + err);
            onError();
        });
}

// Helper method for handling users after finding an account
function handleUsersForFoundAccount(account, onFoundUsers, onError) {
    getUsersForAccount(account)
        .then((res) => {
            const users = res.data;
            onFoundUsers(users);
        })
        .catch((err) => {
            console.log("Error searching users " + err);
            onError();
        });
}

// Given a user's email, search for the user
function searchUserByEmail(email, onFoundUser, onFoundAccount, onError) {
    API.post(API_SEARCH_USER_BY_EMAIL, { query: email })
        .then((res) => {
            const users = res.data;
            if (users.length == 1) {
                let user = users[0];
                onFoundUser(user);
                handleAccountForFoundUser(user, onFoundAccount, onError);
            } else {
                console.log("Too many results!");
                onError();
            }
        })
        .catch((err) => {
            console.log("Error searching user " + err);
            onError();
        });
}

// Given a userId, search for the user
function searchUserById(userId, onFoundUser, onFoundAccount, onError) {
    API.get(API_SEARCH_USER_BY_ID + userId)
        .then((res) => {
            const user = res.data;
            onFoundUser(user);
            handleAccountForFoundUser(user, onFoundAccount, onError);
        })
        .catch((err) => {
            console.log("Error searching user " + err);
            onError();
        });
}

// Helper method for handling users after finding an account
function handleAccountForFoundUser(user, onFoundAccount, onError) {
    getAccountForUser(user)
        .then((res) => {
            const account = res.data;
            onFoundAccount(account);
        })
        .catch((err) => {
            console.log("Error searching account " + err);
            onError();
        });
}

// Helper method for retrieving account information for a single user
function getAccountForUser(user) {
    const accountId = user.accountId;
    return API.get(API_SEARCH_ACCOUNT_BY_ID + accountId);
}

// Helper method for retrieving all users for an account
function getUsersForAccount(account) {
    const accountId = account.id;
    return API.get(API_SEARCH_USER_BY_ACCOUNT_ID + accountId);
}

export default {
    searchAccountByEmail: searchAccountByEmail,
    searchAccountById: searchAccountById,
    searchUserByEmail: searchUserByEmail,
    searchUserById: searchUserById,
};
