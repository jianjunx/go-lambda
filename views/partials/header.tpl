{{define "header"}}
<header>
  <div class="herder-wrap">
  <h1>LOGO</h1>
  <ul>
    <li>
      <a href="/">首页</a>
    </li>
  {{range .books}}
    <li><a href="?book={{.Id}}">{{.Name}}</a></li>
  {{end}}
  </ul>
  </div>
</header>
{{end}}