package testUtils

import (
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func CreateMockAchievement(id uint, year uint, message string) domains.Achieve {
	return domains.Achieve{
		Id:      id,
		Year:    year,
		Message: message,
	}
}

func CreateMockAnnounce(id uint, postedAt time.Time, author string, message string, onHomePage bool) domains.Announce {
	return domains.Announce{
		Id:         id,
		PostedAt:   postedAt,
		Author:     author,
		Message:    message,
		OnHomePage: onHomePage,
	}
}

func CreateMockClass(programId, semesterId, classKey, locationId, times string, startDate, endDate time.Time) domains.Class {
	classId := programId + "_" + semesterId
	if classKey != "" {
		classId += "_" + classKey
	}

	return domains.Class{
		ProgramId:  programId,
		SemesterId: semesterId,
		ClassKey:   domains.NewNullString(classKey),
		ClassId:    classId,
		LocationId: locationId,
		Times:      times,
		StartDate:  startDate,
		EndDate:    endDate,
	}
}

func CreateMockLocation(LocationId string, street string, city string, state string, zipcode string, room string) domains.Location {
	return domains.Location{
		LocationId: LocationId,
		Street:     street,
		City:       city,
		State:      state,
		Zipcode:    zipcode,
		Room:       domains.NewNullString(room),
	}
}

func CreateMockProgram(programId string, name string, grade1 uint, grade2 uint, description string, featured uint) domains.Program {
	return domains.Program{
		ProgramId:   programId,
		Name:        name,
		Grade1:      grade1,
		Grade2:      grade2,
		Description: description,
		Featured:    featured,
	}
}

func CreateMockSemester(semesterId string, title string) domains.Semester {
	return domains.Semester{
		SemesterId: semesterId,
		Title:      title,
	}
}

func CreateMockSession(id uint, classId string, startsAt time.Time, endsAt time.Time, canceled bool, notes string) domains.Session {
	return domains.Session{
		Id:       id,
		ClassId:  classId,
		StartsAt: startsAt,
		EndsAt:   endsAt,
		Canceled: canceled,
		Notes:    domains.NewNullString(notes),
	}
}

func CreateMockUser(id uint, firstName, lastName, middleName, email, phone string, isGuardian bool, familyId uint, notes string) domains.User {
	return domains.User{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: domains.NewNullString(middleName),
		Email:      email,
		Phone:      phone,
		IsGuardian: isGuardian,
		FamilyId:   familyId,
		Notes:      domains.NewNullString(notes),
	}
}

func CreateMockAFH(id uint, title string, date time.Time, timeString string, subject string, locationId string, notes string) domains.AskForHelp {
	return domains.AskForHelp{
		Id:         id,
		Title:      title,
		Date:       date,
		TimeString: timeString,
		Subject:    subject,
		LocationId: locationId,
		Notes:      domains.NewNullString(notes),
	}
}

func CreateMockUserAfh(userId, afhId uint) domains.UserAfh {
	return domains.UserAfh{
		UserId: userId,
		AfhId:  afhId,
	}
}
