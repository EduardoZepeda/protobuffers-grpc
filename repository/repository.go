package repository

import (
	"context"

	"github.com/EduardoZepeda/protobuffers-grpc/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) error
	SetQuestion(ctx context.Context, question *models.Question) error
	SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error
	GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error)
	GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error)
	SetTestAttempt(ctx context.Context, attempt *models.Attempt) (int, error)
	SetAnswer(ctx context.Context, answer *models.Answer) error
	GetScore(ctx context.Context, attemptId string) (int, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementation.GetTest(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) error {
	return implementation.SetTest(ctx, test)
}

func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}

func SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	return implementation.SetEnrollment(ctx, enrollment)
}

func GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	return implementation.GetStudentsPerTest(ctx, testId)
}

func GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	return implementation.GetQuestionsPerTest(ctx, testId)
}

func SetTestAttempt(ctx context.Context, attempt *models.Attempt) (int, error) {
	return implementation.SetTestAttempt(ctx, attempt)
}

func SetAnswer(ctx context.Context, answer *models.Answer) error {
	return implementation.SetAnswer(ctx, answer)
}

func GetScore(ctx context.Context, attemptId string) (int, error) {
	return implementation.GetScore(ctx, attemptId)
}
