/**
 * Copyright (c) 2018 phachon@163.com
 */

var Page = {

    /**
     * ajax Save
     * @param element
     * @returns {boolean}
     */
    ajaxSave: function (element) {

        /**
         * 成功信息条
         * @param message
         * @param data
         */
        function successBox(message, data) {
            Layers.successMsg(message)
        }

        /**
         * 失败信息条
         * @param message
         * @param data
         */
        function failed(message, data) {
            Layers.failedMsg(message)
        }

        /**
         * request success
         * @param result
         */
        function response(result) {
            //console.log(result)
            if (result.code == 0) {
                failed(result.message, result.data);
            }
            if (result.code == 1) {
                successBox(result.message, result.data);
            }
            if (result.redirect.url) {
                var sleepTime = result.redirect.sleep || 3000;
                setTimeout(function () {
                    parent.location.href = result.redirect.url;
                }, sleepTime);
            }
        }

        layer.prompt({
            title: '<i class="fa fa-volume-up"></i> 请输入修改备注',
            formType: 2,
            maxlength: 150,
            value: '',
            area: ['340px', '80px']
        }, function(comment, index, elem){
            if (comment.trim()) {
                layer.close(index);
                var options = {
                    dataType: 'json',
                    success: response,
                    data: {'comment': comment}
                };
                $(element).ajaxSubmit(options);
            }else {
                elem.focus()
            }
        });

        return false;
    },

    /**
     * cancel save
     */
    cancelSave: function (title, url) {
        title = '<i class="fa fa-volume-up"></i> '+title;
        layer.confirm(title, {
            btn: ['是','否'],
            skin: Layers.skin,
            btnAlign: 'c',
            title: "<i class='fa fa-warning'></i><strong> 警告</strong>"
        }, function() {
            parent.location = url
        }, function() {

        });
    }

};