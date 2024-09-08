//go:generate goacc -i generator.go
package generator

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"

	"github.com/kasaikou/goacc/goacc/entity"
	"github.com/kasaikou/goacc/goacc/parser"
	"golang.org/x/tools/imports"
)

func NewGenerator() *Generator {
	g := Generator{}
	g.goaccPreNewHook()
	return &g
}

type Generator struct {
	cache map[string]*entity.FileConfig `goacc:"json"`
}

func (g *Generator) goaccPreNewHook() {
	g.cache = make(map[string]*entity.FileConfig)
}

func (g *Generator) Generate(srcFilename string, generateConfig *entity.GenerateConfig) (destFilename string, b []byte) {

	if !filepath.IsAbs(srcFilename) {
		srcFilename = filepath.Join(generateConfig.WorkingDir(), srcFilename)
	}

	config, exist := g.cache[srcFilename]
	if !exist {
		g.loadFile(srcFilename, generateConfig)
		config = g.cache[srcFilename]
	}

	buffer := bytes.NewBufferString("")
	fprintfln(buffer, "// Code generated by github.com/kasaikou/goacc, DO NOT EDIT.")
	fprintfln(buffer, "// defaultTag=%s", generateConfig.DefaultTag())
	fprintfln(buffer, "package %s", config.PackageName())
	fprintfln(buffer, "")

	empty := buffer.Bytes()

	for _, structConfig := range config.Structs() {
		generateNew(buffer, structConfig)
		generateAccessor(buffer, structConfig)
		generateMarshalJSON(buffer, structConfig)
	}
	destFilename = RenameDestFilename(srcFilename)

	b, err := format.Source(buffer.Bytes())
	if err != nil {
		// debug code
		if writeErr := WriteFile(destFilename, buffer.Bytes()); writeErr != nil {
			panic(err.Error() + ", and " + writeErr.Error())
		}
		panic(err.Error())
	} else if len(b) == len(empty) {
		return destFilename, nil
	}

	b, err = imports.Process(destFilename, b, nil)
	if err != nil {
		panic(err.Error())
	}

	return destFilename, b
}

func (g *Generator) loadFile(srcFilename string, generateConfig *entity.GenerateConfig) {
	dirname := filepath.Dir(srcFilename)
	pkg, err := parser.LoadPackage(parser.NewLoadPackageInputBuilder(
		dirname,
	).Purge())
	if err != nil {
		panic(err.Error())
	}

	fileConfigs := parser.ParsePackage(
		parser.NewParsePackageInputBuilder(
			pkg,
			generateConfig.DefaultTag(),
		).Purge(),
	)

	for _, file := range fileConfigs {
		g.cache[file.Filename()] = &file
	}

	if _, exist := g.cache[srcFilename]; !exist {
		panic("cannot cache: " + srcFilename)
	}
}

func WriteFile(destFilename string, buffer []byte) error {
	if buffer == nil {
		return nil
	}

	file, err := os.Create(destFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(buffer); err != nil {
		return err
	}

	return nil
}
