package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type GetWikiMetaRes struct {
	Code 	int		`json:"code"`
	Msg		string	`json:"msg"`
	Data	GetWikiMetaResDatas
}

type GetWikiMetaResDatas struct {
	Metas		[]GetWikiMetaResDatasMeta			`json:"metas"`
	Failed_List []GetWikiMetaResDatasFailedList		`json:"failed_list"`

}

type GetWikiMetaResDatasMeta struct {
	Doc_token	string	`json:"doc_token"`
	Doc_type	string	`json:"doc_type"`
	Title		string	`json:"title"`
	Owner_id	string	`json:"owner_id"`
	Create_time	string	`json:"create_time"`
	Latest_modify_user	string	`json:"latest_modify_user"`
	Latest_modify_time	string	`json:"latest_modify_time"`
	Url					string	`json:"url"`
}

type GetWikiMetaResDatasFailedList struct {
	Token	string	`json:"token"`
	Code	int		`json:"code"`
}

func GetWikiMeta(wikitoken,usertoken,savepath string) {
	url := "https://open.feishu.cn/open-apis/drive/v1/metas/batch_query?user_id_type=user_id"
	client := &http.Client{}

	str := "{\"request_docs\": [{\"doc_token\": \"" + wikitoken + "\",\"doc_type\": \"wiki\"}],\"with_url\": true}"
	payload := strings.NewReader(str)

	req, err := http.NewRequest("POST", url,  payload)
	if err != nil {
		fmt.Printf("build request error: %s", err.Error())
		return
	}
	req.Header.Set("Authorization", "Bearer " + usertoken)
	req.Header.Set("Content-Type","application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get msg error: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body error: %s", err.Error())
	}

	var gwm GetWikiMetaRes
	err = json.Unmarshal(body, &gwm)
	if err != nil {
		fmt.Println("Unmarshal body error：%s", err.Error())
	}

	for _,val := range gwm.Data.Metas {
		DownLoadWiki(val.Doc_token,val.Title,val.Url,val.Doc_type,usertoken,savepath)
	}
}
