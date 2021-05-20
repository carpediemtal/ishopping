var mainbody = new Vue({
    el: '.mainbody',
    data: {
        all: 8, //总页数
        cur: 1,//当前页码
        evaluation_list: [],
        // [{evaluation_id:1,buyer_id:11,buyer_name:'name1',commodity_id:111,commodity_name:'cname1',content:'Whatever is worth doing is worth doing well.'},
        // {evaluation_id:2,buyer_id:22,buyer_name:'name2',commodity_id:222,commodity_name:'cname2',content:'Happiness is a way station between too much and too little.'},
        // {evaluation_id:3,buyer_id:33,buyer_name:'name3',commodity_id:333,commodity_name:'cname3',content:'I lied when I said I didn’t like you. I lied when I said I didn’t care. I lie every time I try to tell myself I will never fall for you.'},
        // {evaluation_id:4,buyer_id:44,buyer_name:'name4',commodity_id:444,commodity_name:'cname4',content:'If you wish to succeed, you should use persistence as your good friend, experience as your reference, prudence as your brother and hope as your sentry.'}]
    },
    mounted() {
        var that = this
        // http://192.168.1.107:7001/api/admin/evaluation_list?page=1&page_size=10
        axios.get(api_url+'/admin/evaluation_list?page=1&page_size=10')
        .then(function(response){
			if (response.data.code === 0){
				that.evaluation_list = response.data.data.evaluation_list
			}
			else{
				console.log(response.data.msg)
			}
        })
        .catch(function(response){console.log(response.data.msg)})
    },
    methods: {
        deleteEvaluation:function(data){
            console.log(data)
            // http://192.168.1.107:7001/api/admin/delete_evaluation
            axios.post(api_url+'/admin/delete_evaluation',
            {
                evaluation_id:data
            })
            .then(function(response){
                console.log(response)
                alert('Delete evaluation successfully!')
                window.location.reload()
            })
            .catch(function(response){})
        },
        banBuyer:function(data){
            //console.log(data)
            // http://192.168.1.107:7001/api/admin/ban
            axios.post(api_url+'/admin/ban',
            {
                user_id:data
            })
            .then(function(response){
                console.log(response)
                alert('Ban buyer successfully!')
                window.location.reload()
            })
            .catch(function(response){})
        },
        getEvaluationList:function(){
            //console.log(this.cur)
            var that = this
            // 'http://192.168.1.107:7001/api/admin/evaluation_list?page='+that.cur+'&page_size=10'
            axios.get(api_url+'/admin/evaluation_list?page='+that.cur+'&page_size=10')
            .then(function(response){
                console.log(response)
                if(response.data.code){
                    that.cur--
                    alert('This is the last page.')
                    return
                }
                that.evaluation_list = response.data.data.evaluation_list
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
            this.getEvaluationList()
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