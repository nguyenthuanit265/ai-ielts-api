package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"main/component/handlers"
	"main/component/services"
	"main/middleware"
	"main/utils"
	"os"
)

type apiServer struct {
	apiMiddleware   *middleware.ApiMiddleware
	writingHandler  handlers.WritingHandler
	speakingHandler handlers.SpeakingHandler
}

func initConfig() (*utils.Config, error) {
	config := &utils.Config{}
	//ex, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//configPath := fmt.Sprintf("%s/config", filepath.Dir(ex))
	configPath := fmt.Sprintf("%s/config", ".")
	utils.ShowInfoLogs(fmt.Sprintf("Config: %s", configPath))

	dir, err := ioutil.ReadDir(configPath)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range dir {
		file, err := os.Open(fmt.Sprintf("%s/%s", configPath, fileInfo.Name()))
		utils.ShowInfoLogs(fmt.Sprintf("%s/%s", configPath, fileInfo.Name()))
		if err != nil {
			log.Errorf("Error open configPath %v - fileInfo %v, error %v", configPath, fileInfo.Name(), err)
			return nil, err
		}
		defer file.Close()

		d := yaml.NewDecoder(file)

		if err := d.Decode(&config); err != nil {
			return nil, err
		}
	}

	return config, nil
}

func Initialize() *apiServer {
	var err error

	// config
	utils.AppConfig, err = initConfig()
	if err != nil {
		panic(err)
	}

	// init database
	//dbInstance := internal.NewPgDatabase()
	//db := dbInstance.ConnectPgDatabase(utils.AppConfig.Database)
	//dbWh := dbInstance.ConnectPgDatabaseWarehouse(utils.AppConfig.DatabaseWarehouse)

	// Create directory for exporting if not exist
	//utils.CreateDirectory(utils.RootExportFilePath)

	// externals

	// repos

	// services
	writingService := services.NewWritingService()
	speakingService := services.NewSpeakingService()

	// handlers
	writingHandler := handlers.NewWritingHandler(writingService)
	speakingHandler := handlers.NewSpeakingHandler(speakingService)

	// middleware

	return &apiServer{
		writingHandler:  writingHandler,
		speakingHandler: speakingHandler,
	}
}

func (apiServer *apiServer) Start() {
	appConfig := utils.AppConfig
	apiServer.setMode(appConfig.Server.Mode)

	app := gin.New()
	app.SetTrustedProxies(nil)
	app.Use(gin.Logger())
	configs := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}
	app.Use(cors.New(configs))
	// app.Use(...) //middleware

	apiServer.initApiRoutes(app)

	app.Run(fmt.Sprintf(":%d", appConfig.Server.Port))
}

func (apiServer *apiServer) setMode(mode string) {
	switch mode {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	}
}
