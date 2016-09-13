<form class="form-inline" role="form" method="get" action="">
    {{if eq .CurrentUser.Type 1 2 3}}
    {{if .CurrentUser.Major}}
    <div class="form-group">
        <input type="text" class="form-control" name="major_id" placeholder="请输入专业名称" value="{{.CurrentUser.Major.MajorName}}" disabled>
    </div>
    {{end}}
    {{end}}
    {{if eq .CurrentUser.Type 1 4}}
    <div class="form-group">
        <input type="text" class="form-control" name="subject_sender" placeholder="请输入老师姓名" value="{{.SearchTeacher}}">
    </div>
    <button type="submit" class="btn btn-default">搜索</button>
    {{end}}
    {{if ne .endDay 0}}
    {{if eq .CurrentUser.Type 2}}
    <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#add_modal">
        添加课题
    </button>
    {{end}}
    {{end}}

</form>


<hr />
<div class="row">
    <div class="col-md-1">序号</div>
    <div class="col-md-4">标题</div>
    <div class="col-md-2">开题老师</div>
    <div class="col-md-2">状态</div>
    <div class="col-md-2">负责学生</div>
    <div class="col-md-1">研究方向</div>
</div>
<br>
{{range .SubjectList}}
    <div class="row" >
        <div class="col-md-1">{{.Id}}</div>
        <div class="col-md-4"><a href='{{urlfor "SubjectController.Get" ":id" .Id}}'>{{.Title}}</a></div>
        <div class="col-md-2">{{.Sender.Name}}</div>
        <div class="col-md-2">{{.Status.Name}}</div>
        {{if .Student}}
        <div class="col-md-2">{{.Student.Name}}</div>
        {{else}}
        <div class="col-md-2"> - </div>
        {{end}}
        <div class="col-md-1">{{.Sender.Major.MajorName}}</div>
    </div>
    <br>
{{else}}
<h3>没有找到这个结果</h3>
{{end}}

<!-- 编辑Modal -->
<div class="modal fade" id="add_modal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <form role="form" action="" method="post">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal"><span
                                aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                    <h4 class="modal-title" id="myModalLabel">添加课题</h4>
                </div>
                <div class="modal-body">
                    <div class="form-group">

                        <label >课题题目:</label>
                        <input type="text" class="form-control" name="title" placeholder="请输入课题题目">
                        <label >题目描述:</label>
                        <textarea class="form-control" name="detail" placeholder="请输入课题描述" rows="4"></textarea>
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
