package integration_tests

import (
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"time"
)

// Only implemented methods necessary for session integration tests
func createClass(programId string, semesterId string, classId string, locationId string, times string, startDate time.Time, endDate time.Time) domains.Class {
	return domains.Class{
		ProgramId:  programId,
		SemesterId: semesterId,
		ClassId:    classId,
		LocationId: locationId,
		Times:      times,
		StartDate:  startDate,
		EndDate:    endDate,
	}
}
