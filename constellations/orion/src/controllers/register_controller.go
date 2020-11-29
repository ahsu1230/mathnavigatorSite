package controllers

import (
	"context"
	"net/http"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/controllers/utils"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/gin-gonic/gin"
)

type RegisteredUsers struct {
	studentId  uint
	guardianId uint
	accountId  uint
}

func RegisterClass(c *gin.Context) {
	utils.LogControllerMethod(c, "registerController.RegisterClass")
	classId := c.Param("classId")

	var body domains.RegisterBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	studentJson := body.Student
	guardianJson := body.Guardian
	ctx := utils.RetrieveContext(c)
	registeredUsers, err := handleRegisterUsers(ctx, studentJson, guardianJson)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	userClass := domains.UserClass{
		UserId:    registeredUsers.studentId,
		AccountId: registeredUsers.accountId,
		ClassId:   classId,
		State:     domains.USER_CLASS_ENROLLED,
	}

	_, err = repos.UserClassRepo.Insert(ctx, userClass)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}

	c.Status(http.StatusNoContent)
}

func RegisterAfh(c *gin.Context) {
	utils.LogControllerMethod(c, "registerController.RegisterAfh")

	afhId, err := utils.ParseParamId(c, "afhId")
	if err != nil {
		c.Error(appErrors.WrapParse(err, c.Param("afhId")))
		c.Abort()
		return
	}

	var body domains.RegisterBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(appErrors.WrapBindJSON(err, c.Request))
		c.Abort()
		return
	}

	studentJson := body.Student
	guardianJson := body.Guardian
	ctx := utils.RetrieveContext(c)
	registeredUsers, err := handleRegisterUsers(ctx, studentJson, guardianJson)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	userAfh := domains.UserAfh{
		UserId:    registeredUsers.studentId,
		AccountId: registeredUsers.accountId,
		AfhId:     afhId,
	}

	_, err = repos.UserAfhRepo.Insert(ctx, userAfh)
	if err != nil {
		c.Error(appErrors.WrapRepo(err))
		c.Abort()
		return
	}

	c.Status(http.StatusNoContent)
}

func handleRegisterUsers(ctx context.Context, studentJson domains.User, guardianJson domains.User) (RegisteredUsers, error) {
	studentEmail := studentJson.Email
	guardianEmail := guardianJson.Email

	student, errStudent := repos.UserRepo.SelectByEmail(ctx, studentEmail)
	if errStudent != nil {
		return RegisteredUsers{}, errStudent
	}
	guardian, errGuardian := repos.UserRepo.SelectByEmail(ctx, guardianEmail)
	if errGuardian != nil {
		return RegisteredUsers{}, errGuardian
	}

	studentFound := errStudent == nil
	guardianFound := errGuardian == nil

	var studentId uint
	var guardianId uint
	var accountId uint

	if guardianFound {
		guardianId = guardian.Id
		accountId = guardian.AccountId
	}

	if studentFound {
		studentId = student.Id
		accountId = student.AccountId
	}

	if studentFound && guardianFound {
		if student.AccountId != guardian.AccountId {
			logger.Warn("Warning: Student and Guardian have different accountIds", logger.Fields{
				"studentAccountId":  student.AccountId,
				"guardianAccountId": guardian.AccountId,
			})
		}
		return RegisteredUsers{
			studentId:  studentId,
			guardianId: guardianId,
			accountId:  accountId,
		}, nil
	}

	// If both student & guardian emails are not found, create a new account
	if !studentFound && !guardianFound {
		// Create Account
		account := domains.Account{
			PrimaryEmail: guardian.Email,
			Password:     "automatic",
		}
		newAccountId, err := repos.AccountRepo.InsertWithUser(ctx, account, guardianJson)
		if err != nil {
			return RegisteredUsers{}, errGuardian
		}
		accountId = newAccountId
	}

	// If guardian user does not exist,
	// insert guardian as new user with the same accountId as the student's
	if !guardianFound {
		guardianJson.AccountId = accountId
		newGuardianId, err := repos.UserRepo.Insert(ctx, guardianJson)
		if err != nil {
			return RegisteredUsers{}, err
		}
		guardianId = newGuardianId
	}

	// If student user does not exist,
	// insert student as new user with the same accountId as the guardian's
	if !studentFound {
		studentJson.AccountId = accountId
		newStudentId, err := repos.UserRepo.Insert(ctx, studentJson)
		if err != nil {
			return RegisteredUsers{}, err
		}
		studentId = newStudentId
	}

	// Return final results of student & guardian
	result := RegisteredUsers{
		studentId:  studentId,
		guardianId: guardianId,
		accountId:  accountId,
	}
	return result, nil
}
