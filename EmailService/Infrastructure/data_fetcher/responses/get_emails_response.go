package datafetcher

type GetEmailsResponse struct {
	ID      string `json:"ID"`
	From    string `json:"From"`
	To      string `json:"To"`
	Subject string `json:"Subject"`
	Content struct {
		Body string `json:"Body"`
	} `json:"Content"`
}

type GetEmailsRawResponse struct {
	Messages []GetEmailsResponse `json:"items"`
}
