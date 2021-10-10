<!-- /views/admin/master.html -->
<!DOCTYPE html>

<html>
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>{{.title}}</title>
        {{include "layouts/head"}}
    </head>
    <body>
        {{template "header" .}}
        <hr>
        {{template "content" .}}
        <hr>
        {{template "ad" .}}
        <hr>
        {{include "layouts/footer"}}
    </body>
</html>