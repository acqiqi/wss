package app

import (
	"github.com/astaxie/beego/validation"
	"log"
	"wss/lib/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}

func ErrorsGetOne(errors []*validation.Error) error {
	for _, err := range errors {
		log.Println(err.Key)
		logging.Info(err.Key, err.Message)
		return err
	}
	return nil
}
