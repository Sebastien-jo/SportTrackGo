package route

import (
	"time"

	"github.com/Sebastien-jo/SportTrackGo/bootstrap"
	"github.com/Sebastien-jo/SportTrackGo/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
}
