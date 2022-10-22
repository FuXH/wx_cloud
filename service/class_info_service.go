package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wx_cloud/parse"
	"wx_cloud/resposity/wx_cloud"
)

// ParseExcelFile 接收并解析excel文件
func ParseExcelFile(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fileName := r.FormValue("filename")

	// 解析excel
	tclass, tclassInfo, err := parse.ParseExcelFile(fileName)
	if err != nil {
		fmt.Printf("InsertClass fail, err: %v\n", err)
		return
	}

	// 记录课堂信息
	if err := wx_cloud.InsertClass(tclass); err != nil {
		fmt.Printf("InsertClass fail, err: %v\n", err)
		return
	}

	// 记录题目列表
	if err := wx_cloud.InsertClassInfo(tclassInfo); err != nil {
		fmt.Printf("InsertClass fail, err: %v\n", err)
		return
	}
}

// QueryClassInfo 根据classId获取课堂题目
func QueryClassInfo(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	classID := r.FormValue("classId")

	tclassInfo, err := wx_cloud.QueryClassInfo(classID)
	if err != nil {
		fmt.Printf("QueryClassInfo fail, err: %v\n", err)
		return
	}

	rsp, _ := json.Marshal(tclassInfo)
	_, _ = w.Write(rsp)
}

// JudgeAnswer 判断得分，并记录错题
func JudgeAnswer(w http.ResponseWriter, r *http.Request) {

}

// QueryClassWrongInfo 根据classId和openId获取课堂的错题
func QueryClassWrongInfo(w http.ResponseWriter, r *http.Request) {

}
