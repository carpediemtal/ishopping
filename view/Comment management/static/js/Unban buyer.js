var mainbody = new Vue({
    el: '.mainbody',
    data: {
        all: 8, //总页数
        cur: 1,//当前页码
        banneduser_list: [],
        // [{user_id:11,name:'name1'},
        // {user_id:22,name:'name2'},
        // {user_id:33,name:'name3'},
        // {user_id:44,name:'name4'}]
    },
    mounted() {
        var that = this
        // 'http://192.168.1.107:7001/api/admin/banned_user_list?page=1&page_size=20'
        axios.get(api_url+'/admin/banned_user_list?page=1&page_size=20')
        .then(function(response){
            console.log(response)
            that.banneduser_list = response.data.data.list
        })
        .catch(function(response){})
    },
    methods: {
        unbanBuyer:function(data){
            //console.log(data)
            // 'http://192.168.1.107:7001/api/admin/unban'
            axios.post(api_url+'/admin/unban',
            {
                user_id:data
            })
            .then(function(response){
                console.log(response)
                alert('Unban buyer successfully!')
                window.location.reload()
            })
            .catch(function(response){})
        },
        getBannedUserList:function(){
            console.log(this.cur)
            var that = this
            // api_url+'/admin/banned_user_list?page='+that.cur+'&page_size=20'
            axios.get(api_url+'/admin/banned_user_list?page='+that.cur+'&page_size=20')
            .then(function(response){
                console.log(response)
                if(response.data.code){
                    that.cur--
                    alert('This is the last page.')
                    return
                }
                that.banneduser_list = response.data.data.list
            })
            .catch(function(response){})
        },
        btnClick: function(data){
            if(data != this.cur){
                this.cur = data;
            }
            this.getGoodsinfo()
        },
        pageClick: function(){
            console.log('现在在'+this.cur+'页');
            this.getBannedUserList()
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