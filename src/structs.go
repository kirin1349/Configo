package main

type Config struct {
    Platforms []PlatformConfig
}

type PlatformConfig struct {
	Name string
	Ext string
	Output string
	OutputFolder bool
	OutputExt string
	AllInOne bool
	CreateCls bool
	ClsFolder string
	Language string
	LibCls string
	LibPath string
	DataStructs map[string]*PlatformStructConfig
}

type PlatformStructConfig struct {
	LibCls string
	LibPath string
}

type PlatformData struct {
	Name string
	Output string
	OutputFolder bool
	OutputExt string
	Sheets map[string]*SheetData
	AllInOne bool
	CreateCls bool
	ClsFolder string
	Language string
	LibCls string
	LibPath string
	DataStructs map[string]*PlatformStructConfig
}

type SheetData struct {
	Name string
	Keys map[string]*WordData
	KeyIndexes map[int]*WordData
	Values map[int]*ValueData
}

type WordData struct {
	Name string
	Desc string
	DataType string
	ColumnIndex int
}

type ValueData struct {
	RowIndex int
	Values map[string]string
}