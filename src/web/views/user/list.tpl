{{template "../header.tpl" .}}
<body>
<div class="container">
    <div class="panel panel-primary">
        <div class="panel-heading">
            <div class="title"><a href="/link/edit" class="btn btn-primary">添加</a></div>
        </div>
        <div class="panel-body">
            <ul class="list-group">
                {{range .page.List}}
                <li class="list-group-item">
                    {{.Username}}:{{.Password}}
                </li>
                {{end}}
            </ul>
            <ul id="page"></ul>
        </div>
    </div>

    {{template "../footer.tpl" .}}
    <script type="text/javascript" src="/static/js/bootstrap-paginator.min.js"></script>
    <script type="text/javascript">
        $(function () {
            var pageNo=parseInt("{{.page.PageNo}}");
            pageNo=pageNo?pageNo:1
            var totalPage=parseInt("{{.page.TotalPage}}");
            totalPage=totalPage?totalPage:0
            pageNo = pageNo>totalPage?totalPage:1

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