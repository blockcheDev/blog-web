package util

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func LogArchive() {
	job := cron.New()

	// 每天0点执行, 将昨天的日志归档
	job.AddFunc("8-10 3 * * *", func() {
		prev_day := time.Now().AddDate(0, 0, -1).Format("20060102")
		src_file_path := "logs/today.log"
		dst_file_path := fmt.Sprintf("logs/%s.log", prev_day)
		src_file, _ := os.OpenFile(src_file_path, os.O_RDWR, 0666)
		dst_file, err := os.OpenFile(dst_file_path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Error("Failed to open new log file:", err)
			return
		}
		defer func() {
			src_file.Close()
			dst_file.Close()
		}()

		// 归档日志
		_, err = io.Copy(dst_file, src_file)
		if err != nil {
			logrus.Error("Failed to copy log file:", err)
			return
		}

		// 将已归档的日志清空
		// 注意：这里可能会把 从copy完到truncate执行完 这段时间新写入的日志也清掉
		// 但这个时间很短，所以可能丢失的量很小吧
		// 目前还没有想到更好的解决方案
		err = src_file.Truncate(0)
		if err != nil {
			logrus.Error("Failed to truncate log file:", err)
			return
		}

		logrus.Infof("Log file archived from %s to %s", src_file_path, dst_file_path)
	})
	job.Start()

	logrus.Info("job next run time: ", job.Entries()[0].Next)
}
