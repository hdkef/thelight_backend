# The Light
the light is a content management system for blogging and is made with angular framework (frontend) and golang (backend).
The light depends on quill.js for WYSIWYG // What you see is what you get text-editor.
This is the backend of the light project.

#TODO

1. Analytics
for example, article hit counter, geography tracking, user agent tracking, etc

2. Server side rendering
because this is a blogging medium, this web app needs to be exposed to search engine. But,
server side rendering for angular is using node.js and not golang, i'll find a way to make it possible to work with golang. Maybe, i make two separated entities, nodejs for handling the static things (like serving images, html, css, js) and golang for handling data processing.