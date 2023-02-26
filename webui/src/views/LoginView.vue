<script>
export default {
    components: {},
    data: function () {
        return {
            errormsg: null,
            username: "",
        
        }
    },
    methods: {
        async doLogin() {
            if (this.username == "") {
                this.errormsg = "You must enter the username first";
            } else {
            
                    let response = await this.$axios.post("/session", { username: this.username })
                    localStorage.setItem("token", response.data.userId);
                    localStorage.setItem("username", this.username);

                    this.$router.push({ path: '/session' })
            }

        }
    },
    mounted() {

    }

}
</script>

<template>
    <h2></h2>
    <br><br>
    <br><br>
    <div class="login">
    <label class="welcome">Welcome in WasaPhoto</label>
    <br><br>
    <label>Username</label>
    <br><br>
    <input type="text" name="Uname" id="Uname" placeholder="Username" v-model="username">
    <br><br>
    <input type="button" name="log" id="log" value="LogIn" @click="doLogin">
    <ErrorMsg v-if="errormsg" :msg="errormsg" class="errorText"></ErrorMsg>
</div>
</template>


