/**
 * Copyright (c) 2018 phachon@163.com
 * // var treeData = [
 //     { id:10, pId:0, name:"首页", open:true},
 //     { id:11, pId:0, name:"技术部"},
 //     { id:12, pId:0, name:"产品部"},
 //     { id:13, pId:0, name:"运营部", isParent:true},
 //     { id:111, pId:11, name:"前端开发组", isParent:true},
 //     { id:112, pId:11, name:"UED组", isParent:true},
 //     { id:113, pId:11, name:"后台开发组", isParent:true},
 //     { id:114, pId:11, name:"应用开发组", isParent:true},
 //     { id:115, pId:11, name:"基础研发组", isParent:true},
 //
 //     { id:121, pId:111, name:"前端CSS规范"},
 //     { id:122, pId:112, name:"UED 设计标准"},
 //     { id:123, pId:113, name:"PHP 代码规范"},
 //     { id:124, pId:114, name:"广告应用架构"},
 //     { id:125, pId:115, name:"消息中间件的设计"},
 //
 //     { id:121, pId:12, name:"视频组", isParent:true},
 //     { id:122, pId:12, name:"社交组", isParent:true},
 //     { id:123, pId:12, name:"会员组", isParent:true},
 //     { id:124, pId:12, name:"广告组", isParent:true}
 // ];
 */
var Document = {

    /**
     * list nav
     */
    ListTree: function (element, treeData) {

        //配置信息
        var setting = {
            view: {
                showIcon: showIconForTree,
                addHoverDom: addHoverDom,
                removeHoverDom: removeHoverDom
            },
            edit: {
                enable: true,
                showRemoveBtn: false,
                showRenameBtn: false
                // removeTitle: '删除',
                // renameTitle: '修改'
            },
            data: {
                simpleData: {
                    enable: true
                }
            },
            callback: {
                beforeClick: beforeClick,
                onClick: onClick,

                beforeEditName: beforeEditName,

                beforeRemove: beforeRemove,
                onRemove: onRemove,

                beforeRename: beforeRename,
                onRename: onRename,

                beforeDrag: beforeDrag,
                onDrag : onDrag,
                beforeDrop  : beforeDrop,
                onDrop   : onDrop
            },
            drag: {
                isCopy: false,
                isMove: true,
                prev: true,
                inner: true,
                next: true
            }
        };

        function beforeClick(treeId, treeNode) {
            console.log("点击节点前....");
            $("#mainFrame").attr("src", "/page/view?document_id="+treeNode.id);
            // location.href = "/document/index?document_id="+treeNode.id
        }

        function onClick() {
            console.log("点击节点后....");
        }

        function beforeEditName(treeId, treeNode) {
            console.log("开始修改节点...");
            return true;
        }

        function beforeRename(treeId, treeNode, newName) {
            console.log("修改完成提交...");
            return true;
        }

        function onRename(e, treeId, treeNode, isCancel) {
            console.log("修改之后刷新...");
            // setTimeout(function() {
            // 	location.href = listUrl;
            // }, 2000);
            return true;
        }

        function beforeRemove(treeId, treeNode) {
            console.log("删除节点前...");
            return false;
        }

        function onRemove(e, treeId, treeNode) {
            console.log("删除节点...");
            return false;
        }

        function showIconForTree(treeId, treeNode) {
            return true;
        }
        
        function beforeDrag(treeId, treeNodes) {
            console.log("拖拽节点前...");
            return true;
        }

        function onDrag() {
            console.log("拖拽节点中...");
            return true;
        }

        function beforeDrop(treeId, treeNodes, targetNode, moveType) {
            console.log("拖拽节点完成...");
            return true;
        }

        function onDrop() {
            console.log("拖拽节点完成中...");
            return false;
        }

        function addHoverDom(treeId, treeNode) {
            if (treeNode.isParent === false || treeNode.isParent === undefined) {
                return false
            }
            var sObj = $("#" + treeNode.tId + "_span");
            var addBtn = $("#addBtn_"+treeNode.tId);
            var editBtn = $("#editBtn_"+treeNode.tId);
            if (addBtn.length > 0) return;
            if (editBtn.length > 0) return;

            var spanHtml =
                "<span class='button add' id='addBtn_"+treeNode.tId+"' title='新建文档' onfocus='this.blur();'></span>"+
                "<span class='button edit' id='editBtn_"+treeNode.tId+"' title='修改目录' onfocus='this.blur();'></span>";
            sObj.append(spanHtml);

            // bind add
            var addBtn = $("#addBtn_"+treeNode.tId);
            if (addBtn) addBtn.bind("click", function() {
                var content = "/document/add?space_id="+treeNode.spaceId+"&parent_id="+treeNode.id;
                layer.open({
                    type: 2,
                    skin: Layers.skin,
                    title: '<strong>创建文档</strong>',
                    shadeClose: true,
                    shade : 0.6,
                    maxmin: true,
                    area: ["800px", "345px"],
                    content: content,
                    padding:"10px"
                });
                return false;
            });

            // bind edit
            var editBtn = $("#editBtn_"+treeNode.tId);
            if (editBtn) editBtn.bind("click", function() {
                var content = "/document/edit?space_id="+treeNode.spaceId+"&document_id="+treeNode.id;
                layer.open({
                    type: 2,
                    skin: Layers.skin,
                    title: '<strong>修改目录</strong>',
                    shadeClose: true,
                    shade : 0.6,
                    maxmin: true,
                    area: ["800px", "345px"],
                    content: content,
                    padding:"10px"
                });
                return false;
            });
        }

        function removeHoverDom(treeId, treeNode) {
            $("#addBtn_"+treeNode.tId).unbind().remove();
            $("#editBtn_"+treeNode.tId).unbind().remove();
        }

        $(document).ready(function(){
            $.fn.zTree.init($(element), setting, treeData);
        });
    }
};