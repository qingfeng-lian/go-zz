package zz

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HttpGet get请求
func Get(urlStr string, requestData map[string]string, responseData interface{}) error {
	formData := url.Values{}
	for k, v := range requestData {
		formData.Add(k, v)
	}
	data := formData.Encode()
	req, err := http.NewRequest("GET", urlStr, strings.NewReader(data))
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("response StatusCode is not StatusOK")
	}

	fmt.Print("\n======http.Get======\n", string(respData))

	err = json.Unmarshal(respData, responseData)
	if err != nil {
		return err
	}
	return nil
}
