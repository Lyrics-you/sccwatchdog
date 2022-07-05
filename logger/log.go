package logger

import (
	sailor "github.com/Lyrics-you/sail-logrus-formatter/sailor"
	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&sailor.Formatter{
		HideKeys:        true,
		CharStampFormat: "yy-MM-dd HH:mm:ss.SSS",
		Position:        true,
		Colors:          false,
		FieldsColors:    true,
		ShowFullLevel:   true,
	})
	// logpath := `logs`
	// if _, err := os.Stat(logpath); os.IsNotExist(err) {
	// 	if err := os.MkdirAll(logpath, os.ModePerm); err != nil {
	// 		log.Errorf("There is a problem with the log path,err:%v", err)
	// 	}
	// }

	// logfile, err := os.OpenFile(`logs/watchdog.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Errorf("There is a problem with the log path,err:%v", err)
	// }
	// log.SetOutput(logfile)
	return log
}
