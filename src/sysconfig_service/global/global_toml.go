// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package global

import (
	"alphabet/core/toml"
	"alphabet/env"
	"alphabet/log4go"
	"strings"
)

var Config_Store_Folder string = ""

var Config_Tmp_Folder string = ""

type Global_Config_List_Toml struct {
	Stores []Stores_Config_Toml
}

type Stores_Config_Toml struct {
	Storefolder string
	Tmpfolder   string
}

func initGlobalToml() {

	var globalConfigListToml Global_Config_List_Toml

	pathconfig := env.Env_Project_Resource + "/sysconfig_service/" + env.GetSysconfigPath() + "global.toml"
	if _, err := toml.DecodeFile(pathconfig, &globalConfigListToml); err != nil {
		//log4go.ErrorLog(err)
		log4go.ErrorLog("This file (global.toml) is not opened . this path is %s . errinfo is %s .", pathconfig, err.Error())
	} else {
		if globalConfigListToml.Stores != nil && len(globalConfigListToml.Stores) == 1 {
			Config_Store_Folder = globalConfigListToml.Stores[0].Storefolder
			Config_Store_Folder = strings.Replace(Config_Store_Folder, "${project}", env.Env_Project_Root, -1)

			Config_Tmp_Folder = globalConfigListToml.Stores[0].Tmpfolder
			Config_Tmp_Folder = strings.Replace(Config_Tmp_Folder, "${project}", env.Env_Project_Root, -1)
		}
	}

	if Config_Store_Folder == "" {
		log4go.ErrorLog("global.Config_Store_Folder is \"\" . ")
	}

}
