package model

// Input struct represents an input json parameters in HTTP request
type Input struct {
	URLTemplate string `json:"URLTemplate"`
	RecordId    int64  `json:"RecordId"`
}

// Output struct is a representation of a result path for created document
type Output struct {
	URLWord string `json:"URLWord"`
}

// ReturnedData is used to unmarshal json response from sycret API
type ReturnedData struct {
	Result            int    `json:"result"`
	ResultDescription string `json:"resultdescription"`
	ResultData        string `json:"resultdata"`
}
