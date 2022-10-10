package pro

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"os"
	_ "unsafe" // for go:linkname
)

// ReadJson()       : json文件 -> json
// ReadJsonOption() : json文件 -> json单个选项
// json.Marshal()   : obj     -> json
// json.Unmarshal() : json    -> obj
// gjson.Get()      : 获取json单个选项中的内容
// sjson.Set()      : 返回修改单个选项内容后的json字符串

type Json interface {
	readJson()
	readJsonOption()
	encode()
	decoder()
}

//go:linkname readJson github.com/coulsonzero/gopkg/pro.ReadJson
// readJson json文件 -> json字符串
func readJson(filename string) string {
	// 打开json文件
	f, err := os.Open(filename)
	if err != nil {
		panic("error: failed to open the json file")
	}
	defer f.Close()

	// json文件 -> struct
	var obj map[string]interface{}
	if err = json.NewDecoder(f).Decode(&obj); err != nil {
		panic("error: failed to convert the json file to obj(struct/map)")
	}

	// struct -> json
	data, err := json.Marshal(obj)
	if err != nil {
		panic("error: failed to convert the obj to json")
	}
	return string(data)
}

//go:linkname readJsonOption github.com/coulsonzero/gopkg/pro.ReadJsonOption
// readJsonOption 读取json文件中的单个选项
func readJsonOption(filename string, option string) gjson.Result {
	f, _ := os.ReadFile(filename)
	res := gjson.Get(string(f), option)
	return res
}

//go:linkname encode github.com/coulsonzero/gopkg/pro.Encode
// encode object(map/struct) -> json
func encode(obj interface{}) (string, error) {
	bytes, err := json.Marshal(&obj)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//go:linkname decode github.com/coulsonzero/gopkg/pro.Decode
// decode json -> object(map/struct)
func decode(jsonStr []byte, obj interface{}) interface{} {
	err := json.Unmarshal(jsonStr, &obj)
	if err != nil {
		panic("error: 解析失败json -> obj ")
	}
	return obj
}
