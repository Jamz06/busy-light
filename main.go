package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func work() ProfileStruct {
	var prof ProfileStruct
	checkFlag := false
	// Получить заголовок активного окна
	header := check()
	fmt.Println(header)
	// Сравнить с ключевыми словами профиля
	for _, prof = range C.ProfilesList {
		// обойти все ключевые слова в профиле
		for _, subHeader := range prof.HeadersList {
			// сравнить ключевое слово, с заголовком
			if strings.Contains(header, subHeader) {
				checkFlag = true
				return prof
			}
		}
	}
	if checkFlag == false {
		prof.Red = C.Red
		prof.Green = C.Green
		prof.Blue = C.Blue
		prof.Blink = C.Blink
	}
	return prof
}

func main() {
	// инициализация, чтение конфига
	Init()
	// Подключиться к com
	con, err := connect(C.ComPort, C.BaudRate)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Подключено успешно к порту:", C.ComPort)

	for {
		log.Println("Определение активного окна")
		prof := work()

		// Дать команду на включение цвета
		log.Println(prof.Name, prof.Red, prof.Green, prof.Blue, prof.Blink)
		log.Println("Даю команду в МК")

		writeToCom(prof.Red, prof.Green, prof.Blue, prof.Blink, con)
		//
		log.Println("Ожидаю", C.Timeout, "секунд.")
		time.Sleep(time.Duration(C.Timeout * int(time.Second)))

	}

}
