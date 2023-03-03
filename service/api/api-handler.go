package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.POST("/users", rt.wrap(rt.getUserId))
	rt.router.PUT("/users/:userId/username", rt.wrap(rt.setMyUserName))
	rt.router.PUT("/users/:userId/followUser/:userId2", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userId/followUser/:userId2", rt.wrap(rt.unfollowUser))
	rt.router.PUT("/users/:userId/banUser/:userId2", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userId/banUser/:userId2", rt.wrap(rt.unbanUser))
	rt.router.GET("/users/:userId/banUser/:userId2", rt.wrap(rt.checkUserBan))

	rt.router.PUT("/photo/:photoId/like/:userId", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/photo/:photoId/like/:userId", rt.wrap(rt.unlikePhoto))
	rt.router.POST("/photo/:photoId/comment", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/photo/:photoId/comment/:commentId", rt.wrap(rt.uncommentPhoto))

	rt.router.GET("/users/:userId/followers", rt.wrap(rt.getUserFollowers))
	rt.router.GET("/users/:userId/following", rt.wrap(rt.getUserFollowing))

	rt.router.GET("/user/:userId/photo/:photoId", rt.wrap(rt.getPhoto))
	rt.router.GET("/photo/:photoId/likes", rt.wrap(rt.getPhotoLike))
	rt.router.GET("/photo/:photoId/comment", rt.wrap(rt.getPhotoComments))
	rt.router.GET("/users/:userId/photo", rt.wrap(rt.getUserPhotos))
	rt.router.GET("/users/:userId/profile", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users/:userId/stream", rt.wrap(rt.getMyStream))
	rt.router.POST("/users/:userId/photo", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/user/:userId/photo/:photoId", rt.wrap(rt.deletePhoto))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
