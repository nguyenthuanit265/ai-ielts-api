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

type SpeakingService interface {
	GetQuestion(ctx *gin.Context) (models.SpeakingQuestionResponse, *models.AIIeltsError)
	FormatAnswer(ctx *gin.Context, req models.SpeakingFormatAnswerRequest) (models.SpeakingFormatAnswerResponse, *models.AIIeltsError)
	SetIeltsPart(ctx *gin.Context, req models.SpeakingSetIeltsPartRequest) (models.SpeakingSetIeltsPartResponse, *models.AIIeltsError)
	Submit(ctx *gin.Context, req models.SpeakingSubmitAnswerRequest) (models.SpeakingSubmitAnswerResponse, *models.AIIeltsError)
}

type speakingService struct {
}

func NewSpeakingService() SpeakingService {
	return &speakingService{}
}

func (s *speakingService) GetQuestion(ctx *gin.Context) (models.SpeakingQuestionResponse, *models.AIIeltsError) {
	var res models.SpeakingQuestionResponse
	return res, nil
}

func (s *speakingService) FormatAnswer(ctx *gin.Context, req models.SpeakingFormatAnswerRequest) (models.SpeakingFormatAnswerResponse, *models.AIIeltsError) {
	var res models.SpeakingFormatAnswerResponse
	return res, nil
}

func (s *speakingService) SetIeltsPart(ctx *gin.Context, req models.SpeakingSetIeltsPartRequest) (models.SpeakingSetIeltsPartResponse, *models.AIIeltsError) {
	var res models.SpeakingSetIeltsPartResponse
	return res, nil
}

func (s *speakingService) Submit(ctx *gin.Context, req models.SpeakingSubmitAnswerRequest) (models.SpeakingSubmitAnswerResponse, *models.AIIeltsError) {
	var res models.SpeakingSubmitAnswerResponse
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
			log.Errorf("SpeakingService - Submit. Error submit answer, error = %v", utils.LogFull(err))

			// Retry with another model
			bodyGetFeedback := s.buildBodyGetFeedback(req, utils.Gpt3Dot5Turbo)
			_, resp := utils.DoRequest(utils.MethodPost, utils.GptCompletions, header, bodyGetFeedback, utils.RequestBodyJson)
			log.Printf("%v", string(resp))
			err := json.Unmarshal(resp, &respChatCompletion)
			if err != nil {
				log.Errorf("SpeakingService - Submit. Error submit answer, error = %v", utils.LogFull(err))
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
			log.Errorf("SpeakingService - Submit. Error submit answer, error = %v", utils.LogFull(err))

			// Retry with another model
			bodyGetExampleAnswer := s.buildBodyGetExampleAnswer(req, utils.Gpt3Dot5Turbo)
			_, resp := utils.DoRequest(utils.MethodPost, utils.GptCompletions, header, bodyGetExampleAnswer, utils.RequestBodyJson)
			log.Printf("%v", string(resp))
			err := json.Unmarshal(resp, &respChatCompletion)
			if err != nil {
				log.Errorf("SpeakingService - Submit. Error submit answer, error = %v", utils.LogFull(err))
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

func (s *speakingService) buildBodyGetFeedback(req models.SpeakingSubmitAnswerRequest, modelChatGpt string) map[string]interface{} {
	var messages []models.ChatCompletionMessage
	body := make(map[string]interface{})

	messages = append(messages, models.ChatCompletionMessage{
		Role:    "user",
		Content: fmt.Sprintf("%s. %s", req.IeltsQuestion, utils.PROMT_SPEAKING_SUBMIT_1),
	})
	messages = append(messages, models.ChatCompletionMessage{
		Role:    "user",
		Content: req.UserAnswer,
	})
	var reqGpt models.ChatCompletionRequest
	reqGpt.Model = modelChatGpt
	reqGpt.Messages = messages
	body = utils.ConvertStructToMap(reqGpt)
	return body
}

func (s *speakingService) buildBodyGetExampleAnswer(req models.SpeakingSubmitAnswerRequest, modelChatGpt string) map[string]interface{} {
	var messages []models.ChatCompletionMessage
	body := make(map[string]interface{})

	messages = append(messages, models.ChatCompletionMessage{
		Role:    "user",
		Content: fmt.Sprintf("%s. %s", req.IeltsQuestion, utils.PROMT_SPEAKING_SUBMIT_2),
	})
	var reqGpt models.ChatCompletionRequest
	reqGpt.Model = modelChatGpt
	reqGpt.Messages = messages
	body = utils.ConvertStructToMap(reqGpt)
	return body
}
