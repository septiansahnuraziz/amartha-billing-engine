package utils

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot load file .env: ", err)
		panic(err)
	}

	value := GetEnvOrDefault(key, "").(string)
	return value
}

func GetEnvOrDefault(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func WriteStringTemplate(stringTemplate string, args ...interface{}) string {
	return fmt.Sprintf(stringTemplate, args...)
}

func StringToLower(_string string) string {
	return strings.ToLower(_string)
}

func StringToUpper(s string) string {
	return strings.ToUpper(s)
}

func MyCaller(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	details := runtime.FuncForPC(pc)

	if ok && details != nil {
		return details.Name()
	}

	return "failed to identify method caller"
}

// WrapCloser call close and log the error
func WrapCloser(close func() error) {
	if close == nil {
		return
	}
	if err := close(); err != nil {
		logrus.Error(err)
	}
}

func JSONUnmarshal(data []byte, v interface{}) error {
	if err := json.Unmarshal([]byte(data), &v); err != nil {
		return err
	}
	return nil
}

func JSONMarshal(data interface{}) ([]byte, error) {
	result, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ToByte(i interface{}) []byte {
	byte_, _ := JSONMarshal(i)
	return byte_
}

func Dump(i interface{}) string {
	return string(ToByte(i))
}

func ExpectedString(v interface{}) string {
	var result string
	switch v := v.(type) {
	case int, uint, uint64:
		result = fmt.Sprintf("%d", v)
	case float64:
		result = fmt.Sprintf("%f", v)
	case string:
		result = v
	}
	return result
}

func SplitString(value, separator string) []string {
	return strings.Split(value, separator)
}

func ExpectedUint(v interface{}) uint {
	var result uint
	switch v := v.(type) {
	case int:
		result = uint(v)
	case int64:
		result = uint(v)
	case float64:
		result = uint(v)
	case string:
		convertedString, _ := strconv.ParseUint(v, 10, 32)
		result = uint(convertedString)
	case uint:
		result = v
	}
	return result
}
