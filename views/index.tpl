
{{define "content"}}
    <h1 class="hello">This is content!!!!</h1>
    <p>123 - 100= {{sub 123 100}}</p>
    <ul>
    {{range .posts}}
        <li>
        <a href="/p/{{.Slug}}">{{.Title}}</a>
        </li>
    {{end}}
    </ul>
    <hr>
    <ul>
    {{range $i, $v := .pagers}}
        <li data-page="{{$v}}">
        <a href="javascript:void(0)">{{$v}}</a>
        </li>
    {{end}}
    </ul>
{{end}}