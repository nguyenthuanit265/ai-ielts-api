package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/component/models"
	"main/utils"
	"sync"
)

type WritingService interface {
	GetQuestion(ctx *gin.Context) (models.WritingQuestionResponse, *models.AIIeltsError)
	GetQuestionType(ctx *gin.Context) (models.WritingQuestionTypeResponse, *models.AIIeltsError)
	SetIeltsTask(ctx *gin.Context, req models.WritingSetIeltsTaskRequest) (models.WritingSetIeltsTaskResponse, *models.AIIeltsError)
	Submit(ctx *gin.Context, req models.WritingSubmitRequest) (models.WritingSubmitResponse, *models.AIIeltsError)
}
type writingService struct {
}

func NewWritingService() WritingService {
	return &writingService{}
}

func (s *writingService) GetQuestion(ctx *gin.Context) (models.WritingQuestionResponse, *models.AIIeltsError) {
	var res models.WritingQuestionResponse
	return res, nil
}

func (s *writingService) GetQuestionType(ctx *gin.Context) (models.WritingQuestionTypeResponse, *models.AIIeltsError) {
	var res models.WritingQuestionTypeResponse
	return res, nil
}

func (s *writingService) SetIeltsTask(ctx *gin.Context, req models.WritingSetIeltsTaskRequest) (models.WritingSetIeltsTaskResponse, *models.AIIeltsError) {
	var res models.WritingSetIeltsTaskResponse
	return res, nil
}

func (s *writingService) Submit(ctx *gin.Context, req models.WritingSubmitRequest) (models.WritingSubmitResponse, *models.AIIeltsError) {
	var res models.WritingSubmitResponse
	header := make(map[string]string)
	appConfig := utils.AppConfig

	// Build header
	header["Content-Type"] = "application/json"
	header["Authorization"] = fmt.Sprintf("Bearer %s", appConfig.AI.ChatGpt.ApiKey)

	// Build body
	bodyGetFeedback := s.buildBodyGetFeedback(req, utils.Gpt4TurboPreview)
	bodyGetExampleAnswer := s.buildBodyGetExampleAnswer(req, utils.Gpt4TurboPreview)

	// Execute
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		var respChatCompletion models.ChatCompletionResponse
		_, resp := utils.DoRequest(utils.MethodPost, utils.GptCompletions, header, bodyGetFeedback, utils.RequestBodyJson)
		log.Printf("%v", string(resp))
		err := json.Unmarshal(resp, &respChatCompletion)
		if err != nil {
			log.Errorf("WritingService - Submit. Error submit answer, error = %v", utils.LogFull(err))

			// Retry with another model
			bodyGetFeedback := s.buildBodyGetFeedback(req, utils.Gpt3Dot5Turbo)
			_, resp := utils.DoRequest(utils.MethodPost, utils.GptCompletions, header, bodyGetFeedback, utils.RequestBodyJson)
			log.Printf("%v", string(resp))
			err := json.Unmarshal(resp, &respChatCompletion)
			if err != nil {
				log.Errorf("WritingService - Submit. Error submit answer, error = %v", utils.LogFull(err))
			}
		}

		// Mapping
		if len(respChatCompletion.Choices) > 0 {
			res.Feedback = respChatCompletion.Choices[0].Message.Content
		}
	}()

	go func() {
		defer wg.Done()

		var respChatCompletion models.ChatCompletionResponse
		_, resp := utils.DoRequest(utils.MethodPost, utils.GptCompletions, header, bodyGetExampleAnswer, utils.RequestBodyJson)
		log.Printf("%v", string(resp))
		err := json.Unmarshal(resp, &respChatCompletion)
		if err != nil {
			log.Errorf("WritingService - Submit. Error submit answer, error = %v", utils.LogFull(err))

			// Retry with another model
			bodyGetExampleAnswer := s.buildBodyGetExampleAnswer(req, utils.Gpt3Dot5Turbo)
			_, resp := utils.DoRequest(utils.MethodPost, utils.GptCompletions, header, bodyGetExampleAnswer, utils.RequestBodyJson)
			log.Printf("%v", string(resp))
			err := json.Unmarshal(resp, &respChatCompletion)
			if err != nil {
				log.Errorf("WritingService - Submit. Error submit answer, error = %v", utils.LogFull(err))
			}
		}

		// Mapping
		if len(respChatCompletion.Choices) > 0 {
			res.ExampleAnswer = respChatCompletion.Choices[0].Message.Content
		}
	}()

	wg.Wait()

	res.UserAnswer = req.UserAnswer
	return res, nil
}

func (s *writingService) buildBodyGetFeedback(req models.WritingSubmitRequest, modelChatGpt string) map[string]interface{} {
	var messages []models.ChatCompletionMessage
	body := make(map[string]interface{})

	messages = append(messages, models.ChatCompletionMessage{
		Role:    "user",
		Content: fmt.Sprintf("%s", req.IeltsQuestion),
	})
	messages = append(messages, models.ChatCompletionMessage{
		Role:    "user",
		Content: fmt.Sprintf(utils.PROMT_WRITING_SUBMIT_3, req.UserAnswer),
	})
	var reqGpt models.ChatCompletionRequest
	reqGpt.Model = modelChatGpt
	reqGpt.Messages = messages
	body = utils.ConvertStructToMap(reqGpt)
	return body
}

func (s *writingService) buildBodyGetExampleAnswer(req models.WritingSubmitRequest, modelChatGpt string) map[string]interface{} {
	var messages []models.ChatCompletionMessage
	body := make(map[string]interface{})

	messages = append(messages, models.ChatCompletionMessage{
		Role:    "user",
		Content: fmt.Sprintf("%s. %s", req.IeltsQuestion, utils.PROMT_WRITING_SUBMIT_2),
	})
	var reqGpt models.ChatCompletionRequest
	reqGpt.Model = modelChatGpt
	reqGpt.Messages = messages
	body = utils.ConvertStructToMap(reqGpt)
	return body
}
