document.onselectstart = function(){return false;};
var commodity_list = new Vue({
    el:"#app",
    data(){
        return{
            all: 8, //总页数
            cur: 1,//当前页码
            page_size:10,
            page_index:1,
            list:[
                {commodity_id:0,name:"commodityname0",price:100,thumbnail:"...",sales:100},
                {commodity_id:1,name:"commodityname1",price:200,thumbnail:"...",sales:100},
                {commodity_id:2,name:"commodityname2",price:300,thumbnail:"...",sales:100},
                {commodity_id:3,name:"commodityname3",price:400,thumbnail:"...",sales:100},
                {commodity_id:4,name:"commodityname4",price:500,thumbnail:"...",sales:100},
                {commodity_id:5,name:"commodityname5",price:600,thumbnail:"...",sales:100},
                {commodity_id:6,name:"commodityname6",price:700,thumbnail:"...",sales:100},
            ]
        }
    },
    methods:{
        detail:function(id){
            // window.location.replace(`"commodity_detail.html?cid=${id}"`);
            window.location.href = `"commodity_detail.html?cid=${id}"`

        },
        getPage:function(){
            const that = this
            const token = window.localStorage.getItem("ishopping-token")
            axios.defaults.headers.common = { Authorization: `${token}` }
            axios({
                url: api_url + `"/seller/commodity_list?page_index=${that.cur}&page_size=${that.page_size}"`,
                // url: `"http://ishopping.gq/api/seller/commodity_list?page_index=${that.cur}&page_size=${that.page_size}"`,
                method: "get",
            }).then(function (res) {
                    if (res.data.code == 0) {
                        that.list = res.data.data.commodity_list
                    } else {
                        that.cur--
                        alert('This is the last page.')
                        // console.log(res.data.msg)
                        return
                    }
                }).catch(function (error) {console.log(error)})

        },
        btnClick: function(page){
            if(page != this.cur){
                this.cur = page;
            }
            this.getPage()
        },
        pageClick: function(){
            console.log('现在在'+this.cur+'页');
            this.getPage()
        },
    },
    mounted() {
        const that = this
        const token = window.localStorage.getItem("ishopping-token")
        axios.defaults.headers.common = { Authorization: `${token}` }
        axios({
            url: api_url + `"/seller/commodity_list?page_index=${that.page_index}&page_size=${that.page_size}"`,
            // url: `"http://ishopping.gq/api/seller/commodity_list?page_index=${that.page_index}&page_size=${that.page_size}"`,
            method: "get",
        }).then(function (res) {
            that.hasOrder = true
            if (res.data.code == 0) {
                that.list = commodity_list
            } else {
                console.log(res.data.msg)
            }
        }).catch(function (error) {console.log(error)})

    },
    computed: {
        indexs: function(){
          var left = 1;
          var right = this.all;
          var ar = [];
          if(this.all>= 5){
            if(this.cur > 3 && this.cur < this.all - 1){
                    left = this.cur - 2
                    right = this.cur + 2
            }else{
                if(this.cur <= 3){
                    left = 1
                    right = 5
                }else{
                    right = this.all
                    left = this.all - 4
                }
            }
         }
        while (left <= right){
            ar.push(left)
            left ++
        }
        console.log(ar);
        return ar
       }
    }
    
})