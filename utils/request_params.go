package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// ParseRequestParams ...
func ParseRequestParams(c buffalo.Context, paramModel interface{}) (interface{}, error) {
	req := c.Request()
	modelType := reflect.TypeOf(paramModel)
	params := reflect.New(modelType).Interface()

	if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
		log.Printf(" [!] Exception: Internal Server Error: %s: %+v",
			modelType.String(), errors.Wrap(err, "Failed to JSON decode request body"))
		return reflect.New(modelType), err
	}

	log.Println(fmt.Sprintf("%+v", params))
	err := req.Body.Close()
	if err != nil {
		log.Printf(" [!] Exception: BodyCloseWithErrorLog: %+v", err)
		return reflect.New(modelType), err
	}
	return params, nil
}
