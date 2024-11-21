package mappers

import (
	get_email_contract "file-storage/Application.contract/get_email"
	get_emails_contract "file-storage/Application.contract/get_emails"
	send_email_contract "file-storage/Application.contract/send_email"
	send_email "file-storage/Application/commands/send_email"
	get_email "file-storage/Application/queries/get_email"
	get_emails "file-storage/Application/queries/get_emails"
)

func MapToSendEmailCommand(req *send_email_contract.SendEmailRequest) send_email.SendEmailCommand {
	return send_email.SendEmailCommand{
		From:    req.From,
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}
}

func MapToGetEmailsQuery(req *get_emails_contract.GetEmailsRequest) get_emails.GetEmailsQuery {
	return get_emails.GetEmailsQuery{}
}

func MapToGetEmailQuery(req *get_email_contract.GetEmailRequest) get_email.GetEmailQuery {
	return get_email.GetEmailQuery{
		ID: req.ID,
	}
}
