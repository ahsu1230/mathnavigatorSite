const STORAGE_CURRENT_CLASS = "storage_current_class";

export const getCurrentClassId = () => {
    const classId = window.localStorage.getItem(STORAGE_CURRENT_CLASS);
    return JSON.parse(classId);
};

export const setCurrentClassId = (c) => {
    const classId = JSON.stringify(c);
    window.localStorage.setItem(STORAGE_CURRENT_CLASS, classId);
};
