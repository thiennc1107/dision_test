package test_test

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"worker/models"
	"worker/utils"

	"github.com/go-playground/assert/v2"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	Client = &http.Client{Transport: tr}
}

func newRequest(params map[string]string) *http.Request {
	req, err := http.NewRequest("GET", "https://::1:1234/api/v1/test", nil)
	if err != nil {
		log.Fatal("unable to create request")
	}
	req.Header.Add("Accept", "application/json")
	q := req.URL.Query()
	for v, k := range params {
		q.Add(v, k)
	}
	req.URL.RawQuery = q.Encode()
	return req
}

var mockRequest = []struct {
	Name     string
	Query    map[string]string
	Status   int
	Response utils.Res
}{
	{
		Name: "Ok 1",
		Query: map[string]string{
			"a": "1",
			"b": "2",
		},
		Status: http.StatusOK,
		Response: utils.Res{
			Success: true,
			Status:  http.StatusOK,
			Data: models.Data{
				F1: 3,
				F2: 2,
				F3: 20.085536923187668,
				F4: -20.085536923187668,
			},
		},
	},
	{
		Name: "Ok 2",
		Query: map[string]string{
			"a": "1",
			"b": "2",
		},
		Status: http.StatusOK,
		Response: utils.Res{
			Success: true,
			Status:  http.StatusOK,
			Data: models.Data{
				F1: 9,
				F2: 20,
				F3: 8103.083927575383,
				F4: -8103.083927575383,
			},
		},
	},
}

func TestApi(t *testing.T) {
	t.Parallel()
	for _, tc := range mockRequest {
		t.Run(tc.Name, func(t *testing.T) {
			request := newRequest(tc.Query)
			t.Log(request.URL)
			resp, err := Client.Do(request)
			if err != nil {
				fmt.Println(err)
				t.FailNow()
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				t.FailNow()
			}
			responseData := utils.Res{}
			err = json.Unmarshal(body, &responseData)
			if err != nil {
				fmt.Println(err)
				t.FailNow()
			}
			mapData, err := json.Marshal(responseData.Data)
			if err != nil {
				fmt.Println(err)
				t.FailNow()
			}
			data := models.Data{}
			err = json.Unmarshal(mapData, &data)
			if err != nil {
				fmt.Println(err)
				t.FailNow()
			}
			responseData.Data = data
			t.Log(responseData)
			assert.Equal(t, tc.Status, resp.StatusCode)
			assert.Equal(t, responseData, tc.Response)
		})
	}
}
