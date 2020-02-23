
//https://github.com/360EntSecGroup-Skylar/excelize

package main

import(
	"toolky"
	"encoding/json"
	"strings"
	"path/filepath"
	"fmt"
)

var(
	mConfig Config
)

func main() {

	configPath := toolky.GetOSArgByKey("-c")
	filesPath := toolky.GetOSArgByKey("-i")

	result, configStr := toolky.QuickRead(configPath, true)
	
	if !result {
		return
	}
	
	result, filesStr := toolky.QuickRead(filesPath, true)
	
	if !result {
		return
	}

	json.Unmarshal([]byte(configStr), &mConfig)

	filesStr = strings.Replace(filesStr, "\r\n", "\n", 0)
	filesStr = strings.Replace(filesStr, "\r", "\n", 0)
	filesArr := strings.Split(filesStr, "\n")

	for _, filePath := range filesArr {
		filePath = strings.TrimSpace(filePath)
		if filePath == "" {
			continue
		}

		fileName := filepath.Base(filePath)
		index := strings.Index(fileName, ".")
		if index > 0 {
			fileName = fileName[0:index]
		}
		fileNameArr := strings.Split(fileName, "_")

		result, platformDataMap := parseExcel(filePath, &mConfig)
	
		if !result {
			continue
		}
	
		for _, pfData := range (*platformDataMap) {
			err := toolky.CreateFolder("./" + pfData.Name)
			if err != nil {
				fmt.Println("check platform folder failed with error " + err.Error())
				continue
			}
			if pfData.OutputFolder {
				err := toolky.CreateFolder("./" + pfData.Name + "/" + pfData.Output)
				if err != nil {
					fmt.Println("check platform folder failed with error " + err.Error())
					continue
				}
			}
			switch(pfData.Output) {
				case "binary": encodeToBinary(pfData, fileNameArr[1])
				case "json": encodeToJson(pfData, fileNameArr[1])
				default: encodeToXML(pfData, fileNameArr[1])
			}
			if pfData.Language == ""{
				continue
			}
			if pfData.OutputFolder {
				err = toolky.CreateFolder("./" + pfData.Name + "/" + pfData.Language)
				if err != nil {
					fmt.Println("check platform class folder failed with error " + err.Error())
					continue
				}
			} else {
				err = toolky.CreateFolder("./" + pfData.Language)
				if err != nil {
					fmt.Println("check platform class folder failed with error " + err.Error())
					continue
				}
			}
			if pfData.CreateCls {
				err = toolky.CreateFolder("./" + pfData.Name + "/" + pfData.Language + "/" + pfData.ClsFolder)
				if err != nil {
					fmt.Println("check platform class folder failed with error " + err.Error())
					continue
				}
				switch(pfData.Language) {
					case "typescript": encodeToTypeScriptClass(pfData)
					case "csharp": encodeToCSharpClass(pfData)
				}
			}
		}
	}
}
