package testsupport

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/ifrit/sigmon"

	. "github.com/onsi/gomega"
)

type FakeController struct {
	ifrit.Process
	handlerLock  sync.Mutex
	handlers     map[string]*FakeHandler
	handlerFuncs map[string]FakeHandlerFunc
}

type FakeHandlerFunc func(w http.ResponseWriter, r *http.Request)

type FakeHandler struct {
	LastRequestBody []byte
	ResponseCode    int
	ResponseBody    interface{}
}

func (f *FakeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.handlerLock.Lock()
	defer f.handlerLock.Unlock()
	fakeHandlerFunc, ok := f.handlerFuncs[r.URL.Path]
	if ok {
		fakeHandlerFunc(w, r)
		return
	}

	var fakeHandler *FakeHandler
	for route, h := range f.handlers {
		if r.URL.Path == route {
			fakeHandler = h
		}
	}
	if fakeHandler == nil {
		w.WriteHeader(http.StatusTeapot)
		// #nosec G104 - ignore errors when writing HTTP responses so we don't spam our logs during a DoS
		w.Write([]byte(`{}`))
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fakeHandler.LastRequestBody = bodyBytes
	responseBytes, _ := json.Marshal(fakeHandler.ResponseBody)
	w.WriteHeader(fakeHandler.ResponseCode)
	// #nosec G104 - ignore errors when writing HTTP responses so we don't spam our logs during a DoS
	w.Write(responseBytes)
}

func (f *FakeController) SetHandler(route string, handler *FakeHandler) {
	f.handlerLock.Lock()
	defer f.handlerLock.Unlock()
	f.handlers[route] = handler
}

func (f *FakeController) SetHandlerFunc(route string, handlerFunc FakeHandlerFunc) {
	f.handlerLock.Lock()
	defer f.handlerLock.Unlock()
	f.handlerFuncs[route] = handlerFunc
}

func StartServer(serverListenAddr string, tlsConfig *tls.Config) *FakeController {
	fakeServer := &FakeController{
		handlers:     make(map[string]*FakeHandler),
		handlerFuncs: make(map[string]FakeHandlerFunc),
	}

	someServer := http_server.NewTLSServer(serverListenAddr, fakeServer, tlsConfig)

	members := grouper.Members{{
		Name:   "http_server",
		Runner: someServer,
	}}
	group := grouper.NewOrdered(os.Interrupt, members)
	monitor := ifrit.Invoke(sigmon.New(group))

	Eventually(monitor.Ready(), 60*time.Second).Should(BeClosed())
	fakeServer.Process = monitor
	return fakeServer
}

func (f *FakeController) Stop() {
	if f == nil {
		return
	}
	f.Process.Signal(os.Interrupt)
	Eventually(f.Process.Wait()).Should(Receive())
}
