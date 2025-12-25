package courses

import (
	"context"

	repo "github.com/knr1997/quiz-tracker-backend/internal/adapters/postgresql/sqlc"
)

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListCourses(ctx context.Context) ([]repo.Course, error) {
	return s.repo.ListCourses(ctx)
}

func (s *svc) GetCourseByID(ctx context.Context, id int64) (repo.Course, error) {
	return s.repo.FindCourseByID(ctx, id)
}

func (s *svc) CreateCourse(ctx context.Context, tempCourse createCourseParams) (repo.Course, error) {
	return s.repo.CreateCourse(ctx, repo.CreateCourseParams{
		Name: tempCourse.Name,
		Code: tempCourse.Code,
	})
}

func (s *svc) UpdateCourse(ctx context.Context, id int64, tempCourse updateCourseParams) (repo.Course, error) {
	return s.repo.UpdateCourse(ctx, repo.UpdateCourseParams{
		ID:   id,
		Name: tempCourse.Name,
		Code: tempCourse.Code,
	})
}

func (s *svc) DeleteCourse(ctx context.Context, id int64) error {
	return s.repo.DeleteCourse(ctx, id)
}
