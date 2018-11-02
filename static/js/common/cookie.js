
if(!window.localStorage){
    alert("浏览器支持localstorage");
}else{
}

var Cookie = Cookies.withConverter({
    read: function (value, name) {
        return Base64.decode(value);
    },
    write: function (value, name) {
        return Base64.encode(value);
    }
});