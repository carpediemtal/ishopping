const app = new Vue({
    el: "#app",
    data: {
        content: "",
        commodityList: [],
        seen: false,
        price_min: 0.0,
        price_max: 0.0,
        page_index: 0,
        page_size: 2,
        newArr: [],
        i:0
    },
    methods: {
        getDetailUrl: function (commodity) {
            return `commodity_detail.html?cid=${commodity.commodity_id}`
        },
        pageClick: function () {
            const that = this
            console.log("next page")
            if (this.page_index < 0) {
                this.page_index = 0;
                alert("already browse to the first page!")
                return
            }
            console.log(this.page_index)
            // this.seen = true
            //     this.commodityList = [{commodity_id:3,name:"hi",price:1.0,sales:23,
            //     thumbnail:"data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAoHCBUVERISEhIREhISERIRERIREBEREhAPGBQZGRgUGBgcIS4lHB4rHxgYJjgmKy8xNTU1GiQ7QDs0Py40NTEBDAwMEA8QGhISHDQhGCExNDQxMTExMTQxNDQxNDQxNDQ0NDQxNDQ0NDE0NDQ0NDQ/NDQ/Pz80NDQ0NDQ0MTE/P//AABEIAOAA4AMBIgACEQEDEQH/xAAbAAABBQEBAAAAAAAAAAAAAAAAAQIDBAYFB//EADoQAAIBAwIEBAUCBQEJAQAAAAABAgMEEQUhEjFBUQZhcZETIjJCgVKhI3KxwdEUBxVigpKi4fDxM//EABoBAAIDAQEAAAAAAAAAAAAAAAACAQMEBQb/xAAuEQACAQMDAQUIAwEAAAAAAAAAAQIDBBESITFBExQiUWEjMkJScYGRoQUz4bH/2gAMAwEAAhEDEQA/APZgAAAAAAAAAAAAAAAAEZWrXkI85L06kNpckpN8FkDh19divpi5fsUqms1H9OF+5nnd0o9cl8bWpLoahsMmPnfVJc5e2xE6s398vdlLv4dEy3uUvM2uRcmI+LP9Uv3Hxv6seU2Hf4dUw7jLo0bTIGSp61UXPEv2L9DxBF/XHHpuWxvKUuuCuVrUXTJ3wKdvfwn9Ml6ZLSkaU090Z2mtmOATIZJIFAAAAAAAAAAAAAAAAAAAAOVrmuUrWnx1ZbvaMVzm/I6NSeIt9k37GCl4flfVpXF1KUKSk/hw70+jYAcPU/8AaHcTbVKEacM9m5e5zqfji8i8ufF5Szg21bTrWmuClQh2cu5za2gUqnOms90t0ZJ3kISxyaYW02svYhsvHsqmIVcU28LiXJs7MYOe+78+eTB+IvDM7dfEi+Kn36xfmanwJqXxaTpS3nSS/wCjoZ7iGuPaReV/wto1NHha3O3Cz7k8bRFtRFwYtKLnVb6ldUF2QvwV2J+EOEnAuog+CuyGSoLsvYtOInCGCVJlCdpFkFSy6o6jgNcBHFDqo11OFKlKD2yvQ59943lbt04/xJLvtwv16nS8U33wbeUl9Uvlh/Ng8/0TRql1NvlFPM5vlnsjZaw0J1JPEUVXFTX4EtyxdeMbuct6mMvZRXDjy2J7LxjeUnlyco/cpx5/lnbpaHTpbcCljrJZeTu2FWi/krUabT2y4rc0wvqcpaeCmVnOMdXJZ8K+Mad18ksQq4zwt7S9Gaww+qeDqcmq9o/h1YPigl9LfkavS68p0YSmuGWMSX/EtmzYZC6AASAAAAAAAAAABFUngAErVEk88sbmcv8AUHP5IbRW2wuqX/G+GPJPfzKtvRycu5uW3ogdChQUVrnyFChn/J0KdFJCwiksE0cd0Y0sDzqNkVa3U4ShJZjJNNNZ2MZ4PsHRv7qmvpjFY/l4mb2MSvb6fGE6lRY46nN4xt2NEJNRa6Mzy3eSfAuCRRE4StxJyR4DA4BcE5GuI3BINaIDIxoa0SNDWBOTIePqEpUKaW/8Vf0OxpVlGlQp04rZRTfq9y9e2sakOGXqn2Y5JJJdlgedTMFEaKWclepSTWGjlXNs4+h2pNdyGeGsGeSTNMJtFTTNUlTkoyeYPb0Nbb1lKKaaw+xibqhj0LGkai6cuGT+R/szbaXel6JvYrubZTWuHJtUKQ0aikk11JTrnKFAAAAEyBydRvJcXw4fU9slVWrGnHLHhBzeEW7m+hDZvfsjlXF1KquGnCWPT+5gvGmuTpVnb0puMoJOpNc+JrPCu2Fg4mk+Lbq3mpKrOpDPzU6knKMl28vUxSqVJ7SelenP3NkaKjHMd36nqdLR6j58MfV7/sWI6NL9cV6JnS067jWpU60fpqQjNZ5rK5FstVjRx1f3KXdVc/4cGekT6TT/AA0NelVF1NAAOwpev5I7zUM7KyrLll/kb8arHnH3WTSCcK7IR2EV7smie8t8pHBpaguUlgu05qW6eSW40+E+mH3Rx61GVGXdP9zPUp1KO8t4+ZZHRU93aXqdRoYFGopRUv8A3I6QjEGAKyne18fKubK5NIsjFt4RJVrxjzf4RWVzObxTi36LJiPFXiGdGq6NLCnFJzm1lptZ4UuRzdL8c3dGac6nxoZXFCaW8fJrkPCk5PxvC9C1xwvCsv6npqsq8uccerQ+Ok1Oskvzk6lhdRq0qdWH0VIRnHvhrkWjcrCl6v7mR3NT0Rw/9zS/WNejS/Ud4Ce40fL9h3qp5mdnpE/JlOvpclzh7GuASX8fTfDaHjeVEZmyvJU1wyTa8+aO7a3cZrZ7i1rWMucV6nHu7KVJ8cG8f0DNa3W/igvyHs6z+WX6NBkUpadc8cM9Vs/UuG6ElKKkuGZZRcW0+UIcCjNRuW5bbtJ9jQHH1O3g/m5MzXVOTSnHmO5dQkk3F9TzH/aLpMoXU7iK4qVbD4o7qM8JOL7ckZS0t5VJxhTXFOTwl/d+R7E6UJLDmmns01lMdbWFKGeCMI558MFHPsYHWcnnT+zdFaYpZOhpFWnb21Gjx8bp04wbjvmXX9yxLWodIt/nBUjbQXP+o+KprovyWO5rNbYRn7Kn5Nse9c7U/wDuBa33gvcdFQfSPsiT4MX9qF7av8/6DFJfCNhrMOqaJ4apTfVr1IZWkH0IpadHplfkdV7hdUxdFF+aOlG9pv717lLVbiDhhNN+RVel9pY/AsNNSe8s/gKlxVnBxcVuTGnSjJNS4F0xfJ+S3IIRSWEIzPjCwDeZNiM5tV4qpy5ZTOiyKtQUuYss7NDwkk9zznx/oU/9RO7pRdWlVxKTgsuE0sNNduuTJ2NhOtNQpwbbeG2mox82z2n/AHev1v2HQsYrrn8YLXWk99OGWRlFLGc/YXSrunb29KjFyqfDhGHEo4Ta5vfzLE9aX2x9yFW0P05/IqjFfaiXdVntlL6IqVOn5Ng9al0jH9wWsy6xX7iSqJdl+EQyuo+Qnb1fnf4HVOD+D9luOtd4+xNDWIPmmjlyuIdkNcoP7V7jK6rL4v0Dt4P4WjuR1Km/uS9SK8vYOnJJp5XI4MqkMvPEvRZIK1zBL5eJvzRMr6o1hpbhG0jnO539CT+Z9MnZOTpNVOCa5M6qZ07eGimlnJiqy1TbEk9jP61V2x32O7WexldUqZmkRcy00pD28dVRENGI+VZ8l0FoodKj2OLE6uVncRNjiSFLYf8ACLxHJECyT0LqUX3QvwR0LfLJyLJxa3OlCWVkfkjisIXIZMWB7Y3ImRGyHIBWNARsrGEBgxCHsSIxrHNjRRkIVrqrhbc2WGQ14ZQIeOM7nNnNvmMwWXSGfCHya1JFaQmcFl0yOdMVyQ6kitnJBNFpxwQVEZZY6Do7OgV9nHs8I09N7GJ0aeKmO+5sbZ7HobSeqjFnFuo6arQy7lhGVu5ZmzS3z2MtUfzv1K77+sez99lmiW4FSl0LcTlI1zJ4QRL8IjgyaLHRnlkSNIkSEyLkkRgKNyGQIwOEEEyBIrEyI2BBICMGxrYpIDWxWxGAwjGtg2NbFbGQkkRNEkmR5EbHQ1or1WTykVpitlsUQzRWmWplaoIaIiWMsVEbWylsjD0Prj6mz0+WyO3/ABz9k16nLv8A+wXUORl6n1v1NTfrYytf62Pfr2eRbN+0x6FikXISKVJlmT+V+j/oclGyaLUJZ3RMmc7Sq/HTXeD4H6ovpjFDRKmLkjyOyMJgdkMiABAogCZIAUa2DYhBOAEbBsRsgYRjWK2NbAlAxrFbGSYjY6GtjZSFlIhkxCxIZJkcmSMq3tThhJvs1+RVuy2OwcWdyCYtpn4cM8+FZ9RtUhlsSOl9cfU2WnckY+3WZr3Nhp/I7n8cvZN+py79+1LF7HYyV7HFQ2NxHYymq08PPmaLmGqk0UUJaaiZHSkXabOdRkXKbOEjqSRzdHuOC6r0JbZl8SOeqk+X7GiRkPFcJUp0buCbcXwTS5cPdmmsbqNSEZxeVJZLZeZnksltMXIxMXIouB6YuSPIZAjBJkUiyLkAwOyI2MbFbIJwLkaDY1sjJOBWxGxrY1sUbAspEUpCyYxsVjpCSI2xzY1issSGs4XiCu38OjH6qk035RTR2q1RRi5SeEll+hmtCbuLupcS+imuGHZ81kemuZdES+DQuGFhckVajLc3sUJsrLoomsI5mbKwjsZXSIb57mvtI4R6K0hooxRxLmeqq2TVFsZzV6WUzTNHK1GlszQUoylOWC7SkUq8eGeO5LSmcCvT7ObR2ac+0ppl6vbxq05057qawYvQ9RnZV5W1f6G9n29PLc2VKZzfE+hK5p8cNqkF8r7rsFNprSyuWz9DQU6iklJPKa2a6j8nnHh3xJO3l8CvlxTwn1i/8HoFGvGcVKDUovk0LKOkVonAbkMi5IwOyA0CQFyDG5DIoYFyNbEbEbIGSByGSYNjWyGMkDYxsVsYxB0KMkxKk1FZbwurZi/EPiTj/hUOrw5rr6DQg5P0JewviPVnWnG2o5fE8Sa3y/8ABqdNsI29GFNc+cv5upzPCWg/Ch8aqv4k1mKf2o7tSWd2WVZKK0oWO7KtxLoUZ7+pPWnlkdvBymvLmRb0u0qJdC6rPsqbfU7uj0dkaWjHCOXptHCR14o9JwcAcV7mGUWBskAGQ1e16o5lKXf8mt1G3ymZO8pOEs+/oZLuh2kcrlGq1rdnLD4Zap1C7Qq4OTTqdSzTqHDzh5OpKGSp4k8NQuYudPEauOnKXqYyy1S4sqnBJPCe8JZSfoz0mjXaI9S0ujcwxUis9JL6kaozUlhmdpx+hT0jxLRrx+pQn1jLbfy7naUjzPWfCdag+Om3Uh3isSX4K1h4luKHytuSX2z+r3ZDpZ90MJnqzYjZj7LxvTlj4kHT9HxHcttZoT+mpH/maiVOMlygwdNsRyK8LmD5Ti/SSZIpCE4HOQjY1yIp3EVzkl6sjJOCVsRsoXGr0Yc6kPxJNnEvPGVKP/5xc35/KvclQk+EMahs5Wp65Sop8Uk5fpT3MVf+Jq9VuMXwRf2x3fuO0rwzXuJcU1KEesp54seWS2ND5iNQ3UdZrXUlTgmot4jGOcv1NT4Z8KqlirWxKps4x6R/8nX0nQ6NtH5UnPGHNrd/4LlSpn0GnUjFYiLvIbVnn0KN1V6E1eokc2pUMm7ZohHCz0GzmdTSLXLy1v1OdaUnOXka3TbbCR3bK37KGXyzmXdftJYXCOha08ItDIRwh5tMYAAABBWp5RwNUtMmlaK1xQTQAefTi4SfbPsWKczq6lYeRwJpwlh8jmXdnnxw+6OjbXSXgnwdOFQnhUxyOZCqWI1Dk5aN7jlHVhc9/cp3+h29dPipxTf3RSUvcjjUJY1C1Vmil0vIy994D5ulU26Rkm37nCufC1zDfgbS+5YPS4XL9SaNyupcq6EcZo8gdK4htipH+VSBXtyvvre8j2H4sfIZKnB84x9hu0XoRl+R5D/q7l/fWf5mEYXE9v4r9eI9ejTprlGPsPVSK7B2iQZfkeVUPDNzN5VP8yeDuWPgOT3qVEl+lR/ubh3KI53DfISVcnEn0KOn+HbehvGCk+8vmOi66W0f/hBKTfNjJSS5lMqrZKh5j3JvmQVqyW3UirXP6SjOp3ZVu/qXxh1fA+rU333IIwc3hdxkczeEd/TbA61pZ4xOf2RiubrPggT6XZ4S2NHb08IhtaGEi4kdQ5oYFAAAAAAABGhQACncW6fQz2oadnOxrGitXoJgB53Vt5Qba5DqVwn6mpvtPyZ6901rlsZK9nCrutma6N1Kns90EahJGocvinDZrKJadyn1x6nIq21SnyvwdOnXp1OHg6amOUyjGoPVQzFukuqaHKZSVQPiARpL+QyUfiCOo+4EaDocaGSqpFCVQZKYE9n5l2d12K06rfMqzuUvP0IeOU9orBopWtSpwtiqdalT+pPUrpf4QylRlNrOy7dCzaaa3u9zQ2Om46HYt7KFPd7s5te7nU2WyKmnady2NFbW2CShb4LSRrMgJCgAAAAAAAAAAAAAAAjFAAIp00yjcWKfQ6YjQAZW70s41xpfkb6dJMqVbJPoBJ57O0nHk3/Ya6k10NvW05dilV0lFM7elLlFsa9SPDMurp/pfuI7ryZoJ6P5epDLSPIo7hRZd32qcX/VeTFd2/0v3OytI8iWGkeX7ArCiugO9q+Zn/iTfTCHRtJy5t79tjUU9JXLBcpaYljYvhb04cIplXqT5ZmLfSm+h2LXS/I7tKyS6FuFBIuwikoW1ikdCnSSJEhQATAoAAAAAAAAAAH/2Q=="
            // }]
            i = this.page_index
            axios.interceptors.response.use(function (resp) {
                // ???????????????????????????
                if(resp.data.msg=="page out of range"){
                    console.log("try")
                    if(that.page_index!=0&&that.page_index>=i) {
                        alert("out of range")
                        that.page_index--
                    }
                    }
                return resp;
            }, function (error) {
                return resp;
            })
            axios.get(api_url + "/search", {
                params: {
                    content: this.content,
                    price_min: this.price_min,
                    price_max: this.price_max,
                    page_size: this.page_size,
                    page_index: this.page_index
                }
            }).then(function (resp) {
                console.log(resp)
                // handle success
                if (resp.data.code === 0) {
                    that.commodityList = resp.data.data.list
                } else {
                    console.log(resp.data.msg)
                }
            }).catch(function (resp) {
                // handle error
                console.log(resp.data.msg)   
            })
        },
        search: function () {
            const that = this
            that.seen = false
            console.log("search commodity")
            if (this.price_max < this.price_min) {
                alert("invalid prize gap")
                return
            }
            if (this.price_min === this.price_max) {
                this.price_max = 999999
            }
            axios.get(api_url + "/search", {
                params: {
                    content: this.content,
                    price_min: this.price_min,
                    price_max: this.price_max,
                    page_size: this.page_size,
                    page_index: this.page_index
                }
            }).then(function (resp) {
                console.log(resp)
                // handle success
                if (resp.data.code === 0) {
                    that.commodityList = resp.data.data.list
                    that.seen = true
                } else {
                    console.log(resp.data.msg)                           
                }
            }).catch(function (resp) {
                // handle error
				console.log(resp.data.msg)
            })
        },
    },
})