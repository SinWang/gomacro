// this file was generated by gomacro command: import "text/template"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "text/template"
	. "reflect"
)

func init() {
	Binds["text/template"] = map[string]Value{
		"HTMLEscape":	ValueOf(pkg.HTMLEscape),
		"HTMLEscapeString":	ValueOf(pkg.HTMLEscapeString),
		"HTMLEscaper":	ValueOf(pkg.HTMLEscaper),
		"IsTrue":	ValueOf(pkg.IsTrue),
		"JSEscape":	ValueOf(pkg.JSEscape),
		"JSEscapeString":	ValueOf(pkg.JSEscapeString),
		"JSEscaper":	ValueOf(pkg.JSEscaper),
		"Must":	ValueOf(pkg.Must),
		"New":	ValueOf(pkg.New),
		"ParseFiles":	ValueOf(pkg.ParseFiles),
		"ParseGlob":	ValueOf(pkg.ParseGlob),
		"URLQueryEscaper":	ValueOf(pkg.URLQueryEscaper),
	}
	Types["text/template"] = map[string]Type{
		"ExecError":	TypeOf((*pkg.ExecError)(nil)).Elem(),
		"FuncMap":	TypeOf((*pkg.FuncMap)(nil)).Elem(),
		"Template":	TypeOf((*pkg.Template)(nil)).Elem(),
	}
}