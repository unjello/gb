package core

import "fmt"

func getStandard(options map[string]interface{}) string {
	if standard, ok := options["standard"]; ok == true {
		return standard.(string)
	} else {
		return "c++17"
	}
}

func GetCxxOptions(options map[string]interface{}) string {
	return fmt.Sprintf("-Wall -Werror -std=%s", getStandard(options))
}
