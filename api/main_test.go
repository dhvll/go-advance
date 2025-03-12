package api

import (
	"os"
	"testing"
	"time"

	db "github.com/dhvll/go-advance/db/sqlc"
	"github.com/dhvll/go-advance/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmertricKey:  util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server := NewServer(config, store)
	return server
}
func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
