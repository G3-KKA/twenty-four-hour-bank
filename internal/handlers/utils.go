package handlers

import (
	log "bank24/internal/logger"
	"net/http"
	"strconv"
	"strings"
)

// log into debug cores and write precified header
func logError(handler string, r *http.Request, err error) {
	//fmt.Sprintf()
	log.Debug(handler, " failed for", r.RemoteAddr, " as:", r.Method, " to: ", r.URL.Path, " with error:", err)

}

// /url/{id}/continue
// comes after should not include '/', just a word, like accounts
// example: URL /urlname/123/continue comesAfter : urlname
func urlGetId(url, comesAfter string) (integer int, err error) {
	i1 := strings.Index(url, comesAfter)
	r1 := url[i1:]
	i2 := strings.Index(r1, "/")
	r2 := r1[i2+1:]
	i3 := strings.Index(r2, "/")
	integer, err = strconv.Atoi(r2[:i3])
	return

}
