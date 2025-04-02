package util

import (
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func LogArchive() {
	job := cron.New()

	// 每天0点执行, 将昨天的日志归档
	job.AddFunc("0 0 * * *", func() {
		prev_day := time.Now().AddDate(0, 0, -1).Format("20060102")
		os.Rename("logs/today.log", fmt.Sprintf("logs/%s.log", prev_day))
	})
	job.Start()
}
