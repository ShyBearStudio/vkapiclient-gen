package main

import (
	//	"os"
	"bytes"
	"go/format"
	"text/template"
)

const ObjectTemplate = `type {{.Name}} struct {
{{range $value :=  .Properties -}}
{{.Name}} {{$value | GetPropertyType}} // {{.Description}}
{{end -}}
}
`

func GetPropertyType(op ObjectPropertyPrimitive) Type {
	return op.Type.Type()
}

func CreateObject(data ObjectPrimitive) []byte {
	var b bytes.Buffer
	t := template.New("ObjectTemplate")
	t.Funcs(template.FuncMap{"GetPropertyType": GetPropertyType})
	t, err := t.Parse(ObjectTemplate)
	if err != nil {
		panic(err)
	}
	err = t.Execute(&b, data)
	if err != nil {
		panic(err)
	}
	res, err := format.Source(b.Bytes())
	if err != nil {
		panic(err)
	}
	return res
}

/*	*/
