package testUtils

import (
	"fmt"
	"strings"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

func CreateMockAccount(id uint, primary_email string, password string) domains.Account {
	return domains.Account{
		Id:           id,
		PrimaryEmail: primary_email,
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

func CreateMockClass(programId, semesterId, classKey, locationId, times string, pricePerSession, priceLump uint) domains.Class {
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
		Times:           times,
		PricePerSession: domains.NewNullUint(pricePerSession),
		PriceLump:       domains.NewNullUint(priceLump),
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

func CreateMockTransaction(id uint, amount int, paymentType string, paymentNotes string, accountId uint) domains.Transaction {
	return domains.Transaction{
		Id:           id,
		Amount:       amount,
		PaymentType:  paymentType,
		PaymentNotes: domains.NewNullString(paymentNotes),
		AccountId:    accountId,
	}
}

func CreateMockUserAfh(userId, afhId uint) domains.UserAfh {
	return domains.UserAfh{
		UserId: userId,
		AfhId:  afhId,
	}
}
