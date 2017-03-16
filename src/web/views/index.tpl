<!DOCTYPE html>
<html lang="zh-cn">
{{template "common/header_flat.tpl"}}
<title>山顶洞</title>
<body>
{{template "common/navibar.tpl"}}
<div class="container">
    <div class="jumbotron">
        <h1>朋友们</h1>
        <p>这是一个神奇的世界</p>
        <p><a href="https://github.com/jspdba/goweb.git" target="_blank">GitHub:https://github.com/jspdba/goweb.git</a></p>
        <p>
            <a class="btn btn-large btn-primary" type="button" role="button" href="{{urlfor "BookController.List"}}">进入 &raquo;</a>
        </p>
    </div>

    <div class="row">
        <div class="col-lg-10">
            <div class="row">
                <div class="col-xs-12 col-sm-4 col-md-3 col-lg-3">
                    <div class="tile">
                        <img src="/static/Flat-UI/img/icons/svg/compas.svg" alt="二维码" class="tile-image big-illustration">
                        <h3 class="tile-title">二维码</h3>
                        <p></p>
                        <a class="btn btn-primary btn-large btn-block" href="{{urlfor "QrCodeController.Index"}}">进入</a>
                    </div>
                </div>
                <div class="col-xs-12 col-sm-4 col-md-3 col-lg-3">
                    <div class="tile">
                        <img src="/static/Flat-UI/img/icons/svg/paper-bag.svg" alt="二维码解析" class="tile-image big-illustration">
                        <h3 class="tile-title">二维码解析</h3>
                        <p></p>
                        <a class="btn btn-primary btn-large btn-block" href="{{urlfor "QrCodeController.Decode"}}">进入</a>
                    </div>
                </div>
                <div class="col-xs-12 col-sm-4 col-md-3 col-lg-3">
                    <div class="tile">
                        <img src="/static/Flat-UI/img/icons/svg/toilet-paper.svg" alt="收藏" class="tile-image big-illustration">
                        <h3 class="tile-title">收藏</h3>
                        <p></p>
                        <a class="btn btn-primary btn-large btn-block" href="{{urlfor "LinkController.List"}}">进入</a>
                    </div>
                </div>
                <div class="col-xs-12 col-sm-4 col-md-3 col-lg-3">
                    <div class="tile">
                        <img src="/static/Flat-UI/img/icons/svg/ribbon.svg" alt="任务" class="tile-image big-illustration">
                        <h3 class="tile-title">定时任务</h3>
                        <p></p>
                        <a class="btn btn-primary btn-large btn-block" href="{{urlfor "JobController.List"}}">进入</a>
                    </div>
                </div>
            </div>
        </div>

    </div>

</div>

{{template "common/script.tpl"}}
</body>
</html>
