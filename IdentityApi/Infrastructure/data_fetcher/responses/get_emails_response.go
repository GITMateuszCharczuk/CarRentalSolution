package datafetcher

type EmailAddress struct {
	Relays  interface{} `json:"Relays"`
	Mailbox string      `json:"Mailbox"`
	Domain  string      `json:"Domain"`
	Params  string      `json:"Params"`
}

type GetEmailsResponse struct {
	ID      string         `json:"ID"`
	From    EmailAddress   `json:"From"`
	To      []EmailAddress `json:"To"`
	Content struct {
		Headers struct {
			Subject []string `json:"Subject"`
		} `json:"Headers"`
		Body string `json:"Body"`
	} `json:"Content"`
}

type GetEmailsRawResponse struct {
	Messages []GetEmailsResponse `json:"items"`
}
