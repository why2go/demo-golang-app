package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	PingCtrler *pingCtrler
)

type pingCtrler struct {
	logger zerolog.Logger
}

func newPingCtrler() *pingCtrler {
	ctrler := &pingCtrler{
		logger: log.With().Str("ltag", "pingCtrler").Logger(),
	}
	return ctrler
}

func init() {
	PingCtrler = newPingCtrler()
}

func (ctrler *pingCtrler) Ping(c *gin.Context) {
	logger := ctrler.logger.With().Str("request-id", c.GetString("request-id")).Logger()
	logger.Debug().Msg("pong would be sent")
	c.String(http.StatusOK, "%s", "pong")
}
