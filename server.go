package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"infinitylink-go/httpmockable"

	"infinitylink-go/waybackmachine"
)

const defaultPort = "8000"

type ServerMachine interface {
	Server(w http.ResponseWriter, r *http.Request)
}

type ServerMachineImpl struct {
	Waybackmachine waybackmachine.Waybackmachine
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var HttpReal httpmockable.HttpImpl = httpmockable.HttpImpl{}
	var Waybackmachine waybackmachine.WaybackmachineReal = waybackmachine.WaybackmachineReal{Httpmockable: &HttpReal}

	serverMachineImpl := ServerMachineImpl{&Waybackmachine}
	handler := http.HandlerFunc(serverMachineImpl.Server)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func (serverMachineImpl *ServerMachineImpl) Server(w http.ResponseWriter, r *http.Request) {

	if !strings.HasPrefix(r.URL.Path, "/") {
		http.NotFound(w, r)
		return
	}

	url := strings.TrimPrefix(r.URL.Path, "/")

	urlExistsError, urlExists := serverMachineImpl.Waybackmachine.CheckIfUrlExists(url)
	if urlExistsError != nil {
		log.Fatal(urlExistsError)
		http.Error(w, "opsi", http.StatusInternalServerError)
		return
	}

	if urlExists {
		// calling this every time is overkill and a waste - but we don't have a db yet...
		// this is not high traffic after all...
		serverMachineImpl.Waybackmachine.SaveUrlOnWaybackMachine(url)
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

	waybackMachineUrlError, waybackMachineUrlExists, waybackMachineUrl := serverMachineImpl.Waybackmachine.CheckIfUrlExistsOnWaybackmachine(url)
	if waybackMachineUrlError != nil {
		log.Fatal(waybackMachineUrlError)
		http.Error(w, "opsi", http.StatusInternalServerError)
		return
	}

	if waybackMachineUrlExists {
		http.Redirect(w, r, waybackMachineUrl, http.StatusSeeOther)
		return
	}

	// url does not exist AND does not exist on waybackmachine... arg...
	http.NotFound(w, r)

}
