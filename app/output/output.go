package output

import (
	"archive/zip"
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/denisbr/filebin/app/model"
)

func JSONresponse(w http.ResponseWriter, status int, d interface{}, ctx model.Context) {
	dj, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		ctx.Log.Println("Unable to convert response to json: ", err)
		http.Error(w, "Failed while generating a response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	ctx.Log.Println("Response status: " + strconv.Itoa(status))
	io.WriteString(w, string(dj))
}

// This function is a hack. Need to figure out a better way to do this.
func HTMLresponse(w http.ResponseWriter, tpl string, status int, d interface{}, ctx model.Context) {
	box := ctx.TemplateBox
	t := template.New(tpl)

	var templateString string
	var err error

	templateString, err = box.String(tpl + ".html")
	if err != nil {
		ctx.Log.Fatalln(err)
	}
	t, err = t.Parse(templateString)
	if err != nil {
		ctx.Log.Fatalln(err)
	}

	w.WriteHeader(status)
	ctx.Log.Println("Response status: " + strconv.Itoa(status))

	// To send multiple structs to the template
	err = t.Execute(w, map[string]interface{}{
		"Data": d,
		"Ctx":  ctx,
	})
	if err != nil {
		ctx.Log.Fatalln(err)
	}
}

func ZIPresponse(w http.ResponseWriter, status int, tag string, paths []string, ctx model.Context) {
	ctx.Log.Println("Generating zip archive")

	w.Header().Set("Content-Disposition", `attachment; filename="`+tag+`.zip"`)

	zw := zip.NewWriter(w)

	for _, path := range paths {
		// Extract the filename from the absolute path
		fname := filepath.Base(path)
		ctx.Log.Println("Adding to zip archive: " + fname)

		// Get stat info for modtime etc
		info, err := os.Stat(path)
		if err != nil {
			ctx.Log.Println(err)
			return
		}

		// Generate the Zip info header for this file based on the stat info
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			ctx.Log.Println(err)
			return
		}

		ze, err := zw.CreateHeader(header)
		if err != nil {
			ctx.Log.Println(err)
			return
		}

		file, err := os.Open(path)
		if err != nil {
			ctx.Log.Println(err)
			return
		}
		defer file.Close()
		io.Copy(ze, file)
	}

	err := zw.Close()
	if err != nil {
		ctx.Log.Println(err)
		return
	}

	ctx.Log.Println("Zip archive successfully generated")
	return
}
