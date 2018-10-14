package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var client *Client = New()

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/report", getReportHandler)
	http.HandleFunc("/captcha", client.CaptchaHandler)

	log.Println("Start")
	http.ListenAndServe(":8080", nil)
}

const html = `
<html>
	<head>
		<title>店小算代理</title>
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
		<script src="https://cdn.jsdelivr.net/npm/jquery@1.12.4/dist/jquery.min.js"></script>
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
	<head>
	<body>
		<form class="form-horizontal" method="POST">
			<div class="form-group">
				<label for="userName" class="col-sm-2 control-label">账户名</label>
				<div class="col-sm-10">
					<input class="form-control" id="userName" name="userName" value="" placeholder="管理员用户名">
				</div>
			</div>

			<div class="form-group">
				<label for="employeeName" class="col-sm-2 control-label">用户名</label>
				<div class="col-sm-10">
					<input class="form-control" id="employeeName" name="employeeName" value="" placeholder="员工账号">
				</div>
			</div>

			<div class="form-group">
				<label for="password" class="col-sm-2 control-label">密码</label>
				<div class="col-sm-10">
					<input type="password" class="form-control" id="password" name="password" value="" placeholder="密码">
				</div>
			</div>

			<div class="form-group">
				<label for="validateCode" class="col-sm-2 control-label">验证码</label>
				<div class="col-sm-10">
					<input class="form-control" id="validateCode" name="validateCode" value="" placeholder="请识别下图中的验证码">
					<img src="/captcha">
				</div>
			</div>

			<div class="form-group">
				<div class="col-sm-offset-2 col-sm-10">
					<button type="submit" class="btn btn-primary">（重）登陆</button>
				</div>
			</div>
		</form>
	</body>
</html>
`

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	if r.Method == "POST" {
		client.doLogin(r.FormValue("userName"), r.FormValue("employeeName"), r.FormValue("password"), r.FormValue("validateCode"))
	}

	w.Write([]byte(html))

}

func getReportHandler(w http.ResponseWriter, r *http.Request) {
	y, m, d := time.Now().Date()
	end := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	start := end.AddDate(0, 0, -1)

	report, err := client.GetPayReportDetail(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	enc.Encode(report)
}
