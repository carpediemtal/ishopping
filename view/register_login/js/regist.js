var oUpass = document.getElementById("upass")
var eye = document.getElementById("eye")
var toLogin = document.getElementById("toLogin")
var flag = 0
eye.onclick = function() {
    if(flag == 0){
        oUpass.type = 'text';
        eye.innerHTML = '&#xe61b;';
        flag++;
    }
    else{
        oUpass.type = 'password';
        eye.innerHTML = '&#xe61d;';
        flag--;
    }
}
function fnRegist() {
    //获取下拉框的value
    var obj = document.getElementById("droplist")
    var index = obj.selectedIndex;
    var selectOpt = obj.options[index].value
    
    var oUname = document.getElementById("uname")
    var userName = oUname.value 
    var passWord = oUpass.value
    
    //简单错误判断
    var oError = document.getElementById("error_box")
    var isError = true;
    oError.style.fontSize= '5px'
    oError.style.color= '#fff';
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

   
    var user_Type = 0
    if (selectOpt =="seller"){
        user_Type = 1
    }
    var dataObj = {username: userName,password: passWord,user_type:user_Type}
    // console.log(JSON.stringify(dataObj))
    fetch('http://ishopping.gq/api/register',{
            method: 'POST',
            body: JSON.stringify(dataObj)
        })
        .then(res => res.json())
        .then(json =>{
            if (json.code){
                oError.innerHTML = json.msg
                isError = false;
                return;
            }else{
                console.log(json.msg)
                oError.innerHTML = "Successful registration, about to jump"
                //注册成功跳转至登录页面
                setInterval(function(){
                    window.location.href="../login.html"
                },1000)
            }
        })



}