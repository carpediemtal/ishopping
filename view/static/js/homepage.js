var header = new Vue({
    el:".header",
    data:{
        category:["Clothing","Shoe&&Accessories","Make-up&&Cleaning","Electronics","Outdoor&&Sports","Electrical-appliances","Food&&Fresh","Medicine&Healthcare","Kitchenware&&Textiles","Toys"]
    },
    methods: {
        getherf: function(index,item){
            return 'categorypage.html?index='+(index+1)+'&value='+item
        }
    }
})



var mainbody = new Vue({
    el: '.mainbody',
    data: {
        all: 8, //总页数
        cur: 1,//当前页码
        goodsinfo:[]
    },
    watch: {
        cur: function(oldValue , newValue){
            console.log(arguments);
        }
    },    
    mounted() {
       
        //http://ishopping.gq/api/index_category_channel
        var that = this
        axios.get(api_url+'/index_category_channel?category_id=0&page_num=1')
        .then(function(response){
            //console.log(response.data.data.commodity_list)
            that.goodsinfo = response.data.data.commodity_list
        })
        .catch(function(err){})

    },
    methods: {
        getGoodsinfo: function(){       
            //console.log(this.cur)
            var that = this
            axios.get(api_url+'/index_category_channel?category_id=0&page_num='+that.cur)
            .then(function(response){
                //console.log(response.data.data.commodity_list)
                that.goodsinfo = response.data.data.commodity_list
            })
            .catch(function(err){})
        },
        btnClick: function(data){
            if(data != this.cur){
                this.cur = data;
            }
            this.getGoodsinfo()
        },
        pageClick: function(){
            console.log('现在在'+this.cur+'页');
            this.getGoodsinfo()
        },
    },
    computed: {
        indexs: function(){
          var left = 1;
          var right = this.all;
          var ar = [];
          if(this.all>= 5){
          	//这里最大范围从3到6，如果到达7，那么下边加2变成9，已经超过最大的范围值
            if(this.cur > 3 && this.cur < this.all-1){
            	//以4为参考基准，左面加2右边加2
                    left = this.cur - 2
                    right = this.cur + 2
            }else{
                if(this.cur<=3){
                    left = 1
                    right = 5
                }else{
                    right = this.all
                    left = this.all -4
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