package models

////
type ConvertFile struct {
	Apikey       string `json:"apikey"`
	Input        string `json:"input"`
	File         string `json:"file"`
	Filename     string `json:"filename"`
	OutputFormat string `json:"outputformat"`
}

type ConvertFileResponse struct {
	Code   int                     `json:"code"`
	Status string                  `json:"status"`
	Data   ConvertFileResponseData `json:"data"`
}

type ConvertFileResponseData struct {
	Id      string `json:"id"`
	Minutes int    `json:"minutes"`
}

type ConvertFileError struct {
	Code   int    `json:"code"`
	Status string `json:"string"`
	Err    string `json:"error"`
}

/////

type ConvertFileFinal struct {
	Code   int                  `json:"code"`
	Status string               `json:"status"`
	Data   ConvertFileFinalData `json:"data"`
}

type ConvertFileFinalData struct {
	Id   string `json:"id"`
	File string `json:"file"`
	Size string `json:"size"`
}

///

type ReadyFile struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Data   ReadyFileData `json:"data"`
}

type ReadyFileData struct {
	Id          string              `json:"id"`
	Step        string              `json:"step"`
	StepPercent int                 `json:"step_percent"`
	Minutes     int                 `json:"minutes"`
	Output      ReadyFileDataOutput `json:"output"`
}

type ReadyFileDataOutput struct {
	Url  string `json:"url"`
	Size string `json:"size"`
}

///
