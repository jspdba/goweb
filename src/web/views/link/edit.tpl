<!DOCTYPE html>
<html lang="zh-cn">
{{template "common/header_flat.tpl"}}
<title>添加收藏</title>
<body>
{{template "common/navibar.tpl"}}
<div class="container">
    <form id="Link" action="/link/save" role="form" method="post">
        <input id="Id" name="Id" type="hidden" value="">
        <div class="form-group">
            <input type="text" class="form-control" id="Url" name="Url" placeholder="链接地址" value="">
        </div>
        <div class="form-group">
            <input type="text" class="form-control" id="Title" name="Title" placeholder="标题">
        </div>
        <div class="form-group">
            <input type="text" class="form-control" id="Description" name="Description" placeholder="描述">
        </div>
        <div class="form-group">
            <input type="text" class="form-control" id="Tags.Name" name="Tags.Name" placeholder="标签">
        </div>
        <a href="javascript:void(0)" class="btn btn-primary urlInfo">获取信息</a>
        <button type="submit" class="btn btn-primary">提交</button>
    </form>
</div>
{{template "common/script.tpl"}}
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