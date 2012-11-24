package main

import (
    "os"
    "text/template"
    )

func main() {
    t := template.Must(template.New("name").Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))
    t.Execute(os.Stdout, nil)

    t1 := template.Must(template.New("name1").Parse("{{with $x3 := `hola`}}{{$x3}}{{end}}!\n"))
    t1.Execute(os.Stdout, nil)

    t2 := template.Must(template.New("name2").Parse("{{with $x_1 := `hey`}}{{$x_1}} {{.}} {{$x_1}}{{end}}!\n"))
    t2.Execute(os.Stdout, nil)
}
