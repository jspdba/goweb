<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="utf-8">
    <title>{{.entity.Title}}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <!-- Loading Bootstrap -->
    <link href="/static/Flat-UI/dist/css/vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet">

    <!-- Loading Flat UI -->
    <link href="/static/Flat-UI/dist/css/flat-ui.min.css" rel="stylesheet">

    <link rel="shortcut icon" href="/static/Flat-UI/dist/img/favicon.ico">

    <!-- HTML5 shim, for IE6-8 support of HTML5 elements. All other JS at the end of file. -->
    <!--[if lt IE 9]>
    <script src="/static/Flat-UI/dist/js/vendor/html5shiv.js"></script>
    <script src="/static/Flat-UI/dist/js/vendor/respond.min.js"></script>
    <![endif]-->

    <style>
        body {
            padding-top: 70px;
            padding-bottom: 50px;
        }
    </style>
</head>
<body>

<!-- Static navbar -->
<div class="navbar navbar-default navbar-fixed-top" role="navigation">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                <span class="sr-only">Toggle navigation</span>
            </button>
            <a class="navbar-brand">Icon</a>
        </div>
        <div class="navbar-collapse collapse">
            <ul class="nav navbar-nav">
                <li class="active"><a href="/">主页</a></li>
            </ul>
            <ul class="nav navbar-nav navbar-right">
                <li><a>注册</a></li>
                <li class="active"><a>登录</a></li>
            </ul>
        </div><!--/.nav-collapse -->
    </div>
</div>

<div class="container">
    <h4>{{.entity.Title}}</h4>
    <p>
        {{str2html .entity.Content}}
    </p>
    <!--
    <div class="navbar-fixed-bottom">
        <a id="pre" type="button" class="btn btn-primary btn-sm" href="{{urlfor "ChapterController.Detail" ":id" .pre.Id}}">上一页</a>
        <a id="next" type="button" class="btn btn-primary btn-sm" href="{{urlfor "ChapterController.Detail" ":id" .next.Id}}">下一页</a>
    </div>
    -->
    <div class="navbar-fixed-bottom">
        <div class="row">
            <div class="col-md-12">
                <form action="{{urlfor "ChapterController.FindByTitle" ":id" .entity.Book.Id}}" class="form-inline" role="form" method="post">
                    <div class="form-group">
                        <div class="input-group">
                            <input name="Title" type="text" class="form-control" placeholder="搜索">
                            <span class="input-group-btn">
                                <button class="btn btn-default" type="submit">Go!</button>
                                <a id="pre" type="button" class="btn btn-primary btn-sm" href="{{urlfor "ChapterController.Detail" ":id" .pre.Id}}">上一页</a>
                                <a id="next" type="button" class="btn btn-info btn-sm" href="{{urlfor "ChapterController.Detail" ":id" .next.Id}}">下一页</a>
                            </span>
                        </div>
                    </div>
                </form>
            </div><!-- /.col-md-12 -->
        </div><!-- /.row -->
    </div>
</div><!-- /.container -->



<!-- Modal -->
<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title" id="myModalLabel">添加 Tag</h4>
            </div>
            <div class="modal-body">
                <input id="tag" type="text" value="tag1">
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                <button id="useTag" type="button" class="btn btn-primary">使用</button>
            </div>
        </div>
    </div>
</div>


<script src="/static/Flat-UI/dist/js/vendor/jquery.min.js"></script>
<!-- Include all compiled plugins (below), or include individual files as needed -->
<script src="/static/Flat-UI/dist/js/vendor/video.js"></script>
<script src="/static/Flat-UI/dist/js/flat-ui.min.js"></script>
<script src="/static/Flat-UI/docs/assets/js/application.js"></script>
<script src="/static/js/js.cookie.js"></script>

<!--消息-->
<link href="/static/css/toastr.min.css" rel="stylesheet"/>
<script src="/static/js/toastr.min.js"></script>

<script>

    var tag = Cookies.get("__TAG")
    if(!tag){
        $('#myModal').modal('show')
    }
    //使用tag
    $("#useTag").bind("click",function (data) {
        var value = $("#tag").val();
        if(value){
            if($(this).hasClass("sure")){
                Cookies.set('__TAG', value, { expires: 365 });
                $('#myModal').modal('hide')
            }else{
                var url="/log/tag/"+value;
                $.getJSON(url,function (data) {
                    if(data.code==0){
                        if(data.result==0){
                            Cookies.set('__TAG', value, { expires: 365 });
                            $('#myModal').modal('hide')
                        }else{
                            toastr.warning(value+" 已存在！")
                            $("#useTag").html("确定使用")
                            if(!$("#useTag").hasClass("sure")){
                                $("#useTag").addClass("sure")
                            }
                        }
                    }
                })
            }
        }
    });

    var url="{{urlfor "ChapterLogController.Add" ":tag" "tag1" ":bookId" .entity.Book.Id ":index" .entity.Index}}";
    if(url){
      $.getJSON(url,function (data) {
          if(data.code==0){
              console.log("添加日志成功")
          }
      })
    }
</script>
</body>
</html>
