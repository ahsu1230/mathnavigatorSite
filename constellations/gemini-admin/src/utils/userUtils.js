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

export const validateEmail = (email) => {
    return /^[^( @)]+@[^( @)]+\.[^( @)]+$/.test(email);
}

export const validatePhoneNumber = (phone) => {
    return /^[\d\s+.()/-]{3,}$/.test(phone)
}
