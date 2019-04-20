<html>

<body >
  <div style="margin:0 auto;width:800px;font-size:16px">
    <div style="text-align:center;height:40px;padding-top:50px;font-size:18px;color:blue">
      下载数据错误！
    </div>
    <div style="text-align:left;display:inline">
      <div style="width:100px;height:80px;float:left;font-weight:800">
        错误信息:
      </div>
      <div style="width:700px;height:80px;word-break:break-all">
        {{.Datas.error}}
      </div>
    </div>
    <div style="text-align:left;display:inline">
      <div style="width:100px;height:80px;float:left;font-weight:800">
        传入参数:
      </div>
      <div style="width:700px;height:80px;word-break:break-all">
        {{.Datas.paramBean}}
      </div>
    </div>
  </div>

</body>
</html>
