package url

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const hasTricorderUrl = "/has-tricorder-metrics"

var (
	likelyHTTPS = "HTTPS"
	maybeHTTP   = "server gave HTTP response to HTTPS client"
	maybeHTTPS  = "malformed HTTP response"
)

func (p *urlconfig) getURL(path string, useTLS bool) (*http.Response, error) {
	var url string
	if useTLS {
		url = fmt.Sprintf("https://localhost:%d%s", p.urlport, path)
	} else {
		url = fmt.Sprintf("http://localhost:%d%s", p.urlport, path)
	}
	return p.httpClient.Get(url)
}

func (p *urlconfig) probe(useTLS bool) error {
	defer p.httpTransport.CloseIdleConnections()
	res, err := p.getURL(p.urlpath, useTLS)
	if err != nil {
		p.healthy = false
		p.error = err.Error()
		return err
	}
	body, bodyErr := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode == 200 {
		p.statusCode = uint(res.StatusCode)
		p.healthy = true
		p.error = ""
		p.useTLS = useTLS
		if hasTricorderMetrics, err := p.probeTricorder(); err == nil {
			p.hasTricorderMetrics = hasTricorderMetrics
		}
	} else {
		status := res.Status
		if bodyErr == nil {
			if _status := strings.TrimSpace(string(body)); _status != "" {
				status = _status
			}
		}
		if useTLS && strings.Contains(status, maybeHTTP) {
			return p.probe(false)
		} else if !useTLS &&
			(strings.Contains(status, maybeHTTPS) ||
				strings.Contains(status, likelyHTTPS)) {
			return p.probe(true)
		}
		p.statusCode = uint(res.StatusCode)
		p.healthy = false
		p.error = status
	}
	return nil
}

func (p *urlconfig) probeTricorder() (bool, error) {
	res, err := p.getURL(hasTricorderUrl, p.useTLS)
	if err != nil {
		return false, err
	}
	ioutil.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}
