var div=document.querySelector('.mainbody-ad')
var img = document.getElementById('img')
var lis =document.querySelectorAll('.ad-page>a')
var prev=document.querySelector('.prev')
var next=document.querySelector('.next')

// 定义变量表示图片下标志
var index=4;

// 定义函数 功能是切换图片  li的背景颜色和边框颜色
function change() {
    img.src='static/img/'+index+'.png';
    //取消边框颜色
    for (var i=0;i<lis.length;i++){
        lis[i].style.cssText='border:2px solid #fff;border-color: rgba(255,255,255,0.3);background-color: #f5f5f5;background-color:rgba(0,0,0,0.5);';
    }

    lis[index].style.cssText='background-color: #fff;background-color:rgba(255,255,255,0.4);border-color: #757575;border-color: rgba(0,0,0,0.4);';
}

//自动切换图片
function autoChange() {
    index++;
    if (index>4){
        index=0;
    }
    change();
}
// 启动定时器
var timer=setInterval(autoChange,5000);

// 设置鼠标移入效果
div.onmouseover=function () {
    clearInterval(timer);
};
div.onmouseout=function () {
    timer=setInterval(autoChange,5000);
}

prev.onclick=function () {
    index--;
    if(index<0)
    {
        index=4;
    }
    change();

}
next.onclick=function () {
    index++;
    if(index>4)
    {
        index=0;
    }
    change();

}
for (var i=0;i<lis.length;i++)
{
    lis[i].id=i;
    lis[i].onclick=function () {
        index=this.id;
        change();
    }
}