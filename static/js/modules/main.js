/**
 * main.js
 * Copyright (c) 2018 phachon@163.com
 */

var Main = {
    
    Search:function (searchField) {

        $(searchField).bind('input propertychange', function() {
            var _keywords = $(this).val();
            searchLazy(_keywords);
        });

        var timeoutId = null;

        var metaChar = '[\\[\\]\\\\\^\\$\\.\\|\\?\\*\\+\\(\\)]'; //js meta characters
        var rexMeta = new RegExp(metaChar, 'gi');//regular expression to match meta characters

        function textFilter(_keywords) {

            $("#collection_documents > li > a > span").each(function () {
                var docName = $(this).text();
                if (_keywords) {
                    if (docName.indexOf(_keywords) !== -1) {
                        var newKeywords = _keywords.replace(rexMeta,function(matchStr){
                            return '\\' + matchStr;
                        });
                        var rexGlobal = new RegExp(newKeywords, 'gi');
                        var newDocName = docName.replace(rexGlobal, function(originalText){
                            var highLightText =
                                '<span class="collect_search_text" style="color: whitesmoke;background-color: darkred;">'
                                + originalText
                                +'</span>';
                            return 	highLightText;
                        });
                        $(this).html(newDocName);
                        $(this).parents('li').show()
                    }else {
                        $(this).parents('li').hide()
                    }
                }else {
                    $(this).html(docName);
                    $(this).parents('li').show()
                }
            });
            console.log(_keywords)
        }

        function searchLazy(_keywords) {
            if (timeoutId) {
                clearTimeout(timeoutId);
            }
            timeoutId = setTimeout(function() {
                textFilter(_keywords);
                $(searchField).focus();
            }, 500);
        }
    }
};