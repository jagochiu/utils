package japi

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jagochiu/utils"
)

func NewRouter(name string) *gin.Engine {
	dir := `/var/log/app/`
	err := utils.CreateDir(dir)
	if err != nil {
		log.Fatalf(`[LOG] create directory fail - %v`, err)
	}
	f, err := os.Create(dir + name + ".log")
	if err != nil {
		log.Fatalf(`[LOG] create log file fail - %v`, err)
	}
	gin.DefaultWriter = io.MultiWriter(f)
	gin.ForceConsoleColor()
	r := gin.Default()
	r.Use(Gzip(BestCompression))
	return r
}
