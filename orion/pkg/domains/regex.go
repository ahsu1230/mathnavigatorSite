package domains

const (
	// Alphanumeric characters separated by underscores
	REGEX_PROGRAM_ID = `^[[:alnum:]]+(_[[:alnum:]]+)*$`

	// In the form "year_season" i.e. 2020_fall
	REGEX_SEMESTER_ID = `^[1-9]\d{3,}_((spring)|(summer)|(fall)|(winter))$`

	/* Starts with a capital letter or number. Words consist of alphanumeric characters and dashes, spaces, and underscores
	separate words. Words can have parentheses around them and number signs must be followed by numbers. */
	REGEX_TITLE = `^[A-Z0-9][[:alnum:]-]*([- _]([(]?#\d[)]?|&|([(]?[[:alnum:]]+[)]?)))*$`

	// Ensures at least one uppercase or lowercase letter
	REGEX_LETTER = `[A-Za-z]+`

	// There are no spaces and there must be characters before and after the @
	REGEX_EMAIL = `^[^ ]+@[^ ]+$`

	// US phone numbers
	REGEX_PHONE = `^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]\d{3}[\s.-]\d{4}$`
)
