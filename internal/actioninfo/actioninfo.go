package actioninfo

import (
	"fmt"
	"log"
	"os"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, dataString := range dataset {
		err := dp.Parse(dataString)
		if err != nil {
			SetLog(fmt.Errorf("Ошибка парсинга: %w", err))
		}
		info, err := dp.ActionInfo()
		if err != nil {
			SetLog(fmt.Errorf("Ошибка вывода: %w", err))
		}
		fmt.Print(info)
	}
}

func SetLog(dataErr error) {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print(dataErr)
}
