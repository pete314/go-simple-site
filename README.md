# Introduction

This repository contains some basic code for a simple website built with Golang  and [Macaron framework](https://go-macaron.com/)

## Steps
1. Create a git repository for the following problems. Make sure your repository has a
readme file, a license file and a gitignore file.
2. Setup your environment so that you can create a web application. It is recommended
that you install a Go web application framework.
3. Create a basic web application. Add a route to serve a resource which simply contains
the text “Hello, world!”. Test your application so far.
4. Add a static HTML page resource to your web application. Have it served as the
root resource at /. Include any necessary CSS and JS files.
5. Add a textbox and a label to the HTML page. Use JavaScript to have the reverse of
the string in the text box output to the label. The label should update every time
the user changes the text in the text box.
6. Add a route using the GET method to the web server at /reverse. This route should
accept a single variable, and it should respond with the variable’s value reversed as
the resource.
7. Change the JavaScript to use the server-side string reversal, instead of the client-side.
8. Change the /reverse route to use the POST method instead, and update the JavaScript
accordingly.
References

