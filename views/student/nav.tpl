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
            <li class="dropdown"> <a class="navbar-brand" href="/">高校论文管理平台<small>学生端</small></a>  </li>
        </ul>

    </div>
    <!-- Top Menu Items -->
    <ul class="nav navbar-right top-nav">
        <li class="dropdown">
            <a class="navbar-brand" href="{{urlfor "TeacherController.Profile"}}">{{.CurrentUser.Name}}</a>
        </li>
        <li class="dropdown">
            <a href="{{urlfor "LoginController.Logout"}}"><i class="fa fa-fw fa-power-off"></i> 登出</a>
        </li>
    </ul>
    <!-- Sidebar Menu Items - These collapse to the responsive navigation menu on small screens -->
    <div class="collapse navbar-collapse navbar-ex1-collapse">
        <ul class="nav navbar-nav side-nav">
            <li {{if eq .pageKey "index"}}  class="active" {{end}}  >
                <a href="{{urlfor "TeacherController.Index"}}"><i class="fa fa-fw fa-dashboard"></i> 主页</a>
            </li>
            <li {{if eq .pageKey "profile"}}  class="active" {{end}} >
                <a href="{{urlfor "TeacherController.Profile"}}"><i class="fa fa-fw fa-bar-chart-o"></i> 用户资料</a>
            </li>
            <li {{if eq .pageKey "subject"}}  class="active" {{end}} >
                <a href="{{urlfor "TeacherController.Subject"}}"><i class="fa fa-fw fa-table"></i> 课题</a>
            </li>
            <li {{if eq .pageKey "assignment"}}  class="active" {{end}} >
                <a href="{{urlfor "TeacherController.Assignment"}}"><i class="fa fa-fw fa-edit"></i> 任务书</a>
            </li>
            <li {{if eq .pageKey "student"}}  class="active" {{end}}>
                <a href="{{urlfor "TeacherController.Student"}}"><i class="fa fa-fw fa-desktop"></i> 学生</a>
            </li>
            <li {{if eq .pageKey "paper"}}  class="active" {{end}}>
                <a href="{{urlfor "TeacherController.Paper"}}"><i class="fa fa-fw fa-wrench"></i> 论文</a>
            </li>
            <li {{if eq .pageKey "message"}}  class="active" {{end}}>
                <a href="{{urlfor "TeacherController.Message"}}"><i class="fa fa-fw fa-file"></i> 答辩通知</a>
            </li>
            <li {{if eq .pageKey "score"}}  class="active" {{end}}>
                <a href="{{urlfor "TeacherController.Score"}}"><i class="fa fa-fw fa-dashboard"></i> 成绩</a>
            </li>
        </ul>
    </div>
    <!-- /.navbar-collapse -->
</nav>