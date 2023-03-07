<script>
export default {
	data: function () {
		return {
			errormsg: null,
			detailedmsg: null,
			limit: 10,
			startIndex: 0,
			photoComment:[
					{	userId: 0,
						name: "",
						comment: "",
						commentId: 0,
						date: "",
						photoId: 0,
					}
			],
			photoStream: [
					{
						photoId: 0,
						userId: 0,
						name: "",
						likes: 0,
						comments: 0,
						upladTime: "",
					}
				],
			c: 0,
			photoId: 0,
			nCoomments: 0,
			userPhoto: null,
			username: localStorage.getItem('username'),
			token: localStorage.getItem('token'),
			photoLike: [],
			
		}
	},
	methods: {
		async refresh() {
			this.getStream()
		},
		
		async doLogout() {
			localStorage.removeItem("token")
			localStorage.removeItem("username")
			this.$router
            	.push({ path: '/' })
				.then(() => { this.$router.go() })
		},
		async searchUser() {
			if (this.searchUsername != this.username){

				try{

					localStorage.setItem('viewName', this.searchUsername);
					let viewId = await this.$axios.get("/users?username="+this.searchUsername , {
							headers: {
								Authorization: this.token
							}
					})
					localStorage.setItem('viewId', viewId.data.userId)
					this.$router
						.push({ path: '/users/' + this.searchUsername + '/view' })
						.then(() => { this.$router.go() })
				}catch(e){
					alert(e.response.data)

				}

			}else{
				this.viewProfile()
			}
			
		},
		
		async getStream() {

			try{
				let response = await this.$axios.get("users/" + localStorage.getItem("token") + "/stream?limit=" + this.limit + "&startIndex=" + this.startIndex , {
					headers: {
						Authorization:  this.token
					}
				})
				this.startIndex = this.startIndex + this.limit;
				this.photoStream =  response.data;
				this.userPhoto = new Map();
				
				for(var i = 0; i < (response.data).length; i++){
					try{
						let fotoFile = await this.$axios.get("user/" + (response.data)[i].userId + "/photo/" + (response.data)[i].photoId,{
								headers: {
									Authorization:  this.token
								}
							})
							this.userPhoto.set((response.data)[i].photoId,fotoFile.data.image );
						//-- Costruisco array con gli id degli utenti che hanno messo mi piace alla foto 
						
							let users = await this.$axios.get("photo/" + (response.data)[i].photoId  +"/likes", {
								headers: {
									Authorization:  this.token
								}
							})
							if ( users.data ){
								for(var x = 0; x < (users.data).length; x++){
									if ( users.data[x].userId == this.token){
										
										this.photoLike.push((response.data)[i].photoId)
									}
								}
							}
					}catch (error) {
							continue;
					}
					
				 }

			}catch(e){
				alert(e.response.data)
			}
		},
		
		async viewProfile() {
			this.$router
            	.push({ path: '/users/' + this.username + '/profile' })
				.then(() => { this.$router.go() })
		},
		async goToProfile(val, viewId){
			try{
				let response = await this.$axios.get("users/" + viewId + "/profile", {
							headers: {
								Authorization: this.token
							}
						})
						localStorage.setItem("viewId", viewId);
						localStorage.setItem("viewName", val);
				this.$router
						.push({ path: '/users/' + val + '/view' })
						.then(() => { this.$router.go() })
			}catch (e){
				alert(e.response.data)
			}
		},
		async likePost(val){
			try{
				let response = await this.$axios.put("photo/" + val + "/like/" + this.token , {},  {
						headers: {
								Authorization: this.token
						}
				})
				location.reload();
			}catch(e){
				alert(e.response.data)
			}
		},
		async unlikePost(val){
			try{
				let response = await this.$axios.delete("photo/" + val + "/like/" + this.token, {
										headers: {
											Authorization: this.token
										}
							})
				location.reload();
			}catch(e){
				alert(e.response.data)
			}
		},
		async showComment(val){
			let response = await this.$axios.get("photo/" + val + "/comment", {
						headers: {
							Authorization: this.token
						}
					})
			this.photoComment = response.data;
			if(response.data == null){
				this.c = 0;
			}else{
				this.c = 1;
			}
			this.photoId = val;
			document.getElementById("commentForm").style.display = "block";
		},
		async closeComment(){
			document.getElementById("commentForm").style.display = "none";
			this.photoComment = [];
		},
		async deleteComment(commentId, photoId){
			let response = await this.$axios.delete("photo/" + photoId + "/comment/"+commentId, {
						headers: {
							Authorization: this.token
						}
					})
			location.reload();
		},
		async postComment(pId){
			let response = await this.$axios.post("photo/" + pId + "/comment", {text: this.inputCommentText}, {
						headers: {
							Authorization: this.token
						}
			})
			location.reload();
		},
		async loadMore(){
			let response = await this.$axios.get("users/" + localStorage.getItem("token") + "/stream?limit=" + this.limit + "&startIndex=" + this.startIndex , {
								headers: {
									Authorization:  this.token
								}
							})
							this.startIndex = this.startIndex + this.limit;
			for(var i = 0; i < (response.data).length; i++){
				this.photoStream.push( (response.data)[i]);
			}
			for(var i = 0; i < (response.data).length; i++){
								
				try{
						let fotoFile = await this.$axios.get("user/" + (response.data)[i].userId + "/photo/" + (response.data)[i].photoId,{
							headers: {
										Authorization:  this.token
							}
						})
						this.userPhoto.set((response.data)[i].photoId,fotoFile.data.image );
						//-- Costruisco array con gli id degli utenti che hanno messo mi piace alla foto 
									
						let users = await this.$axios.get("photo/" + (response.data)[i].photoId  +"/likes", {
							headers: {
										Authorization:  this.token
										}
						})
						if ( users.data ){
							for(var x = 0; x < (users.data).length; x++){
								if ( users.data[x].userId == this.token){
									this.photoLike.push((response.data)[i].photoId)
								}
							}
						}
					}catch (error) {
						continue;
					}
								
							}
		},
	},
	mounted() {
		this.getStream()
	},
	loadImages(){
		var  userId = document.getElementById('userId').innerHTML;
		var  photoId = document.getElementById('photoId').innerHTML;
	}
}
</script>
<template>
	<div>
		<!-- Icona bottone di ricerca-->
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
		<div class="topnav">
		<a class="active" href="#session" >WasaPhoto</a>
		<a class="leftItem" @click="viewProfile">Profile</a>
		<a class="leftItem" @click="doLogout">LogOut</a>
  		<input type="text" class="searchBar"  placeholder="Username" v-model="searchUsername">
        <button class="searchButton" @click="searchUser"><i class="fa fa-search"></i></button>
		</div>
		<!-- Messaggio se non ci sono post da visualizzare -->
		<div v-if="(photoStream[0].photoId == 0)" class="noPost">
			No Post
		</div>
		<!-- Controllo prima che ci siano post da visualizzare -->
		<div v-if="(photoStream[0].photoId != 0)" class="wrapper">
			<div v-for="post in photoStream" :key="post.photoId" class="card">
				<button v-on:click="goToProfile(post.name, post.userId)" value=post.name type="button" class="invisibleButton" >{{post.name}}</button>
				<br>
				<hr class="divUsername">
				<img alt="Image" :src="'data:image/jpeg;base64,'+userPhoto.get(post.photoId)" class="imageStandard">   
				<div>
					<table class="infoSection"> 
								<tr >
									<th ><button class="unlikeButton" id="likeButton" v-if="photoLike.indexOf(post.photoId)== -1"  @click="likePost(post.photoId)"><i class="fa fa-heart" aria-hidden="true"></i></button><label id="nLike"  class="showNumber" v-if="photoLike.indexOf(post.photoId) == -1">{{post.likes}}</label></th>
									<th ><button class="likeButton" id="likeButton" v-if="photoLike.indexOf(post.photoId)!== -1"  @click="unlikePost(post.photoId)"><i class="fa fa-heart" aria-hidden="true"></i></button><label id="nLike"  class="showNumber" v-if="photoLike.indexOf(post.photoId)!= -1">{{post.likes}}</label></th>
									<th class="commentInfo" ><button class="commentButton" id="commentButton"  @click="showComment(post.photoId)"><i class="fa fa-comment" aria-hidden="true"></i></button><label id="nComment" class="nComment" >{{post.comments}}</label></th>
								</tr>
					</table>
					<br>
					<label id="date" class="date" >{{post.upladTime}}</label><br>
				</div>

			</div>
		<button type="submit"  class="moreButton" @click="loadMore"><i class="fa fa-plus-circle" aria-hidden="true"></i></button>

		</div>


			<!-- Popup usato per mostrare i commenti e commentare un post -->
			<div class="commentPopup">
				<div class="formPopup" id="commentForm">
					<div class="formContainer">
						<label for="javascript" class="commentLabel">Comment</label>
						<button type="button" class="btn cancel" @click="closeComment"><i class="fa fa-times" aria-hidden="true"></i></button>
					<br>
					<div style="overflow-y:scroll; height:400px;" >
						<div v-if="c == 0" class="noPost">
							No Comment
						</div>
						<div v-if="c != 0" v-for="comment in photoComment" :key="comment.commentId" class="commentSec">
								<div class="userComment">
								<label >{{comment.name}}</label>
								</div>
								<label class="dateComment">{{comment.date}}</label>
								<button class="cancelComment" v-if="comment.userId == token" @click="deleteComment(comment.commentId, comment.photoId)"><i class="fa fa-trash"></i></button>
								<br>
								<label class="commentText">{{comment.comment}}</label>
								
						</div>
					</div>
				
					<input type="text" id="inputCommentText" name="inputCommentText" v-model="inputCommentText" class="inputCommentText">
					<button type="submit"  class="btn" @click="postComment(photoId)">Post Comment</button>
					</div>
				</div>
    		</div>



	



	 
            
	</div>
</template>
<style>
</style>