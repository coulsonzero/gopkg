package pro

import "github.com/tidwall/gjson"

type Json interface {
	ReadJson()
	ReadJsonOption()
	Encode()
	Encoder()
}

func ReadJson(filename string) string

func readJsonOption(filename string, option string) gjson.Result

func Encode(obj interface{}) (string, error)

func Decode(jsonStr []byte, obj interface{}) interface{}
