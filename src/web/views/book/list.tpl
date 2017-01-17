<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>图书列表</title>
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
            min-height: 2000px;
            padding-top: 70px;
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
            <a class="navbar-brand" href="/">家</a>
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
    <div class="panel panel-primary">
        <div class="panel-heading">
            <div class="title">
                <a href="{{urlfor "BookController.Edit"}}" class="btn btn-primary">添加</a>
                <a id="check" class="btn btn-primary">检查更新</a>
            </div>
        </div>
        <div class="panel-body">
            <div class="table-responsive">
                <table class="table table-bordered">
                    {{range .page.List}}
                        <tr>
                            <td colspan="2"><a href="{{urlfor "BookController.Edit" ":id" .Id}}" title="{{.Name}}">{{.Name}}</a></td>
                            <td>
                                <div class="btn-group pull-right">
                                    <a type="button" class="btn btn-primary" href="{{urlfor "ChapterController.List" ":id" .Id}}">查看章节</a>
                                    <a type="button" class="btn btn-info updateChapter" link="{{urlfor "BookController.TaskUpdate" ":id" .Id}}">更新章节</a>
                                    <a type="button" class="btn btn-info" id="toRead" href="{{urlfor "ChapterController.ListByLog" ":tag" "tag1" ":id" .Id}}">继续阅读</a>
                                </div>
                            </td>
                            <td class="news" bid="{{.Id}}" data-url="{{urlfor "ChapterController.HasNewChapter" ":id" .Id}}">
                            </td>
                        </tr>
                    {{end}}
                </table>
            </div>
            <ul id="page"></ul>
        </div>
</div>

<script src="/static/Flat-UI/dist/js/vendor/jquery.min.js"></script>
<!-- Include all compiled plugins (below), or include individual files as needed -->
<script src="/static/Flat-UI/dist/js/vendor/video.js"></script>
<script src="/static/Flat-UI/dist/js/flat-ui.min.js"></script>
    <script>
        $(".updateChapter").bind("click",function () {
            var url=$(this).attr("link")
            console.log(url)
            url && $.getJSON(url,function (data) {
                console.log(data)
                if (data.code==0){
                    alert("success")
                }
            })
        });

        $("#check").bind("click",function () {
            $(".news").each(function () {
                var url=$(this).data("url");
                var that=$(this);
                $.getJSON(url,function (data) {
                    if(data.code==0){
                        $(that).html(data.data)
                    }
                })
            })
        })
    </script>
</body>
</html>
