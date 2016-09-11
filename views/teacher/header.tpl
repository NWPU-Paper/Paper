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
        <a class="navbar-brand" href="{{urlfor "IndexController.Get"}}">{{.CurrentUser.Name}}</a>
    </div>
    <!-- Top Menu Items -->
    <ul class="nav navbar-right top-nav">

        <li class="dropdown">
            <a href="{{urlfor "LoginController.Logout"}}"><i class="fa fa-fw fa-power-off"></i> 登出</a>
        </li>
    </ul>
    <!-- Sidebar Menu Items - These collapse to the responsive navigation menu on small screens -->
    <div class="collapse navbar-collapse navbar-ex1-collapse">
        <ul class="nav navbar-nav side-nav">
            <li  class="active">
                <a href="index.html"><i class="fa fa-fw fa-dashboard"></i> 主页</a>
            </li>
            <li>
                <a href="charts.html"><i class="fa fa-fw fa-bar-chart-o"></i> 用户资料</a>
            </li>
            <li>
                <a href="tables.html"><i class="fa fa-fw fa-table"></i> 课题</a>
            </li>
            <li>
                <a href="forms.html"><i class="fa fa-fw fa-edit"></i> 任务书</a>
            </li>
            <li>
                <a href="bootstrap-elements.html"><i class="fa fa-fw fa-desktop"></i> 学生</a>
            </li>
            <li>
                <a href="paper.html"><i class="fa fa-fw fa-wrench"></i> 论文</a>
            </li>
            <li>
                <a href="blank-page.html"><i class="fa fa-fw fa-file"></i> 答辩通知</a>
            </li>
            <li>
                <a href="grades.html"><i class="fa fa-fw fa-dashboard"></i> 成绩</a>
            </li>
        </ul>
    </div>
    <!-- /.navbar-collapse -->
</nav>