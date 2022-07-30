package routes

import (
	"api/src/controller"
	"net/http"
)

var postsRoutes = []route{
	{
		URI:         "/posts",
		Method:      http.MethodGet,
		HandlerFunc: controller.GetPosts,
		NeedAuth:    true,
	},
	{
		URI:         "/posts/{id}",
		Method:      http.MethodGet,
		HandlerFunc: controller.FindPost,
		NeedAuth:    true,
	},
	{
		URI:         "/posts",
		Method:      http.MethodPost,
		HandlerFunc: controller.CreatePosts,
		NeedAuth:    true,
	},
	{
		URI:         "/posts/{id}",
		Method:      http.MethodPut,
		HandlerFunc: controller.UpdatePost,
		NeedAuth:    true,
	},
	{
		URI:         "/posts/{id}",
		Method:      http.MethodDelete,
		HandlerFunc: controller.DeletePost,
		NeedAuth:    true,
	},
	{
		URI:         "/posts/{id}/like",
		Method:      http.MethodPost,
		HandlerFunc: controller.LikePost,
		NeedAuth:    true,
	},
	{
		URI:         "/posts/{id}/unlike",
		Method:      http.MethodPost,
		HandlerFunc: controller.UnlikePost,
		NeedAuth:    true,
	},
	{
		URI:         "/users/{id}/posts",
		Method:      http.MethodGet,
		HandlerFunc: controller.GetPostByUser,
		NeedAuth:    true,
	},
}
