package wx_cloud

import (
	"encoding/json"
	"fmt"

	"wx_cloud/entity"
)

const (
	// 学生等级信息
	studentLevalTable = "t_student_level"
	// 学生信息
	studentTable = "t_student"
)

// InsertStudentLeval 记录学生等级
func InsertStudentLeval(info *entity.TStudentLevel) error {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databaseadd?access_token=%s", GetAccessToken())
	body, _ := json.Marshal(info)
	req := map[string]interface{}{
		"env": "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"%s\").add({data:[%s]})`,
			studentLevalTable, string(body)),
	}
	rsp, _ := clientPost(url, req)
	fmt.Printf("InsertWrongInfo rsp: %v\n", string(rsp))

	return nil
}

// QueryStudentLevel 查询学生的等级
func QueryStudentLevel(school, grade, class string, student string) (*entity.TStudentLevel, error) {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databasequery?access_token=%s", GetAccessToken())
	req := map[string]interface{}{
		"env": "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"%s\").where({school:\"%s\",grade:\"%s\",class:\"%s\",student:\"%s\"}).get()`,
			studentLevalTable, school, grade, class, student),
	}
	body, _ := clientPost(url, req)

	resData := &ResData{}
	if err := json.Unmarshal(body, &resData); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	res := &entity.TStudentLevel{}
	_ = json.Unmarshal([]byte(resData.Data[0]), res)

	return res, nil
}

// QueryStudentInfo 查询学生的信息
func QueryStudentInfo(openID string) (*entity.TStudent, error) {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databasequery?access_token=%s", GetAccessToken())
	sql := fmt.Sprintf(`db.collection(\"%s\").where({
		_openid:\"%s\"
	}).get()`,
		studentTable, openID)
	fmt.Println("sql: ", sql)
	req := map[string]interface{}{
		"env":   "cloud1-4g2pzysxb452412a",
		"query": sql,
	}
	body, _ := clientPost(url, req)
	fmt.Println("bpdy: ", string(body))
	resData := &ResData{}
	if err := json.Unmarshal(body, &resData); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	res := &entity.TStudent{}
	_ = json.Unmarshal([]byte(resData.Data[0]), res)

	return res, nil
}
