// console.log(window.location.href.split("index=")[1].split("&")[0])
// console.log(window.location.href.split("&value=")[1].split("?")[0])
var title = new Vue({
    el: '.title',
    data: {
        thetitle: window.location.href.split("&value=")[1].split("?")[0]
    }
})

var mainbody = new Vue({
    el: '.mainbody',
    data: {
        all: 8, //总页数
        cur: 1,//当前页码
        goodsinfo: []
    },
    watch: {
        cur: function (oldValue, newValue) {
            console.log(arguments);
        }
    },
    mounted() {
        // this.goodsinfo = [{pic:"../img/goods1.jpg",commodity_id:"1",name:"Toothbrush",price:"170.00 yuan"},
        // {pic:"../img/goods1.jpg",commodity_id:"1",name:"Toothbrush",price:"170.00 yuan"},
        // {pic:"../img/goods1.jpg",commodity_id:"1",name:"Toothbrush",price:"170.00 yuan"},
        // {pic:"../img/goods1.jpg",commodity_id:"1",name:"Toothbrush",price:"170.00 yuan"}]
        var that = this
        axios.get(api_url + '/index_category_channel?category_id=' + window.location.href.split("index=")[1].split("&")[0] + '&page_num=1')
            .then(function (response) {
                //console.log(response.data.data.commodity_list)
                that.goodsinfo = response.data.data.commodity_list
            })
            .catch(function (err) {
            })
    },
    methods: {
		jumpToDetail:function(cid){
			console.log(cid);
			window.location.replace("commodity_detail.html?cid="+cid);
		},		
		jumpToPurchase:function(cid){
			console.log(cid);
            const token = window.localStorage.getItem("ishopping-token")
			if (token){
				axios.defaults.headers.common = {Authorization: `${token}`}
				axios({
                    method: "post",
                    url: api_url + '/userType'
                })
				.then(function (resp) {
					// handle success
					if (resp.data.data.type == 0)
						window.location.replace("buyer/confirm_purchase.html");
					if (resp.data.data.type == 1)
						alert("please sign in as a buyer!");
				}).catch(function (resp) {
					// handle error
					console.log(resp.msg);
				})
			}
			else{
				alert("please sign in first!");
				window.location.replace("signin.html");
			}
		},
        getGoodsinfo: function () {
            console.log(this.cur)
            var that = this
            axios.get(api_url + '/index_category_channel?category_id=' + window.location.href.split("index=")[1].split("&")[0] + '&page_num=' + that.cur)
                .then(function (response) {
                    //console.log(response.data.data.commodity_list)
                    that.goodsinfo = response.data.data.commodity_list
                })
                .catch(function (err) {
                })
        },
        btnClick: function (data) {
            if (data != this.cur) {
                this.cur = data;
            }
            this.getGoodsinfo()
        },
        pageClick: function () {
            console.log('现在在' + this.cur + '页');
            this.getGoodsinfo()
        },
    },
    computed: {
        indexs: function () {
            var left = 1;
            var right = this.all;
            var ar = [];
            if (this.all >= 5) {
                //这里最大范围从3到6，如果到达7，那么下边加2变成9，已经超过最大的范围值
                if (this.cur > 3 && this.cur < this.all - 1) {
                    //以4为参考基准，左面加2右边加2
                    left = this.cur - 2
                    right = this.cur + 2
                } else {
                    if (this.cur <= 3) {
                        left = 1
                        right = 5
                    } else {
                        right = this.all
                        left = this.all - 4
                    }
                }
            }
            while (left <= right) {
                ar.push(left)
                left++
            }
            console.log(ar);
            return ar
        }
    }
})