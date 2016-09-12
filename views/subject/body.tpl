<h2>{{.Subject.Title}}</h2>
<p class="lead">{{.Subject.Presentation}}</p>
<hr/>
<h4>当前状态</h4>
<p>{{.Subject.Status.Name}}</p>

{{if eq .Subject.Status.Id 11 15 16}}
<h4>当前已选择学生</h4>


{{end}}

<div class="btn-group">
    {{if eq .Subject.Status.Id 16 20 22}}
    <button type="button" class="btn btn-default" data-toggle="modal" data-target="#task_modal">上传计划书</button>
    {{end}}
    {{if eq .CurrentUser.Type 1}}
    {{if eq true .Subject.Status.Id  21 30 32}}
    <button type="button" class="btn btn-default" data-toggle="modal" data-target="#primary_modal">上传开题报告</button>
    {{end}}
    {{if eq true .Subject.Status.Id  31 40 42}}
    <button type="button" class="btn btn-default" data-toggle="modal" data-target="#paper_modal">上传中期报告</button>
    {{end}}
    {{if eq true .Subject.Status.Id  41 50 52}}
    <button type="button" class="btn btn-default" data-toggle="modal" data-target="#interim_modal">上传论文</button>
    {{end}}
    {{if eq true .Subject.Status.Id  51 60 62}}
    <button type="button" class="btn btn-default" data-toggle="modal" data-target="#translate_modal">上传论文译文</button>
    {{end}}
    {{end}}
</div>

<hr/>

{{if .Subject.Task}}
<div class="row">
    <div class="col-md-2">任务书</div>
    <div class="col-md-6"><a href="{{print .Subject.Task.Path .Subject.Task.Name}}">{{.Subject.Task.Name}}</a></div>
    <div class="col-md-2">{{dateformat .Subject.Task.Time "2006-01-02 15:04:05 "}}</div>
    {{if eq .CurrentUser.Type 3}}
    <div class="col-md-2">
        <div class="btn-group btn-group-xm">
            <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown">
                操作 <span class="caret"></span>
            </button>
            <ul class="dropdown-menu" role="menu">
                <li><a href="#">审核通过</a></li>
                <li><a href="#">审核失败</a></li>
            </ul>
        </div>
    </div>
    {{end}}
</div>
<hr/>
{{end}}

{{if .Subject.Primary}}
<div class="row">
    <div class="col-md-2">开题报告</div>
    <div class="col-md-6">文件名称</div>
    <div class="col-md-2">上传时间</div>
    {{if eq .CurrentUser.Type 2}}
    <div class="col-md-2">
        <div class="btn-group btn-group-xm">
            <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown">
                操作 <span class="caret"></span>
            </button>
            <ul class="dropdown-menu" role="menu">
                <li><a href="#">审核通过</a></li>
                <li><a href="#">审核失败</a></li>
            </ul>
        </div>
    </div>
    {{end}}
</div>
<hr/>
{{end}}

{{if .Subject.Interim}}
<div class="row">
    <div class="col-md-2">中期报告</div>
    <div class="col-md-6">文件名称</div>
    <div class="col-md-2">上传时间</div>
    {{if eq .CurrentUser.Type 2}}
    <div class="col-md-2">
        <div class="btn-group btn-group-xm">
            <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown">
                操作 <span class="caret"></span>
            </button>
            <ul class="dropdown-menu" role="menu">
                <li><a href="#">审核通过</a></li>
                <li><a href="#">审核失败</a></li>
            </ul>
        </div>
    </div>
    {{end}}
</div>
<hr/>
{{end}}

{{if .Subject.Paper}}
<div class="row">
    <div class="col-md-2">论文</div>
    <div class="col-md-6">文件名称</div>
    <div class="col-md-2">上传时间</div>
    {{if eq .CurrentUser.Type 2}}
    <div class="col-md-2">
        <div class="btn-group btn-group-xm">
            <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown">
                操作 <span class="caret"></span>
            </button>
            <ul class="dropdown-menu" role="menu">
                <li><a href="#">审核通过</a></li>
                <li><a href="#">审核失败</a></li>
            </ul>
        </div>
    </div>
    {{end}}
</div>
<hr/>
{{end}}

