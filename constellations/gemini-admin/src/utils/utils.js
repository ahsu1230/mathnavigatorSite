export const generatePassword = () => {
    var length = 8,
        charset =
            "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
        retVal = "";
    for (var i = 0, n = charset.length; i < length; ++i) {
        retVal += charset.charAt(Math.floor(Math.random() * n));
    }
    return retVal;
};

export const getFullName = (user) => {
    var fullName = user.firstName + " ";
    fullName += user.middleName
        ? user.middleName + " " + user.lastName
        : user.lastName;

    return fullName;
};
