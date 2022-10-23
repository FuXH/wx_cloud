package wx_cloud

import (
	"encoding/json"
	"fmt"
	"wx_cloud/entity"
)

const (
	// 作业信息
	homeworkTable = "t_homework"
)

// InsertHomework 记录布置的作业信息
func InsertHomework(homework *entity.THomework) error {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databaseadd?access_token=%s", GetAccessToken())
	body, _ := json.Marshal(homework)
	sql := fmt.Sprintf(`db.collection(\"%s\").add({data:[%s]})`,
		homeworkTable, string(body))
	fmt.Println("sql: ", sql)
	req := map[string]interface{}{
		"env":   "cloud1-4g2pzysxb452412a",
		"query": sql,
	}
	rsp, err := clientPost(url, req)
	fmt.Printf("InsertWrongInfo rsp: %v, err: %v\n", string(rsp), err)

	return nil
}

// QueryHomework 查询学生的作业列表
func QueryHomework(openID string) ([]*entity.THomework, error) {
	// 查询学生信息
	studentInfo, _ := QueryStudentInfo(openID)

	// 查询作业列表
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databasequery?access_token=%s", GetAccessToken())
	req := map[string]interface{}{
		"env": "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"%s\").where({school:\"%s\",grade:\"%s\",class:\"%s\"}).get()`,
			homeworkTable, studentInfo.School, studentInfo.Grade, studentInfo.Class),
	}
	body, _ := clientPost(url, req)

	resData := &ResData{}
	if err := json.Unmarshal(body, &resData); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	res := make([]*entity.THomework, len(resData.Data))
	for i, val := range resData.Data {
		tmpHomework := &entity.THomework{}
		_ = json.Unmarshal([]byte(val), tmpHomework)
		res[i] = tmpHomework
	}

	// 查询学生等级
	levelInfo, _ := QueryStudentLevel(studentInfo.School, studentInfo.Grade, studentInfo.Class, studentInfo.StudentName)

	for _, val := range res {
		var realHomework *entity.HomeInfo
		for _, tmp := range val.Datas {
			if levelInfo.Level == tmp.Level {
				realHomework = tmp
			}
		}
		val.Datas = []*entity.HomeInfo{realHomework}
	}

	return res, nil
}
