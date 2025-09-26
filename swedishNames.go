package swedishNames

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GetFnamn fetches data from a specific SCB API endpoint and returns the response body as a string.
func GetFnamn() string {
	resp, _ := http.Get("https://api.scb.se/OV0104/v1/doris/sv/ssd/START/BE/BE0001/BE0001G/BE0001T06AR")
	defer resp.Body.Close()
	x, _ := io.ReadAll(resp.Body)

	return string(x)
}
func GetFnamnV2(gender string) ([]string, error) {
	var nameList struct {
		Title     string `json:"title"`
		Variables []struct {
			Values []string `json:"values"`
		} `json:"variables"`
	}

	resp, err := http.Get("https://api.scb.se/OV0104/v1/doris/sv/ssd/START/BE/BE0001/BE0001G/BE0001T06AR")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	x, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(x, &nameList)
	names := nameList.Variables[0].Values

	if gender == "male" {
		var maleList []string
		for _, v := range names {
			if strings.HasPrefix(v, "1") && !strings.HasPrefix(v, "10") {
				v, _ := strings.CutPrefix(v, "1")
				//fmt.Printf("Name=%s\n", v)
				maleList = append(maleList, v)
			}
		}
		return maleList, nil
	}
	if gender == "female" {
		var femaleList []string
		for _, v := range names {
			if strings.HasPrefix(v, "2") && !strings.HasPrefix(v, "20") {
				v, _ := strings.CutPrefix(v, "2")
				//fmt.Printf("Name=%s\n", v)
				femaleList = append(femaleList, v)
			}
		}
		return femaleList, nil

	}
	return nil, errors.New(fmt.Sprintf("gender %s is not valid", gender))
}
