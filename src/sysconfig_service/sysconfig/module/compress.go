// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package module

import (
	"alphabet/core/utils"
	"fmt"
	"math/rand"
	"sysconfig_service/global"
	"time"
)

func SysconfigCompress(projectid string, aliaskey string, appname string) (string, string, error) {
	rand.Seed(time.Now().UnixNano())
	destFileName := fmt.Sprintf("%s_%s_%s_%02d%02d.zip", projectid, aliaskey, appname, rand.Intn(99), rand.Intn(99))
	sysconfigFolder := "sysconfig." + aliaskey
	srcRoot := global.Config_Store_Folder + "/" + projectid + "/" + aliaskey + "/" + appname
	dest := global.Config_Tmp_Folder + "/" + destFileName

	if !utils.ExistFile(srcRoot) {
		return "", "", fmt.Errorf("When zip file, source file is not exist. filepath: %s .", srcRoot)
	}

	if !utils.ExistFile(global.Config_Tmp_Folder) {
		return "", "", fmt.Errorf("When zip file, dest folder is not exist. filepath: %s .", global.Config_Tmp_Folder)
	}

	utils.CompressZip(srcRoot, dest, sysconfigFolder, []string{"."})

	return dest, destFileName, nil

}

func SysconfigDeCompress(projectid string, aliaskey string, appname string, srcFile string) {
	dest := global.Config_Tmp_Folder + "/" + fmt.Sprintf("%s_%s_%s", projectid, aliaskey, appname)
	utils.DeCompressZip(srcFile, dest)
}
