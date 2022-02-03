package main

import (
	"io"
	"log"
	"time"

	"github.com/tarm/serial"
	"golang.org/x/sys/windows/registry"
)

func connect(port string, baud int) (io.ReadWriteCloser, error) {
	config := &serial.Config{Name: port, Baud: baud}
	connection, err := serial.OpenPort(config)
	time.Sleep(3 * time.Second)
	return connection, err
}

func writeToCom(red, green, blue, blink int, con io.ReadWriteCloser) {

	n, err := con.Write([]byte{byte(red)})
	if err != nil {
		log.Panicln("Ошибка отправки данных red: ", err)
	}
	log.Println("Отправлено байт:", n)

	n, err = con.Write([]byte{byte(green)})
	if err != nil {
		log.Panicln("Ошибка отправки данных green: ", err)
	}
	n, err = con.Write([]byte{byte(blue)})
	if err != nil {
		log.Panicln("Ошибка отправки данных blue: ", err)
	}
	n, err = con.Write([]byte{byte(blink)})
	if err != nil {
		log.Panicln("Ошибка отправки данных blink: ", err)
	}
	// con.Write([]byte("\n"))

	// buf := make([]byte, 4)
	// binary.PutUvarint(buf, uint64(250))
	// binary.PutUvarint(buf, uint64(green))
	// binary.PutUvarint(buf, uint64(blue))
	// binary.PutUvarint(buf, uint64(blink))
	// log.Println(buf)
	// n, err := con.Write(buf)
	// if err != nil {
	// 	log.Panicln("Ошибка отправки данных red: ", err)
	// }
	// log.Println("Отправлено байт:", n)

	time.Sleep(1 * time.Second)

}

// Получить список COM-портов  в windows
func readPorts() []string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "HARDWARE\\DEVICEMAP\\SERIALCOMM", registry.QUERY_VALUE)

	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	ki, err := k.Stat()

	if err != nil {
		log.Fatal(err)

	}

	// fmt.Printf("Subkey %d ValueCount %d\n", ki.SubKeyCount, ki.ValueCount)
	s, err := k.ReadValueNames(int(ki.ValueCount))
	if err != nil {
		log.Fatal(err)
	}
	kvalue := make([]string, ki.ValueCount)

	for i, test := range s {
		q, _, err := k.GetStringValue(test)
		if err != nil {
			log.Fatal(err)
		}
		kvalue[i] = q
	}

	// fmt.Printf("%s \n", kvalue)
	return kvalue

}
