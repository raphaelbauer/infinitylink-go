package waybackmachine

import (
	"encoding/json"
	"io/ioutil"

	"infinitylink-go/httpmockable"
)

type Waybackmachine interface {
	CheckIfUrlExists(url string) (error, bool)
	CheckIfUrlExistsOnWaybackmachine(url string) (error, bool, string)
	SaveUrlOnWaybackMachine(url string)
}

type WaybackmachineReal struct {
	Httpmockable httpmockable.HttpMockable
}

func (waybackmachineReal *WaybackmachineReal) CheckIfUrlExists(url string) (error, bool) {
	res, err := waybackmachineReal.Httpmockable.Get(url)

	if err != nil {
		return err, false
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return nil, true
	} else {
		return nil, false
	}
}

// https://archive.org/help/wayback_api.php
//{
//    "archived_snapshots": {
//        "closest": {
//            "available": true,
//            "url": "http://web.archive.org/web/20130919044612/http://example.com/",
//            "timestamp": "20130919044612",
//            "status": "200"
//        }
//    }
//}

type WaybackMachineResponseEnvelope struct {
	Archived_snapshots WaybackMachineResponseClosest `json:"archived_snapshots"`
}

type WaybackMachineResponseClosest struct {
	Closest WaybackMachineResponseSingle
}

type WaybackMachineResponseSingle struct {
	Available bool
	Url       string
	Timestamp string
	Status    string
}

// ///////////////////////////////////////////////////////////////////////////
func (waybackmachineReal *WaybackmachineReal) CheckIfUrlExistsOnWaybackmachine(url string) (error, bool, string) {
	res, err := waybackmachineReal.Httpmockable.Get("http://archive.org/wayback/available?url=" + url)
	if err != nil {
		return err, false, ""
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return err, false, ""
	}

	waybackMachineResponseEnvelope := WaybackMachineResponseEnvelope{}
	jsonErr := json.Unmarshal(body, &waybackMachineResponseEnvelope)

	if jsonErr != nil {
		return err, false, ""
	}

	return nil, waybackMachineResponseEnvelope.Archived_snapshots.Closest.Available, waybackMachineResponseEnvelope.Archived_snapshots.Closest.Url
}

func (waybackmachineReal *WaybackmachineReal) SaveUrlOnWaybackMachine(url string) {
	waybackmachineReal.Httpmockable.Get("https://web.archive.org/save/" + url)
}
