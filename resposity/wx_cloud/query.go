package wx_cloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"wx_cloud/entity"
)

// ResData
type ResData struct {
	Errcode int      `json:"errcode"`
	Errmsg  string   `json:"errmsg"`
	Data    []string `json:"data"`
}

// QueryClassList 查询课堂信息
func QueryClassList(school, grade, class string) ([]*entity.TClass, error) {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databasequery?access_token=%s", GetAccessToken())
	req := map[string]interface{}{
		"env": "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"t_class\").where({school:\"%s\",grade:\"%s\",class:\"%s\"}).get()`,
			school, grade, class),
	}
	body, _ := clientPost(url, req)

	resData := &ResData{}
	if err := json.Unmarshal(body, &resData); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	res := make([]*entity.TClass, len(resData.Data))
	for i, val := range resData.Data {
		tmpClass := &entity.TClass{}
		_ = json.Unmarshal([]byte(val), tmpClass)
		res[i] = tmpClass
	}

	return res, nil
}

// QueryClassInfo 查询课堂题目
func QueryClassInfo(classID string) (*entity.TClassInfo, error) {
	url := fmt.Sprintf("http://api.weixin.qq.com/tcb/databasequery?access_token=%s", GetAccessToken())
	req := map[string]interface{}{
		"env":   "cloud1-4g2pzysxb452412a",
		"query": fmt.Sprintf(`db.collection(\"t_class_info\").where({classId:\"%s\"}).get()`, classID),
	}
	body, _ := clientPost(url, req)

	resData := &ResData{}
	if err := json.Unmarshal(body, &resData); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	if len(resData.Data) == 0 {
		return nil, nil
	}

	fmt.Println("resData: ", resData.Data[0])
	tmp := &struct {
		ClassID   string `json:"classId"`
		Questions string `json:"questions"`
	}{}
	_ = json.Unmarshal([]byte(resData.Data[0]), tmp)
	res := &entity.TClassInfo{
		ClassID:   classID,
		Questions: []*entity.QuestionInfo{},
	}
	_ = json.Unmarshal([]byte(tmp.Questions), &res.Questions)

	return res, nil
}

func clientPost(url string, req map[string]interface{}) ([]byte, error) {
	client := &http.Client{}

	respdata, _ := json.Marshal(req)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(respdata))
	resp, err := client.Do(request)
	fmt.Println("err: ", err)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("err: ", err)

	return body, nil
}
