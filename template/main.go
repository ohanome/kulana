package template

import (
	"kulana/output"
)

var url string
var hostname string
var port int
var status int
var time float64
var destination string
var contentLength int64
var ipAddress string
var mxRecords []string
var pingSuccessful int
var icmpCode int
var cname string
var content string
var foreignId string
var certificateValid int
var certificateValidUntil string
var certificateIssuer string

var noColor bool

func Render(t string, o output.Output, nc bool) string {
	noColor = nc

	url = o.Url
	hostname = o.Hostname
	port = o.Port
	status = o.Status
	time = o.Time
	destination = o.Destination
	contentLength = o.ContentLength
	ipAddress = o.IpAddress
	mxRecords = o.MXRecords
	pingSuccessful = o.PingSuccessful
	cname = o.CNAME
	icmpCode = o.ICMPCode
	content = o.Content
	foreignId = o.ForeignID
	certificateValid = o.CertificateValid
	certificateValidUntil = o.CertificateValidUntil
	certificateIssuer = o.CertificateIssuer

	switch t {
	case "json":
		return RenderJSON()

	case "csv":
		return RenderCSV()

	default:
		return RenderDefault()
	}
}
