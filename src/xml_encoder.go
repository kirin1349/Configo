package main

import(
	"toolky"
)

func encodeToXML(pfData *PlatformData, fileName string) {
	if pfData.AllInOne {
		if len(pfData.Sheets) > 0 {
			xmlContent := "<config>" + "\n"
			for _, sheetData := range pfData.Sheets {
				if len(sheetData.Values) > 0 {
					xmlList := make([](map[string]string), 0)
					for _, valueData := range sheetData.Values {
						if len(valueData.Values) > 0 {
							xmlData := make(map[string]string)
							for nameStr, valueStr := range valueData.Values {
								xmlData[nameStr] = valueStr
							}
							xmlList = append(xmlList, xmlData)
						}
					}
					xmlContent += formatIndent("<" + sheetData.Name + "s>", 1) + "\n"
					xmlContent += encodeToXMLString(sheetData.Name, xmlList, 2)
					xmlContent += formatIndent("</" + sheetData.Name + "s>", 1) + "\n"
				}
			}
			xmlContent += "</config>"
			filePath := formatPlatformFilePath(pfData.Name, pfData.Output, fileName, pfData.OutputFolder, pfData.OutputExt)
			result := toolky.QuickWrite(filePath, string(xmlContent), true)
			if result {
				toolky.PrintInfo("write " + fileName + " success")
			} else {
				toolky.PrintError("write " + fileName + " failed")
			}
		} else {
			toolky.PrintInfo("skip empty file [" + pfData.Name + ": " + fileName + "]")
		}
	} else {
		for _, sheetData := range pfData.Sheets {
			if len(sheetData.Values) > 0 {
				xmlContent := ""
				xmlList := make([](map[string]string), 0)
				for _, valueData := range sheetData.Values {
					if len(valueData.Values) > 0 {
						xmlData := make(map[string]string)
						for nameStr, valueStr := range valueData.Values {
							xmlData[nameStr] = valueStr
						}
						xmlList = append(xmlList, xmlData)
					}
				}
				xmlContent += "<" + sheetData.Name + "s>" + "\n"
				xmlContent += encodeToXMLString(sheetData.Name, xmlList, 1)
				xmlContent += "</" + sheetData.Name + "s>" + "\n"
				filePath := formatPlatformFilePath(pfData.Name, pfData.Output, sheetData.Name, pfData.OutputFolder, pfData.OutputExt)
				result := toolky.QuickWrite(filePath, string(xmlContent), true)
				if result {
					toolky.PrintInfo("write " + sheetData.Name + " success")
				} else {
					toolky.PrintError("write " + sheetData.Name + " failed")
				}
			} else {
				toolky.PrintInfo("skip empty sheet " + sheetData.Name)
			}
		}
	}
}

func encodeToXMLString(nodeName string, nodeList [](map[string]string), indent int) (string) {
	xmlStrs := ""
	for i := range(nodeList) {
		xmlStr := formatIndent("<" + nodeName + " ", indent)
		for xmlKey, xmlValue := range nodeList[i] {
			xmlStr += xmlKey + "=\"" + xmlValue + "\" "
		}
		xmlStr += "/>"
		xmlStrs += xmlStr + "\n"
	}
	return xmlStrs
}