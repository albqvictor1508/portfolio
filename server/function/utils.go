package function

import "net/url"

func isValidURL(urlStr string) bool {
	if urlStr == "" {
		return true
	}
	_, err := url.ParseRequestURI(urlStr)
	return err == nil
}

