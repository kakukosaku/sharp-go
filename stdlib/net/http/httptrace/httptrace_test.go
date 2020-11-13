// Author: kaku
// Date: 2020/11/13
//
// GitHub:
//	https://github.com/kakukosaku
//
package httptrace_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

func Example() {
	req, _ := http.NewRequest("GET", "http://localhost:80", nil)
	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info err: %+v\n", dnsInfo.Err)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn was idle: %+v\n", connInfo.WasIdle)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		log.Fatal(err)
	}
	// Output:
	// DNS Info err: <nil>
	// Got Conn was idle: false
}
