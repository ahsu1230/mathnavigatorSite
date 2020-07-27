package domains

const (
	// Ensures at least one alphanumeric character
	REGEX_ALPHA = `[[:alnum:]]`

	// Ensures at least one uppercase or lowercase letter
	REGEX_LETTER = `[A-Za-z]`

	// Finds characters that are not alphabetic, whitespace, a period, or a hyphen
	REGEX_SCHOOLS = `[^a-zA-Z\s.-]`

	// Ensures at least one number
	REGEX_NUMBER = `[0-9]`

	// Alphanumeric characters separated by underscores
	REGEX_GENERIC_ID = `^[[:alnum:]]+(_[[:alnum:]]+)*$`

	// In the form "year_season" i.e. 2020_fall
	REGEX_SEMESTER_ID = `^[1-9]\d{3,}_((spring)|(summer)|(fall)|(winter))$`

	/* Starts with a capital letter or number. Words consist of alphanumeric characters. Dashes, spaces, and underscores
	separate words. Words can also have parentheses around them. All number signs must be followed by numbers. */
	REGEX_TITLE = `^[A-Z0-9][[:alnum:]-]*([- _]([(]?#\d[)]?|&|([(]?[[:alnum:]]+[)]?)))*$`

	// There are no spaces and there must be characters before and after the @
	REGEX_EMAIL = `^[^ ]+@[^ ]+$`

	// At least 3 characters and there can be digits, spaces, pluses, periods, parentheses, slashes, and dashes
	REGEX_PHONE = `^[\d\s+.()/-]{3,}$`

	// Ensures formatting matches a valid street address
	REGEX_STREET = `^[1-9][0-9]*( [A-Z][a-z]*){2,}$`

	// One or more capitalized words separated by spaces or dashes
	REGEX_CITY = `^([A-Z][a-z]*)([ -][A-Z][a-z]*)*$`

	// All U.S. state and possession addresses used by the USPS
	REGEX_STATE = `^(A[A|E|K|L|P|R|S|Z]|C[A|O|T]|D[C|E]|F[L|M]|G[A|U]|HI|I[A|D|L|N]|K[S|Y]|LA|M[A|D|E|H|I|N|O|P|S|T]` +
		`|N[C|D|E|H|J|M|V|Y]|O[H|K|R]|P[A|R|W]|RI|S[C|D]|T[N|X]|UT|V[A|I|T]|W[A|I|V|Y])$`

	// Ensures formatting matches a valid ZIP code
	REGEX_ZIPCODE = `^[0-9]{5}(-[0-9]{4})?$`
)
