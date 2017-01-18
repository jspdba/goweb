<!DOCTYPE html>
<html lang="zh-cn">
{{template "common/header_flat.tpl"}}
<title>章节列表</title>
<body>
{{template "common/navibar.tpl"}}
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
{{template "common/script.tpl"}}
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
