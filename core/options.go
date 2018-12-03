package core

import "fmt"

func getStandard(options map[string]interface{}) string {
	if standard, ok := options["standard"]; ok == true {
		return standard.(string)
	} else {
		return "c++17"
	}
}

func getConcepts(options map[string]interface{}) string {
	if features, ok := options["features"]; ok == true {
		for _, v := range features.([]interface{}) {
			if v.(string) == "concepts" {
				return "-fconcepts "
			}
		}
	}

	return ""
}

func GetCxxOptions(options map[string]interface{}) string {
	return fmt.Sprintf("-Wall -Werror -std=%s %s", getStandard(options), getConcepts(options))
}
