/**
 * Copyright (c) 2017 phachon
 */
var Page = {

    /**
     * list nav
     */
    ListTree: function (element) {

        //配置信息
        var setting = {
            view: {
                showIcon: showIconForTree,
                addHoverDom: addHoverDom,
                removeHoverDom: removeHoverDom
            },
            edit: {
                // enable: true,
                // showRemoveBtn: true,
                // showRenameBtn: true,
                removeTitle: '删除',
                renameTitle: '修改'
            },
            data: {
                simpleData: {
                    enable: true
                }
            },
            callback: {
                beforeEditName: beforeEditName,
                beforeRemove: beforeRemove,
                beforeRename: beforeRename,
                onRemove: onRemove,
                onRename: onRename
            },
            drag: {
                isCopy: true,
                isMove: true,
                prev: true,
                inner: true,
                next: true
            }
        };

        var listUrl = '';
        // var treeData =[];
        var treeData = [
            { id:10, pId:0, name:"首页", open:true},
            { id:11, pId:0, name:"技术部"},
            { id:12, pId:0, name:"产品部"},
            { id:13, pId:0, name:"运营部", isParent:true},
            { id:111, pId:11, name:"前端开发组", isParent:true},
            { id:112, pId:11, name:"UED组", isParent:true},
            { id:113, pId:11, name:"后台开发组", isParent:true},
            { id:114, pId:11, name:"应用开发组", isParent:true},
            { id:115, pId:11, name:"基础研发组", isParent:true},

            { id:121, pId:111, name:"前端CSS规范"},
            { id:122, pId:112, name:"UED 设计标准"},
            { id:123, pId:113, name:"PHP 代码规范"},
            { id:124, pId:114, name:"广告应用架构"},
            { id:125, pId:115, name:"消息中间件的设计"},

            { id:121, pId:12, name:"视频组", isParent:true},
            { id:122, pId:12, name:"社交组", isParent:true},
            { id:123, pId:12, name:"会员组", isParent:true},
            { id:124, pId:12, name:"广告组", isParent:true}
        ];
        // $.ajax({
        //     async: false,
        //     url: '/slong/index_header/ajaxGetAllIndexHeader',
        //     type: 'get',
        //     dataType: 'json',
        //     data: {},
        //     success: function(response) {
        //         var headerNavs = response.data;
        //         for (var i = 0; i < headerNavs.length; i++) {
        //             treeData.push({'id': parseInt(headerNavs[i].index_header_id), 'pId': parseInt(headerNavs[i].parent_id), 'name': headerNavs[i].name});
        //         }
        //     },
        //     error: function(response) {
        //         console.log(response.messages);
        //         Common.errorAlert(response.messages);
        //     }
        // });

        /**
         * 开始修改节点
         * @param  string treeId id
         * @param  object treeNode 节点信息
         * @return boolean
         */
        function beforeEditName(treeId, treeNode) {
            //console.log("开始修改节点...");
            var headerNavId = treeNode.id;
            layer.open({
                type: 2,
                skin: 'layui-layer-lan',
                title: '修改导航信息',
                shadeClose: true,
                shade: 0.6,
                maxmin: true,
                area: ['1000px', '500px'],
                content: '/slong/index_header/edit?index_header_id=' + headerNavId
            });
            return true;
        }

        /**
         * 修改完成
         * @param  string  treeId
         * @param  object  treeNode
         * @param  string  newName
         * @param  boolean isCancel
         * @return boolean
         */
        function beforeRename(treeId, treeNode, newName) {
            //console.log("修改完成提交...");
            return true;
        }

        /**
         * 修改之后
         */
        function onRename(e, treeId, treeNode, isCancel) {
            // console.log("修改之后刷新...");
            // setTimeout(function() {
            // 	location.href = listUrl;
            // }, 2000);
            return true;
        }

        /**
         * 删除节点前
         * @param  string  treeId
         * @param  object  treeNode
         * @return boolean
         */
        function beforeRemove(treeId, treeNode) {
            if(("children" in treeNode)) {
                Common.warningAlert("");
                return false;
            }
            swal({
                    title: "警告",
                    text: "<h4>确定要删除？</h4>",
                    html: true,
                    type: "warning",
                    showCancelButton: true,
                    confirmButtonClass: "btn-danger",
                    confirmButtonColor: "#d9534f",
                    confirmButtonText: "是",
                    cancelButtonText: "否",
                    closeOnConfirm: false
                },
                function() {
                    $.ajax({
                        async: false,
                        url: '',
                        type: 'get',
                        dataType: 'json',
                        data: {index_header_id: treeNode.id},
                        success: function(response) {
                            if(response.code == 1) {
                                Common.successAlert(response.messages);
                                setTimeout(function() {
                                    location.href = listUrl;
                                }, 2000);
                            }else {
                                console.log(response.messages);
                                Common.errorAlert(response.messages);
                            }
                        },
                        error: function(response) {
                            console.log(response.messages);
                            Common.errorAlert(response.messages);
                        }
                    });
                });
            return false;
        }

        function onRemove(e, treeId, treeNode) {
            console.log("删除节点...");
            return false;
        }

        function showIconForTree(treeId, treeNode) {
            return true;
        }

        var newCount = 1;

        function addHoverDom(treeId, treeNode) {
            if (treeNode.isParent == false || treeNode.isParent == undefined) {
                return false
            }
            var sObj = $("#" + treeNode.tId + "_span");
            if (treeNode.editNameFlag || $("#addBtn_"+treeNode.tId).length > 0) return;
            var addStr = "<span class='button add' id='addBtn_" + treeNode.tId
                + "' title='新建' onfocus='this.blur();'></span>";
            sObj.after(addStr);
            var btn = $("#addBtn_"+treeNode.tId);
            if (btn) btn.bind("click", function(){
                var zTree = $.fn.zTree.getZTreeObj("dir_tree");
                zTree.addNodes(treeNode, {id:(100 + newCount), pId:treeNode.id, name:"新建文件" + (newCount++)});
                return false;
            });
        }

        function removeHoverDom(treeId, treeNode) {
            $("#addBtn_"+treeNode.tId).unbind().remove();
        }

        $(document).ready(function(){
            $.fn.zTree.init($(element), setting, treeData);
        });
    }
};