package repos

import (
	"database/sql"
)

func SetupRepos(db *sql.DB) {
	ProgramRepo.Initialize(db)
	ClassRepo.Initialize(db)
	LocationRepo.Initialize(db)
	AnnounceRepo.Initialize(db)
	AchieveRepo.Initialize(db)
	SemesterRepo.Initialize(db)
	SessionRepo.Initialize(db)
	UserRepo.Initialize(db)
	FamilyRepo.Initialize(db)
	AskForHelpRepo.Initialize(db)
	// AccountRepo.Initialize(db)
}