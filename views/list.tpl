<form class="form-inline" role="form" method="get" action="">
    {{if eq .CurrentUser.Type 1 2 3}}
    <div class="form-group">
        <input type="text" class="form-control" name="major_id" placeholder="请输入专业名称" value="{{.CurrentUser.Major.MajorName}}" disabled>
    </div>
    {{end}}

    {{if eq .CurrentUser.Type 1 2 4}}
    <div class="form-group">
        <input type="text" class="form-control" name="subject_sender" placeholder="请输入老师姓名" value="{{.SearchTeacher}}">
    </div>
    {{end}}
    <button type="submit" class="btn btn-default">搜索</button>
</form>

<hr />
<div class="row">
    <div class="col-md-1">序号</div>
    <div class="col-md-4">标题</div>
    <div class="col-md-2">开题老师</div>
    <div class="col-md-2">状态</div>
    <div class="col-md-2">负责学生</div>
</div>
<br>
{{range .SubjectList}}
    <div class="row">
        <div class="col-md-1">{{.Id}}</div>
        <div class="col-md-4"><a href='{{urlfor "SubjectController.Get" ":id" .Id}}'>{{.Title}}</a></div>
        <div class="col-md-2">{{.Sender.Name}}</div>
        <div class="col-md-2"><span class='label {{if eq .Status.Id 12 22 32 42 52 62 71}}label-warning{{else}}label-info{{end}}'>{{.Status.Name}}</span></div>
        {{if .Student}}
        <div class="col-md-2">{{.Student.Name}}</div>
        {{else}}
        <div class="col-md-2"> - </div>
        {{end}}
    </div>
    <br>
{{else}}
<h3>没有找到这个结果</h3>
{{end}}