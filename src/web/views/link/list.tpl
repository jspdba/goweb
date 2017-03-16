<!DOCTYPE html>
<html lang="zh-cn">
{{template "common/header_flat.tpl"}}
<title>收藏列表</title>
<body>
{{template "common/navibar.tpl"}}
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
                    {{range .Tags}}
                        <span class="tag badge pull-right">{{.Name}}</span>
                    {{end}}
                </div>
                {{end}}
            </div>
        <ul id="page"></ul>
        </div>
</div>
{{template "common/script.tpl"}}
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