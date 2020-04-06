package utils

import "strconv"

func MapInterface2String(inputData map[string]interface{}) map[string]string {
	outputData := map[string]string{}
	for key, value := range inputData {
		switch value.(type) {
		case string:
			outputData[key] = value.(string)
		case int:
			tmp := value.(int)
			outputData[key] = strconv.Itoa(tmp)
		case int64:
			tmp := value.(int64)
			outputData[key] = strconv.FormatInt(tmp, 10)
		}
	}
	return outputData
}

func MapString2Interface(fromMaps []map[string]string) (toMaps []map[string]interface{}) {
	var dest []map[string]interface{}

	for fromIndex := range fromMaps {
		fromMap := fromMaps[fromIndex]
		toMap := make(map[string]interface{})

		for key, value := range fromMap {
			toMap[key] = interface{}(value)
		}

		dest = append(dest, toMap)
	}

	return dest
}
