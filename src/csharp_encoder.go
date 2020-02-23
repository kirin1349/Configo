package main

import (
	"toolky"
	"strconv"
)

func encodeToCSharpClass(pfData *PlatformData) {
	for _, sheetData := range pfData.Sheets {
		if len(sheetData.Values) > 0 {
			clsName := formatCamelCaseName(sheetData.Name) + "ConfigData"
			filePath := formatPlatformScriptFilePath(pfData.Name, pfData.Language, clsName, pfData.OutputFolder, pfData.OutputExt, pfData.ClsFolder)
			err := toolky.RemoveFile(filePath)
			if err != nil {
				toolky.PrintError("remove file path " + filePath + " with error " + err.Error())
				continue
			}
			clsContent := ""
			clsContent += "\n"
			clsContent += "public class " + clsName + " : " + pfData.LibCls + " \n"
			clsContent += "{ " + "\n"
			clsContent += formatIndent("protected void onParseData()" + "\n", 1)
			clsContent += formatIndent("{" + "\n", 1)
			importMap := make(map[string]int)
			getterContent := ""
			index := 0
			for true {
				keyData := sheetData.KeyIndexes[index]
				if keyData == nil {
					break
				}
				clsContent += formatIndent(encodeCSharpKeyDataParseString(keyData) + "// ColumnIndex: " + strconv.Itoa(keyData.ColumnIndex) + "\n", 2)
				if index > 0 {
					getterContent += "\n"
				}
				getterCellContent, mainType := encodeCSharpKeyDataGetterString(keyData, &pfData.DataStructs)
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
			content := getClsHeader()
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

func encodeCSharpKeyDataParseString(keyData *WordData) (string) {
	mainType, preType := parseDataTypes(keyData.DataType)
	funcName := formatCamelCaseName(mainType)
	funcName = "read" + funcName
	if preType != "" {
		if preType != "list" {
			toolky.PrintError("unsupport pre type " + preType + " [" + keyData.Name + ", " + strconv.Itoa(keyData.ColumnIndex) + "]")
			return ""
		}
		return "setConfigValue(\"" + keyData.Name + "\", readList(this." + funcName + ")());"
	} else {
		return "setConfigValue(\"" + keyData.Name + "\", " + funcName + "());"
	}
}

func encodeCSharpKeyDataGetterString(keyData *WordData, dataStructs *(map[string]*PlatformStructConfig)) (string, string) {
	mainType, preType := parseDataTypes(keyData.DataType)
	finalType := ""
	content := "public "
	if preType != "" {
		if preType != "list" {
			toolky.PrintError("unsupport pre type " + preType + " [" + keyData.Name + ", " + strconv.Itoa(keyData.ColumnIndex) + "]")
			return "", ""
		}
		finalType += "List<"
		typeCls := (*dataStructs)[mainType]
		if typeCls != nil {
			finalType += (*dataStructs)[mainType].LibCls
		} else {
			switch(mainType) {
				case "int":
					finalType += "int";
				default:
					finalType += "string";
			}
		}
		finalType += ">"
	} else {
		typeCls := (*dataStructs)[mainType]
		if typeCls != nil {
			finalType += (*dataStructs)[mainType].LibCls
		} else {
			switch(mainType) {
				case "int":
					finalType += "int";
				default:
					finalType += "string";
			}
		}
	}
	content += finalType + " "
	content += "get" + formatCamelCaseName(keyData.Name) + "()"
	content = formatIndent(content + "\n", 1)
	content += formatIndent("{" + "\n", 1);
	content += formatIndent("return getConfigValue" + "<" + finalType + ">(\"" + keyData.Name + "\");" + "\n", 2)
	content += formatIndent("}" + "\n", 1);
	return content, mainType
}