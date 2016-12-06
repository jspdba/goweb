<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>edit</title>

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
    <h1>添加/编辑</h1>
    <form id="Link" action="/link/save" role="form" method="post">
        <input id="Id" name="Id" type="hidden" value="">
        <div class="form-group">
            <label for="url">链接地址</label>
            <input type="text" class="form-control" id="Url" name="Url" placeholder="http(s)://" value="">
        </div>
        <div class="form-group">
            <label for="title">标题</label>
            <input type="text" class="form-control" id="Title" name="Title" placeholder="标题">
        </div>
        <div class="form-group">
            <label for="description">描述</label>
            <input type="text" class="form-control" id="Description" name="Description" placeholder="描述">
        </div>
        <div class="form-group">
            <label for="tags">标签</label>
            <input type="text" class="form-control" id="Tags.Name" name="Tags.Name" placeholder="标签">
        </div>
        <a href="javascript:void(0)" class="btn btn-default urlInfo">获取信息</a>
        <button type="submit" class="btn btn-default">提交</button>
    </form>
</div>
<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
<script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
<!-- Include all compiled plugins (below), or include individual files as needed -->
<script src="/static/js/bootstrap.min.js"></script>
<script>
    $(".urlInfo").bind("click",function(){
        var url="/link/info";
        var data={
            url:$("#Url").val()
        }

        if(data.url){
            $.getJSON(url,data,function (data) {
                if(data && data.code==0){
                    data.data.title && $("#Title").val(data.data.title);
                    data.data.description && $("#Description").val(data.data.description);
                }else{
                    console.log(data.msg,data.code)
                }
            })
        }
    })
</script>
</body>
</html>