package entity

// TClassInfo
type TClassInfo struct {
	ClassID   string          `json:"classId"`
	Questions []*QuestionInfo `json:"questions"`
}

type QuestionInfo struct {
	QuestionID int           `json:"questionId"`
	Subject    string        `json:"subject"`
	Options    []*OptionInfo `json:"options"`
	Right      string        `json:"right"`
}

type OptionInfo struct {
	Value   string `json:"value"`
	Name    string `json:"name"`
	Checked string `json:"checked"`
}
