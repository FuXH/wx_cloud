package main

import (
	"log"
	"net/http"

	"wx_cloud/resposity/wx_cloud"
	"wx_cloud/service"
)

func main() {
	wx_cloud.InitWxCloudAPI()

	http.HandleFunc("/", service.GetIndex)
	http.HandleFunc("/parse_excel_file", service.ParseExcelFile)
	http.HandleFunc("/query_class_info", service.QueryClassInfo)
	http.HandleFunc("/record_answer", service.RecordAnswer)
	http.HandleFunc("/query_class_wrong_info", service.QueryClassWrongInfo)

	http.HandleFunc("/parse_homework", service.RecordHomework)
	http.HandleFunc("/query_homework", service.QueryHomework)

	http.HandleFunc("/parse_student_level", service.ParseStudentLevel)

	log.Fatal(http.ListenAndServe(":80", nil))
}
