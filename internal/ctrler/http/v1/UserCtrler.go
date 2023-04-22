package v1

import (
	"demo/internal/svc"
	"demo/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	UserCtrler *userCtrler
)

type userCtrler struct {
	logger zerolog.Logger
}

func newUserCtrler() *userCtrler {
	ctrler := &userCtrler{
		logger: log.With().Str("ltag", "userCtrler").Logger(),
	}
	return ctrler
}

func (ctrler *userCtrler) GetDemoUser(c *gin.Context) {
	ctx := util.CtxWithRequestId(c)
	user, err := svc.UserService.RetrieveDemoUser(ctx)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errNo":  "not specified",
			"errMsg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
