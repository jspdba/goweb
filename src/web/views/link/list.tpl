<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>收藏列表</title>

    <!-- Bootstrap -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="http://cdn.bootcss.com/html5shiv/3.7.2/html5shiv.min.js"></script>
    <script src="http://cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
</head>
<body>
<div class="container">
    <div class="panel panel-primary">
        <div class="panel-heading">
            <div class="title"><a href="/link/edit" class="btn btn-primary">添加</a><a href="{{urlfor "LinkController.Import"}}" class="btn btn-primary">远程导入</a></div>
        </div>
        <div class="panel-body">
            <div class="list-group">
                {{range .page.List}}
                <div class="list-group-item">
                    <a href="{{.Url}}" title="{{.Description | html}}" target="_blank">{{.Title}}</a>
                    <a href="{{urlfor "LinkController.Delete" ":id" .Id}}">
                        <span class="badge pull-right">删除</span></a>
                </div>
                {{end}}
            </div>
        <ul id="page"></ul>
        </div>
</div>

<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
<script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
<!-- Include all compiled plugins (below), or include individual files as needed -->
<script src="/static/js/bootstrap.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-paginator.min.js"></script>
<script type="text/javascript">
    $(function () {
        var pageNo=parseInt("{{.page.PageNo}}");
        pageNo=pageNo?pageNo:1
        var totalPage=parseInt("{{.page.TotalPage}}");
        totalPage=totalPage?totalPage:0
        pageNo = pageNo>totalPage?totalPage:pageNo

        $("#page").bootstrapPaginator({
            currentPage: pageNo,
            totalPages: totalPage,
            bootstrapMajorVersion: 3,
            size: "small",
            onPageClicked: function(e,originalEvent,type,page){
                window.location.href = "/link/list?PageNo=" + page
            }
        });
    });
</script>
</body>
</html>