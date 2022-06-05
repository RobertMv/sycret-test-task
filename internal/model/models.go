package model

type Input struct {
	URLTemplate string `json:"URLTemplate"`
	RecordId    int64  `json:"RecordId"`
}

type Output struct {
	URLWord string `json:"URLWord"`
}

type ReturnedData struct {
	Result            int    `json:"result"`
	ResultDescription string `json:"resultdescription"`
	ResultData        string `json:"resultdata"`
}
