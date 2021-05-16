 var nav = new Vue({	
    el: '#navigation',
	data:{
		isSignedIn:3
	},
	methods:{
		signOut:function(){
			window.localStorage.removeItem("ishopping-token");
			location.reload();
		},
		initNavigation:function(){
			var that = this;
            const token = window.localStorage.getItem("ishopping-token");
			if (token){
                axios.defaults.headers.common = { Authorization: `${token}` }
                axios({
                method: "get",
                url: `${api_url}/userType`,
                })
				.then(function (resp) {
					// handle success
					that.isSignedIn = resp.data.data.type;
				}).catch(function (resp) {
					// handle error
					console.log(resp.msg);
				})
			}
			else{
				that.isSignedIn = 3;
			}
		}
	}    ,
	mounted() {
		this.initNavigation();
    },
})
