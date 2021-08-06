package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"io"
	"os"
	"sort"
	"strings"
)

var exprTypes = ExprTypes{
	{
		Name:    `BinaryExpr`,
		Fields:  `Left Expression, Operator token.Token, Right Expression`,
		Comment: "",
	},
	{
		Name:    `GroupingExpr`,
		Fields:  `Expr Expression`,
		Comment: "",
	},
	{
		Name:    `LiteralExpr`,
		Fields:  `Value interface{}`,
		Comment: "",
	},
	{
		Name:    `UnaryExpr`,
		Fields:  `Operator token.Token, Right Expression`,
		Comment: "",
	},
}

var output = flag.String("o", "exprTypes.generated.go", "output file path")

func main() {
	flag.Parse()

	var err error
	defer func() {
		if err != nil {
			fmt.Println("fatal: " + err.Error())
			os.Exit(65)
		}
	}()

	var g = new(Generator)

	g.WriteHeader()

	sort.Sort(exprTypes)
	err = g.WriteExprTypes(exprTypes)
	if err != nil {
		err = fmt.Errorf("generating content via exprTypes: %w", err)
		return
	}

	err = g.Format()
	if err != nil {
		err = fmt.Errorf("formating generated code: %w", err)
		return
	}

	f, err := os.Create(*output)
	if err != nil {
		err = fmt.Errorf("creating new file %q: %w", *output, err)
		return
	}
	defer f.Close()
	defer f.Sync() // nolint: errcheck

	_, err = g.WriteTo(f)
	if err != nil {
		err = fmt.Errorf("writing file: %w", err)
		return
	}
}

type Type struct {
	// type name, must start with uppercase letter
	Name string
	// comma separated
	Fields string
	// optional, comment for this type
	Comment string
}

type ExprTypes []Type

func (t ExprTypes) Len() int {
	return len(t)
}

func (t ExprTypes) Less(i, j int) bool {
	return t[i].Name < t[j].Name
}

func (t ExprTypes) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type Generator struct {
	buf bytes.Buffer
}

func (g *Generator) WriteHeader() {
	g.buf.WriteString(`// Code generated by gen-ast-exprTypes. DO NOT EDIT.

package ast

import (
	"errors"
	"github.com/nanmu42/bluelox/token"
)

type Expression interface {
	Accept(visitor ExprVisitor) (result interface{}, err error)
}

`)
}

func (g *Generator) Format() (err error) {
	formatted, err := format.Source(g.buf.Bytes())
	if err != nil {
		err = fmt.Errorf("gofmt generated code: %w", err)
		return
	}
	g.buf.Reset()
	g.buf.Write(formatted)

	return
}

func (g *Generator) WriteTo(writer io.Writer) (int64, error) {
	return io.Copy(writer, &g.buf)
}

func (g *Generator) WriteExprTypes(types ExprTypes) (err error) {
	err = g.checkExprTypes(types)
	if err != nil {
		err = fmt.Errorf("checking exprTypes: %w", err)
		return
	}

	g.writeExprVisitor(types)

	for _, item := range types {
		g.writeExprType(item)
	}

	return
}

func (g *Generator) checkExprTypes(types ExprTypes) (err error) {
	if len(types) == 0 {
		err = errors.New("provided 0 type")
		return
	}

	var nameSet = make(map[string]struct{}, len(types))
	for index, item := range types {
		if strings.TrimSpace(item.Name) == "" {
			err = fmt.Errorf("empty type name at index %d", index)
			return
		}
		if !ast.IsExported(item.Name) {
			err = fmt.Errorf("name %q at index %d is not exported", item.Name, index)
			return
		}
		if !strings.HasSuffix(item.Name, "Expr") {
			err = fmt.Errorf("name %q at index %d does not end with Expr", item.Name, index)
			return
		}
		if _, ok := nameSet[item.Name]; ok {
			err = fmt.Errorf("duplicated name %s at index %d", item.Name, index)
			return
		}
		nameSet[item.Name] = struct{}{}

		if item.Fields == "" {
			err = fmt.Errorf("fields are empty for type %q", item.Name)
			return
		}
		fields := strings.Split(item.Fields, ",")
		for fieldIdx, field := range fields {
			field = strings.TrimSpace(field)
			if field == "" {
				err = fmt.Errorf("empty field at index %d from type %q", fieldIdx, item.Name)
				return
			}
			if !ast.IsExported(field) {
				err = fmt.Errorf("field %q of type %q is not exported", field, item.Name)
				return
			}
		}
	}

	return
}

func (g *Generator) linebreak() {
	g.buf.WriteByte('\n')
}

func (g *Generator) writeExprVisitor(types ExprTypes) {
	g.buf.WriteString(`type ExprVisitor interface {`)
	g.linebreak()

	for _, item := range types {
		g.buf.WriteByte('\t')
		_, _ = fmt.Fprintf(&g.buf, "Visit%s(v *%s) (result interface{}, err error)", item.Name, item.Name)
		g.linebreak()
	}

	g.buf.WriteString(`}`)
	g.linebreak()
	g.linebreak()

	// stub visitor
	g.buf.WriteString(`type StubExprVisitor struct{}

var _ ExprVisitor = StubExprVisitor{}`)
	g.linebreak()
	g.linebreak()

	for _, item := range types {
		_, _ = fmt.Fprintf(&g.buf, `func (s StubExprVisitor) Visit%s(_ *%s) (interface{}, error) {
	return nil, errors.New("visit func for %s is not implemented")
}`, item.Name, item.Name, item.Name)
		g.linebreak()
		g.linebreak()
	}
}

func (g *Generator) writeExprType(item Type) {
	if item.Comment != "" {
		g.buf.WriteString("// ")
		g.buf.WriteString(item.Comment)
		g.linebreak()
	}
	_, _ = fmt.Fprintf(&g.buf, "type %s struct {", item.Name)
	g.linebreak()

	fields := strings.Split(item.Fields, ",")
	g.buf.WriteString(strings.Join(fields, "\n"))

	g.buf.WriteString(`}`)
	g.linebreak()
	g.linebreak()

	_, _ = fmt.Fprintf(&g.buf, "var _ Expression = (*%s)(nil)", item.Name)
	g.linebreak()
	g.linebreak()

	_, _ = fmt.Fprintf(&g.buf, `func (b *%s) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.Visit%s(b)
}`, item.Name, item.Name)
	g.linebreak()
	g.linebreak()
}
