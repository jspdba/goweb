<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="utf-8">
    <title>图书编辑</title>
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
    <h4>添加/修改图书</h4>
    <form class="form-horizontal" role="form" action="/book/save" method="post">
        <input type="hidden" name="Id" value="{{.entry.Id}}">
        <div class="form-group">
            <div class="col-sm-10">
                <input class="form-control" type="text" id="Name" name="Name" placeholder="标题" value="{{.entry.Name}}">
            </div>
        </div>
        <div class="form-group">
            <div class="col-sm-10">
                <input class="form-control" type="text" id="Maker" name="Maker" placeholder="作者" value="{{.entry.Maker}}">
            </div>
        </div>
        <div class="form-group">
            <div class="col-sm-10">
                <input class="form-control" type="text" id="ChapterRules" name="ChapterRules" placeholder="章节规则" value="{{.entry.ChapterRules}}">
            </div>
        </div>
        <div class="form-group">
            <div class="col-sm-10">
                <input class="form-control" type="text" id="ContentRules" name="ContentRules" placeholder="内容规则" value="{{.entry.ContentRules}}">
            </div>
        </div>
        <div class="form-group">
            <div class="col-sm-10">
                <input class="form-control" type="text" id="Url" name="Url" placeholder="图书地址" value="{{.entry.Url}}">
            </div>
        </div>
        <div class="form-group">
            <div class="col-sm-10">
                <textarea class="form-control" rows="3" id="Content" name="Content" placeholder="描述">{{.entry.Content}}</textarea>
            </div>
        </div>
        <button type="submit" class="btn btn-primary">提交</button>
    </form>
</div><!-- /.container -->

<script src="/static/Flat-UI/dist/js/vendor/jquery.min.js"></script>
<!-- Include all compiled plugins (below), or include individual files as needed -->
<script src="/static/Flat-UI/dist/js/vendor/video.js"></script>
<script src="/static/Flat-UI/dist/js/flat-ui.min.js"></script>
<script src="/static/Flat-UI/docs/assets/js/application.js"></script>

<script>

</script>
</body>
</html>
