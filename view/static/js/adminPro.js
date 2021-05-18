document.write("<script src='../static/js/common.js'></script>");
var app = new Vue({	
    el: '#container',
	data:{
		times:0,
		issignin:false,
	},
	methods:{
		JumpToSignIn:function(){
			alert("Please sign in first.");
			window.location.replace("../index.html");
		},
		init:function(){
			var that = this;
            const token = window.localStorage.getItem("ishopping-token")
			if (!token)
				that.JumpToSignIn();
		},
	},
	mounted(){this.$nextTick(() => {
		setTimeout(() => {
			this.init();
		}, 100)
		})
	}
})