{{if .Subject.Translate}}
<div class="row">
    <div class="col-md-2">译文</div>
    <div class="col-md-6">文件名称</div>
    <div class="col-md-2">上传时间</div>

    {{if eq .CurrentUser.Type 2}}
    <div class="col-md-2">
        <div class="btn-group btn-group-xm">
            <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown">
                操作 <span class="caret"></span>
            </button>
            <ul class="dropdown-menu" role="menu">
                <li><a href="#">审核通过</a></li>
                <li><a href="#">审核失败</a></li>
            </ul>
        </div>
    </div>
    {{end}}

</div>
<hr/>
{{end}}

<h3>得分</h3>
<div class="well">
    <div class="row">
        <div class="col-md-4">论文成绩：{{.Subject.PaperGrade}}</div>
        <div class="col-md-4">答辩成绩：{{.Subject.DefenceGrade}}</div>
        <div class="col-md-4">总分：{{.Subject.FinalGrade}}</div>
    </div>
</div>
<!-- Large modal -->
<div class="row">

    {{if eq .CurrentUser.Type 2}}
    {{if gt .Subject.Status.Id  50}}
    <div class="col-md-2">
        <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#paper_model">
            录入论文分数
        </button>
    </div>
    {{end}}

    {{if eq .CurrentUser.Type 4}}
    <div class="col-md-2">
        <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#defence_model">
            录入答辩分数
        </button>
    </div>
    {{end}}
    {{end}}
</div>
<hr/>


<!-- Modal -->
<div class="modal fade" id="paper_model" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <form role="form" action="#" method="post">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal"><span
                                aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                    <h4 class="modal-title" id="myModalLabel">论文成绩</h4>
                </div>
                <div class="modal-body">
                    <div class="form-group">
                        <input type="text" class="form-control" name="paper_grade" placeholder="请输入论文成绩">
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
<!-- 答辩成绩Modal -->
<div class="modal fade" id="defence_model" tabindex="-1" role="dialog" aria-labelledby="myModalLabel"
     aria-hidden="true">
    <form role="form" action="#" method="post">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal"><span
                                aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                    <h4 class="modal-title" id="myModalLabel">答辩成绩</h4>
                </div>
                <div class="modal-body">
                    <div class="form-group">
                        <input type="text" class="form-control" name="defence_grade" placeholder="请输入答辩成绩">
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
<!-- 上传任务书 -->
<div class="modal fade" id="task_modal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <form role="form" action="#" method="post">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal"><span
                                aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                    <h4 class="modal-title">上传任务书</h4>
                </div>
                <div class="modal-body">
                    <div class="form-group">
                        <input type="file" name="task_upload_doc">
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
<!-- 上传开题报告 -->
<div class="modal fade" id="primary_modal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel"
     aria-hidden="true">
    <form role="form" action="#" method="post">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal"><span
                                aria-hidden="true">&times;</span><span class="sr-only">关闭</span></button>
                    <h4 class="modal-title">上传开题报告</h4>
                </div>
                <div class="modal-body">
                    <div class="form-group">
                        <input type="file" name="primary_upload_doc">
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
<!-- 上传中期报告 -->
<div class="modal fade" id="interim_modal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel"
     aria-hidden="true">
    <form role="form" action="#" method="post">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal"><span
                                aria-hidden="true">&times;</span><span class="sr-only">关闭</span></button>
                    <h4 class="modal-title">上传中期报告</h4>
                </div>
                <div class="modal-body">
                    <div class="form-group">
                        <input type="file" name="interim_upload_doc">
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
<!-- 上传论文 -->
<div class="modal fade" id="paper_modal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <form role="form" action="#" method="post">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal"><span
                                aria-hidden="true">&times;</span><span class="sr-only">关闭</span></button>
                    <h4 class="modal-title">上传论文</h4>
                </div>
                <div class="modal-body">
                    <div class="form-group">
                        <input type="file" name="paper_upload_doc">
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
<!-- 上传论文译文 -->
<div class="modal fade" id="translate_modal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel"
     aria-hidden="true">
    <form role="form" action="#" method="post">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal"><span
                                aria-hidden="true">&times;</span><span class="sr-only">关闭</span></button>
                    <h4 class="modal-title">上传论文译文</h4>
                </div>
                <div class="modal-body">
                    <div class="form-group">
                        <input type="file" name="translate_upload_doc">
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


