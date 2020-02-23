package main

import (
	"strings"
	"toolky"
	"os"
	"strconv"
)

func encodeToBinary(pfData *PlatformData, fileName string) {
	if pfData.AllInOne {
		toolky.PrintError("binary do not support all in one yet");
	} else {
		for _, sheetData := range pfData.Sheets {
			if len(sheetData.Values) > 0 {
				filePath := formatPlatformFilePath(pfData.Name, pfData.Output, sheetData.Name, pfData.OutputFolder, pfData.OutputExt)
				err := toolky.RemoveFile(filePath)
				if err != nil {
					toolky.PrintError("remove file path " + filePath + " with error " + err.Error())
					continue
				}
				file, err := os.Create(filePath)
				if err != nil{
					toolky.PrintError("create file " + filePath + " failed with error " + err.Error())
					continue
				}
				for _, valueData := range sheetData.Values {
					index := 0
					for true {
						keyData := sheetData.KeyIndexes[index]
						if keyData == nil {
							break
						}
						encodeValue(file, keyData, valueData.Values[keyData.Name])
						index++
					}
				}
				file.Close()
			} else {
				toolky.PrintInfo("skip empty sheet " + sheetData.Name)
			}
		}
	}
}

func encodeValue(file *os.File, keyData *WordData, value string) {
	mainType, preType := parseDataTypes(keyData.DataType)
	if preType != "" {
		if preType != "list" {
			toolky.PrintError("unsupport pre type " + preType + " [" + keyData.Name + ", " + strconv.Itoa(keyData.ColumnIndex) + "]")
			return
		}
		if len(value) == 0 {
			file.Write(toolky.IntToLittleEndianBytes(0))
		} else {
			valueArr := strings.Split(value, ",")
			valueLen := len(valueArr)
			file.Write(toolky.IntToLittleEndianBytes(valueLen))
			if valueLen > 0 {
				i := 0
				for i < valueLen {
					encodeMainValue(file, keyData, mainType, valueArr[i])
					i++
				}
			}
		}
	} else {
		encodeMainValue(file, keyData, mainType, value)
	}
}

func encodeMainValue(file *os.File, keyData *WordData, mainType string, value string) {
	switch(mainType) {
		case "int":
			writeIntToFile(file, value, keyData.Name, keyData.ColumnIndex)
		case "int2":
			intValues := strings.Split(value, "|")
			writeIntsToFile(file, intValues, 2, keyData.Name, keyData.ColumnIndex)
		case "int3":
			intValues := strings.Split(value, "|")
			writeIntsToFile(file, intValues, 3, keyData.Name, keyData.ColumnIndex)
		case "int4":
			intValues := strings.Split(value, "|")
			writeIntsToFile(file, intValues, 4, keyData.Name, keyData.ColumnIndex)
		case "int5":
			intValues := strings.Split(value, "|")
			writeIntsToFile(file, intValues, 5, keyData.Name, keyData.ColumnIndex)
		default:
			ascValue := strconv.QuoteToASCII(value)
			ascValue = strings.Replace(ascValue, "\"", "", -1)
			bytes := []byte(ascValue)
			file.Write(toolky.IntToLittleEndianBytes(len(bytes)))
			file.Write(bytes)
	}
}

func writeIntToFile(file *os.File, intStr string, keyName string, columnIndex int) {
	if intStr == "" {
		file.Write(toolky.IntToLittleEndianBytes(0))
		return
	}
	intValue, err := strconv.Atoi(intStr)
	if err != nil {
		toolky.PrintError("parse value [" + keyName + ", " + strconv.Itoa(columnIndex) + "] with error " + err.Error() + " and write 0 instand of")
		file.Write(toolky.IntToLittleEndianBytes(0))
		return
	}
	file.Write(toolky.IntToLittleEndianBytes(intValue))
}

func writeIntsToFile(file *os.File, intStrList []string, destLen int, keyName string, columnIndex int) {
	i := 0
	strLen := len(intStrList)
	for i < destLen {
		if i < strLen {
			writeIntToFile(file, intStrList[i], keyName, columnIndex)
		} else {
			file.Write(toolky.IntToLittleEndianBytes(0))
		}
		i++
	}
}