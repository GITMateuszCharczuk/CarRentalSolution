package mappers

import (
	create_contract "identity-api/Application.contract/BlogPosts/CreateBlogPost"
	update_contract "identity-api/Application.contract/BlogPosts/UpdateBlogPost"
	create_commands "identity-api/Application/command_handlers/create_blog_post"
	delete_commands "identity-api/Application/command_handlers/delete_blog_post"
	update_commands "identity-api/Application/command_handlers/update_blog_post"
	get_queries "identity-api/Application/query_handlers/get_blog_post"
	get_tags_queries "identity-api/Application/query_handlers/get_tags"
)

func MapToCreateBlogPostCommand(request *create_contract.CreateBlogPostRequest) create_commands.CreateBlogPostCommand {
	return create_commands.CreateBlogPostCommand{
		Title:    request.Title,
		Content:  request.Content,
		Tags:     request.Tags,
		Visible:  request.Visible,
		JwtToken: request.JwtToken,
	}
}

func MapToUpdateBlogPostCommand(request *update_contract.UpdateBlogPostRequest) update_commands.UpdateBlogPostCommand {
	return update_commands.UpdateBlogPostCommand{
		Id:       request.Id,
		Title:    request.Title,
		Content:  request.Content,
		Tags:     request.Tags,
		Visible:  request.Visible,
		JwtToken: request.JwtToken,
	}
}

func MapToDeleteBlogPostCommand(id string, jwtToken string) delete_commands.DeleteBlogPostCommand {
	return delete_commands.DeleteBlogPostCommand{
		ID:       id,
		JwtToken: jwtToken,
	}
}

func MapToGetBlogPostQuery(id string) get_queries.GetBlogPostQuery {
	return get_queries.GetBlogPostQuery{
		ID: id,
	}
}

func MapToGetTagsQuery() get_tags_queries.GetTagsQuery {
	return get_tags_queries.GetTagsQuery{}
}
