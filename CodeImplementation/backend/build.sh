#!/bin/bash/env bash
go build -o ./app ./main.go ./stats.go ./routes/auth_routes.go ./routes/blog_routes.go ./routes/profile_routes.go ./routes/feed_routes.go ./profile/create_user.go ./profile/helper_functions.go ./profile/profile_classes.go ./profile/profile_search.go ./profile/retrieve_user.go ./profile/update_user.go ./middleware/crud_auth.go ./middleware/require_auth.go ./mail/mail_smtp.go ./feed/add_tags.go ./feed/feed_classes.go ./feed/helper_functions.go ./feed/reload_feed.go ./feed/send_feed.go ./database/db.go ./blog/blog_classes.go ./blog/blog_search.go ./blog/comment.go ./blog/create_post.go ./blog/delete_post.go ./blog/edit_post.go ./blog/helper_functions.go ./blog/react/go ./blog/retrieve_blog.go ./blog/retrieve_post.go ./authn/auth_classes.go ./authn/forgot_password.go ./authn/helper_functions.go ./authn/login.go ./authn/logout.go ./authn/signup.go