//go:generate goversioninfo -icon=bl.ico

package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
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
	com := C.ComPort
	var con io.ReadWriteCloser
	var flagConnnect bool = false
	var err error
	for flagConnnect == false {
		log.Println("Подключение к порту", com)
		con, err = connect(com, C.BaudRate)
		if err != nil {
			log.Println("Не удалось подключиться к порту", com)
			log.Println("Выберите другой порт, из этих:")
			serials := readPorts()
			var num int
			for i, name := range serials {
				fmt.Println(i, "-", name)
			}

			fmt.Scan(&num)
			com = serials[num]
			fmt.Println("Выбран порт", com)
			viper.Set("ComPort", com)
			viper.WriteConfig()
		} else {
			flagConnnect = true
		}

	}
	// Подключиться к com
	// con, err := connect(com, C.BaudRate)

	log.Println("Подключено успешно к порту:", com)

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
