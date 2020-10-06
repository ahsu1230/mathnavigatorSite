package domains

const (
	SUBJECT_MATH        = "math"
	SUBJECT_ENGLISH     = "english"
	SUBJECT_PROGRAMMING = "programming"
)

var ALL_SUBJECTS = []string{SUBJECT_MATH, SUBJECT_ENGLISH, SUBJECT_PROGRAMMING}

func validateSubject(subject string) bool {
	return subject == SUBJECT_MATH || subject == SUBJECT_ENGLISH || subject == SUBJECT_PROGRAMMING
}
