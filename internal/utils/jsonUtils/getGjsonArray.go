package utils_json

import "github.com/tidwall/gjson"

func GetGJsonStringArray(data gjson.Result, path string) []string {
	var res []string
	vals := data.Get(path).Array()
	for _, item := range vals {
		res = append(res, item.Str)
	}
	return res
}
