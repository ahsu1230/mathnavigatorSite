package domains

// Alphanumeric characters separated by underscores
const REGEX_PROGRAM_ID = `^[[:alnum:]]+(_[[:alnum:]]+)*$`

// In the form "year_season" i.e. 2020_fall
const REGEX_SEMESTER_ID = `^[1-9]\d{3,}_((spring)|(summer)|(fall)|(winter))$`

/* Starts with a capital letter or number. Words consist of alphanumeric characters and dashes, spaces, and underscores
separate words. Words can have parentheses around them and number signs must be followed by numbers. */
const REGEX_NAME = `^[A-Z0-9][[:alnum:]-]*([- _]([(]?#\d[)]?|&|([(]?[[:alnum:]]+[)]?)))*$`

// Ensures at least one uppercase or lowercase letter
const REGEX_LETTER = `[A-Za-z]+`

/* Starts with a capital letter or number. Words consist of alphanumeric characters and dashes, spaces, and underscores
separate words. Words can have parentheses around them and number signs must be followed by numbers. */
const REGEX_TITLE = `^[A-Z0-9][[:alnum:]]*([- _]([(]?#\d[)]?|&|([(]?[[:alnum:]]+[)]?)))*$`

// There are no spaces and there must be characters before and after the @
const REGEX_EMAIL = `^[^ ]+@[^ ]+$`

// US phone numbers
const REGEX_PHONE = `^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]\d{3}[\s.-]\d{4}$`
