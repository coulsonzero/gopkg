package gopkg

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"os"
)

// ReadJson()       : json文件 -> json
// ReadJsonOption() : json文件 -> json单个选项
// json.Marshal()   : obj     -> json
// json.Unmarshal() : json    -> obj
// gjson.Get()      : 获取json单个选项中的内容
// sjson.Set()      : 返回修改单个选项内容后的json字符串

// ReadJson json文件 -> json字符串
func ReadJson(filename string) string {
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

// ReadJsonOption 读取json文件中的单个选项
func ReadJsonOption(filename string, option string) gjson.Result {
	f, _ := os.ReadFile(filename)
	res := gjson.Get(string(f), option)
	return res
}

// Encode object(map/struct) -> json
func Encode(obj interface{}) (string, error) {
	bytes, err := json.Marshal(&obj)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Decode json -> object(map/struct)
func Decode(jsonStr []byte, obj interface{}) interface{} {
	err := json.Unmarshal(jsonStr, &obj)
	if err != nil {
		panic("error: 解析失败json -> obj ")
	}
	return obj
}
