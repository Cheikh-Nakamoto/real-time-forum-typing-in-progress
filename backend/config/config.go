package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const configFilePath = ".env"

type ConfigValue string

var Conf = make(map[string]ConfigValue)

func init(){
	LoadEnv()
}

func LoadEnv() error {
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return err
	}
	var scanner = bufio.NewScanner(configFile)
	var lineNumber = 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		if strings.Contains(line, "#") {
			line = strings.TrimSpace(line[:strings.Index(line, "#")])
		}
		s := strings.Split(line, "=")
		if len(s) == 2 {
			key, value := strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
			Conf[key] = ConfigValue(strings.Trim(value, "\""))
		} else {
			continue
		}
		lineNumber++
	}

	return nil
}

func Get(key string) ConfigValue {
	if Conf[key] == "" {
		fmt.Printf("error: no value found for the key %v\n", key)
		os.Exit(0)
	}
	return Conf[key]
}

func (v ConfigValue) ToInt() int {
	valToInt, err := strconv.Atoi(string(v))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return valToInt
}

func (v ConfigValue) ToString() string {
	return string(v)
}

func (v ConfigValue) ToBool() (bool, error) {
	vToLower := strings.ToLower(string(v))
	if vToLower == "true" || vToLower == "1" {
		return true, nil
	} else if vToLower == "false" || vToLower == "0" {
		return false, nil
	}

	return false, fmt.Errorf("error: could not get %v as bool", v)
}
