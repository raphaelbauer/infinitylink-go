package waybackmachine

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	. "infinitylink-go/hamcrest"
)

// ////////////////////////////////////////////
// Mock
// ////////////////////////////////////////////

type HttpMock struct {
	httpStatusCode int
	body           string
	GetUrl         string
}

func (h *HttpMock) Get(url string) (resp *http.Response, err error) {
	h.GetUrl = url
	r := ioutil.NopCloser(bytes.NewReader([]byte(h.body)))

	return &http.Response{
		StatusCode: h.httpStatusCode,
		Body:       r,
	}, nil
}

// end mock

func TestServerImpl(t *testing.T) {
	t.Run("CheckIfUrlExistsOnWaybackmachine works when real webpage is available", func(t *testing.T) {
		// given
		json := `{
			    "archived_snapshots": {
			        "closest": {
			            "available": true,
			            "url": "http://web.archive.org/web/20130919044612/http://example.com/",
			            "timestamp": "20130919044612",
			            "status": "200"
			        }
			    }
			}`
		httpMock := HttpMock{
			200,
			json,
			"",
		}

		waybackmachineReal := WaybackmachineReal{&httpMock}

		// when
		var error, exists, url = waybackmachineReal.CheckIfUrlExistsOnWaybackmachine("http://myurl.com")

		// then
		Expect(error, t).ToEqual(nil)
		Expect(exists, t).ToEqual(true)
		Expect(url, t).ToEqual("http://web.archive.org/web/20130919044612/http://example.com/")
	})

	t.Run("CheckIfUrlExistsOnWaybackmachine works when real webpage is NOT available", func(t *testing.T) {
		// given
		json := `{
			    "archived_snapshots": {
			        "closest": {
			            "available": false
			        }
			    }
			}`
		httpMock := HttpMock{
			200,
			json,
			"",
		}

		waybackmachineReal := WaybackmachineReal{&httpMock}

		// when
		var error, exists, url = waybackmachineReal.CheckIfUrlExistsOnWaybackmachine("http://myurl.com")

		// then
		Expect(error, t).ToEqual(nil)
		Expect(exists, t).ToEqual(false)
		Expect(url, t).ToEqual("")
	})

	t.Run("CheckIfUrlExists works when real webpage is available", func(t *testing.T) {
		// given
		httpMock := HttpMock{
			200,
			"",
			"",
		}
		waybackmachineReal := WaybackmachineReal{&httpMock}

		// when
		var error, exists = waybackmachineReal.CheckIfUrlExists("http://myurl.com")

		// then
		Expect(error, t).ToEqual(nil)
		Expect(exists, t).ToEqual(true)
	})

	t.Run("CheckIfUrlExists works when real webpage is NOT available", func(t *testing.T) {
		// given
		httpMock := HttpMock{
			404,
			"",
			"",
		}
		waybackmachineReal := WaybackmachineReal{&httpMock}

		// when
		var error, exists = waybackmachineReal.CheckIfUrlExists("http://myurl.com")

		// then
		Expect(error, t).ToEqual(nil)
		Expect(exists, t).ToEqual(false)
	})

	t.Run("SaveUrlOnWaybackMachine works", func(t *testing.T) {
		// given
		httpMock := HttpMock{
			200,
			"",
			"",
		}
		waybackmachineReal := WaybackmachineReal{&httpMock}

		// when
		waybackmachineReal.SaveUrlOnWaybackMachine("http://myurl.com")

		// then
		Expect(httpMock.GetUrl, t).ToEqual("https://web.archive.org/save/http://myurl.com")
	})
}
