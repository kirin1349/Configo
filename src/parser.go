package main

import (
	"toolky"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
)

func parseExcel(path string, config *Config) (result bool, pfDataMap *(map[string]*PlatformData)) {
	result, sheetDataMap := parseSheets(path)

	if !result {
		return false, nil
	}

	platformDataMap := make(map[string]*PlatformData)

	for _, pfConfig := range config.Platforms {
		platformData := createPlatformData(sheetDataMap, &pfConfig)
		platformDataMap[platformData.Name] = platformData
	}

	return true, &platformDataMap;
}

func parseSheets(path string) (result bool, optData *(map[string]*SheetData)) {

    xlsx, err := excelize.OpenFile(path)
    if err != nil {
		toolky.PrintError("open sheet failed with error " + err.Error())
        return false, nil
	}

	sheetDataMap := make(map[string]*SheetData)

	sheetMap := xlsx.GetSheetMap()
	for _, sheetName := range sheetMap{
		rows := xlsx.GetRows(sheetName)

		sheetData := SheetData{}
		sheetData.Name = sheetName
		sheetData.Keys = make(map[string]*WordData)
		sheetData.Values = make(map[int]*ValueData)

		// rowIndex == 0 
		// rowIndex == 1 desc
		// rowIndex == 2 name
		// rowIndex == 3 data_type
		row := rows[2]
		for colIndex, colCell := range row {
			if colCell == "" {
				continue
			}
			wordData := WordData{}
			wordData.ColumnIndex = colIndex
			wordData.Name = colCell
			wordData.Desc = rows[1][colIndex]
			wordData.DataType = rows[3][colIndex]
			sheetData.Keys[wordData.Name] = &wordData
		}

		for rowIndex, row := range rows {
			if rowIndex < 4 {
				continue
			}
			valueData := ValueData{}
			valueData.RowIndex = rowIndex
			valueData.Values = make(map[string]string)
			for colIndex, colCell := range row {
				if rows[2][colIndex] == "" {
					continue
				}
				valueData.Values[rows[2][colIndex]] = colCell
			}
			sheetData.Values[valueData.RowIndex] = &valueData
		}

		sheetDataMap[sheetName] = &sheetData
	}

	return true, &sheetDataMap
}

func createPlatformData(srcMap *(map[string]*SheetData), pfConfig *PlatformConfig) (optMap *PlatformData) {
	platformData := PlatformData{}
	platformData.Name = pfConfig.Name
	platformData.AllInOne = pfConfig.AllInOne
	platformData.CreateCls = pfConfig.CreateCls
	platformData.ClsFolder = pfConfig.ClsFolder
	platformData.Output = pfConfig.Output
	platformData.OutputFolder = pfConfig.OutputFolder
	platformData.OutputExt = pfConfig.OutputExt
	platformData.Language = pfConfig.Language
	platformData.LibCls = pfConfig.LibCls
	platformData.LibPath = pfConfig.LibPath
	platformData.DataStructs = pfConfig.DataStructs
	platformData.Sheets = make(map[string]*SheetData)
	for _, srcData := range (*srcMap) {
		if !checkExtMatch(srcData.Name, pfConfig.Ext) {
			continue
		}
		platformData.Sheets[srcData.Name] = filterSheet(srcData, pfConfig.Ext)
	}
	return &platformData
}

func filterSheet(srcData *SheetData, ext string) (optData *SheetData) {
	if ext == "" {
		return srcData
	}
	sheetData := SheetData{}
	sheetData.Name = getSheetName(srcData.Name);
	sheetData.Keys = make(map[string]*WordData)
	sheetData.KeyIndexes = make(map[int]*WordData)
	sheetData.Values = make(map[int]*ValueData)
	for _, keyData := range srcData.Keys {
		if !checkExtMatch(keyData.Name, ext) {
			continue
		}
		wordData := WordData{}
		wordData.ColumnIndex = keyData.ColumnIndex
		wordData.Name = keyData.Name
		wordData.Desc = keyData.Desc
		wordData.DataType = keyData.DataType
		sheetData.Keys[keyData.Name] = &wordData
		sheetData.KeyIndexes[keyData.ColumnIndex] = &wordData
	}
	for _, vData := range srcData.Values {
		valueData := ValueData{}
		valueData.RowIndex = vData.RowIndex
		valueData.Values = make(map[string]string)
		for vKey, vValue := range vData.Values {
			if !checkExtMatch(vKey, ext) {
				continue
			}
			valueData.Values[vKey] = vValue
		}
		sheetData.Values[valueData.RowIndex] = &valueData
	}
	return &sheetData
}

func checkExtMatch(txt string, ext string) (result bool) {
	txtLen := len(txt)
	if txtLen <= 2 {
		return true
	}
	lastTxt := txt[txtLen-2:txtLen]
	if len(lastTxt) < 2 {
		return true
	}
	tmpTxt := lastTxt[0:1]
	if tmpTxt != "_" {
		return true
	}
	tmpTxt = lastTxt[1:2]
	if strings.ToUpper(tmpTxt) == tmpTxt {
		if tmpTxt != ext {
			return false
		}
	}
	return true
}

func getSheetName(txt string) (result string) {
	txtLen := len(txt)
	if txtLen <= 2 {
		return txt
	}
	lastTxt := txt[txtLen-2:txtLen]
	if len(lastTxt) < 2 {
		return txt
	}
	tmpTxt := lastTxt[0:1]
	if tmpTxt != "_" {
		return txt
	}
	tmpTxt = lastTxt[1:2]
	if strings.ToUpper(tmpTxt) == tmpTxt {
		return txt[0:len(txt)-2]
	}
	return txt
}