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

       async getUserPhoto() {
				let response = await this.$axios.get("users/" + this.token + "/photo", {
					headers: {
						Authorization:  this.token
					}
				})
				this.photoStream = response.data
		},

		async changeUsername(){

			let response = await this.$axios.put("users/" + this.token + "/username", { username: this.newUsername }, {
						headers: {
							Authorization:  this.token
						}
		    })
			alert("Username cambiato in " + response.data.username)
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

        <label id="photoId" for="photoId" >{{post.name}}</label><br>		
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