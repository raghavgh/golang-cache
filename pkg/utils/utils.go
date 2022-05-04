package utils

import (
	"encoding/json"
	"github.com/raghavgh/bookmanagement/pkg/models"
	"io/ioutil"
	"net/http"
	"reflect"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// SetResponse prepares response with body and all other needed data which needs to be sent to client
func SetResponse(status int, body interface{}, headers interface{}, w *http.ResponseWriter, error error) {

	if reflect.TypeOf(headers).Kind() == reflect.Map {
		for key, val := range headers.(map[string]interface{}) {
			(*w).Header().Set(key, val.(string))
		}
	}
	if error != nil {
		body = map[string]interface{}{"error": error}
	}
	res, _ := json.Marshal(body)
	(*w).Write(res)
	(*w).WriteHeader(status)
}

func CopyBookModelIfNotEmpty(source *models.Book, dest *models.Book) {
	if source.Name != "" {
		dest.Name = source.Name
	}
	if source.Author != "" {
		dest.Author = source.Author
	}
	if source.Publication != "" {
		dest.Publication = source.Publication
	}
}
