const STORAGE_CURRENT_CLASS = "storage_current_class";
const STORAGE_CURRENT_ACCOUNT = "storage_current_account";
const STORAGE_CURRENT_USER_SEARCH = "storage_current_user_search";

export const getCurrentClassId = () => {
    const classId = window.localStorage.getItem(STORAGE_CURRENT_CLASS);
    return JSON.parse(classId);
};

export const setCurrentClassId = (id) => {
    const classId = JSON.stringify(id);
    window.localStorage.setItem(STORAGE_CURRENT_CLASS, classId);
};

export const getCurrentAccountId = () => {
    const accountId = window.localStorage.getItem(STORAGE_CURRENT_ACCOUNT);
    return parseInt(JSON.parse(accountId));
};

export const setCurrentAccountId = (id) => {
    const accountId = JSON.stringify(id);
    window.localStorage.setItem(STORAGE_CURRENT_ACCOUNT, accountId);
};

export const getCurrentUserSearch = () => {
    const userSearch = window.localStorage.getItem(STORAGE_CURRENT_USER_SEARCH);
    return JSON.parse(userSearch);
};

export const setCurrentUserSearch = (query) => {
    const userSearch = JSON.stringify(query);
    window.localStorage.setItem(STORAGE_CURRENT_USER_SEARCH, userSearch);
};
