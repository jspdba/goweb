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
        <button type="submit" class="btn btn-primary">提交</button>
    </form>
</div><!-- /.container -->

{{template "common/script.tpl"}}
<script>
</script>
</body>
</html>
