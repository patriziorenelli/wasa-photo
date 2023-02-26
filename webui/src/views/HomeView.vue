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
			

			username: localStorage.getItem('username'),
			token: localStorage.getItem('token'),
			
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
							Authorization: "userAuth " + localStorage.getItem("token")
						}
					})
					// Manca il controllo su esito api 
					this.$router
					.push({ path: '/users/' + this.searchUsername + '/view' })
					.then(() => { this.$router.go() })
				
			}
		},
		


		async getStream() {
				let response = await this.$axios.get("users/" + localStorage.getItem("token") + "/stream?limit=" + this.limit + "&startIndex=" + this.startIndex , {
					headers: {
						Authorization: "userAuth " + localStorage.getItem("token")
					}
				})
				this.photoStream = response.data
		},

		
		async viewProfile() {
			this.$router
            	.push({ path: '/users/' + this.username + '/profile' })
				.then(() => { this.$router.go() })
		},

		async goToProfile(){
			let response = await this.$axios.get("users/" + this.searchUsername + "/profile", {
						headers: {
							Authorization: "userAuth " + localStorage.getItem("token")
						}
					})
					this.$router
					.push({ path: '/users/' + this.searchUsername + '/view' })
					.then(() => { this.$router.go() })
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

			<!-- Per andare al profilo dell'utente che ha postato la foto-->
			<button v-on:click="goToProfile" value=post.name type="button" class="invisibleButton" >{{post.name}}</button>
			<br>

  			<img src="img_avatar.png" alt="Avatar" style="width:100%">


				<div class="container">
				<br>
				  	<label id="photoId" for="photoId" >PhotoId:{{post.photoId}}</label><br>

				  	<label id="userId"  >userId:{{post.userId}}</label><br>
				</div>

			</div>
		</div>
		








	 



            


	</div>

</template>

<style>

</style>
