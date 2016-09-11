<!DOCTYPE >
<head lang="zh-CN">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <meta name="viewport" content="width=device-width,initial-scale=1,maximum-scale=1.0"/>
    <title>论文管理系统::{{.Title}}</title>
    <link rel="shortcut icon" href="/favicon.ico" type="image/x-icon"/>
    <link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
    {{.HtmlHead}}
</head>

<body>
<div id="wrapper">
    {{.Nav}}
    <div id="page-wrapper">
        <div class="container-fluid">
            <!-- Page Heading -->

            {{.LayoutContent}}

            <!-- /.row -->
        </div>
        <!-- /.container-fluid -->
    </div>


</div>
<script src="https://cdn.bootcss.com/jquery/3.1.0/jquery.min.js"></script>
<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
{{.Script}}
</body>
</html>