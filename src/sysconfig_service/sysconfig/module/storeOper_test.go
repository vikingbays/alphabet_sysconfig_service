// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package module

import (
	"sysconfig_service/global"
	"testing"
)

//go test -v -test.run Test_Compress
func Test_Compress(t *testing.T) {
	global.Config_Store_Folder = "/Users/vikingbays/golang/AlphabetwebProject/sysconfig_service/store"
	global.Config_Tmp_Folder = "/Users/vikingbays/golang/AlphabetwebProject/sysconfig_service/tmp"

	SysconfigCompress("octopus", "base", "sample_octopus_frontend")
}

//go test -v -test.run Test_DeCompress
func Test_DeCompress(t *testing.T) {
	global.Config_Store_Folder = "/Users/vikingbays/golang/AlphabetwebProject/sysconfig_service/store"
	global.Config_Tmp_Folder = "/Users/vikingbays/golang/AlphabetwebProject/sysconfig_service/tmp"

	SysconfigDeCompress("octopus", "base", "sample_octopus_frontend", global.Config_Tmp_Folder+"/octopus_base_sample_octopus_frontend_0327.zip")
}
