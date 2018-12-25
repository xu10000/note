package help

type Result struct {
	Success bool        `json: name`
	Data    interface{} `json: data`
	Message interface{} `json: message`
}
