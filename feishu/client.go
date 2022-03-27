//基础飞书请求器

package feishu

import (
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func NewClient() *FeishuClient {
	return &FeishuClient{
		HttpClient: http.DefaultClient,
	}
}

type FeishuClient struct {
	HttpClient *http.Client
}

func (client *FeishuClient) Do(req *http.Request, accessToken ...string) ([]byte, error) {
	token := ""
	if len(accessToken) > 0 {
		token = accessToken[0]
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", Content_Type)
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("User-Agent", "goodsman")

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		logrus.Error("response status: %d", response.StatusCode)
	}

	return resp, nil
}
