<script>
import LogModal from "../components/Logmodal.vue";

export default {
	components: { LogModal },
	data: function () {
		return {
			errormsg: null,
			detailedmsg: null,
			limit: 10,
			images: null,
			imageD: null,
			startIndex: 0,
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
			photoLike: [],
			photoComment:[
					{	userId: 0,
						name: "",
						comment: "",
						commentId: 0,
						date: "",
						photoId: 0,
					}
			],
			userPhoto: null,
			comments: 0,
			photoId: 0,
			username: localStorage.getItem('username'),
			token: localStorage.getItem('token'),
			profile: {
				userId: 0,
				userName: "",
				numPosts: 0,
				numFollower: 0,
				numFollowing: 0,
			},
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
			}else{
				
				this.viewProfile()
			}

			
			

			
		},

    
        async searcUserInfo(){
            let response = await this.$axios.get("users/" + this.token + "/profile", {
						headers: {
							Authorization:  localStorage.getItem("token")
						}
		    })
            this.profile = response.data
        },

		
       async getUserPhoto() {
				let response = await this.$axios.get("users/" + this.token + "/photo", {
					headers: {
						Authorization:  this.token
					}
				})
				this.photoStream =  response.data;

				this.userPhoto = new Map();
				
				for(var i = 0; i < (response.data).length; i++){
					
					try{ 
						let fotoFile = await this.$axios.get("user/" + this.token + "/photo/" + (response.data)[i].photoId,{
								headers: {
									Authorization:  this.token
								}
							})
							this.userPhoto.set((response.data)[i].photoId,fotoFile.data.image );

							// Costruisco array con gli id delle foto a cui l'utente ha messo mi piace in modo da cambiare il tasto like / unlike 
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


		async changeUsername(){

			let response = await this.$axios.put("users/" + this.token + "/username", { username: this.newUsername }, {
						headers: {
							Authorization:  this.token
						}
		    })
			alert("Username cambiato in " + response.data.username)
			this.$router
					.push({ path: '/users/' + response.data.username + '/profile' })
					.then(() => { this.$router.go() })
		},

		async uploadFile() {
			this.images = this.$refs.file.files[0]
		},
		
		async uploadPhoto() {
			if (this.images === null) {
				this.errormsg = "Please select a file to upload."
			} else {
				
					let response = await this.$axios.post("users/" + this.token + "/photo" , this.images, {
						headers: {
							Authorization: this.token
						}
					})

			}

			this.viewProfile()
		},
		async deletePhoto(val){
			let response = await this.$axios.delete("user/" + this.token + "/photo/" + val, {
						headers: {
							Authorization: this.token
						}
			})
			this.viewProfile();

		},

		async likePost(val){
			let response = await this.$axios.put("photo/" + val + "/like/" + this.token , {},  {
					headers: {
							Authorization: this.token
					}
		    })
			this.viewProfile();

		},

		async unlikePost(val){
			let response = await this.$axios.delete("photo/" + val + "/like/" + this.token, {
									headers: {
										Authorization: this.token
									}
						})
			this.viewProfile();

		},
		

		async viewProfile() {
			this.$router
            	.push({ path: '/users/' + this.username + '/profile' })
				.then(() => { this.$router.go() })
		},

		async showComment(val){
			let response = await this.$axios.get("photo/" + val + "/comment", {
						headers: {
							Authorization: this.token
						}
					})

			this.photoComment = response.data;
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
			this.viewProfile();


		},

		async postComment(photoId){
			
			let response = await this.$axios.post("photo/" + photoId + "/comment", {text: this.inputCommentText}, {
						headers: {
							Authorization: this.token
						}
			})
			this.viewProfile();
		},


		
	},


	mounted() {
		this.searcUserInfo()
        this.getUserPhoto()

	},




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


        <br>
        <hr>
        <table > 
            <tr class="firstPart">
                <th ><input type="text" class="newUsername"  placeholder="newUsername" v-model="newUsername"></th>
				<th class="firstPartR"><button class="changeButton" @click="changeUsername"><i class="fa fa-paper-plane-o"></i></button></th>

                <th class="rowFirstPart">{{profile.userName}}</th>
                <th class="lFirstPart"><input type="file" accept="image/*" @change="uploadFile" ref="file"></th>
                <th class="l2FistPart"><button class="changeButton" @click="uploadPhoto"><i class="fa fa-paper-plane-o"></i></button></th>
            </tr>
        </table>
        <hr>

        <br><br>


        <table class="secondPart">
            <tr class="labelProfile">
                <th>Post</th>
                <th>Follower</th>
                <th>Following</th>
            </tr>
            <tr>
                <th>{{profile.numPosts}}</th>
                <th>{{profile.numFollower}}</th>
                <th>{{profile.numFollowing}}</th>
            </tr>
        </table>


		<!-- Messaggio se non ci sono post da visualizzare -->
		<div v-if="(photoStream[0].photoId == 0)" class="noPost">
			No Post
		</div>


		<!-- Controllo prima che ci siano post da visualizzare e visualizzo tutta la lista di foto dell'utente -->
		<div v-if="(photoStream[0].photoId != 0)" class="wrapper">
			<div v-for="post in photoStream" :key="post.photoId" class="card">

			<label id="photoAuthor" for="photoAuthor" class="usPhoto">{{post.name}}</label>
			<button class="deleteButton" @click="deletePhoto(post.photoId)"><i class="fa fa-trash"></i></button>

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
		</div>


		<!-- Popup usato per mostrare i commenti e commentare un post -->
		<div class="commentPopup">
				<div class="formPopup" id="commentForm">
					<div class="formContainer">
						<label for="javascript" class="commentLabel">Comment</label>
						<button type="button" class="btn cancel" @click="closeComment"><i class="fa fa-times" aria-hidden="true"></i></button>

					<br>
					<div style="overflow-y:scroll; height:400px;" >
						<div v-if="photoComment.length == 0" class="noPost">
							No Comment
						</div>

							<div v-if="photoComment.length != 0" v-for="comment in photoComment" :key="comment.commentId" class="commentSec">
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

					<button type="submit" class="btn" @click="postComment(photoId)" >Post Comment</button>
					</div>
				</div>
    		</div>





	</div>

</template>