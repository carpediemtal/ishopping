const app = new Vue({
            el: "#container",
            data: {
				commodity_id : 0,
				count : 0,
                detail: {
                    name: "commodity name",
                    inventory: "inventory",
                    sales: "sales",
                    price: "price",
                    introduction:
                        "Placeholder content for the tab panel. This one relates to the home tab. Takes you miles high, so high, 'cause she’s got that one international smile. There's a stranger in my bed, there's a pounding in my head. Oh, no. In another life I would make you stay. ‘Cause I, I’m capable of anything. Suiting up for my crowning battle. Used to steal your parents' liquor and climb to the roof. Tone, tan fit and ready, turn it up cause its gettin' heavy. Her love is like a drug. I guess that I forgot I had a choice.",
                    thumbnail:"",
                    images: [],
                },
				comments:[]
            },
            created: function () {
                let url = new URL(window.location.href)
                const cid = url.searchParams.has("cid") ? parseInt(url.searchParams.get("cid")) : 0
				this.commodity_id = cid;
                const that = this
                axios({
                    method: "get",
                    url: `${api_url}/commodity_detail?commodity_id=${cid}`,
                })
                    .then(function (resp) {
                        if (resp.data.code === 0) {
                            that.detail = resp.data.data
                        } else {
                            alert("commodity not found:" + resp.data.msg)
                        }
                    })
                    .catch(function (resp) {
                        console.log(resp)
                        alert(resp.data.msg)
                    })
				axios({
                    method: "get",
                    url: `${api_url}/commodity_evaluation?commodity_id=${cid}`,
                })
                    .then(function (resp) {
                        that.comments = resp.data.data.list;
                    })
                    .catch(function (resp) {
                        console.log(resp)
                        alert(resp.data.msg)
                    })
            },
            methods: {
				addToCart:function(){
					var that = this;
					const token = window.localStorage.getItem("ishopping-token")
					if (token){
						axios.defaults.headers.common = {Authorization: `${token}`}
						axios({
							method: "get",
							url: api_url + '/userType'
						})
						.then(function (resp) {
							// handle success
							if (resp.data.data.type == 0){
								axios.defaults.headers.common = { Authorization: `${token}` }
								axios({
									method: "post",
									url: api_url + "/buyer/cart_add",
									data: {
										'commodity_id' : that.commodity_id,
										'count' : that.count
									},
								})
									.then(function (resp) {
										// handle success
										if (resp.data.code === 0) {
											alert("OK!");
											location.reload();
										} else {
											console.log(resp.data.msg);
											alert(resp.data.msg);
											location.reload();
										}
									})
									.catch(function (resp) {
										// handle error
										console.log(resp.msg);
										alert(resp.msg);
									})
							}
							else if (resp.data.data.type == 1)
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
            },
        })