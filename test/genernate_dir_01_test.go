package test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var stRootDir string
var stSeparator string
var iJsonData map[string]any

const stJsonName = "dirLib.json"

func loadJson() {
	// 获取路径分隔符
	stSeparator = string(filepath.Separator)
	// 获取当前路径
	stWorkDir, _ := os.Getwd()
	// 获取根路径
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)]

	gnJsonBytes, _ := os.ReadFile(stWorkDir + stSeparator + stJsonName)
	// 将读取的字节数据转换为Json
	err := json.Unmarshal(gnJsonBytes, &iJsonData)
	if err != nil {
		panic("Load Json Data Error: " + err.Error())
	}
}

// 解析字典
func parseMap(mpData map[string]any, stParentDir string) {
	for k, v := range mpData {
		// 判断值类型, 是字符串则创建
		switch v.(type) {
		case string:
			{
				// 断言 v 是 string 类型
				path, _ := v.(string)
				if path == "" {
					continue
				}

				if stParentDir != "" {
					path = stParentDir + stSeparator + path
					if k == "text" {
						stParentDir = path
					}
				} else {
					stParentDir = path
				}

				CreateDir(path)
			}
		case []any:
			{
				parseArray(v.([]any), stParentDir)
			}
		}
	}
}

// 解析数组
func parseArray(giJsonData []any, stParentDir string) {
	for _, v := range giJsonData {
		// 断言 v 是 字典 类型
		mapV, _ := v.(map[string]any)
		parseMap(mapV, stParentDir)
	}
}

// 创建文件
func CreateDir(path string) {
	if path == "" {
		return
	}

	err := os.MkdirAll(stRootDir+stSeparator+path, fs.ModePerm)
	if err != nil {
		panic("Create Directory Error: " + err.Error())
	}
}

func TestGenerateDir01(t *testing.T) {
	loadJson()
	parseMap(iJsonData, "")
}
