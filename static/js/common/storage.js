
if(!window.localStorage){
    alert("浏览器支持localstorage");
}

var Storage = {

    get: function (key) {
        var v = localStorage.getItem(key);
        if (v == null || v == undefined) {
            return ""
        }
        return Base64.decode(v);
    },

    set: function (key, value) {
        localStorage.setItem(key, Base64.encode(value));
    },

    remove: function (key) {
        localStorage.removeItem(key)
    }
};