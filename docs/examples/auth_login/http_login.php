<?php
// 获取数据
$username = isset($_POST["username"]) ? $_POST["username"] : "";
$password = isset($_POST["password"]) ? $_POST["password"] : "";
$extData = isset($_POST["ext_data"]) ? $_POST["ext_data"] : "";

if (!$username || !$password) {
echo json_encode(['message'=>'参数错误', 'data'=>[]], JSON_UNESCAPED_UNICODE);
exit();
}
// 1. ext_data 可用于接口安全验证
if ($extData != "login api token") {
$result = [
    'message' => '登录接口验证失败',
    'data' => [],
];
echo json_encode($result, JSON_UNESCAPED_UNICODE);
exit();
}

// 2. 验证用户名密码是否正确
// ....
// ....

// 3. 成功返回
$result = [
'message' => '',
'data' => [
    'given_name' => '王哈哈',
    'mobile' => '111111111111',
    'phone' => '010-9929921',
    'email' => 'root@mmWiki.com',
    'department' => '广告事业部.技术部.系统开发组',
    'position' => '高级JAVA开发工程师',
    'location' => 'B座32层E区1002',
    'im' => 'QQ：12211',
],
];
echo json_encode($result, JSON_UNESCAPED_UNICODE);