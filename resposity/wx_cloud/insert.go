package wx_cloud

import (
	"encoding/json"
	"fmt"

	"wx_cloud/entity"
)

// InsertClass 记录课堂信息
func InsertClass(class *entity.TClass) error {
	url := fmt.Sprintf("https://api.weixin.qq.com/tcb/databaseadd?access_token=%s", GetAccessToken())
	body, _ := json.Marshal(class)
	req := map[string]interface{}{
		"env":   "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"t_class\").add({data:[%s]})`, string(body)),
	}
	rsp, _ := clientPost(url, req)

	resData := &ResData{}
	if err := json.Unmarshal(rsp, &resData); err != nil {
		fmt.Println("err: ", err)
		return err
	}
	if resData.Errcode != 0 {
		err := fmt.Errorf("InsertClass fail, rsp: %v", resData)
		return err
	}

	return nil
}

// InsertClassInfo 记录课堂题目
func InsertClassInfo(classInfo *entity.TClassInfo) error {
	url := fmt.Sprintf("https://api.weixin.qq.com/tcb/databaseadd?access_token=%s", GetAccessToken())
	questions, _ := json.Marshal(classInfo.Questions)
	tmp := &struct {
		ClassID   string `json:"classId"`
		Questions string `json:"questions"`
	}{
		ClassID:   classInfo.ClassID,
		Questions: string(questions),
	}
	body, _ := json.Marshal(tmp)
	req := map[string]interface{}{
		"env":   "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"t_class\").add({data:[%s]})`, string(body)),
	}
	rsp, _ := clientPost(url, req)

	resData := &ResData{}
	if err := json.Unmarshal(rsp, &resData); err != nil {
		fmt.Println("err: ", err)
		return err
	}
	if resData.Errcode != 0 {
		err := fmt.Errorf("InsertClass fail, rsp: %v", resData)
		return err
	}

	return nil
}
