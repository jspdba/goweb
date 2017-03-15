<!-- Static navbar -->
<div class="navbar navbar-default navbar-fixed-top" role="navigation">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                <span class="sr-only">Toggle navigation</span>
            </button>
            <a class="navbar-brand" href="/">家</a>
        </div>
        <div class="navbar-collapse collapse">
            <!--<ul class="nav navbar-nav">
                <li class="active"><a href="/">主页</a></li>
            </ul>-->
            <ul class="nav navbar-nav navbar-right">
                <li><a href="{{urlfor "QrCodeController.Index"}}">二维码</a></li>
                <li><a href="{{urlfor "QrCodeController.Decode"}}">二维码解析</a></li>
                <li><a href="{{urlfor "LinkController.List"}}">收藏</a></li>
                <li><a href="{{urlfor "JobController.List"}}">任务</a></li>
                <!--<li class="active"><a>登录</a></li>-->
            </ul>
        </div><!--/.nav-collapse -->
    </div>
</div>