{{define "header"}}
<ul>
<li>
<a href="/">首页</a>
</li>
{{range .books}}
<li><a href="?book={{.Id}}">{{.Name}}</a></li>
{{end}}
</ul>
{{end}}