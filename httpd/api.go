package httpd

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/Cloud-Foundations/Dominator/lib/html"
	"github.com/Cloud-Foundations/Dominator/lib/log"
	"github.com/Cloud-Foundations/Dominator/lib/net/reverseconnection"
	"github.com/Cloud-Foundations/tricorder/go/tricorder"
)

type HtmlWriter interface {
	WriteHtml(writer io.Writer)
}

type RequestHtmlWriter interface {
	HtmlWriter
	RequestWriteHtml(writer io.Writer, req *http.Request)
}

var htmlWriters []HtmlWriter

func StartServer(portNum uint, logger log.DebugLogger) error {
	var listener net.Listener
	if l, err := reverseconnection.Listen("tcp", portNum, logger); err != nil {
		logger.Printf("error creating reverse listener, falling back: %s\n",
			err)
		address := fmt.Sprintf(":%d", portNum)
		if listener, err = net.Listen("tcp", address); err != nil {
			return err
		}
	} else {
		listener = l
		err = l.RequestConnections(tricorder.CollectorServiceName)
		if err != nil {
			return fmt.Errorf("error requesting reverse connections: %s", err)
		}
	}
	html.HandleFunc("/", statusHandler)
	html.HandleFunc("/favicon.ico", func(http.ResponseWriter, *http.Request) {})
	go http.Serve(listener, nil)
	return nil
}

func AddHtmlWriter(htmlWriter HtmlWriter) {
	htmlWriters = append(htmlWriters, htmlWriter)
}
