package wx_cloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	accessToken = ""
	appid       = "wx1c33302ebfbc9daf"
	secret      = "191b8dc68f0197c4d2cf6fdfc4c2cd34"
)

// GetAccessToken
func GetAccessToken() string {
	fmt.Printf("GetAccessToken: %s\n", accessToken)
	return accessToken
}

// InitWxCloudAPI
func InitWxCloudAPI() {
	if err := initAccessToken(); err != nil {
		fmt.Printf("initAccessToken fail, err: %v\n", err)
		return
	}

	ticker := time.NewTicker(90 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := initAccessToken(); err != nil {
					fmt.Printf("initAccessToken fail, err: %v\n", err)
				}
			}
		}
	}()
}

func initAccessToken() error {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		appid, secret)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	rsp := &struct {
		AccessToken string `json:"access_token"`
	}{}
	if err := json.Unmarshal(body, rsp); err != nil {
		return err
	}
	accessToken = rsp.AccessToken
	return nil
}
