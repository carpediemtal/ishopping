var oUpass = document.getElementById("upass")
var flag = 0
eye.onclick = function () {
    if (flag == 0) {
        oUpass.type = 'text';
        eye.innerHTML = '&#xe61b;';
        flag++;
    } else {
        oUpass.type = 'password';
        eye.innerHTML = '&#xe61d;';
        flag--;
    }
}

function fnLogin() {

    var oUname = document.getElementById("uname")
    var userName = oUname.value
    var passWord = oUpass.value

    //简单错误判断
    var oError = document.getElementById("error_box")
    var isError = true;
    oError.style.fontSize = '5px'
    oError.style.color = '#fff';
    if (oUname.value.length > 12 || oUname.value.length < 3) {
        oError.innerHTML = "Please enter 3-12 characters for user name";
        isError = false;
        return;
    }

    if (oUpass.value.length > 20 || oUpass.value.length < 6) {
        oError.innerHTML = "Please enter 6-20 characters for password"
        isError = false;
        return;
    }

    //设置 cookie 的函数
    function setCookie(cname, cvalue, exdays) {
        var d = new Date();
        d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
        var expires = "expires=" + d.toUTCString();
        document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
    }

    var dataObj = {username: userName, password: passWord}

    // fetch('http://ishopping.gq/api/login',{
    fetch('http://localhost:7001/api/login', {
        method: 'POST',
        // headers: {
        //     'Content-Type': 'application/json'
        // },
        body: JSON.stringify(dataObj)
    })
        .then(res => res.json())
        .then(json => {
            if (json.code) {
                oError.innerHTML = json.msg
                isError = false;
                return;
            } else {
                //token设置
                setCookie('jwt', json.data.token, 1)

                oError.innerHTML = 'Login successful, about to jump'
                //登陆成功跳转至首页
                setInterval(function () {
                    window.location.href = "/"
                }, 1000)
            }
        })
}

