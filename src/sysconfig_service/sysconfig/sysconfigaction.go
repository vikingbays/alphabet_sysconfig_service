// Copyright 2019 The VikingBays(in Nanjing , China) . All rights reserved.
// Released under the Apache license : http://www.apache.org/licenses/LICENSE-2.0 .
//
// authors:   VikingBays
// email  :   vikingbays@gmail.com

package db

import (
	"alphabet/web"
	"fmt"
	"os"
	"sysconfig_service/global"
	"sysconfig_service/sysconfig/module"
)

type ParamBean struct {
	Projectid string //
	Aliaskey  string
	Appname   string
	Level     int //所处级别： 1 表示查询 project信息 ， 2表示查询 aliasKey信息 ， 3 表示查询 appName信息 ， 4表示查询 打包文件列表
	Path      string
}

func Filelist(paramBean *ParamBean, context *web.Context) {
	context.Return.Forward("view_filelist", nil)
}

func FilelistJson(paramBean *ParamBean, context *web.Context) {
	htmls := ""
	if paramBean == nil || paramBean.Level == 0 || paramBean.Level == 1 { //
		datafileList := module.GetProjectList()
		html := "<li class='list-group-item' onclick='clickProjectId(this)' path='%s' filename='%s'><span class='glyphicon glyphicon-unchecked' aria-hidden='true'></span><span style='padding-left:5px'>%s</span></li>"
		for _, flb := range datafileList {

			htmls = htmls + fmt.Sprintf(html, flb.Path[len(global.Config_Store_Folder):], flb.Filename, flb.Filename)
		}
	} else if paramBean.Level == 2 {
		datafileList := module.GetAliasKeyList(global.Config_Store_Folder + paramBean.Path)
		html := "<li class='list-group-item' onclick='clickAliasKey(this)' path='%s' filename='%s'><span class='glyphicon glyphicon-unchecked' aria-hidden='true'></span><span style='padding-left:5px'>%s</span></li>"
		for _, flb := range datafileList {
			htmls = htmls + fmt.Sprintf(html, flb.Path[len(global.Config_Store_Folder):], flb.Filename, flb.Filename)
		}
	} else if paramBean.Level == 3 {
		datafileList := module.GetAppNameList(global.Config_Store_Folder + paramBean.Path)
		html := "<li class='list-group-item' onclick='clickAppName(this)' path='%s' filename='%s'><span class='glyphicon glyphicon-unchecked' aria-hidden='true'></span><span style='padding-left:5px'>%s</span></li>"
		for _, flb := range datafileList {
			htmls = htmls + fmt.Sprintf(html, flb.Path[len(global.Config_Store_Folder):], flb.Filename, flb.Filename)
		}
	} else if paramBean.Level == 4 {
		datafileList := module.GetSysconfigFileList(global.Config_Store_Folder + paramBean.Path)
		//margin-left:20px
		html_file := "<li class='list-group-item' style='color:#555555; %s'><span class='glyphicon glyphicon-file' aria-hidden='true'></span><span style='padding-left:5px'>%s</span></li>"
		html_folder := "<li class='list-group-item' style='color:#999999; %s'><span class='glyphicon glyphicon-folder-open' aria-hidden='true'></span><span style='padding-left:5px'>%s</span></li>"
		for _, flb := range datafileList {
			if flb.IsDir {
				htmls = htmls + fmt.Sprintf(html_folder, "", flb.Filename)
				if flb.Next != nil {
					htmls = genHtmlForFilelistJson(flb.Next, html_file, html_folder, 1, htmls)
				}
			} else {
				htmls = htmls + fmt.Sprintf(html_file, "", flb.Filename)
			}

		}
	}
	context.Return.Json(htmls)
}

func genHtmlForFilelistJson(datafileList []*module.FileListBean, html_file_template string, html_folder_template string, deep int, htmls string) string {
	for _, flb := range datafileList {
		if flb.IsDir {
			htmls = htmls + fmt.Sprintf(html_folder_template, fmt.Sprintf("margin-left:%dpx", deep*20), flb.Filename)
			if flb.Next != nil {
				htmls = genHtmlForFilelistJson(flb.Next, html_file_template, html_folder_template, (deep + 1), htmls)
			}
		} else {
			htmls = htmls + fmt.Sprintf(html_file_template, fmt.Sprintf("margin-left:%dpx", deep*20), flb.Filename)
		}
	}
	return htmls
}

func Downzip(paramBean *ParamBean, context *web.Context) {

	if paramBean.Projectid != "" && paramBean.Aliaskey != "" && paramBean.Appname != "" {
		desc, filename, err := module.SysconfigCompress(paramBean.Projectid, paramBean.Aliaskey, paramBean.Appname)

		if err == nil {
			defer os.RemoveAll(desc)
			context.Return.DownloadFile(desc, filename)
		} else {
			context.Return.Forward("view_downzip_err", map[string]interface{}{"error": err.Error(), "paramBean": paramBean})
		}

	} else {
		context.Return.Forward("view_downzip_err", map[string]interface{}{"error": "params has null .", "paramBean": paramBean})
	}

}
