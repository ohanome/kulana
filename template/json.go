package template

import (
	"fmt"
	"strconv"
	"strings"
)

func RenderJSON() string {
	var template []string

	if url != "" {
		template = append(template, "\"url\":\""+url+"\"")
	}

	if hostname != "" {
		template = append(template, "\"hostname\":\""+hostname+"\"")
	}

	if port > 0 {
		template = append(template, "\"port\":\""+strconv.Itoa(port)+"\"")
	}

	if ipAddress != "" {
		template = append(template, "\"ip\":\""+ipAddress+"\"")
	}

	if status != 0 {
		template = append(template, "\"status\":\""+strconv.Itoa(status)+"\"")
	}

	if time > -1 {
		template = append(template, "\"time\":\""+fmt.Sprintf("%f", time)+"\"")
	}

	if destination != "" && status < 400 && status >= 300 {
		template = append(template, "\"destination\":\""+destination+"\"")
	}

	if contentLength > -1 {
		template = append(template, "\"length\":\""+fmt.Sprintf("%d", contentLength)+"\"")
	}

	if len(mxRecords) > 0 {
		var mx []string
		for _, m := range mxRecords {
			mx = append(mx, m)
		}

		template = append(template, fmt.Sprintf("\"mx_records\":[\"%s\"]", strings.Join(mx, "\",\"")))
	}

	if icmpCode >= 0 {
		template = append(template, fmt.Sprintf("\"icmp_code\":%d", icmpCode))
	}

	if pingSuccessful > -1 {
		success := false
		if pingSuccessful == 1 {
			success = true
		}
		template = append(template, fmt.Sprintf("\"ping_successful\":%v", success))
	}

	if pingError != "" {
		template = append(template, fmt.Sprintf("\"ping_error\":\"%s\"", pingError))
	}

	if cname != "" {
		template = append(template, fmt.Sprintf("\"cname\":\"%s\"", cname))
	}

	// Content will not be output for now.
	//if content != "" {
	//	c := strings.ReplaceAll(content, "\"", "'")
	//	c = strings.ReplaceAll(c, "\n", "")
	//	c = strings.ReplaceAll(c, "\r", "")
	//	template = append(template, "\"content\":\""+c+"\"")
	//}

	if foreignId != "" {
		template = append(template, fmt.Sprintf("\"foreign_id\":\"%v\"", foreignId))
	}

	if certificateValid > -1 {
		isValid := false
		if certificateValid == 1 {
			isValid = true
		}
		template = append(template, fmt.Sprintf("\"certificate_valid\":\"%v\"", isValid))
	}

	if certificateValidUntil != "" {
		template = append(template, fmt.Sprintf("\"certificate_valid_until\":\"%v\"", certificateValidUntil))
	}

	if certificateIssuer != "" {
		template = append(template, fmt.Sprintf("\"certificate_issuer\":\"%v\"", certificateIssuer))
	}

	t := "{" + strings.Join(template, ",") + "}"

	return fmt.Sprint(t)
}
