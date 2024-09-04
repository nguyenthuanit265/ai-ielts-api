package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (apiServer *apiServer) initApiRoutes(app *gin.Engine) {
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiWriting := app.Group("/api/v1/IELTS_writing")
	{
		apiWriting.POST("", apiServer.writingHandler.Submit)
		apiWriting.GET("/next_question", apiServer.writingHandler.GetQuestion)
		apiWriting.POST("/set-ielts-task", apiServer.writingHandler.SetIeltsTask)
		apiWriting.POST("/get_question_type", apiServer.writingHandler.GetQuestionType)
	}

	apiSpeaking := app.Group("/api/v1/IELTS_speaking")
	{
		apiSpeaking.POST("", apiServer.speakingHandler.Submit)
		apiSpeaking.GET("/next_question", apiServer.speakingHandler.GetQuestion)
		apiSpeaking.POST("/set-ielts-part", apiServer.speakingHandler.SetIeltsPart)
		apiSpeaking.POST("/format", apiServer.speakingHandler.Format)
	}
}
