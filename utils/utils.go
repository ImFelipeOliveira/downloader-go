package utils

import (
	"net/url"
	"strings"
)

func VerifyExtension(link, extensions string) bool {
	for _, ext := range strings.Split(extensions, ",") {
		if strings.HasSuffix(link, strings.TrimSpace(ext)) {
			return true
		}
	}
	return false
}

func FormatUrl(base, link string) string {
	if strings.HasPrefix(link, "http") {
		return link
	}
	u, err := url.Parse(base)
	if err != nil {
		return link
	}

	return u.Scheme + "://" + u.Host + "/" + strings.TrimPrefix(link, "/")
}
