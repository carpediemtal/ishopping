document.write("<script src='../static/js/common.js'></script>");
var app = new Vue({	
    el: '#container',
	data:{
		times:0,
		isModifying:false,
		user:{
			UID:0,
			name:"",
			telephone:"",
			address:""
		}
	},
	methods:{
		JumpToSignIn:function(){
			alert("Please sign in first.");
			window.location.replace("../index.html");
		},
		initUserInfo:function(){
			var that = this;
			that.isModifying = false;
            const token = window.localStorage.getItem("ishopping-token")
			if (token){
				axios.defaults.headers.common = {Authorization: `${token}`}
				axios.get(api_url + "/buyer/information")
				.then(function (resp) {
					// handle success
					if (resp.data.code === 0) {
						that.user.UID = resp.data.data.uid;
						that.user.name = resp.data.data.name;
						that.user.address = resp.data.data.address;
						that.user.telephone = resp.data.data.phone_num;
					} else {
						console.log(resp.msg);
						that.JumpToSignIn();
					}
				}).catch(function (resp) {
					// handle error
					console.log(resp.msg);
			})
		}
		else{
			that.JumpToSignIn();
		}
	},
		Submit:function(){
			if (!this.user.name) {
				alert("please input your name!");
				return;
			} else if (!this.user.telephone) {
				alert("please input your telephone number!");
				return;
			} else if (!this.user.address) {
				alert("please input your address!");
				return;
			}
			var that = this;
               const token = window.localStorage.getItem("ishopping-token")
			if(token){
				axios.defaults.headers.common = {Authorization: `${token}`}
				axios({
					method: 'post',
					url: api_url + "/buyer/information_modify",
					data: {
						'uid' : that.user.UID,
						'name': that.user.name,
						'address' : that.user.address,
						'phone_num': that.user.telephone
					}
				}).then(function (resp) {
					if (resp.data.code === 0) {
						console.log("OK!");
						that.initUserInfo();
					} 
					else {
						console.log(resp.msg);
					}
				}).catch(function (resp) {
					console.log(resp.msg);
				})
			}
			else{
				that.JumpToSignIn();
			}
		},
		Modify:function(){
			this.isModifying = true;
		},
		Cancel:function(){
			this.initUserInfo();
		}
	},
	mounted(){this.$nextTick(() => {
		setTimeout(() => {
			this.initUserInfo();
		}, 100)
		})
	}
})