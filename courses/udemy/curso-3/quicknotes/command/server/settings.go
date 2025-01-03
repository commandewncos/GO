package main

import (
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

type Settings struct {
	ServerPort         string `env:"SERVER_PORT,5000"`
	DatabaseConnection string `env:"DATABASE_URL,required"`
	LevelLog           string `env:"LEVEL_LOG,info"`
}

func (s Settings) GetLevelLog() slog.Level {
	switch strings.ToLower(s.LevelLog) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}

}

func (s Settings) Sprint() (envs string) {

	value := reflect.ValueOf(s)
	valueType := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := valueType.Field(i)
		envTag := strings.Split(field.Tag.Get("env"), ",")
		envName := envTag[0]
		envValue := envTag[1]
		envs += fmt.Sprintf("%s - %s\n", envName, envValue)
	}

	return
}

func (s Settings) loadSettingsFromEnv() (settings Settings) {
	value := reflect.ValueOf(s)
	valueType := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := valueType.Field(i)
		envTag := strings.Split(field.Tag.Get("env"), ",")
		envName := envTag[0]
		defaultValue := envTag[1]
		envValue := os.Getenv(envName)

		if envValue == "" && defaultValue != "required" {
			f := reflect.ValueOf(&settings).Elem().FieldByName(field.Name)
			f.SetString(defaultValue)

		} else {
			f := reflect.ValueOf(&settings).Elem().FieldByName(field.Name)
			f.SetString(envValue)

		}
	}

	return
}

func (s Settings) validate() {
	var validationMsg string
	value := reflect.ValueOf(s)
	valueType := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		envTag := strings.Split(valueType.Field(i).Tag.Get("env"), ",")
		envName := envTag[0]
		envValue := envTag[1]

		if envValue == "required" && field.String() == "" {
			validationMsg += fmt.Sprintf("%s is required\n", envName)
		}

	}
	if len(validationMsg) != 0 {
		panic(validationMsg)

	}
}

func loadSettings() Settings {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	settings := Settings{}
	settings = settings.loadSettingsFromEnv()
	settings.validate()
	fmt.Println(settings.loadSettingsFromEnv().ServerPort)
	fmt.Println(settings.loadSettingsFromEnv().DatabaseConnection)

	return settings
}
