package commands

import (
	"io/ioutil"

	realUtil "go-cx/pkg/lib/fyne/fyne/internal/util"
)

type packagerUtil interface {
	Exists(path string) bool
	CopyFile(source string, target string) error
	CopyExeFile(src, tgt string) error
	WriteFile(target string, data []byte) error
	EnsureSubDir(parent, name string) string

	RequireAndroidSDK() error
	AndroidBuildToolsPath() string

	IsAndroid(os string) bool
	IsIOS(os string) bool
	IsMobile(os string) bool
}

type defaultUtil struct{}

func (d defaultUtil) Exists(path string) bool {
	return realUtil.Exists(path)
}

func (d defaultUtil) CopyFile(source string, target string) error {
	return realUtil.CopyFile(source, target)
}

func (d defaultUtil) CopyExeFile(src, tgt string) error {
	return realUtil.CopyExeFile(src, tgt)
}

func (d defaultUtil) WriteFile(target string, data []byte) error {
	return ioutil.WriteFile(target, data, 0644)
}

func (d defaultUtil) EnsureSubDir(parent, name string) string {
	return realUtil.EnsureSubDir(parent, name)
}

func (d defaultUtil) RequireAndroidSDK() error {
	return realUtil.RequireAndroidSDK()
}

func (d defaultUtil) AndroidBuildToolsPath() string {
	return realUtil.AndroidBuildToolsPath()
}

func (d defaultUtil) IsAndroid(os string) bool {
	return realUtil.IsAndroid(os)
}

func (d defaultUtil) IsIOS(os string) bool {
	return realUtil.IsIOS(os)
}

func (d defaultUtil) IsMobile(os string) bool {
	return realUtil.IsMobile(os)
}

var util packagerUtil

func init() {
	util = defaultUtil{}
}
