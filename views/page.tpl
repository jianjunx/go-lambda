<!-- /views/page.html -->
<!doctype html>

<html>
    <head>
        <title>{{.title}}</title>
        {{include "layouts/head"}}
    </head>

    <body>
        {{template "ad" .}}
        <div class="content">
        <h1>{{.detail.Title}}</h1>
        {{innerHtml .detail.BodyHtml}}
        </div>
        <hr>
        {{include "layouts/footer"}}
    </body>
</html>