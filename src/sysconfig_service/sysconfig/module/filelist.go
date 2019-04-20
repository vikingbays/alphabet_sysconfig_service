// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package module

import (
	"alphabet/log4go"
	"fmt"
	"io/ioutil"
	"strings"
	"sysconfig_service/global"
)

type FileListBean struct {
	Filename string          // 文件名称
	Path     string          // 文件存储全路径
	IsDir    bool            // 是否是文件夹
	Level    int             // 所处级别： 1 表示 projectId ， 2表示 aliasKey ， 3 表示 appName ， 4表示 打包文件列表
	Next     []*FileListBean // 下一层目录
}

func GetProjectList() []*FileListBean {
	fileList := make([]*FileListBean, 0, 1)

	rd, err := ioutil.ReadDir(global.Config_Store_Folder)
	if err != nil {
		log4go.ErrorLog(err)
	} else {
		for _, fi := range rd {
			if fi.IsDir() {
				flb := new(FileListBean)
				flb.Filename = fi.Name()
				flb.Path = fmt.Sprintf("%s/%s", global.Config_Store_Folder, fi.Name())
				flb.IsDir = true
				flb.Level = 1
				fileList = append(fileList, flb)
			}
		}
	}
	return fileList
}

func GetAliasKeyList(pathForProjectId string) []*FileListBean {
	fileList := make([]*FileListBean, 0, 1)

	rd, err := ioutil.ReadDir(pathForProjectId)
	if err != nil {
		log4go.ErrorLog(err)
	} else {
		for _, fi := range rd {
			if fi.IsDir() {
				flb := new(FileListBean)
				flb.Filename = fi.Name()
				flb.Path = fmt.Sprintf("%s/%s", pathForProjectId, fi.Name())
				flb.IsDir = true
				flb.Level = 1
				fileList = append(fileList, flb)
			}
		}
	}
	return fileList
}

func GetAppNameList(pathForAliasKey string) []*FileListBean {
	fileList := make([]*FileListBean, 0, 1)

	rd, err := ioutil.ReadDir(pathForAliasKey)
	if err != nil {
		log4go.ErrorLog(err)
	} else {
		for _, fi := range rd {
			if fi.IsDir() {
				flb := new(FileListBean)
				flb.Filename = fi.Name()
				flb.Path = fmt.Sprintf("%s/%s", pathForAliasKey, fi.Name())
				flb.IsDir = true
				flb.Level = 1
				fileList = append(fileList, flb)
			}
		}
	}
	return fileList
}

func GetSysconfigFileList(pathForAppName string) []*FileListBean {
	fileList := make([]*FileListBean, 0, 1)

	rd, err := ioutil.ReadDir(pathForAppName)
	if err != nil {
		log4go.ErrorLog(err)
	} else {
		for _, fi := range rd {
			if !strings.HasPrefix(fi.Name(), ".") {
				flb := new(FileListBean)
				flb.Filename = fi.Name()
				flb.Path = fmt.Sprintf("%s/%s", pathForAppName, fi.Name())
				flb.Level = 1
				fileList = append(fileList, flb)
				if fi.IsDir() {
					flb.IsDir = true
					flb.Next = GetSysconfigFileList(flb.Path)
				} else {
					flb.IsDir = false
				}
			}
		}
	}
	return fileList
}
