package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Demo4() {
	client := &http.Client{}

	body := ioutil.NopCloser(strings.NewReader("page_size=50&page_num=1&session_id=9&sdate=2021-11-07 00:00:00&edate=2021-11-07 23:59:59&player_id=52428800029&Applicant=&server_ids"))
	req, err := http.NewRequest("POST", "http://10.16.168.68:7771/game/mail_list/mail_list", body)

	if err != nil {
		fmt.Println(err)
	}

	req_body, err := ioutil.ReadAll(req.Body)
	fmt.Println(string(req_body))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "sessionid=m91j9iqb43f4c2sivdzapykb11qowqiq")
	resp, err := client.Do(req)
	if resp != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Body)

}
