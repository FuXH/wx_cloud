package entity

type THomework struct {
	Date string `json:"date"`
	School string `json:"school"`
	Grade string `json:"grade"`
	Class string `json:"class"`
	Infos []*HomeInfo `json:"infos"`
}

type HomeInfo struct {
	Content string `json:"content"`
	Level int `json:"level"`
}