package main

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
	jq := gojsonq.New(gojsonq.SetDecoder(&yamlDecoder{})).File("go语言每日一库/day14 gojsonq/3 高级查询/3.5 其他格式/data.yaml")
	jq.From("items").Where("price", "<=", 500)
	fmt.Printf("%v\n", jq.First())
}

type yamlDecoder struct {
}

func (i *yamlDecoder) Decode(data []byte, v interface{}) error {
	bb, err := yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bb, &v)
}
