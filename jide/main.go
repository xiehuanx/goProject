package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Data struct {
	ASSIGNEE_     string `json:"ASSIGNEE_"`
	BIZID__       string `json:"BIZID__"`
	BJ            string `json:"BJ"`
	CREATOR__     string `json:"CREATOR__"`
	FDY           string `json:"FDY"`
	ID__          string `json:"ID__"`
	LXFS          string `json:"LXFS"`
	PID__         string `json:"PID__"`
	PROC_INST_ID_ string `json:"PROC_INST_ID_"`
	QJJSSJ        string `json:"QJJSSJ"`
	QJSY          string `json:"QJSY"`
	QSH           string `json:"QSH"`
	RN            int    `json:"RN"`
	SZSS          string `json:"SZSS"`
	TASK_ID_      string `json:"TASK_ID_"`
	TASK_NAME_    string `json:"TASK_NAME_"`
	XH            string `json:"XH"`
	XM            string `json:"XM"`
	XY            string `json:"XY"`
}

type Response struct {
	PageIndex  int    `json:"pageIndex"`
	PageSize   int    `json:"pageSize"`
	TotalCount int    `json:"totalCount"`
	Items      []Data `json:"data"`
}

func main() {
	fmt.Println("Hello World")
	formData := url.Values{}
	formData.Add("page", "1")
	formData.Add("limit", "30")
	//tbName: XSQJSQ
	//colNames: XH,XM,XY,BJ,FDY,SZSS,QSH,LXFS,QJKSSJ,QJJSSJ,QJSY,ID__,TASK_NAME_,PROC_INST_ID_,TASK_ID_,ASSIGNEE_,BIZID__,CREATOR__,PID__
	//userId: 01317
	//orgId: 01108
	//formId: e8346d3bd76548b6b903c2bcfcb18d45
	//tabId: _db_krl_
	//orderCol: []
	//	_search: false
	//	nd: 1712804087997
	formData.Add("tbName", "XSQJSQ")
	formData.Add("colNames", "XH,XM,XY,BJ,FDY,SZSS,QSH,LXFS,QJKSSJ,QJJSSJ,QJSY,ID__,TASK_NAME_,PROC_INST_ID_,TASK_ID_,ASSIGNEE_,BIZID__,CREATOR__,PID__")
	formData.Add("userId", "01317")
	formData.Add("orgId", "01108")
	formData.Add("formId", "e8346d3bd76548b6b903c2bcfcb18d45")
	formData.Add("tabId", "_db_krl_")
	formData.Add("orderCol", "[]")
	formData.Add("_search", "false")
	formData.Add("nd", "1712804087997")

	req, err := http.NewRequest("POST", "http://xs.cqie.edu.cn/xg/formQueryModel/loadFromData?random=0.6026306205852199", strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "JSESSIONID=841F1BF9D273C3DF28D0116CA04F85F6")
	//Accept:
	//
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 处理响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 使用 response 对象
	fmt.Println(response.PageIndex)
	fmt.Println(response.PageSize)
	fmt.Println(response.Items)
}
