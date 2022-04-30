package helpers

import (
	"fmt"
	"os"
	"time"
)

func Logger(event string, message string, isPanic bool) {
	// Load Timezone
	location, err := time.LoadLocation(os.Getenv("APP_TIMEZONE"))

	if err != nil {
		// throw err
		panic(fmt.Errorf("cannot load time for the timezone: %v", os.Getenv("APP_TIMEZONE")))
	}

	layout := "2006-01-02 15:04:05 Z07:00"
	date := time.Now().In(location).Format(layout)

	logMessage := fmt.Sprintf("[%s][%s] %s", date, event, message)
	if isPanic {
		panic(logMessage)
	} else {
		fmt.Println(logMessage)
	}
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func WriteFile(location string, data []byte) error {
	f, err := os.Create(location)

	if err != nil {
		msg := fmt.Sprintf("Cannot create file: %v", location)
		Logger("WRITE FILE", msg, false)
		return fmt.Errorf(msg)
	}

	defer f.Close()

	_, err2 := f.Write(data)

	if err2 != nil {
		msg := fmt.Sprintf("Cannot writing file: %v", location)
		Logger("WRITE FILE", msg, false)
		return fmt.Errorf(msg)
	}

	return nil
}

func ReadFile(location string) (file []byte, err error) {
	file, err = os.ReadFile(location)
	if err != nil {
		return nil, err
	}
	return file, nil
}
