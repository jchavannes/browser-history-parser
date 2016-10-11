package bhp

import (
	"time"
	"io/ioutil"
	"encoding/json"
)

type BrowserHistory struct {
	Favicon  string `json:"favicon_url"`
	Title    string
	Url      string
	TimeUsec int64 `json:"time_usec"`
}

func (b BrowserHistory) Time() time.Time {
	return time.Unix(0, b.TimeUsec * int64(time.Microsecond))
}

type BrowserHistoryJsonData struct {
	BrowserHistory []BrowserHistory `json:"Browser History"`
}

func LoadFromFile(filename string) ([]BrowserHistory, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jsonData BrowserHistoryJsonData
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData.BrowserHistory, nil
}
