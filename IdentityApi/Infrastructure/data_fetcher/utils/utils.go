package datafetcher

import (
	responses "identity-api/Infrastructure/data_fetcher/responses"
	"strings"
)

func ExtractFirstString(arr []string) string {
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func FormatToAddresses(toAddresses []responses.EmailAddress) string {
	var addresses []string
	for _, addr := range toAddresses {
		addresses = append(addresses, addr.Mailbox+"@"+addr.Domain)
	}
	return joinAddresses(addresses)
}

func joinAddresses(addresses []string) string {
	return strings.Join(addresses, ", ")
}
