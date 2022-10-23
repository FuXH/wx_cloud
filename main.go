package main

import (
	"log"
	"net/http"

	"wx_cloud/resposity/wx_cloud"
	"wx_cloud/service"
)

func main() {
	wx_cloud.InitWxCloudAPI()

	//tmp, _ := wx_cloud.QueryClassList("龙华区实验学校", "五年级", "四班")
	//for _, val := range tmp {
	//	fmt.Println(val)
	//}

	//data := &entity.TClass{
	//	ClassID:   "3",
	//	ClassName: "数学课第一单元",
	//	School:    "龙华区实验学校",
	//	Grade:     "五年级",
	//	Class:     "四班",
	//	Teacher:   "方正",
	//}
	//err := wx_cloud.InsertClass(data)
	//fmt.Println(err)

	//tclass, tclassInfo, err := parse.ParseExcelFile("test_01.xlsx")
	//fmt.Println(tclass)
	//fmt.Println(tclassInfo.ClassID)
	//for _, val := range tclassInfo.Questions {
	//	fmt.Println(val)
	//}
	//fmt.Println(err)

	http.HandleFunc("/", service.GetIndex)
	http.HandleFunc("/parse_excel_file", service.ParseExcelFile)
	http.HandleFunc("/query_class_info", service.QueryClassInfo)
	http.HandleFunc("/record_answer", service.RecordAnswer)
	http.HandleFunc("/query_class_wrong_info", service.QueryClassWrongInfo)
	//
	log.Fatal(http.ListenAndServe(":80", nil))
}
