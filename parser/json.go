package parser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/johanliu/Vidar/constant"
)

type JSONParser struct{}

func (jp *JSONParser) PluginName() string {
	return "JSONParser"
}

func (jp *JSONParser) Parse(obj interface{}, req *http.Request) error {
	if !strings.HasPrefix(req.Header.Get(constant.HeaderContentType), constant.MIMEApplicationJSON) {
		return constant.UnsupportedMediaTypeError
	}

	fmt.Println(req.Body)

	if err := json.NewDecoder(req.Body).Decode(obj); err != nil {
		return constant.BadRequestError
	}

	return nil
}
