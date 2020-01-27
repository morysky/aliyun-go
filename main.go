package main

import (
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	router := gin.Default()
	router.GET("/aliyun/ecs/start", StartECSInstance)
	router.GET("/aliyun/ecs/stop", StopECSInstance)

	router.Run(":10089")
}

// StdResult ...
type StdResult struct {
	Errno      int         `json:"err_no"`
	ErrMessage string      `json:"err_message"`
	Data       interface{} `json:"data"`
}

// StartECSInstance ...
func StartECSInstance(c *gin.Context) {
	result := StdResult{}

	client, err := ecs.NewClientWithAccessKey(os.Getenv("Region"), os.Getenv("AccessKey"), os.Getenv("AccessSecret"))

	request := ecs.CreateStartInstanceRequest()
	request.Scheme = "https"

	request.InstanceId = os.Getenv("InstanceID")

	response, err := client.StartInstance(request)

	if err == nil {
		result.Data = response
	} else {
		result.ErrMessage = err.Error()
		result.Errno = 1
	}

	c.JSON(200, result)
}

// StopECSInstance ...
func StopECSInstance(c *gin.Context) {
	result := StdResult{}

	client, err := ecs.NewClientWithAccessKey(os.Getenv("Region"), os.Getenv("AccessKey"), os.Getenv("AccessSecret"))

	request := ecs.CreateStopInstanceRequest()
	request.ConfirmStop = requests.NewBoolean(false)
	request.ForceStop = requests.NewBoolean(false)
	request.StoppedMode = "StopCharging"
	request.Scheme = "https"

	request.InstanceId = os.Getenv("InstanceID")

	response, err := client.StopInstance(request)

	if err == nil {
		result.Data = response
	} else {
		result.ErrMessage = err.Error()
		result.Errno = 1
	}

	c.JSON(200, result)
}
