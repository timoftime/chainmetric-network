package main

import (
	"github.com/gin-gonic/gin"
	"github.com/timoth-y/chainmetric-contracts/shared"
	"github.com/timoth-y/chainmetric-contracts/src/users/api"
	"github.com/timoth-y/chainmetric-contracts/src/users/usecase/identity"
)

func init() {
	shared.InitCore()
	shared.InitMongoDB()
	shared.MustExecute(identity.Init, "failed to initialize identity package")
}

func main() {
	engine := gin.Default()
	api.Setup(engine)

	_ = engine.Run(":8080")
}