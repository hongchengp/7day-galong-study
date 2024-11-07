package geecache

import (
	"log"
	"net/http"
	"strings"
	"fmt"
)

const defaultBasePath = "/_geecache/"

type HTTPPool struct {
	self     string
	basePath string
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (h *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", h.self, fmt.Sprintf(format, v...))
}

func (h *HTTPPool) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if !strings.HasPrefix(req.URL.Path, h.basePath) {
		panic("HTTPPool serving unexpected path: " + req.URL.Path)
	}

	h.Log("%s %s", req.Method, req.URL.Path)
	parts := strings.SplitN(req.URL.Path[len(h.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	name := parts[0]
	key := parts[1]

	group := groups[name] 
	if group == nil {
		http.Error(w, "name is bad", http.StatusBadRequest)
	}

	value, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(value.b)
}

