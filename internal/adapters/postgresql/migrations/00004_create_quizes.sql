-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quizzes (
  id BIGSERIAL PRIMARY KEY,
  course_id BIGINT NOT NULL,
  week_number INT,
  date_time TIMESTAMPTZ,
  status VARCHAR(10) NOT NULL
    CHECK (status IN ('pending', 'postponed', 'completed')),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT fk_course FOREIGN KEY (course_id) REFERENCES courses(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quizes;
-- +goose StatementEnd
