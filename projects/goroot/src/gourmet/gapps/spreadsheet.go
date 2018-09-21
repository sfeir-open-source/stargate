/*
 Copyright (c) 2018, SFEIR OSPO <ospo@sfeir.com>

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package apps

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"

	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Sheet struct {
	Version string `json:"version"`
	ReqId   string `json:"reqId"`
	Status  string `json:"status"`
	Sig     string `json:"sig"`
	Table   Table  `json:"table"`
}

type Table struct {
	Cols []Col `json:"cols"`
	Rows []Row `json:"rows"`
}

type Col struct {
	Id    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
}

type Row struct {
	C []C `json:"c"`
}

type C struct {
	V string `json:"v"`
}

var (
	// Can't load Spreadsheet
	ErrBadUrl = errors.New("Bad Url")

	// ErrMovedPermanently is returned when a 301/302 is returned.
	ErrBadBody = errors.New("Bad Body")
)

func GetSpreadsheet(c context.Context, spreadsheetsID string, sheetID string) (*Sheet, error) {
	var result *Sheet
	transport := &oauth2.Transport{
		Source: google.AppEngineTokenSource(c, "https://spreadsheets.google.com/feeds"),
		Base:   &urlfetch.Transport{Context: c},
	}

	client := &http.Client{Transport: transport}

	url := "https://docs.google.com/spreadsheets/d/" + spreadsheetsID + "/gviz/tq?gid=" + sheetID

	resp, err := client.Get(url)
	if err != nil {
		//		http.Error(w, err.Error(), http.StatusInternalServerError)
		return result, ErrBadUrl
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		//		http.Error(w, err.Error(), http.StatusInternalServerError)
		return result, ErrBadBody
	}

	jsonString := strings.Replace(string(contents), "/*O_o*/", "", -1)
	jsonString = strings.Replace(string(jsonString), "google.visualization.Query.setResponse(", "", -1)
	jsonString = strings.Replace(jsonString, ");", "", -1)

	result = &Sheet{}
	json.Unmarshal([]byte(jsonString), &result)

	return result, nil

}
