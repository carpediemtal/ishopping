document.write("<script src='../static/js/common.js'></script>");
 nav = new Vue({	
    el: '#navigation',
	data:{
		isSignedIn:2
	},
	methods:{
		initUserInfo:function(){
			var that = this;
            const token = window.localStorage.getItem("ishopping-token")
			if (token){
				axios.defaults.headers.common = {Authorization: `${token}`}
				axios({
                    method: "post",
                    url: api_url + '/userType',
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
				that.isSignedIn = 2;
			}
		}
	},
	mounted(){this.$nextTick(() => {
		setTimeout(() => {
			this.initUserInfo();
		}, 100)
		})
	}
})