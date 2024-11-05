package main

import (
	"fmt"
	"go/ast"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			// if x, ok := n.(*ast.File); ok {
			// 	fmt.Printf("%s\n", x.Name.Name)
			// }

			if n != nil {
				fmt.Printf("%T\n", n)
			}
			// if x, ok := n.(*ast.ImportSpec); ok {
			// 	// fmt.Printf("%+v\n", x.Path)
			// }

			if x, ok := n.(*ast.Ident); ok {
				if x != nil && x.Name == "Id" {
					pass.Reportf(x.Pos(), "call of %s(...)", x.Name)
				}

			}
			if call, ok := n.(*ast.CallExpr); ok {
				var id *ast.Ident
				switch fun := call.Fun.(type) {
				case *ast.Ident:
					id = fun
				case *ast.SelectorExpr:
					id = fun.Sel
				}
				if id != nil && !pass.TypesInfo.Types[id].IsType() && id.Name == "Id" {
					pass.Reportf(call.Lparen, "call of %s(...)", id.Name)
				}
			}
			return true
		})
	}

	// inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	// nodeFilter := []ast.Node{(*ast.Ident)(nil)}
	// inspect.Preorder(nodeFilter, func(n ast.Node) {
	// 	switch n := n.(type) {
	// 	case *ast.Ident:
	// 		if n.Name == "Id" {
	// 			pass.Reportf(n.Pos(), "NG") // エラーを出力
	// 		}
	// 	}
	// })

	return nil, nil
}

func isRepositoryPkg(s string) bool {
	return strings.HasSuffix(strings.Trim(s, "\""), "repository")
}

// Analyzer provides static analysis for layered architecture.
var Analyzer = &analysis.Analyzer{
	Name: "importcheck",
	Doc:  Doc,
	Run:  run,
	// Requires: []*analysis.Analyzer{
	// 	inspect.Analyzer, // 依存するAnalyzer
	// },
}

type depchecker struct {
	v int
}

// Doc writes description this analyzer.
const Doc = "importcheck confirms clean architecture."

func (d *depchecker) run(pass *analysis.Pass) (interface{}, error) {
	fmt.Println(d.v)
	return nil, nil
}
func trimQuotes(s string) string {
	if len(s) >= 2 {
		switch {
		case s[0] == '"' && s[len(s)-1] == '"':
			return s[1 : len(s)-1]
		case s[0] == '\'' && s[len(s)-1] == '\'':
			return s[1 : len(s)-1]
		}
	}
	return s
}
func wrapRun(appConfig *appConfig) func(*analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		fmt.Println("PKG PATH: " + pass.Pkg.Path())
		if packageConfig, ok := appConfig.packageConfigs[pass.Pkg.Path()]; ok {

			for _, f := range pass.Files {
				ast.Inspect(f, func(n ast.Node) bool {
					// if structTy, ok := n.(*ast.StructType); ok {
					// 	for _, fl := range structTy.Fields.List {
					// 		fmt.Printf("%s\n", fl.Type)
					// 	}
					// }
					// if x, ok := n.(*ast.File); ok {
					// 	// fmt.Printf("%s\n", x.Name.Name)
					// 	// fmt.Printf("%s\n", x.Pos())

					// 	fmt.Print(x.FileStart)
					// 	fmt.Println()
					// 	return true
					// }

					// if n != nil {
					// 	fmt.Printf("%T\n", n)
					// }
					// if x, ok := n.(*ast.ImportSpec); ok {
					// 	// fmt.Printf("%+v\n", x.Path)
					// }

					// if x, ok := n.(*ast.Ident); ok {
					// 	if x != nil && x.Name == "Id" {
					// 		// d := analysis.Diagnostic{
					// 		// 	Pos:     x.Pos(),
					// 		// 	Message: "integer addition found",
					// 		// }
					// 		// fmt.Println(d.Message)
					// 		// fmt.Println(pass.Pkg.Path())
					// 		pass.Reportf(x.Pos(), "call of %s(...)", x.Name)
					// 	}
					// }
					if x, ok := n.(*ast.ImportSpec); ok {
						importedPackage := trimQuotes(x.Path.Value)

						fmt.Println(x.Path.Value)
						fmt.Println(packageConfig.forbiddenPackages)
						if _, ok := packageConfig.forbiddenPackages[importedPackage]; ok {
							fmt.Println("TRUE")
							pass.Reportf(x.Pos(), "imported forbidden package: %s", importedPackage)
						}
						// fmt.Println(packageConfig)
					}
					return true
				})

				// inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

				// nodeFilter := []ast.Node{
				// 	(*ast.RangeStmt)(nil),
				// }

				// inspect.Preorder(nodeFilter, func(n ast.Node) {
				// 	if x, ok := n.(*ast.Ident); ok {
				// 		if x != nil && x.Name == "Id" {
				// 			pass.Reportf(x.Pos(), "call of %s(...)", x.Name)
				// 		}

				// 	}

				// 	// ここに君だけの最強の静的解析処理を書こう
				// })
			}
		}
		return nil, nil
	}
	// for _, f := range pass.Files {
	// 	for _, d := range f.Decls {
	// 		fmt.Printf("%T\n", d)
	// 	}
	// 	ast.Inspect(f, func(n ast.Node) bool {
	// 		if structTy, ok := n.(*ast.StructType); ok {
	// 			for _, fl := range structTy.Fields.List {
	// 				fmt.Printf("%s\n", fl.Type)
	// 			}
	// 		}
	// 		// if x, ok := n.(*ast.File); ok {
	// 		// 	// fmt.Printf("%s\n", x.Name.Name)
	// 		// 	// fmt.Printf("%s\n", x.Pos())

	// 		// 	fmt.Print(x.FileStart)
	// 		// 	fmt.Println()
	// 		// 	return true
	// 		// }

	// 		if n != nil {
	// 			fmt.Printf("%T\n", n)
	// 		}
	// 		// if x, ok := n.(*ast.ImportSpec); ok {
	// 		// 	// fmt.Printf("%+v\n", x.Path)
	// 		// }

	// 		if x, ok := n.(*ast.Ident); ok {
	// 			if x != nil && x.Name == "Id" {
	// 				pass.Reportf(x.Pos(), "call of %s(...)", x.Name)
	// 			}

	// 		}
	// 		if call, ok := n.(*ast.CallExpr); ok {
	// 			var id *ast.Ident
	// 			switch fun := call.Fun.(type) {
	// 			case *ast.Ident:
	// 				id = fun
	// 			case *ast.SelectorExpr:
	// 				id = fun.Sel
	// 			}
	// 			if id != nil && !pass.TypesInfo.Types[id].IsType() && id.Name == "Id" {
	// 				pass.Reportf(call.Lparen, "call of %s(...)", id.Name)
	// 			}
	// 		}
	// 		return true
	// 	})
	// }

	// return nil, nil
	// }
}

