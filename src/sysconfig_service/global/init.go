// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package global

import (
	"alphabet/log4go"
)

func InitAll() {
	initGlobalToml()
	log4go.InfoLog("Store Folder is %s .", Config_Store_Folder)
	log4go.InfoLog("Tmp Folder is %s .", Config_Tmp_Folder)
}
