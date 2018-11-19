const fs = require("fs");
const { JSDOM } = require("jsdom");
const { Script } = require("vm");

const bundleFile = fs.readFileSync("./build/elm.js", "utf8")

const dom = new JSDOM(
  `<!DOCTYPE HTML>
  <html>
  <head>
    <meta charset="UTF-8">
    <title>Main</title>
    <script>${bundleFile}</script>
  </head>
  <body>
    <script>
      console.log("init Elm app on Server");
      var app = Elm.Main.init({ flags: false });
    </script>
  </body>
  </html>`,
  { 
    runScripts: "dangerously",
    pretendToBeVisual: true,
    url: "http://localhost:8080/"
  }
);

console.log(dom.window.document.body.innerHTML);
