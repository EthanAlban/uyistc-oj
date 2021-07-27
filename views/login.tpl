<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>用户登录</title>
</head>
<body>
<div>
    <h1>login</h1>
    <form action="" method="post">
        <input type="text" name="username" id="name" placeholder="用户名">
        <input type="password" name="password" id="password" placeholder="密码">
        <button class="but" name="button" onclick="sub()">登陆</button>
    </form>
</div>
<script>
    function sub(){
        console.log("ok")
        $ajax({

        })
    }
</script>
</body>
</html>