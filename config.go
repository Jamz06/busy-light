package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const confName string = "busy_light"
const confType string = "json"

// Структуры
type ProfileStruct struct {
	Name        string
	HeadersList []string
	Blink       int
	Red         int
	Green       int
	Blue        int
}

type ConfigStruct struct {
	Timeout      int
	ComPort      string
	BaudRate     int
	Red          int
	Green        int
	Blue         int
	Blink        int
	ProfilesList []ProfileStruct `mapstructure:"profilesList"`
}

var C ConfigStruct

func readConf() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println(fmt.Errorf("Ошибка чтения конфигурации %s", err))
		}
	}

	// var profiles []profileStruct
	err := viper.UnmarshalExact(&C)
	// err := viper.Unmarshal(&C)
	if err != nil {
		fmt.Println("Не могу распарсить конфигурацию!", err)
	}
}

func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigName(confName)
	viper.SetConfigType(confType)

	readConf()

	// // Создать канал, для передачи сигнала в горутину
	// c := make(chan os.Signal, 1)
	// // Получить сигнал на reload от системы
	// signal.Notify(c, syscall.SIGHUP)

	// // Создатить горутину, которая будет следить за сигналом
	// go func() {
	// 	for sig := range c {
	// 		println(sig)
	// 		fmt.Println("Получен сигнал к перезагрузке конфигураци")
	// 		readConf()
	// 	}
	// }()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Конфигурация поменялась:", e.Name)
		readConf()
	})

}
