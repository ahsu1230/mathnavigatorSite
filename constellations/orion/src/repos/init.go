package repos

import (
	"context"
	"database/sql"
)

func SetupRepos(ctx context.Context, db *sql.DB) {
	ProgramRepo.Initialize(ctx, db)
	ClassRepo.Initialize(ctx, db)
	LocationRepo.Initialize(ctx, db)
	AnnounceRepo.Initialize(ctx, db)
	AchieveRepo.Initialize(ctx, db)
	SemesterRepo.Initialize(ctx, db)
	SessionRepo.Initialize(ctx, db)
	UserRepo.Initialize(ctx, db)
	UserClassRepo.Initialize(ctx, db)
	AccountRepo.Initialize(ctx, db)
	AskForHelpRepo.Initialize(ctx, db)
	UserAfhRepo.Initialize(ctx, db)
	TransactionRepo.Initialize(ctx, db)
}
