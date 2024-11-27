package mappers

import (
	get_email_contract "email-service/Application.contract/get_email"
	get_emails_contract "email-service/Application.contract/get_emails"
	send_email_contract "email-service/Application.contract/send_email"
	send_email "email-service/Application/commmand_handlers/send_email"
	get_email "email-service/Application/query_handlers/get_email"
	get_emails "email-service/Application/query_handlers/get_emails"
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
