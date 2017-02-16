<!DOCTYPE html>
<html lang="zh-cn">
{{template "common/header_flat.tpl"}}
<title>job编辑</title>
<body>
{{template "common/navibar.tpl"}}

<div class="container">
    <h4>添加/修改job</h4>
    <form class="form-horizontal" role="form" action="/job/save" method="post">
        <input type="hidden" name="Id" value="{{.entity.Id}}">
        <div class="form-group">
            <div class="col-sm-10">
                <input class="form-control" type="text" id="Name" name="Name" placeholder="名称" value="{{.entity.Name}}">
            </div>
        </div>
        <div class="form-group">
            <div class="col-sm-10">
                <select class="form-control" id="BookId" name="BookId">
                    {{range .page.List}}
                        {{if $.entity.BookId}}
                            {{if eq $.entity.BookId .Id}}
                                <option value="{{.Id}}" selected="selected">{{.Name}}</option>
                            {{else}}
                                <option value="{{.Id}}">{{.Name}}</option>
                            {{end}}
                        {{else}}
                            <option value="{{.Id}}">{{.Name}}</option>
                        {{end}}
                    {{end}}
                </select>
            </div>
        </div>
        <div class="form-group">
            <div class="col-sm-10">
                {{if .entity.Cron}}
                <input class="form-control" type="text" id="Cron" name="Cron" placeholder="cron" value="{{.entity.Cron}}">
                {{else}}
                <input class="form-control" type="text" id="Cron" name="Cron" placeholder="cron" value="0 0 8 * * 1-5">
                {{end}}
            </div>
        </div>
        <div class="form-group">
            <div class="col-sm-10">
                <textarea class="form-control" rows="3" id="Content" name="Content" placeholder="描述">{{.entity.Content}}</textarea>
            </div>
        </div>
        <button type="submit" class="btn btn-primary">提交</button>
    </form>
</div><!-- /.container -->

{{template "common/script.tpl"}}
<script>
    $("#BookId").bind("change",function () {
        $("#Name").val($(":selected").html()+"_job");
    })
</script>
</body>
</html>
