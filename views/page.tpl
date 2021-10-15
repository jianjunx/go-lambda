<!-- /views/page.html -->
<!doctype html>

<html>
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
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
    <script src="/public/js/index.js"></script>
</html>