func findFilesWithWalkDir(root string) ([]string, error) {
	findList := []string{}

	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed filepath.WalkDir")
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".go") {
			findList = append(findList, path)
		}
		return nil
	})
	return findList, err
}

type layerConfig struct {
	packagePath       string
	forbiddenPackages []string
}

type packageConfig struct {
	forbiddenPackages map[string]bool
}

type appConfig struct {
	layers         map[string]layerConfig
	packageConfigs map[string]packageConfig
}

func main() {
	appConfig := appConfig{
		layers: map[string]layerConfig{
			"domain": {
				packagePath: "github.com/pecolynx/golang-structure/src/domain",
				forbiddenPackages: []string{
					"net/http",
				},
			},
		},
		packageConfigs: map[string]packageConfig{
			"github.com/pecolynx/golang-structure/src/domain/model": {
				forbiddenPackages: map[string]bool{
					"net/http": true,
				},
			},
		},
	}
	fmt.Println(appConfig)

	// cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedSyntax}
	// // pkgs, err := packages.Load(cfg, flag.Args()...)
	// pkgs, err := packages.Load(cfg, "./...")
	// for _, pkg := range pkgs {
	// 	fmt.Println(pkg.ID, pkg.GoFiles)
	// }

	// os.Exit(0)
	// dc := depchecker{
	// 	v: 123,
	// }

	singlechecker.Main(&analysis.Analyzer{
		Name: "importcheck",
		Doc:  Doc,
		Run:  wrapRun(&appConfig),
		// Run:  dc.run,
		// Requires: []*analysis.Analyzer{
		// 	inspect.Analyzer, // 依存するAnalyzer
		// },
		// Requires: []*analysis.Analyzer{
		// 	inspect.Analyzer,
		// },
	})

	findFilesWalkDir, err := findFilesWithWalkDir("/home/hiroto/pecolynx/golang-structure")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dirs := make(map[string]struct{})
	for _, f := range findFilesWalkDir {
		// fmt.Println(f)

		dir := filepath.Dir(f)
		dirs[dir] = struct{}{}
	}
	for k, _ := range dirs {
		fmt.Println(k)

	}
}
