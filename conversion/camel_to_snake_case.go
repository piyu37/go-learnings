package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func camel_to_snake_case() {
	jsonFile, err := os.Open("request.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON: %v", err)
		return
	}

	convertedData := convertFieldNames(data)

	// Marshal the converted data back to JSON
	convertedJSON, err := json.Marshal(convertedData)
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %v", err)
		return
	}

	err = ioutil.WriteFile("response.json", convertedJSON, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func convertFieldNames(data interface{}) interface{} {
	value := reflect.ValueOf(data)
	switch value.Kind() {
	case reflect.Map:
		newMap := make(map[string]interface{})
		iter := value.MapRange()
		for iter.Next() {
			key := iter.Key().String()
			if strings.Contains(key, "CustomInfo") {
				continue
			}

			if strings.Contains(key, "centAmount") {
				ca := iter.Value().Interface().(string)
				centAmt, err := strconv.Atoi(ca)
				if err != nil {
					return ca
				}
				fieldName := camelToSnakeCase(key)
				var centAmtInt interface{} = centAmt
				newMap[fieldName] = convertFieldNames(centAmtInt)
				continue
			}

			fieldName := camelToSnakeCase(key)
			newMap[fieldName] = convertFieldNames(iter.Value().Interface())
		}
		return newMap
	case reflect.Slice:
		newSlice := make([]interface{}, value.Len())
		for i := 0; i < value.Len(); i++ {
			newSlice[i] = convertFieldNames(value.Index(i).Interface())
		}
		return newSlice
	case reflect.Struct:
		newStruct := make(map[string]interface{})
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fieldName := value.Type().Field(i).Name
			if strings.Contains(fieldName, "CustomInfo") {
				continue
			}
			fieldValue := convertFieldNames(field.Interface())
			newStruct[camelToSnakeCase(fieldName)] = fieldValue
		}
		return newStruct
	default:
		return data
	}
}

// Function to convert camel case to snake case
func camelToSnakeCase(s string) string {
	var result strings.Builder

	for i, ch := range s {
		if i > 0 && (ch >= 'A' && ch <= 'Z') {
			result.WriteByte('_')
		}
		result.WriteRune(ch)
	}

	return strings.ToLower(result.String())
}
