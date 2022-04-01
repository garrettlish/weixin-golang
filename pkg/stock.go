package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	url = "http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData?symbol=%s&scale=%d&ma=no&datalen=1"

	todayData   = make(map[string]map[string]float64)
	curData     = make(map[string]float64)
	lastUpdated time.Time
)

func fetchPrice(id string, scale int) (string, string, string, error) {
	resp, err := http.Get(fmt.Sprintf(url, id, scale))
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", err
	}
	if resp.StatusCode >= 300 {
		return "", "", "", fmt.Errorf("invalid request (%d): %s", resp.StatusCode, string(b))
	}
	var result = make([]interface{}, 0)
	err = json.Unmarshal(b, &result)
	if err != nil {
		return "", "", "", err
	}
	if len(result) == 0 {
		return "", "", "", fmt.Errorf("empty result")
	}
	prices := result[0].(map[string]interface{})
	return prices["day"].(string), prices["open"].(string), prices["close"].(string), nil
}

func FetchLastPrice(ids ...string) ([]string, error) {
	var (
		result  []string
		today   = Today()
		updated = false
	)
	for _, id := range ids {
		var v float64
		_, ok := todayData[today]
		if !ok {
			todayData[today] = make(map[string]float64)
		}
		if _, ok := todayData[today][id]; !ok {
			_, _, c, err := fetchPrice(id, 240)
			if err != nil {
				return nil, err
			}
			v, _ = strconv.ParseFloat(c, 64)
			todayData[today][id] = v
		}

		_, ok = curData[id]
		if !ok || time.Now().Sub(lastUpdated).Seconds() > 60 {
			_, _, c, err := fetchPrice(id, 5)
			if err != nil {
				return nil, err
			}
			curData[id], _ = strconv.ParseFloat(c, 64)
			updated = true
		}
		result = append(result, fmt.Sprintf("%.2f, %.2f (%.2f%%)", v, curData[id], (curData[id]-v)/v*100))
	}
	if updated {
		lastUpdated = time.Now()
	}
	return result, nil
}

func Today() string {
	now := time.Now()
	return fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
}
