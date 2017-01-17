<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>章节列表</title>
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
    <div class="panel panel-primary">
        <div class="panel-heading">
            <div class="title">
            </div>
        </div>
        <div class="panel-body">
            <ul class="list-group">
                {{range .page.List}}
                <li class="list-group-item">
                    <a href="{{urlfor "ChapterController.Detail" ":id" .Id}}" >{{.Title}}</a>
                </li>
                {{end}}
            </ul>
            <ul id="page"></ul>
        </div>
    </div>
</div>

<script src="/static/Flat-UI/dist/js/vendor/jquery.min.js"></script>
<!-- Include all compiled plugins (below), or include individual files as needed -->
<script src="/static/Flat-UI/dist/js/vendor/video.js"></script>
<script src="/static/Flat-UI/dist/js/flat-ui.min.js"></script>

<script src="/static/Flat-UI/docs/assets/js/application.js"></script>

<script type="text/javascript" src="/static/js/bootstrap-paginator.min.js"></script>
    <script type="text/javascript">
        $(function () {
            var pageNo=parseInt("{{.page.PageNo}}");
            pageNo=pageNo?pageNo:1
            var totalPage=parseInt("{{.page.TotalPage}}");
            totalPage=totalPage?totalPage:0
            pageNo = pageNo>totalPage?totalPage:pageNo

            totalPage && $("#page").bootstrapPaginator({
                currentPage: pageNo,
                totalPages: totalPage,
                bootstrapMajorVersion: 3,
                size: "small",
                onPageClicked: function(e,originalEvent,type,page){
                    var thisUrl=window.location.href
                    if(thisUrl.indexOf("?")>-1){
                        window.location.href = thisUrl.substr(0,thisUrl.indexOf("?"))+"?PageNo="+page
                    }else{
                        window.location.href = thisUrl+"?PageNo="+page
                    }
                }
            });
        });
    </script>
</body>
</html>
