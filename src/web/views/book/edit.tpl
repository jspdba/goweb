<!DOCTYPE html>
<html lang="zh-cn">
{{template "common/header_flat.tpl"}}
<title>图书编辑</title>
<body>
{{template "common/navibar.tpl"}}

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
        <button id="urlInfo" type="button" class="btn btn-primary" link="{{urlfor "BookController.UrlInfo"}}">链接搜索</button>
        <button id="search" type="button" class="btn btn-primary" link="{{urlfor "BookController.Search"}}">标题搜索</button>
        <button type="submit" class="btn btn-primary">提交</button>
    </form>
</div><!-- /.container -->

{{template "common/script.tpl"}}
<script>
    function trim(str){ //删除左右两端的空格
        return str.replace(/[作者：]/g,"").replace(/(^\s*)|(\s*$)/g, "");
    }
    $("#urlInfo").bind("click",function () {
        var url=$("#Url").val();
        if(url && url.indexOf("http://www.biquge")>=0){
            var link=$(this).attr("link")
            $.getJSON(link,{Url:url},function (data) {
                if(data.code==0){
                    $("#Name").val(trim(data.result.Name))
                    $("#Maker").val(trim(data.result.Maker))
                    $("#ChapterRules").val(trim(data.result.ChapterRules))
                    $("#ContentRules").val(trim(data.result.ContentRules))
                    $("#Content").val(trim(data.result.Content))
                }
            })
        }
    })
    $("#search").bind("click",function () {
        var name=$("#Name").val();
        if(name){
            var link=$(this).attr("link")
            $.getJSON(link,{Name:name},function (data) {
                if(data.code==0){
                    $("#Name").val(trim(data.result.Name))
                    $("#Maker").val(trim(data.result.Maker))
                    $("#ChapterRules").val(trim(data.result.ChapterRules))
                    $("#ContentRules").val(trim(data.result.ContentRules))
                    $("#Content").val(trim(data.result.Content))
                    $("#Url").val(trim(data.result.Url))
                }
            })
        }
    })
</script>
</body>
</html>
