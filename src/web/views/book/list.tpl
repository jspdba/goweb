<!DOCTYPE html>
<html lang="zh-cn">
{{template "common/header_flat.tpl"}}
<title>图书列表</title>
<body>
{{template "common/navibar.tpl"}}
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
                                    <a type="button" class="btn btn-primary" id="toRead" href="{{urlfor "ChapterController.ListByLog" ":tag" "tag1" ":id" .Id}}">继续阅读</a>
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
    {{template "common/script.tpl"}}
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
