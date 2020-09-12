package testUtils

import (
	"fmt"
	"strings"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func CreateMockAccount(id uint, primaryEmail string, password string) domains.Account {
	return domains.Account{
		Id:           id,
		PrimaryEmail: primaryEmail,
		Password:     password,
	}
}

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

func CreateMockClass(programId, semesterId, classKey, locationId, timesStr string, pricePerSession, priceLump uint) domains.Class {
	classId := programId + "_" + semesterId
	if classKey != "" {
		classId += "_" + classKey
	}

	return domains.Class{
		ProgramId:       programId,
		SemesterId:      semesterId,
		ClassKey:        domains.NewNullString(classKey),
		ClassId:         classId,
		LocationId:      locationId,
		TimesStr:        timesStr,
		PricePerSession: domains.NewNullUint(pricePerSession),
		PriceLumpSum:    domains.NewNullUint(priceLump),
	}
}

func CreateMockLocation(LocationId string, title string, street string, city string, state string, zipcode string, room string) domains.Location {
	return domains.Location{
		LocationId: LocationId,
		Title:      title,
		Street:     domains.NewNullString(street),
		City:       domains.NewNullString(city),
		State:      domains.NewNullString(state),
		Zipcode:    domains.NewNullString(zipcode),
		Room:       domains.NewNullString(room),
	}
}

func CreateMockProgram(programId string, title string, grade1 uint, grade2 uint, description string, featured string) domains.Program {
	return domains.Program{
		ProgramId:   programId,
		Title:       title,
		Grade1:      grade1,
		Grade2:      grade2,
		Description: description,
		Featured:    featured,
	}
}

func CreateMockSemester(season string, year uint) domains.Semester {
	return domains.Semester{
		SemesterId: fmt.Sprintf("%d_%s", year, season),
		Season:     season,
		Year:       year,
		Title:      strings.Title(fmt.Sprintf("%s %d", season, year)),
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

func CreateMockUser(id uint, firstName, lastName, middleName, email, phone string, isGuardian bool, accountId uint, notes string) domains.User {
	return domains.User{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: domains.NewNullString(middleName),
		Email:      email,
		Phone:      phone,
		IsGuardian: isGuardian,
		AccountId:  accountId,
		Notes:      domains.NewNullString(notes),
	}
}

func CreateMockUserClasses(id uint, userId uint, classId string, accountId uint, state uint) domains.UserClasses {
	return domains.UserClasses{
		Id:        id,
		UserId:    userId,
		ClassId:   classId,
		AccountId: accountId,
		State:     state,
	}
}

func CreateMockAFH(id uint, startsAt time.Time, endsAt time.Time, title string, subject string, locationId string, notes string) domains.AskForHelp {
	return domains.AskForHelp{
		Id:         id,
		StartsAt:   startsAt,
		EndsAt:     endsAt,
		Title:      title,
		Subject:    subject,
		LocationId: locationId,
		Notes:      domains.NewNullString(notes),
	}
}

func CreateMockTransaction(id uint, amount int, transactionType string, notes string, accountId uint) domains.Transaction {
	return domains.Transaction{
		Id:        id,
		Amount:    amount,
		Type:      transactionType,
		Notes:     domains.NewNullString(notes),
		AccountId: accountId,
	}
}

func CreateMockUserAfh(userId, afhId, accountId uint) domains.UserAfh {
	return domains.UserAfh{
		UserId:    userId,
		AfhId:     afhId,
		AccountId: accountId,
	}
}
