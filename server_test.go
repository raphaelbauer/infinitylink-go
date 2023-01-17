package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "infinitylink-go/hamcrest"
)

// ////////////////////////////////////////////
// Mock
// ////////////////////////////////////////////
type WaybackmachineMock struct {
	checkIfUrlExistsReturn bool
	checkIfUrlExistsCalled int

	checkIfUrlExistsOnWaybackmachineUrl     string
	checkIfUrlExistsOnWaybackmachineReturn1 bool
	checkIfUrlExistsOnWaybackmachineReturn2 string
	checkIfUrlExistsOnWaybackmachineCalled  int

	saveUrlOnWaybackMachineUrl    string
	saveUrlOnWaybackMachineCalled int
}

func (w *WaybackmachineMock) CheckIfUrlExists(url string) bool {
	w.checkIfUrlExistsCalled++
	return w.checkIfUrlExistsReturn
}

func (w *WaybackmachineMock) CheckIfUrlExistsOnWaybackmachine(url string) (bool, string) {
	w.checkIfUrlExistsOnWaybackmachineCalled++
	return w.checkIfUrlExistsOnWaybackmachineReturn1, w.checkIfUrlExistsOnWaybackmachineReturn2
}
func (w *WaybackmachineMock) SaveUrlOnWaybackMachine(url string) {
	w.saveUrlOnWaybackMachineCalled++
}

// end mock

func TestServerImpl(t *testing.T) {
	t.Run("works when real webpage is available", func(t *testing.T) {
		//given
		request, _ := http.NewRequest(http.MethodGet, "/https://www.google.de", nil)
		response := httptest.NewRecorder()

		waybackmachine := WaybackmachineMock{
			checkIfUrlExistsReturn: true,
		}
		serverMachineImpl := ServerMachineImpl{&waybackmachine}

		// when
		serverMachineImpl.Server(response, request)

		//then
		Expect(response.Body.String(), t).ToStartWith(`<a href="https://www.google.de">See Other</a>.`)
		Expect(response.Code, t).ToEqual(303)
		Expect(waybackmachine.checkIfUrlExistsCalled, t).ToEqual(1)
		Expect(waybackmachine.saveUrlOnWaybackMachineCalled, t).ToEqual(1)

	})

	t.Run("redirects to waybackmachine when real webpage is offline", func(t *testing.T) {
		//given
		request, _ := http.NewRequest(http.MethodGet, "/https://www.google.de", nil)
		response := httptest.NewRecorder()

		waybackmachine := WaybackmachineMock{
			checkIfUrlExistsReturn: false,

			checkIfUrlExistsOnWaybackmachineReturn1: true,
			checkIfUrlExistsOnWaybackmachineReturn2: "http://wayback.com/url",
		}
		serverMachineImpl := ServerMachineImpl{&waybackmachine}

		// when
		serverMachineImpl.Server(response, request)

		//then
		Expect(response.Body.String(), t).ToStartWith(`<a href="http://wayback.com/url">See Other</a>.`)
		Expect(response.Code, t).ToEqual(303)
		Expect(waybackmachine.saveUrlOnWaybackMachineCalled, t).ToEqual(0)
	})

	t.Run("returns 404 when site is not available and NOT available on waybackmachine", func(t *testing.T) {
		//given
		request, _ := http.NewRequest(http.MethodGet, "/https://www.google.de", nil)
		response := httptest.NewRecorder()

		waybackmachine := WaybackmachineMock{
			checkIfUrlExistsReturn: false,

			checkIfUrlExistsOnWaybackmachineReturn1: false,
			checkIfUrlExistsOnWaybackmachineReturn2: "",
		}
		serverMachineImpl := ServerMachineImpl{&waybackmachine}

		// when
		serverMachineImpl.Server(response, request)

		//then
		Expect(response.Body.String(), t).ToStartWith(`404 page not found`)
		Expect(response.Code, t).ToEqual(404)
		Expect(waybackmachine.saveUrlOnWaybackMachineCalled, t).ToEqual(0)
	})

	t.Run("works when user does not specify any path (was a bug)", func(t *testing.T) {
		//given
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		waybackmachine := WaybackmachineMock{
			checkIfUrlExistsReturn: false,

			checkIfUrlExistsOnWaybackmachineReturn1: false,
			checkIfUrlExistsOnWaybackmachineReturn2: "",
		}
		serverMachineImpl := ServerMachineImpl{&waybackmachine}

		// when
		serverMachineImpl.Server(response, request)

		//then
		Expect(response.Body.String(), t).ToStartWith(`404 page not found`)
		Expect(response.Code, t).ToEqual(404)
		Expect(waybackmachine.saveUrlOnWaybackMachineCalled, t).ToEqual(0)
	})
}
