<script>
import LogModal from "../components/Logmodal.vue";
import SuccessMsg from "../components/SuccessMsg.vue";

export default {
	components: { LogModal, SuccessMsg },
	data: function () {
		return {
			errormsg: null,
			successmsg: null,
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
			
			userPhoto: null,
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
			if (this.searchUsername === this.username) {
				this.errormsg = "You can't search yourself"
			} else if (this.searchUsername === "") {
				this.errormsg = "The username to be searched must be inserted ..."
			} else {
				
					let response = await this.$axios.get("users/" + this.searchUsername + "/profile", {
						headers: {
							Authorization: localStorage.getItem("token")
						}
					})
					this.$router
					.push({ path: '/users/' + this.searchUsername + '/view' })
					.then(() => { this.$router.go() })
				
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

		// Prende correttamente la stessa foto ma condividono variabile -> mostra solo una foto 
       async getUserPhoto() {
				let response = await this.$axios.get("users/" + this.token + "/photo", {
					headers: {
						Authorization:  this.token
					}
				})
				this.photoStream =  response.data;

				this.userPhoto = new Map();
				
				for(var i = 0; i < (response.data).length; i++){

					let fotoFile = await this.$axios.get("user/" + this.token + "/photo/" + (response.data)[i].photoId,{
						headers: {
							Authorization:  this.token
						}
					})
					this.userPhoto.set((response.data)[i].photoId,fotoFile.data.image );
					//this.imageD = fotoFile.data.image

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
					this.successmsg = "Photo uploaded successfully."

			}
		},


		async deletePhoto(val){
			alert(val)
		}







		
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


		<!-- Controllo prima che ci siano post da visualizzare -->
		<div v-if="(photoStream[0].photoId != 0)" class="wrapper">
			<div v-for="post in photoStream" :key="post.photoId" class="card">

        <label id="photoAuthor" for="photoAuthor" class="usPhoto">{{post.name}}</label>
		<button class="deleteButton" @click="deletePhoto(post.photoId)"><i class="fa fa-trash"></i></button>

		<br>

        <hr class="divUsername">



			<img alt="Image" :src="'data:image/jpeg;base64,'+userPhoto.get(post.photoId)" class="imageStandard">   
  			
			<div class="container">
				<label id="nLike"  class="showNumber">Like:{{post.likes}}</label> <label id="nComment"  >Comment:{{post.comments}}     </label><br>
				<br><br>
				<label id="date" class="date" >{{post.upladTime}}</label><br>
			</div>


		</div>
	</div>






	</div>

</template>