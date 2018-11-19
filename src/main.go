package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"text/template"
)

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
		<script>{{.BundleFile}}</script>
	</head>
	<body>
		{{.Content}}
		<script>
			console.log("init Elm app in Browser");
			var app = Elm.Main.init({ flags: true });
		</script>
	</body>
</html>`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bundleFile, err := ioutil.ReadFile("build/elm.js")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cmd := exec.Command("yarn", "run", "-s", "render")
		stdout, err := cmd.Output()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t, err := template.New("webpage").Parse(tpl)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := struct {
			Title      string
			BundleFile string
			Content    string
		}{
			Title:      "Hello from Go",
			BundleFile: string(bundleFile),
			Content:    string(stdout),
		}

		if err := t.Execute(w, data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "text/html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
