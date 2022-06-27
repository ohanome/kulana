package fetcher

import (
	"io"
	"io/ioutil"
	"kulana/hostinfo"
	"kulana/misc"
	"kulana/output"
	"net/http"
)

func CreateHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func FetchHTTPHost(url string, foreignId string, checkSSLCert bool) output.Output {
	client := CreateHTTPClient()
	start := misc.MicroTime()
	resp, err := client.Get(url)
	end := misc.MicroTime()
	defer client.CloseIdleConnections()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	if err != nil {
		panic(err)
	}

	statusCode := resp.StatusCode
	responseTime := (end - start) * 1000
	responseTimeRounded := misc.Round(responseTime, 0.000001)

	var destination string
	location, err := resp.Location()
	if err != nil {
		destination = url
	} else {
		destination = location.String()
	}

	contentLength := resp.ContentLength
	body, err := ioutil.ReadAll(resp.Body)

	response := output.Output{
		Url:                   url,
		Status:                statusCode,
		Time:                  responseTimeRounded,
		Destination:           destination,
		ContentLength:         contentLength,
		Content:               string(body),
		ForeignID:             foreignId,
		IpAddress:             hostinfo.ResolveIPAddress(url),
		CertificateValid:      -1,
		CertificateValidUntil: "",
	}

	if checkSSLCert {
		certIsValid, validUntil, issuer := hostinfo.CheckCertificate(url)
		if certIsValid {
			response.CertificateValid = 1
			response.CertificateValidUntil = validUntil.Format("2006-01-02 15:04:05")
			response.CertificateIssuer = issuer
		} else {
			response.CertificateValid = 0
		}
	}

	return response
}
