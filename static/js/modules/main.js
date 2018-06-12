/**
 * main.js
 * Copyright (c) 2018 phachon@163.com
 */

$(function () {
   Main.loadCollectionPage()
});


var Main = {

    loadCollectionPage: function() {
      $("a[name='collection_page']").each(function () {
          $(this).bind('click', function (e) {
              var url = this.href || $(this).attr("data-link");
              $.get(url, function(data, status){
                  $("#main-right").html(data);
              });
          })
      })
    },

    GetPage: function (pageId) {
        $.get("/page/info?page_id="+pageId, function(data, status){
            $("#main-right").html(data)
        }, 'html');
    },

    EditPage: function (pageId) {
        $.get("/page/edit?page_id="+pageId, function(data, status){
            $("#main-right").html(data)
        }, 'html');
    }
};

