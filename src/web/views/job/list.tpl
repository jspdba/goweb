<!DOCTYPE html>
<html lang="zh-cn">
{{template "common/header_flat.tpl"}}
<title>job列表</title>
<body>
{{template "common/navibar.tpl"}}
<div class="container">
    <div class="panel panel-primary">
        <div class="panel-heading">
            <div class="title">
                <a href="{{urlfor "JobController.Edit"}}" class="btn btn-primary">添加</a>
            </div>
        </div>
        <div class="panel-body">
            <div class="table-responsive">
                <table class="table table-bordered">
                    {{range .page.List}}
                    <tr>
                        <td colspan="2"><a href="{{urlfor "JobController.Edit" ":id" .Id}}" title="{{.Name}}">{{.Name}}</a></td>
                        <td>
                            <div class="btn-group pull-right">
                                {{if eq 0 .State }}
                                <a type="button" class="btn btn-primary" href="{{urlfor "JobController.Start" ":id" .Id}}">开启任务</a>
                                {{else}}
                                <a type="button" class="btn btn-primary" href="{{urlfor "JobController.Pause" ":id" .Id}}">终止任务</a>
                                {{end}}
                                <a type="button" class="btn btn-primary" href="{{urlfor "JobController.Delete" ":id" .Id}}">删除任务</a>
                            </div>
                        </td>
                        <td>{{if eq 0 .State }} 停止 {{else}} 运行 {{end}}</td>
                    </tr>
                    {{end}}
                </table>
            </div>
            <ul id="page"></ul>
        </div>
    </div>
    {{template "common/script.tpl"}}
</body>
</html>
