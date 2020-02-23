package main

import (
	"strings"
)

func formatIndent(content string, indent int) (string) {
	for indent > 0 {
		indent--
		content = "    " + content
	}
	return content
}

func formatPlatformFilePath(platformName string, output string, fileName string, outputFolder bool, outputExt string) (string) {
	fileName = formatFileName(output, fileName, outputExt)
	filePath := "./"
	if outputFolder {
		filePath += formatFilePath(platformName, output)
		filePath = formatFilePath(filePath, fileName)
	} else {
		filePath += formatFilePath(platformName, fileName)
	}
	return filePath
}

func formatPlatformScriptFilePath(platformName string, output string, fileName string, outputFolder bool, outputExt string, clsFolder string) (string) {
	fileName = formatFileName(output, fileName, outputExt)
	filePath := "./"
	if outputFolder {
		filePath += formatFilePath(platformName, output)
		filePath = formatFilePath(filePath, clsFolder)
		filePath = formatFilePath(filePath, fileName)
	} else {
		filePath += formatFilePath(platformName, fileName)
	}
	return filePath
}

func formatFilePath(folder string, file string) (string) {
	filePath := folder
	if file != "" {
		len := len(folder)
		if folder[len-1:len] != "/" {
			if file[0:1] != "/" {
				filePath += "/"
			}
		}
		filePath += file
	}
	return filePath
}

func formatFileName(outputType string, fileName string, outputExt string) (string) {
	if fileName != "" {
		switch(outputType) {
			case "binary":
				if outputExt != "" {
					fileName += "." + outputExt
				} else {
					fileName += ".bytes"
				}
			case "json":
				if outputExt != "" {
					fileName += "." + outputExt
				} else {
					fileName += ".json"
				}
			case "xml": fileName += ".xml"
			case "typescript": fileName += ".ts"
			case "javascript": fileName += ".js"
			case "csharp": fileName += ".cs"
		}
	}
	return fileName
}

func formatCamelCaseName(name string) (string) {
	newName := ""
	nameArr := strings.Split(name, "_")
	for i := range nameArr {
		nameElement := nameArr[i]
		if nameElement == "" {
			continue
		}
		newName += strings.ToUpper(nameElement[0:1]) + strings.ToLower(nameElement[1:len(nameElement)])
	}
	return newName
}

func parseDataTypes(dataType string) (string, string) {
	dataTypeLen := len(dataType)
	preType := ""
	if dataType[dataTypeLen-2:dataTypeLen] == "[]" {
		dataType = dataType[0:dataTypeLen-2]
		preType = "list"
	}
	if dataType == "float" {
		dataType = "string"
	}
	return dataType, preType
}

func getClsHeader() (string) {
	content :=  "////////////////////////////////////////////////////////////////\n"
	content += "//\n"
	content += "// 此文件由Configo自动生成\n"
	content += "// 请勿随意修改，以免造成不必要的损失\n"
	content += "//\n"
	content += "////////////////////////////////////////////////////////////////\n"
	content += "\n"
	return content
}