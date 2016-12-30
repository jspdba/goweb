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
          <h1 class="text-center text-primary">注册</h1>
        </div>
        <div class="modal-body">
          {{template "common/flash_error.tpl" .}}
          <div class="panel panel-default">
            <div class="panel-heading">注册</div>
            <div class="panel-body">
              <form action="/register" method="post">
                <div class="form-group">
                  <label for="username">用户名</label>
                  <input type="text" id="username" name="username" class="form-control" placeholder="用户名">
                </div>
                <div class="form-group">
                  <label for="password">密码</label>
                  <input type="password" id="password" name="password" class="form-control" placeholder="密码">
                </div>
                <input type="submit" class="btn btn-sm btn-default" value="注册"> <a href="/login">去登录</a>
              </form>
            </div>
          </div>
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

