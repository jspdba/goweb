{{template "../header.tpl"}}
<style>
    body{
        font-family: 'microsoft yahei',Arial,sans-serif;
        margin:0;
        padding:0;
    }

</style>
<body>
<div class="container">
    <div id="loginModal" class="modal show">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close">x</button>
                    <h1 class="text-center text-primary">登录</h1>
                </div>
                <div class="modal-body">
                    {{template "common/flash_error.tpl" .}}
                    <form action="/login" method="post" role="form" class="form col-md-12 center-block">
                        <div class="form-group">
                            <input type="text" name="username" class="form-control input-lg" placeholder="用户名">
                        </div>
                        <div class="form-group">
                            <input name="password" type="password" class="form-control input-lg" placeholder="登录密码">
                        </div>
                        <div class="form-group">
                            <button class="btn btn-primary btn-lg btn-block">立刻登录</button>
                            <span>
                                <a href="#">找回密码</a>
                                <label>
                                    <input name="remberme" type="checkbox"> Remember me
                                </label>
                            </span>
                            <span><a href="/register" class="pull-right">注册</a></span>
                        </div>
                    </form>
                </div>
                <div class="modal-footer">
                </div>
            </div>
        </div>
    </div>
</div>

{{template "../footer.tpl"}}
</body>
</html>
