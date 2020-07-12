const STORAGE_CURRENT_CLASS = "storage_current_class";
const STORAGE_CURRENT_ACCOUNT = "storage_current_account";

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
    return JSON.parse(accountId);
};

export const setCurrentAccountId = (id) => {
    const accountId = JSON.stringify(id);
    window.localStorage.setItem(STORAGE_CURRENT_ACCOUNT, accountId);
};
