package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func DownLoadWiki(docToken,name,url,Doctype,usertoken,savepath string) {
	TICKET,msg,code := CreateTask(docToken,Doctype,usertoken)
	if code != 0 {
		fmt.Println(msg,name,docToken,url)
		return
	}
	//fmt.Println(TICKET,docToken)
	var fileToken string
	for {
		fileToken = SearchTask(TICKET,docToken,usertoken)
		//fmt.Println(fileToken)
		if fileToken != "" {
			break
		}
	}
	//fmt.Println(fileToken)
	DownloadTask(fileToken,name,Doctype,usertoken,savepath)
}

type CreateTaskRes struct {
	Code 	int		`json:"code"`
	Msg		string	`json:"msg"`
	Data	CreateTaskResdata
}

type CreateTaskResdata struct {
	Ticket	string	`json:"ticket"`
}

func CreateTask(docToken,doctype,usertoken string) (ticket,msg string,code int) {
	url := "https://open.feishu.cn/open-apis/drive/v1/export_tasks"
	client := &http.Client{}
	var str string

	if doctype == "doc" || doctype == "wiki" || doctype == "docx"  {
		str = "{\"file_extension\": \"docx\",\"token\":\"" + docToken + "\",\"type\": \"" + doctype + "\"}"
	}
	if doctype == "sheet" || doctype == "bitable"{
		str = "{\"file_extension\": \"xlsx\",\"token\":\"" + docToken + "\",\"type\": \"" + doctype + "\"}"
	}
	if doctype == "pdf" || doctype == "file" {
		str = "{\"file_extension\": \"pdf\",\"token\":\"" + docToken + "\",\"type\": \"docx\"}"
	}

	payload := strings.NewReader(str)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", "Bearer " + usertoken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var ctr CreateTaskRes
	err = json.Unmarshal(body, &ctr)
	if err != nil {
		fmt.Println(err)
	}

	return ctr.Data.Ticket,ctr.Msg,ctr.Code
}

type SearchTaskRes struct {
	Code 	int		`json:"code"`
	Msg		string	`json:"msg"`
	Data	SearchTaskResdata
}

type SearchTaskResdata struct {
	Result	ExportTask
}

type ExportTask struct {
	File_Extension	string	`json:"file_extension"`
	Type			string	`json:"type"`
	File_Name		string	`json:"file_name"`
	File_Token		string	`json:"file_token"`
	File_Size		int		`json:"file_size"`
	Job_Error_Msg	string	`json:"job_error_msg"`
	Job_Status		int		`json:"job_status"`
	//Extra			string	`json:"extra"`
}

func SearchTask(ticket,docToken,usertoken string) (fileToken string){
	url := "https://open.feishu.cn/open-apis/drive/v1/export_tasks/" + ticket + "?token=" + docToken
	//fmt.Println(url)
	client := &http.Client {}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer " + usertoken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))

	var str SearchTaskRes
	err = json.Unmarshal(body, &str)
	if err != nil {
		fmt.Println(err)
	}

	fileToken = str.Data.Result.File_Token
	return
}

func DownloadTask(fileToken,name,doctype,usertoken,savepath string) {
	url := "https://open.feishu.cn/open-apis/drive/v1/export_tasks/file/" + fileToken + "/download"

	client := &http.Client {}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer " + usertoken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if doctype == "doc" || doctype == "wiki" || doctype == "docx"  {
		ioutil.WriteFile(savepath + "/" + name + ".docx",body,os.ModePerm)
	}
	if doctype == "sheet" || doctype == "bitable" {
		ioutil.WriteFile(savepath + "/" + name + ".xlsx",body,os.ModePerm)
	}
	if doctype == "file" {
		ioutil.WriteFile(savepath + "/" + name + ".pdf",body,os.ModePerm)
	}
}