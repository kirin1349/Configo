package main

import (
	"toolky"
	"strconv"
	"strings"
)

func encodeToTypeScriptClass(pfData *PlatformData) {
	for _, sheetData := range pfData.Sheets {
		if len(sheetData.Values) > 0 {
			clsName := formatCamelCaseName(sheetData.Name) + "ConfigData"
			filePath := formatPlatformScriptFilePath(pfData.Name, pfData.Language, clsName, pfData.OutputFolder, pfData.OutputExt, pfData.OutputExt)
			err := toolky.RemoveFile(filePath)
			if err != nil {
				toolky.PrintError("remove file path " + filePath + " with error " + err.Error())
				continue
			}
			clsContent := ""
			clsContent += "export class " + clsName + " extends " + pfData.LibCls + " \n"
			clsContent += "{ " + "\n"
			clsContent += formatIndent("protected onParseData()" + "\n", 1)
			clsContent += formatIndent("{" + "\n", 1)
			importMap := make(map[string]int)
			getterContent := ""
			index := 0
			for true {
				keyData := sheetData.KeyIndexes[index]
				if keyData == nil {
					break
				}
				clsContent += formatIndent(encodeTypeScriptKeyDataParseString(pfData.Output, keyData) + "// ColumnIndex: " + strconv.Itoa(keyData.ColumnIndex) + "\n", 2)
				if index > 0 {
					getterContent += "\n"
				}
				getterCellContent, mainType := encodeTypeScriptKeyDataGetterString(keyData, &pfData.DataStructs)
				getterContent += getterCellContent
				if mainType != "" {
					if importMap[mainType] != 1 {
						importMap[mainType] = 1
					}
				}
				index++
			}
			clsContent += formatIndent("}" + "\n", 1)
			clsContent += "\n"
			clsContent += getterContent
			clsContent += "} " + "\n"
			importContent := ""
			index = strings.Index(pfData.LibCls, ".")
			if index <= 0 {
				importContent += "import { " + pfData.LibCls + " } from \"" + pfData.LibPath + "/" + pfData.LibCls + "\";\n"
			}
			for importType, _ := range importMap {
				structConfig := pfData.DataStructs[importType]
				if structConfig == nil {
					continue
				}
				index = strings.Index(structConfig.LibPath, ".")
				if index <= 0 {
					importContent += "import { " + structConfig.LibCls + " } from \"" + structConfig.LibPath + "/" + structConfig.LibCls + "\";\n"
				}
			}
			content := getClsHeader()
			if importContent != "" {
				content += importContent
				content += "\n"
			}
			content += clsContent
			result := toolky.QuickWrite(filePath, string(content), true)
			if result {
				toolky.PrintInfo("write " + clsName + " success")
			} else {
				toolky.PrintError("write " + clsName + " failed")
			}
		} else {
			toolky.PrintInfo("skip empty sheet " + sheetData.Name)
		}
	}
}

func encodeTypeScriptKeyDataParseString(output string, keyData *WordData) (string) {
	mainType, preType := parseDataTypes(keyData.DataType)
	funcName := formatCamelCaseName(mainType)
	funcName = "read" + funcName
	if preType != "" {
		if preType != "list" {
			toolky.PrintError("unsupport pre type " + preType + " [" + keyData.Name + ", " + strconv.Itoa(keyData.ColumnIndex) + "]")
			return ""
		}
		if output == "json" {
			return "this.setConfigValue(\"" + keyData.Name + "\", this.readList(this.getSrcValue(\"" + keyData.Name + "\"), this." + funcName + ", this));"
		}
		return "this.setConfigValue(\"" + keyData.Name + "\", this.readList(this." + funcName + ", this));"
	} else {
		if output == "json" {
			return "this.setConfigValue(\"" + keyData.Name + "\", this." + funcName + "(this.getSrcValue(\"" + keyData.Name + "\")));"
		}
		return "this.setConfigValue(\"" + keyData.Name + "\", this." + funcName + "());"
	}
}

func encodeTypeScriptKeyDataGetterString(keyData *WordData, dataStructs *(map[string]*PlatformStructConfig)) (string, string) {
	mainType, preType := parseDataTypes(keyData.DataType)
	content := "get" + formatCamelCaseName(keyData.Name)
	if preType != "" {
		if preType != "list" {
			toolky.PrintError("unsupport pre type " + preType + " [" + keyData.Name + ", " + strconv.Itoa(keyData.ColumnIndex) + "]")
			return "", ""
		}
		content += "List(): "
		typeCls := (*dataStructs)[mainType]
		if typeCls != nil {
			content += (*dataStructs)[mainType].LibCls
		} else {
			switch(mainType) {
				case "int":
					content += "number";
				default:
					content += "string";
			}
		}
		content += "[]"
	} else {
		content += "(): "
		typeCls := (*dataStructs)[mainType]
		if typeCls != nil {
			content += (*dataStructs)[mainType].LibCls
		} else {
			switch(mainType) {
				case "int":
					content += "number";
				default:
					content += "string";
			}
		}
	}
	content = formatIndent(content + "\n", 1)
	content += formatIndent("{" + "\n", 1);
	content += formatIndent("return this.getConfigValue(\"" + keyData.Name + "\");" + "\n", 2)
	content += formatIndent("}" + "\n", 1);
	return content, mainType
}