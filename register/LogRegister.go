package register

import (
	"encoding/json"
	_ "github.com/beego/beego/v2/client/cache/redis"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"os"
	"path/filepath"
)

// RegisterLogger 注册日志
func LogRegister() {
	logs.Info("正在初始化文件日志配置.")
	logs.SetLogFuncCall(true)
	_ = logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)

	if web.AppConfig.DefaultBool("log_is_async", true) {
		logs.Async(1e3)
	}
	var log string
	logPath, err := filepath.Abs(web.AppConfig.DefaultString("log_path", ""))
	if err == nil {
		log = logPath
	}

	logPath = filepath.Join(log, web.AppConfig.DefaultString("appname", "log") + ".log")

	logs.Info("日志文件位置: " + logPath)

	if _, err := os.Stat(log); os.IsNotExist(err) {
		_ = os.MkdirAll(log, 0755)
	}

	config := make(map[string]interface{}, 1)

	config["filename"] = logPath
	config["perm"] = "0755"
	config["rotate"] = true

	if maxLines := web.AppConfig.DefaultInt("log_maxlines", 1000000); maxLines > 0 {
		config["maxLines"] = maxLines
	}
	if maxSize := web.AppConfig.DefaultInt("log_maxsize", 1<<28); maxSize > 0 {
		config["maxsize"] = maxSize
	}
	if !web.AppConfig.DefaultBool("log_daily", true) {
		config["daily"] = false
	}
	if maxDays := web.AppConfig.DefaultInt("log_maxdays", 7); maxDays > 0 {
		config["maxdays"] = maxDays
	}
	if level := web.AppConfig.DefaultString("log_level", "Trace"); level != "" {
		switch level {
		case "Emergency":
			config["level"] = logs.LevelEmergency
		case "Alert":
			config["level"] = logs.LevelAlert
		case "Critical":
			config["level"] = logs.LevelCritical
		case "Error":
			config["level"] = logs.LevelError
		case "Warning":
			config["level"] = logs.LevelWarning
		case "Notice":
			config["level"] = logs.LevelNotice
		case "Informational":
			config["level"] = logs.LevelInformational
		case "Debug":
			config["level"] = logs.LevelDebug
		}
	}
	b, err := json.Marshal(config)
	if err != nil {
		logs.Error("初始化文件日志时出错 ->", err)
		_ = logs.SetLogger("file", `{"filename":"`+logPath+`"}`)
	} else {
		_ = logs.SetLogger(logs.AdapterFile, string(b))
	}

	logs.SetLogFuncCall(true)
	logs.Info("文件日志初始化完成.")
}

//注册错误处理方法.
func ErrorRegister() {

}

//goland:noinspection SpellCheckingInspection
func Execute() {
	LogRegister()
	ErrorRegister()
}
