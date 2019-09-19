package genres

import (
	"flag"
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
resAt=`
package {{.PkgName}}

type (
 Action struct {
	ActName string
	ActDes string
	ActMeta string
	ActPath string
	ActVerb string
}
 ResActions struct {
	ResName string
	ResDes string
	ModuleName string
   Actions []Action
}
)

var resActionsMap map[string]ResActions

func init()  {
	SetResActions()
}

// To set the resources and actions for crius
func SetResActions()  {
	resActionsMap=map[string]ResActions{}{{range $k,$v:=.Data}}
    resActions{{$v.Name}}:=ResActions{}
	resActions{{$v.Name}}.ResName="{{goify $v.Name true }}Controller"
	resActions{{$v.Name}}.ResDes="{{$v.Description}}"{{$metaR:=index $v.Metadata "module"}}{{ if $metaR}}
    resActions{{$v.Name}}.ModuleName="{{  index $metaR 0 }}"{{end}}{{range $v.Actions}}{{$name:=replace .Name "-" ""  }}
	{{$v.Name}}Action{{$name}}:=Action{}
	{{$v.Name}}Action{{$name}}.ActDes="{{.Description}}"{{$metaA:=index .Metadata "action"}}{{ if $metaA}}
	{{$v.Name}}Action{{$name}}.ActMeta="{{index $metaA 0}}"{{end}}
	{{$v.Name}}Action{{$name}}.ActName="{{.Name}}"
	{{$v.Name}}Action{{$name}}.ActPath="{{(index .Routes 0 ).Path}}"
	{{$v.Name}}Action{{$name}}.ActVerb="{{(index .Routes 0 ).Verb}}"
	resActions{{$v.Name}}.Actions=append(resActions{{$v.Name}}.Actions,{{$v.Name}}Action{{$name}}){{ end }}
	resActionsMap[resActions{{$v.Name}}.ResName]=resActions{{$v.Name}}{{ end }}
}
// To get the resources and actions from crius
func GetResActions() map[string]ResActions {
	return resActionsMap
}`
)

func Generate() ([]string, error) {
	var (
		ver    string
		outDir string
	)
	set := flag.NewFlagSet("app", flag.PanicOnError)
	set.String("design", "", "") // Consume design argument so Parse doesn't complain
	set.StringVar(&ver, "version", "", "")
	set.StringVar(&outDir, "out", "", "")
	set.Parse(os.Args[2:])
	// First check compatibility
	if err := codegen.CheckVersion(ver); err != nil {
		return nil, err
	}
	return WriteNames(design.Design, outDir)
}

type DData struct {
	Data map[string]*design.ResourceDefinition
	PkgName string
}

// WriteNames creates the names.txt file.
func WriteNames(api *design.APIDefinition, outDir string) ([]string, error) {
	// Now iterate through the resources to gather their names
	//names := make([]string, len(api.Resources))
	//i := 0
	//api.IterateResources(func(res *design.ResourceDefinition) error {
	//	if n, ok := res.Metadata["pseudo"]; ok {
	//		names[i] = n[0]
	//	} else {
	//		names[i] = res.Name
	//		res.Actions
	//	}
	//	i++
	//	return nil
	//})

//	res:=[]string{}
	//for k,v:=range api.Resources {
	//
	//	fmt.Println("----the resource---%v-->",k)
	//	fmt.Println("----the resource---%v--moduleName--%v->",k,v.Metadata["module"])
	//
	//	res=append(res,fmt.Sprintf("----the resource---%v-->",k))
	//	res=append(res,fmt.Sprintf("----the resource---%v--moduleName--%v->",k,v.Metadata["module"]))
	//
	//	for ak,av:=range v.Actions {
	//
	//		fmt.Println("----the action---%s-->",ak)
	//		fmt.Println("----the action-name--%s-->",av.Name)
	//		fmt.Println("----the action---%s-Description--%s-->",av.Name,av.Description)
	//		fmt.Println("----the action---%s-Routes---Path--%s-->",av.Name,av.Routes[0].Path)
	//		fmt.Println("----the action---%s-Routes---Verb--%s-->",av.Name,av.Routes[0].Verb)
	//
	//		res=append(res,fmt.Sprintf("----the action---%s-->",ak))
	//		res=append(res,fmt.Sprintf("----the action-name--%s-->",av.Name))
	//		res=append(res,fmt.Sprintf("----the action---%s-Description--%s-->",av.Name,av.Description))
	//
	//		if len(av.Metadata["action"])>0 {
	//			fmt.Println("----the action---%s-Metadata--action---%s-->",av.Name,av.Metadata["action"][0])
	//			res=append(res,fmt.Sprintf("----the action---%s-Metadata--%s-->",av.Name,av.Metadata["action"][0]))
	//		}
	//	}
	//}

	data:=DData{}
	data.Data=api.Resources
	data.PkgName=outDir[strings.LastIndex(outDir,"/")+1:]

	outputFile := filepath.Join(outDir, "res_actions.go")
	f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return nil,err
	}
	funcMap:=template.FuncMap{}
	funcMap["replace"]= func(s,old,new string) string {
		return strings.ReplaceAll(s,old,new)
	}
	funcMap["goify"]=codegen.Goify
	//tstr:=strings.ReplaceAll(resAt,"\n","")
	t:=template.Must(template.New("res").Funcs(funcMap).Parse(resAt))
	t.Execute(f,data)
	return []string{outputFile}, nil
}
