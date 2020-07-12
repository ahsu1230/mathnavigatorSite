const STORAGE_CURRENT_CLASS = "storage_current_class";
const STORAGE_CURRENT_ACCOUNT = "storage_current_account";

export const getCurrentClassId = () => {
    const classId = window.localStorage.getItem(STORAGE_CURRENT_CLASS);
    return JSON.parse(classId);
};

export const setCurrentClassId = (c) => {
    const classId = JSON.stringify(c);
    window.localStorage.setItem(STORAGE_CURRENT_CLASS, classId);
};

export const getCurrentAccountId = () => {
    const accountId = window.localStorage.getItem(STORAGE_CURRENT_ACCOUNT);
    return JSON.parse(accountId);
};

export const setCurrentAccountId = () => {
    const accountId = JSON.stringify(c);
    window.localStorage.setItem(STORAGE_CURRENT_ACCOUNT);
};
