{{template "../header.tpl"}}
<body>
<form class="form-horizontal" role="form">
    <div class="form-group">
        <label for="email" class="col-sm-2 control-label">邮箱: </label>
        <div class="col-sm-10">
            <input type="email" class="form-control" id="email" placeholder="Email">
        </div>
    </div>
    <div class="form-group">
        <label for="pwd" class="col-sm-2 control-label">密码: </label>
        <div class="col-sm-10">
            <input type="password" class="form-control" id="pwd" placeholder="Password">
        </div>
    </div>
    <div class="form-group">
        <div class="col-sm-offset-2 col-sm-10">
            <div class="checkbox">
                <label>
                    <input name="remberme" type="checkbox"> Remember me
                </label>
            </div>
        </div>
    </div>
    <div class="form-group">
        <div class="col-sm-offset-2 col-sm-10">
            <button type="submit" class="btn btn-default">登录</button>
        </div>
    </div>
</form>
{{template "../footer.tpl"}}
</body>
</html>
