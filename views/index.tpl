
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
    <div class="pagination">
    {{if .showPrev}}
        <div>上一页</div>
    {{end}}
    <ul class="pagination-page">
    {{range $i, $v := .pages}}
        <li data-page="{{$v}}">
            <a href="javascript:void(0)">{{$v}}</a>
        </li>
    {{end}}
    </ul>
    {{if .showNext}}
        <div>下一页</div>
    {{end}}

    </div>
{{end}}