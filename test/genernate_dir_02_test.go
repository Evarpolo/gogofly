package test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type node struct {
	Text     string `json:"text"`
	Children []node `json:"children"`
}

var stRootDir02 string
var stSeparator02 string
var iRootNode node

const stJsonName02 = "dirLib.json"

func loadJson02() {
	stSeparator02 = string(filepath.Separator)
	stWorkDir, _ := os.Getwd()
	stRootDir02 = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator02)]

	gnJsonFileDytes, _ := os.ReadFile(stWorkDir + stSeparator02 + stJsonName02)
	err := json.Unmarshal(gnJsonFileDytes, &iRootNode)
	if err != nil {
		panic("Load Json Data Error: " + err.Error())
	}
}

func parseNode(iNode node, stParentDir string) {
	if iNode.Text != "" {
		CreateDir02(iNode, stParentDir)
	}

	if stParentDir != "" {
		stParentDir = stParentDir + stSeparator02
	}

	if iNode.Text != "" {
		stParentDir += iNode.Text
	}

	for _, v := range iNode.Children {
		parseNode(v, stParentDir)
	}
}

func CreateDir02(iNode node, stParentDir string) {
	stDirPath := stRootDir02 + stSeparator02
	if stParentDir != "" {
		stDirPath = stDirPath + stParentDir + stSeparator02
	}
	stDirPath = stDirPath + iNode.Text

	err := os.MkdirAll(stDirPath, fs.ModePerm)
	if err != nil {
		panic("Create Directory Error: " + err.Error())
	}
}

func TestGenerateDir02(t *testing.T) {
	loadJson02()
	parseNode(iRootNode, "")
}
