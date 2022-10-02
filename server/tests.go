package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/EduardoZepeda/protobuffers-grpc/models"
	"github.com/EduardoZepeda/protobuffers-grpc/repository"
	"github.com/EduardoZepeda/protobuffers-grpc/studentpb"
	"github.com/EduardoZepeda/protobuffers-grpc/testpb"
)

type TestServer struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{repo: repo}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := &models.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}
	err := s.repo.SetTest(ctx, test)
	if err != nil {
		return nil, err
	}
	return &testpb.SetTestResponse{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		// Error caused by the client when the connection is closed
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}
		question := &models.Question{
			Id:       msg.GetId(),
			Answer:   msg.GetAnswer(),
			Question: msg.GetQuestion(),
			TestId:   msg.GetTestId(),
		}
		err = s.repo.SetQuestion(context.Background(), question)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) EnrollStudents(stream testpb.TestService_EnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()
		//
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}
		enrollment := &models.Enrollment{
			StudentId: msg.GetStudentId(),
			TestId:    msg.GetTestId(),
		}
		err = s.repo.SetEnrollment(context.Background(), enrollment)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) GetStudentsPerTest(req *testpb.GetStudentsPerTestRequest, stream testpb.TestService_GetStudentsPerTestServer) error {
	students, err := s.repo.GetStudentsPerTest(context.Background(), req.GetTestId())
	if err != nil {
		return err
	}
	for _, student := range students {
		student := &studentpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}
		err := stream.Send(student)
		time.Sleep(2 * time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *TestServer) TakeTest(stream testpb.TestService_TakeTestServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}
		// stream variable can receive and send, using Recv y Send methods, respectively
		// creating the bidirectional stream
		questions, err := s.repo.GetQuestionsPerTest(context.Background(), msg.GetTestId())
		if err != nil {
			log.Println(err)
			return err
		}
		i := 0
		correctAnswers := 0
		var currentQuestion = &models.Question{}
		var currentAnswers []*models.Answer
		currentAnswer := &models.Answer{
			QuestionId: msg.GetQuestionId(),
			StudentId:  msg.GetStudentId(),
			Answer:     msg.GetAnswer(),
			TestId:     msg.GetTestId(),
		}
		currentAnswers = append(currentAnswers, currentAnswer)
		for {
			fmt.Println(currentAnswers)
			if i < len(questions) {
				currentQuestion = questions[i]
			}
			if i <= len(questions) {
				questionToSend := &testpb.Question{
					Id:       currentQuestion.Id,
					Question: currentQuestion.Question,
				}
				err := stream.Send(questionToSend)
				if err != nil {
					log.Println(err)
					return err
				}
				i++
			}
			answer, err := stream.Recv()
			fmt.Println("Current answer:", currentAnswer.Answer, "Current question:", currentQuestion.Answer)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				log.Println(err)
				return err
			}
			currentAnswer := &models.Answer{
				QuestionId: answer.GetQuestionId(),
				StudentId:  answer.GetStudentId(),
				Answer:     answer.GetAnswer(),
				TestId:     answer.GetTestId(),
			}
			currentAnswers = append(currentAnswers, currentAnswer)
			if i >= len(questions)-1 {
				fmt.Println("Result:", correctAnswers, len(questions))
				score := 100 * correctAnswers / (len(questions) - 1)
				attempt := &models.Attempt{
					StudentId: msg.GetStudentId(),
					TestId:    msg.GetTestId(),
					Score:     score,
				}
				// Score test and save it in database
				lastInsertedId, err := s.repo.SetTestAttempt(context.Background(), attempt)
				if err != nil {
					log.Println(err)
					return err
				}
				for _, answer := range currentAnswers {
					answer.AttemptId = lastInsertedId
					err := s.repo.SetAnswer(context.Background(), answer)
					if err != nil {
						return err
					}
				}
				return nil
			}
		}
	}
}

func (s *TestServer) GetScore(ctx context.Context, req *testpb.GetScoreRequest) (*testpb.GetScoreResponse, error) {
	score, err := s.repo.GetScore(ctx, req.GetAttemptId())
	if err != nil {
		return nil, err
	}
	return &testpb.GetScoreResponse{
		Score: int32(score),
	}, nil
}
