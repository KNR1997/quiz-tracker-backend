package quizzes

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/knr1997/quiz-tracker-backend/internal/adapters/postgresql/sqlc"
)

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListQuizzes(ctx context.Context) ([]repo.Quiz, error) {
	return s.repo.ListQuizzes(ctx)
}

func (s *svc) GetQuizByID(ctx context.Context, id int64) (repo.Quiz, error) {
	return s.repo.FindQuizByID(ctx, id)
}

func (s *svc) CreateQuiz(ctx context.Context, tempQuiz createQuizParams) (repo.Quiz, error) {
	return s.repo.CreateQuiz(ctx, repo.CreateQuizParams{
		CourseID: tempQuiz.CourseID,
		WeekNumber: pgtype.Int4{
			Int32: int32(tempQuiz.WeekNumber),
			Valid: true,
		},
		Status: tempQuiz.Status,
	})
}

func (s *svc) UpdateQuiz(ctx context.Context, id int64, tempQuiz updateQuizParams) (repo.Quiz, error) {
	return s.repo.UpdateQuiz(ctx, repo.UpdateQuizParams{
		ID: id,
		WeekNumber: pgtype.Int4{
			Int32: int32(tempQuiz.WeekNumber),
			Valid: true,
		},
		// DateTime: pgtype.Timestamptz{
		// 	Int32: int32(tempQuiz.date_time),
		// 	Valid: true,
		// },
		Status: tempQuiz.Status,
	})
}

func (s *svc) DeleteQuiz(ctx context.Context, id int64) error {
	return s.repo.DeleteQuiz(ctx, id)
}
