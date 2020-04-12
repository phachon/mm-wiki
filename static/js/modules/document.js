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
    ListTree: function (element, treeData, defaultId, isEditor, isDelete) {

        //配置信息
        var setting = {
            view: {
                showIcon: showIconForTree,
                addHoverDom: addHoverDom,
                removeHoverDom: removeHoverDom
            },
            edit: {
                enable: true,
                showRemoveBtn: true,
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
                onDrag: onDrag,
                beforeDrop: beforeDrop,
                onDrop: onDrop
            },
            drag: {
                isCopy: false,
                isMove: true,
                prev: true,
                inner: true,
                next: true
            }
        };

        if (isDelete == false) {
            setting.edit.showRemoveBtn = false;
        }

        function beforeClick(treeId, treeNode) {
            console.log("点击节点前....");
            // $("#mainFrame").attr("src", "/page/view?document_id="+treeNode.id);
            location.href = "/document/index?document_id=" + treeNode.id
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
            console.log(treeNode);
            if (treeNode.isParent) {
                if (treeNode.children && treeNode.children.length > 0) {
                    Layers.failedMsg("请先删除或移动目录下所有文档！");
                    return false;
                }
            }

            var title = '<i class="fa fa-volume-up"></i> 确定要删除文档吗？';
            layer.confirm(title, {
                btn: ['是', '否'],
                skin: Layers.skin,
                btnAlign: 'c',
                title: "<i class='fa fa-warning'></i><strong> 警告</strong>"
            }, function () {
                Common.ajaxSubmit("/document/delete?document_id=" + treeNode.id);
                // location.href = "/document/index?document_id="+moveNode.id;
            }, function () {

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

        function beforeDrag(treeId, treeNodes) {
            console.log("拖拽节点前...");
            if (isEditor == false) {
                return false;
            }
            // if (treeNodes[0].isParent) {
            //     return false;
            // }
            return true;
        }

        function onDrag() {
            console.log("拖拽节点中...");
            return true;
        }

        function beforeDrop(treeId, treeNodes, targetNode, moveType) {
            console.log("拖拽数据:", treeId, treeNodes, targetNode, moveType);
            console.log("拖拽节点完成...");
            if (isEditor == false) {
                return false;
            }
            var moveNode = treeNodes[0];
            // 文档当前层级排序
            if (moveType === "prev" || moveType === "next") {
                let moveUrl = "/document/move?move_type=" + moveType
                    + "&document_id=" + moveNode.id
                    + "&target_id=" + targetNode.id;
                Common.ajaxSubmit(moveUrl, moveUrl);
                return false;
            }

            // 移动文档到目录中
            if (moveNode.isParent) {
                return false;
            }
            if (!targetNode.isParent) {
                return false;
            }
            
            var title = '<i class="fa fa-volume-up"></i> 确定要移动文档吗？';
            layer.confirm(title, {
                btn: ['是', '否'],
                skin: Layers.skin,
                btnAlign: 'c',
                title: "<i class='fa fa-warning'></i><strong> 警告</strong>"
            }, function () {
                Common.ajaxSubmit("/document/move?document_id=" + moveNode.id + "&target_id=" + targetNode.id);
            }, function () {

            });

            return false;
        }

        function onDrop(treeId, treeNodes, targetNode, moveType) {
            console.log("拖拽节点完成中...");
            return false;
        }

        function addHoverDom(treeId, treeNode) {
            if (isEditor == false) {
                return false;
            }
            if (treeNode.isParent === false || treeNode.isParent === undefined) {
                return false
            }
            var sObj = $("#" + treeNode.tId + "_span");
            var addBtn = $("#addBtn_" + treeNode.tId);
            if (addBtn.length > 0) return;

            var spanHtml = "<span class='button add' id='addBtn_" + treeNode.tId + "' title='新建文档' onfocus='this.blur();'></span>";
            sObj.append(spanHtml);

            // bind add
            var addBtn = $("#addBtn_" + treeNode.tId);
            if (addBtn) addBtn.bind("click", function () {
                var content = "/document/add?space_id=" + treeNode.spaceId + "&parent_id=" + treeNode.id;
                layer.open({
                    type: 2,
                    skin: Layers.skin,
                    title: '<strong>创建文档</strong>',
                    shadeClose: true,
                    shade: 0.6,
                    maxmin: true,
                    area: ["800px", "345px"],
                    content: content,
                    padding: "10px"
                });
                return false;
            });
        }

        function removeHoverDom(treeId, treeNode) {
            $("#addBtn_" + treeNode.tId).unbind().remove();
        }

        $(document).ready(function () {
            $.fn.zTree.init($(element), setting, treeData);
            var zTreeMenu = $.fn.zTree.getZTreeObj("dir_tree");
            var node = zTreeMenu.getNodeByParam("id", defaultId);
            zTreeMenu.selectNode(node, true);
            zTreeMenu.expandNode(node, true, false);
            //initialize fuzzysearch function
            fuzzySearch("dir_tree", '#document_search', null, true);
        });
    }
};