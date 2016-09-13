<!-- Navigation -->
<nav class="navbar navbar-inverse navbar-fixed-top" role="navigation">
    <!-- Brand and toggle get grouped for better mobile display -->
    <div class="navbar-header">
        <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-ex1-collapse">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
        </button>
        <ul class="nav top-nav">
            <li class="dropdown"> <a class="navbar-brand" href="/">高校论文管理平台 <small>{{.Module}}端</small></a>  </li>
            <li class="dropdown">
                <form class="navbar-form navbar-left" action="{{urlfor "BaseController.R"}}" method="post">
                    <div class="form-group">
                        <input type="tel" name="id" class="form-control" placeholder="请输入课题ID">
                        <input type="text" name="operate" style="display: none" value="9">
                    </div>
                    <button type="submit" class="btn btn-default">快速跳转</button>
                </form>
            </li>
        </ul>

    </div>
    <!-- Top Menu Items -->
    <ul class="nav navbar-right top-nav">

        <li class="dropdown">
            <a href="#">还有{{.endDay}}天结束</a>
        </li>
        {{if eq .CurrentUser.Type 5}}
        <li class="dropdown">
            <a class="navbar-brand" href="#" data-toggle="modal" data-target="#edit_time_modal">修改截至日期</a>
        </li>
        {{end}}
        <li class="dropdown">
            <a class="navbar-brand" href="{{urlfor "ProfileController.Get"}}">{{.CurrentUser.Name}}</a>
        </li>
        <li class="dropdown">
            <a class="navbar-brand" href="{{urlfor "LoginController.Logout"}}"><i class="fa fa-fw fa-power-off"></i> 登出</a>
        </li>
    </ul>
    <!-- Sidebar Menu Items - These collapse to the responsive navigation menu on small screens -->
    <div class="collapse navbar-collapse navbar-ex1-collapse">
        <ul class="nav navbar-nav side-nav">
            <li {{if eq .pageKey "index"}}  class="active" {{end}}  >
                <a href="{{urlfor "AdminController.Get"}}"><i class="fa fa-fw fa-dashboard"></i> 主页</a>
            </li>
            <li {{if eq .pageKey "profile"}}  class="active" {{end}} >
                <a href="{{urlfor "ProfileController.Get"}}"><i class="fa fa-fw fa-bar-chart-o"></i> 用户资料</a>
            </li>
            <li {{if eq .pageKey "list"}}  class="active" {{end}} >
                <a href="{{urlfor "ListController.Get"}}"><i class="fa fa-fw fa-table"></i> 课题列表</a>
            </li>
            {{if eq .CurrentUser.Type 1}}
            <li {{if eq .pageKey "mylist"}}  class="active" {{end}} >
                <a href="{{urlfor "ListController.Select"}}"><i class="fa fa-fw fa-edit"></i> 我的课题</a>
            </li>
            {{end}}
        </ul>
    </div>
    <!-- /.navbar-collapse -->
</nav>


<!-- 编辑Modal -->
<div class="modal fade" id="edit_time_modal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <form role="form" action="{{urlfor "AdminController.Post"}}" method="post">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal"><span
                                aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                    <h4 class="modal-title" id="myModalLabel">编辑截至日期</h4>
                </div>
                <div class="modal-body">
                    <div class="form-group">
                        <input type="text" class="form-control" name="end_date" placeholder="请输入结束日期:(年-月-日)" value="{{dateformat .endDate.Date "2006-01-02"}}">
                    </div>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
                    <button type="submit" class="btn btn-primary">保存</button>
                </div>
            </div>
        </div>
    </form>
</div>