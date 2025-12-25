package courses

import (
	"context"

	repo "github.com/knr1997/quiz-tracker-backend/internal/adapters/postgresql/sqlc"
)

type createCourseParams struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type updateCourseParams struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Service interface {
	ListCourses(ctx context.Context) ([]repo.Course, error)
	GetCourseByID(ctx context.Context, id int64) (repo.Course, error)
	CreateCourse(ctx context.Context, tempCourse createCourseParams) (repo.Course, error)
	UpdateCourse(ctx context.Context, id int64, tempCourse updateCourseParams) (repo.Course, error)
	DeleteCourse(ctx context.Context, id int64) error
}
