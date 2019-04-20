// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package module

import (
	"fmt"
	"sysconfig_service/global"
	"testing"
)

//go test -v -test.run Test_GetProjectList
func Test_GetProjectList(t *testing.T) {
	global.Config_Store_Folder = "/Users/vikingbays/golang/AlphabetwebProject/sysconfig_service/store"

	fileList := GetProjectList()
	for _, flb := range fileList {
		fmt.Println("projectId:::: ", flb)
		file2List := GetAliasKeyList(flb.Path)
		for _, flb2 := range file2List {
			fmt.Println("aliasKey:::: ", flb2)
			file3List := GetAppNameList(flb2.Path)
			for _, flb3 := range file3List {
				fmt.Println("appName:::: ", flb3)
				file4List := GetSysconfigFileList(flb3.Path)
				for _, flb4 := range file4List {
					fmt.Println("files:::: ", flb4)
					if flb4.Next != nil {
						printFile4List(flb4.Next)
					}
				}
			}
		}
	}

}

func printFile4List(fileList []*FileListBean) {
	for _, flb := range fileList {
		fmt.Println("files:::: ", flb)
		if flb.Next != nil {
			printFile4List(flb.Next)
		}
	}
}
