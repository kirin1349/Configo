package main

import(
	"toolky"
	"encoding/json"
)

func encodeToJson(pfData *PlatformData, fileName string) {
	if pfData.AllInOne {
		if len(pfData.Sheets) > 0 {
			jsonMap := make(map[string]([](map[string]string)))
			for _, sheetData := range pfData.Sheets {
				if len(sheetData.Values) > 0 {
					jsonList := make([](map[string]string), 0)
					for _, valueData := range sheetData.Values {
						if len(valueData.Values) > 0 {
							jsonData := make(map[string]string)
							for nameStr, valueStr := range valueData.Values {
								jsonData[nameStr] = valueStr
							}
							jsonList = append(jsonList, jsonData)
						}
					}
					jsonMap[sheetData.Name] = jsonList
				}
			}
			jsonStr, err := json.MarshalIndent(jsonMap, "", "    ")
			if err != nil {
				toolky.PrintError("write " + fileName + "  with error " + err.Error())
			} else {
				filePath := formatPlatformFilePath(pfData.Name, pfData.Output, fileName, pfData.OutputFolder, pfData.OutputExt)
				result := toolky.QuickWrite(filePath, string(jsonStr), true)
				if result {
					toolky.PrintInfo("write " + fileName + " success")
				} else {
					toolky.PrintError("write " + fileName + " failed")
				}
			}
		} else {
			toolky.PrintInfo("skip empty file [" + pfData.Name + ": " + fileName + "]")
		}
	} else {
		for _, sheetData := range pfData.Sheets {
			if len(sheetData.Values) > 0 {
				jsonList := make([](map[string]string), 0)
				for _, valueData := range sheetData.Values {
					if len(valueData.Values) > 0 {
						jsonData := make(map[string]string)
						for nameStr, valueStr := range valueData.Values {
							jsonData[nameStr] = valueStr
						}
						jsonList = append(jsonList, jsonData)
					}
				}
				jsonStr, err := json.MarshalIndent(jsonList, "", "    ")
				if err != nil {
					toolky.PrintError("write " + fileName + "  with error " + err.Error())
				} else {
					filePath := formatPlatformFilePath(pfData.Name, pfData.Output, sheetData.Name, pfData.OutputFolder, pfData.OutputExt)
					result := toolky.QuickWrite(filePath, string(jsonStr), true)
					if result {
						toolky.PrintInfo("write " + sheetData.Name + " success")
					} else {
						toolky.PrintError("write " + sheetData.Name + " failed")
					}
				}
			} else {
				toolky.PrintInfo("skip empty sheet " + sheetData.Name)
			}
		}
	}
}