package dao

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type reflectStruct func(k, v, valueType string)

// Generate object atributes
func mapAttributes(m interface{}, rs reflectStruct) {
	typ := reflect.TypeOf(m)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
	}

	val := reflect.ValueOf(m).Elem()
	t := reflect.TypeOf(m).Elem()

	for i := 0; i < val.NumField(); i++ {

		//name
		p := t.Field(i)

		// value
		valueField := val.Field(i)
		f := valueField.Interface()
		val := reflect.ValueOf(f)

		rs(p.Name, val.String(), valueField.Type().String())
	}

}

// getItems passed a function to mapAttributes to create the items
// that are then passed to dynamodb.
// It will  return a list of parameters that are acceptable input for dynamodb request
// it uses mapAttributes and reflect to determine what public fields are available and
// what values do they have. This way we only add the existing values to the structure
// and not everything. It will also work for any kind of structure
// Example:
//     "Items": [
//        {
//            "Field1": {
//                "S": "value1"
//            },
//            "Id": {
//                "S": "id1"
//            }
//        }
//    ],

func getItems(i interface{}) map[string]*dynamodb.AttributeValue {

	items := map[string]*dynamodb.AttributeValue{}
	mapAttributes(i, func(k, v, vType string) {
		if vType == "string" && v != "" {
			items[k] = &dynamodb.AttributeValue{S: aws.String(v)}
		}
	})
	return items
}
