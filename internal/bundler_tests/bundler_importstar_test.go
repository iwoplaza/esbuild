package bundler_tests

import (
	"regexp"
	"testing"

	"github.com/evanw/esbuild/internal/config"
)

var importstar_suite = suite{
	name: "importstar",
}

func TestImportStarUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportImportStarUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportImportStarNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportImportStarCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarAsUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarAsNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarAsCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './bar'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './bar'
				let foo = 234
				console.log(foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './bar'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './bar'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
			"/bar.js": `
				export * from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarCommonJSUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarCommonJSCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarCommonJSNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarAndCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				const ns2 = require('./foo')
				console.log(ns.foo, ns2.foo)
			`,
			"/foo.js": `
				export const foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarNoBundleUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarNoBundleCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarNoBundleNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarMangleNoBundleUnused(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			MinifySyntax:  true,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarMangleNoBundleCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			MinifySyntax:  true,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarMangleNoBundleNoCapture(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				let foo = 234
				console.log(ns.foo, ns.foo, foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			MinifySyntax:  true,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarExportStarOmitAmbiguous(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './common'
				console.log(ns)
			`,
			"/common.js": `
				export * from './foo'
				export * from './bar'
			`,
			"/foo.js": `
				export const x = 1
				export const y = 2
			`,
			"/bar.js": `
				export const y = 3
				export const z = 4
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportExportStarAmbiguousError(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {x, y, z} from './common'
				console.log(x, y, z)
			`,
			"/common.js": `
				export * from './foo'
				export * from './bar'
			`,
			"/foo.js": `
				export const x = 1
				export const y = 2
			`,
			"/bar.js": `
				export const y = 3
				export const z = 4
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: ERROR: Ambiguous import "y" has multiple matching exports
foo.js: NOTE: One matching export is here:
bar.js: NOTE: Another matching export is here:
`,
	})
}

func TestImportExportStarAmbiguousWarning(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './common'
				console.log(ns.x, ns.y, ns.z)
			`,
			"/common.js": `
				export * from './foo'
				export * from './bar'
			`,
			"/foo.js": `
				export const x = 1
				export const y = 2
			`,
			"/bar.js": `
				export const y = 3
				export const z = 4
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "y" will always be undefined because there are multiple matching exports
foo.js: NOTE: One matching export is here:
bar.js: NOTE: Another matching export is here:
`,
	})
}

func TestReExportStarNameCollisionNotAmbiguousImport(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {x, y} from './common'
				console.log(x, y)
			`,
			"/common.js": `
				export * from './a'
				export * from './b'
			`,
			"/a.js": `
				export * from './c'
			`,
			"/b.js": `
				export {x} from './c'
			`,
			"/c.js": `
				export let x = 1, y = 2
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarNameCollisionNotAmbiguousExport(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from './a'
				export * from './b'
			`,
			"/a.js": `
				export * from './c'
			`,
			"/b.js": `
				export {x} from './c'
			`,
			"/c.js": `
				export let x = 1, y = 2
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarNameShadowingNotAmbiguous(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {x} from './a'
				console.log(x)
			`,
			"/a.js": `
				export * from './b'
				export let x = 1
			`,
			"/b.js": `
				export let x = 2
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarNameShadowingNotAmbiguousReExport(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {x} from './a'
				console.log(x)
			`,
			"/a.js": `
				export * from './b'
			`,
			"/b.js": `
				export * from './c'
				export let x = 1
			`,
			"/c.js": `
				export let x = 2
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportStarOfExportStarAs(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as foo_ns from './foo'
				console.log(foo_ns)
			`,
			"/foo.js": `
				export * as bar_ns from './bar'
			`,
			"/bar.js": `
				export const bar = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportOfExportStar(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {bar} from './foo'
				console.log(bar)
			`,
			"/foo.js": `
				export * from './bar'
			`,
			"/bar.js": `
				// Add some statements to increase the part index (this reproduced a crash)
				statement()
				statement()
				statement()
				statement()
				export const bar = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportOfExportStarOfImport(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {bar} from './foo'
				console.log(bar)
			`,
			"/foo.js": `
				// Add some statements to increase the part index (this reproduced a crash)
				statement()
				statement()
				statement()
				statement()
				export * from './bar'
			`,
			"/bar.js": `
				export {value as bar} from './baz'
			`,
			"/baz.js": `
				export const value = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfIIFE(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfIIFEWithName(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
			GlobalName:    []string{"someName"},
		},
	})
}

func TestExportSelfES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfCommonJSMinified(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				module.exports = {foo: 123}
				console.log(require('./entry'))
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:              config.ModeBundle,
			MinifyIdentifiers: true,
			OutputFormat:      config.FormatCommonJS,
			AbsOutputFile:     "/out.js",
		},
	})
}

func TestImportSelfCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				exports.foo = 123
				import {foo} from './entry'
				console.log(foo)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfAsNamespaceES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * as ns from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportExportSelfAsNamespaceES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				import * as ns from './entry'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportOtherFileExportSelfAsNamespaceES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from './foo'
			`,
			"/foo.js": `
				export const foo = 123
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportOtherFileImportExportSelfAsNamespaceES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from './foo'
			`,
			"/foo.js": `
				export const foo = 123
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestOtherFileExportSelfAsNamespaceUnusedES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export {foo} from './foo'
			`,
			"/foo.js": `
				export const foo = 123
				export * as ns from './foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestOtherFileImportExportSelfAsNamespaceUnusedES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export {foo} from './foo'
			`,
			"/foo.js": `
				export const foo = 123
				import * as ns from './foo'
				export {ns}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfAsNamespaceCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				export * as ns from './entry'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfAndRequireSelfCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export const foo = 123
				console.log(require('./entry'))
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportSelfAndImportSelfCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as x from './entry'
				export const foo = 123
				console.log(x)
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportOtherAsNamespaceCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as ns from './foo'
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportExportOtherAsNamespaceCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				export {ns}
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestNamespaceImportMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns, ns.foo)
			`,
			"/foo.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "foo" will always be undefined because there is no matching export in "foo.js"
`,
	})
}

func TestExportOtherCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export {bar} from './foo'
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestExportOtherNestedCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export {y} from './bar'
			`,
			"/bar.js": `
				export {x as y} from './foo'
			`,
			"/foo.js": `
				exports.foo = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestNamespaceImportUnusedMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns.foo)
			`,
			"/foo.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "foo" will always be undefined because there is no matching export in "foo.js"
`,
	})
}

func TestNamespaceImportMissingCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns, ns.foo)
			`,
			"/foo.js": `
				exports.x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestNamespaceImportUnusedMissingCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns.foo)
			`,
			"/foo.js": `
				exports.x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportNamespaceImportMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './foo'
				console.log(ns, ns.foo)
			`,
			"/foo.js": `
				export * as ns from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "foo" will always be undefined because there is no matching export in "bar.js"
`,
	})
}

func TestReExportNamespaceImportUnusedMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import {ns} from './foo'
				console.log(ns.foo)
			`,
			"/foo.js": `
				export * as ns from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "foo" will always be undefined because there is no matching export in "bar.js"
`,
	})
}

func TestNamespaceImportReExportMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns, ns.foo)
			`,
			"/foo.js": `
				export {foo} from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `foo.js: ERROR: No matching export in "bar.js" for import "foo"
foo.js: ERROR: No matching export in "bar.js" for import "foo"
`,
	})
}

func TestNamespaceImportReExportUnusedMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns.foo)
			`,
			"/foo.js": `
				export {foo} from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `foo.js: ERROR: No matching export in "bar.js" for import "foo"
foo.js: ERROR: No matching export in "bar.js" for import "foo"
`,
	})
}

func TestNamespaceImportReExportStarMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns, ns.foo)
			`,
			"/foo.js": `
				export * from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "foo" will always be undefined because there is no matching export in "foo.js"
`,
	})
}

func TestNamespaceImportReExportStarUnusedMissingES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as ns from './foo'
				console.log(ns.foo)
			`,
			"/foo.js": `
				export * from './bar'
			`,
			"/bar.js": `
				export const x = 123
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "foo" will always be undefined because there is no matching export in "foo.js"
`,
	})
}

func TestExportStarDefaultExportCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from './foo'
			`,
			"/foo.js": `
				export default 'default' // This should not be picked up
				export let foo = 'foo'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestIssue176(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as things from './folders'
				console.log(JSON.stringify(things))
			`,
			"/folders/index.js": `
				export * from "./child"
			`,
			"/folders/child/index.js": `
				export { foo } from './foo'
			`,
			"/folders/child/foo.js": `
				export const foo = () => 'hi there'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarExternalIIFE(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
			GlobalName:    []string{"mod"},
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarExternalES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarExternalCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarIIFENoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
			GlobalName:    []string{"mod"},
		},
	})
}

func TestReExportStarES6NoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarCommonJSNoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarAsExternalIIFE(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
			GlobalName:    []string{"mod"},
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarAsExternalES6(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarAsExternalCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"foo": true,
				}},
			},
		},
	})
}

func TestReExportStarAsIIFENoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatIIFE,
			AbsOutputFile: "/out.js",
			GlobalName:    []string{"mod"},
		},
	})
}

func TestReExportStarAsES6NoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportStarAsCommonJSNoBundle(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * as out from "foo"
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeConvertFormat,
			OutputFormat:  config.FormatCommonJS,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestImportDefaultNamespaceComboIssue446(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/external-default2.js": `
				import def, {default as default2} from 'external'
				console.log(def, default2)
			`,
			"/external-ns.js": `
				import def, * as ns from 'external'
				console.log(def, ns)
			`,
			"/external-ns-default.js": `
				import def, * as ns from 'external'
				console.log(def, ns, ns.default)
			`,
			"/external-ns-def.js": `
				import def, * as ns from 'external'
				console.log(def, ns, ns.def)
			`,
			"/external-default.js": `
				import def, * as ns from 'external'
				console.log(def, ns.default)
			`,
			"/external-def.js": `
				import def, * as ns from 'external'
				console.log(def, ns.def)
			`,
			"/internal-default2.js": `
				import def, {default as default2} from './internal'
				console.log(def, default2)
			`,
			"/internal-ns.js": `
				import def, * as ns from './internal'
				console.log(def, ns)
			`,
			"/internal-ns-default.js": `
				import def, * as ns from './internal'
				console.log(def, ns, ns.default)
			`,
			"/internal-ns-def.js": `
				import def, * as ns from './internal'
				console.log(def, ns, ns.def)
			`,
			"/internal-default.js": `
				import def, * as ns from './internal'
				console.log(def, ns.default)
			`,
			"/internal-def.js": `
				import def, * as ns from './internal'
				console.log(def, ns.def)
			`,
			"/internal.js": `
				export default 123
			`,
		},
		entryPaths: []string{
			"/external-default2.js",
			"/external-ns.js",
			"/external-ns-default.js",
			"/external-ns-def.js",
			"/external-default.js",
			"/external-def.js",
			"/internal-default2.js",
			"/internal-ns.js",
			"/internal-ns-default.js",
			"/internal-ns-def.js",
			"/internal-default.js",
			"/internal-def.js",
		},
		options: config.Options{
			Mode:         config.ModeBundle,
			AbsOutputDir: "/out",
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{Exact: map[string]bool{
					"external": true,
				}},
			},
		},
		expectedCompileLog: `internal-def.js: WARNING: Import "def" will always be undefined because there is no matching export in "internal.js"
internal-ns-def.js: WARNING: Import "def" will always be undefined because there is no matching export in "internal.js"
`,
	})
}

func TestImportDefaultNamespaceComboNoDefault(t *testing.T) {
	// Note: "entry-dead.js" checks that this warning doesn't happen for dead code
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry-default-ns-prop.js": `import def, * as ns from './foo'; console.log(def, ns, ns.default)`,
			"/entry-default-ns.js":      `import def, * as ns from './foo'; console.log(def, ns)`,
			"/entry-default-prop.js":    `import def, * as ns from './foo'; console.log(def, ns.default)`,
			"/entry-default.js":         `import def from './foo'; console.log(def)`,
			"/entry-prop.js":            `import * as ns from './foo'; console.log(ns.default)`,
			"/entry-dead.js":            `import * as ns from './foo'; 0 && console.log(ns.default)`,
			"/entry-typo.js":            `import * as ns from './foo'; console.log(ns.buton)`,
			"/entry-typo-indirect.js":   `import * as ns from './indirect'; console.log(ns.buton)`,
			"/foo.js":                   `export let button = {}`,
			"/indirect.js":              `export * from './foo'`,
		},
		entryPaths: []string{
			"/entry-default-ns-prop.js",
			"/entry-default-ns.js",
			"/entry-default-prop.js",
			"/entry-default.js",
			"/entry-prop.js",
			"/entry-dead.js",
			"/entry-typo.js",
			"/entry-typo-indirect.js",
		},
		options: config.Options{
			Mode:         config.ModeBundle,
			AbsOutputDir: "/out",
		},
		expectedCompileLog: `entry-default-ns-prop.js: ERROR: No matching export in "foo.js" for import "default"
entry-default-ns-prop.js: WARNING: Import "default" will always be undefined because there is no matching export in "foo.js"
entry-default-ns.js: ERROR: No matching export in "foo.js" for import "default"
entry-default-prop.js: ERROR: No matching export in "foo.js" for import "default"
entry-default-prop.js: WARNING: Import "default" will always be undefined because there is no matching export in "foo.js"
entry-default.js: ERROR: No matching export in "foo.js" for import "default"
entry-prop.js: WARNING: Import "default" will always be undefined because there is no matching export in "foo.js"
entry-typo-indirect.js: WARNING: Import "buton" will always be undefined because there is no matching export in "indirect.js"
foo.js: NOTE: Did you mean to import "button" instead?
entry-typo.js: WARNING: Import "buton" will always be undefined because there is no matching export in "foo.js"
foo.js: NOTE: Did you mean to import "button" instead?
`,
	})
}

func TestImportNamespaceUndefinedPropertyEmptyFile(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry-nope.js": `
				import * as js from './empty.js'
				import * as mjs from './empty.mjs'
				import * as cjs from './empty.cjs'
				console.log(
					js.nope,
					mjs.nope,
					cjs.nope,
				)
			`,

			// Note: For CommonJS-style modules, we automatically assign the exports
			// object to the "default" property if there is no property named "default".
			// This is for compatibility with node. So this test intentionally behaves
			// differently from the test above.
			"/entry-default.js": `
				import * as js from './empty.js'
				import * as mjs from './empty.mjs'
				import * as cjs from './empty.cjs'
				console.log(
					js.default,
					mjs.default,
					cjs.default,
				)
			`,

			"/empty.js":  ``,
			"/empty.mjs": ``,
			"/empty.cjs": ``,
		},
		entryPaths: []string{
			"/entry-nope.js",
			"/entry-default.js",
		},
		options: config.Options{
			Mode:         config.ModeBundle,
			AbsOutputDir: "/out",
		},
		expectedCompileLog: `entry-default.js: WARNING: Import "default" will always be undefined because there is no matching export in "empty.mjs"
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "empty.js" has no exports
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "empty.mjs" has no exports
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "empty.cjs" has no exports
`,
	})
}

func TestImportNamespaceUndefinedPropertySideEffectFreeFile(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry-nope.js": `
				import * as js from './foo/no-side-effects.js'
				import * as mjs from './foo/no-side-effects.mjs'
				import * as cjs from './foo/no-side-effects.cjs'
				console.log(
					js.nope,
					mjs.nope,
					cjs.nope,
				)
			`,

			// Note: For CommonJS-style modules, we automatically assign the exports
			// object to the "default" property if there is no property named "default".
			// This is for compatibility with node. So this test intentionally behaves
			// differently from the test above.
			"/entry-default.js": `
				import * as js from './foo/no-side-effects.js'
				import * as mjs from './foo/no-side-effects.mjs'
				import * as cjs from './foo/no-side-effects.cjs'
				console.log(
					js.default,
					mjs.default,
					cjs.default,
				)
			`,

			"/foo/package.json":        `{ "sideEffects": false }`,
			"/foo/no-side-effects.js":  `console.log('js')`,
			"/foo/no-side-effects.mjs": `console.log('mjs')`,
			"/foo/no-side-effects.cjs": `console.log('cjs')`,
		},
		entryPaths: []string{
			"/entry-nope.js",
			"/entry-default.js",
		},
		options: config.Options{
			Mode:         config.ModeBundle,
			AbsOutputDir: "/out",
		},
		expectedCompileLog: `entry-default.js: WARNING: Import "default" will always be undefined because there is no matching export in "foo/no-side-effects.mjs"
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "foo/no-side-effects.js" has no exports
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "foo/no-side-effects.mjs" has no exports
entry-nope.js: WARNING: Import "nope" will always be undefined because the file "foo/no-side-effects.cjs" has no exports
`,
	})
}

// Failure case due to a bug in https://github.com/evanw/esbuild/pull/2059
func TestReExportStarEntryPointAndInnerFile(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				export * from 'a'
				import * as inner from './inner.js'
				export { inner }
			`,
			"/inner.js": `
				export * from 'b'
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:         config.ModeBundle,
			AbsOutputDir: "/out",
			OutputFormat: config.FormatCommonJS,
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{
					Exact: map[string]bool{
						"a": true,
						"b": true,
					},
				},
			},
		},
	})
}

// https://github.com/evanw/esbuild/issues/1420
func TestReExportedNamespacePropertyAccess(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot, tgpu.slot())
				new tgpu.Klass
			`,
			"/index.js": `
				export * as tgpu from './tgpu.js'
			`,
			"/tgpu.js": `
				import { heavy } from './heavy.js'
				export function slot() {}
				export class Klass {}
				export function init() { return heavy() }
			`,
			"/heavy.js": `
				export function heavy() { return 'heavy' }
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportedNamespaceImportStarThenExport(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot)
			`,
			"/index.js": `
				import * as tgpu from './tgpu.js'
				export { tgpu }
			`,
			"/tgpu.js": `
				export function slot() {}
				export function init() {}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportedNamespaceUsedAsValue(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot, tgpu)
			`,
			"/index.js": `
				export * as tgpu from './tgpu.js'
			`,
			"/tgpu.js": `
				export function slot() {}
				export function init() {}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportedNamespaceMutation(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/assign.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot)
				tgpu.slot = 1
			`,
			"/delete.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot)
				delete tgpu.slot
			`,
			"/update.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot)
				tgpu.slot++
			`,
			"/destructure.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot);
				[tgpu.slot] = []
			`,
			"/optional.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot, tgpu?.init)
			`,
			"/index.js": `
				export * as tgpu from './tgpu.js'
			`,
			"/tgpu.js": `
				export function slot() {}
				export function init() {}
			`,
		},
		entryPaths: []string{"/assign.js", "/delete.js", "/update.js", "/destructure.js", "/optional.js"},
		options: config.Options{
			Mode:         config.ModeBundle,
			OutputFormat: config.FormatESModule,
			AbsOutputDir: "/out",
		},
	})
}

func TestReExportedNamespaceDefault(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import tgpu from './index.js'
				console.log(tgpu.slot)
			`,
			"/index.js": `
				export * as default from './tgpu.js'
			`,
			"/tgpu.js": `
				export function slot() {}
				export function init() {}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportedNamespaceTwoLevels(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import { b } from './index.js'
				console.log(b.a.x, b.a.x(), b.y)
			`,
			"/index.js": `
				export * as b from './mid.js'
			`,
			"/mid.js": `
				export * as a from './a.js'
				export function y() {}
				export function unusedMid() {}
			`,
			"/a.js": `
				export function x() {}
				export function unusedA() {}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportedNamespaceTwoLevelsInnerUsedAsValue(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import { b } from './index.js'
				console.log(b.a.x, b.a, b.a.missing, b.y.prop)
			`,
			"/index.js": `
				export * as b from './mid.js'
			`,
			"/mid.js": `
				export * as a from './a.js'
				export const y = { prop: 1 }
				export function unusedMid() {}
			`,
			"/a.js": `
				export function x() {}
				export function unusedA() {}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "missing" will always be undefined because there is no matching export in "a.js"
`,
	})
}

func TestReExportedNamespaceCommonJS(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot)
			`,
			"/index.js": `
				export * as tgpu from './tgpu.cjs'
			`,
			"/tgpu.cjs": `
				exports.slot = function() {}
				exports.init = function() {}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportedNamespaceRequire(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				const { tgpu } = require('./index.js')
				console.log(tgpu.slot)
			`,
			"/index.js": `
				export * as tgpu from './tgpu.js'
			`,
			"/tgpu.js": `
				export function slot() {}
				export function init() {}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportedNamespaceRequiredElsewhere(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot, require('./tgpu.js'))
			`,
			"/index.js": `
				export * as tgpu from './tgpu.js'
			`,
			"/tgpu.js": `
				export function slot() {}
				export function init() {}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportedNamespaceSplitting(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/a.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot)
			`,
			"/b.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot, tgpu.init)
			`,
			"/index.js": `
				export * as tgpu from './tgpu.js'
			`,
			"/tgpu.js": `
				export function slot() {}
				export function init() {}
				export function unused() {}
			`,
		},
		entryPaths: []string{"/a.js", "/b.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			CodeSplitting: true,
			AbsOutputDir:  "/out",
		},
	})
}

func TestReExportedNamespaceNotDevirtualized(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/index-access.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot, tgpu['init'])
			`,
			"/jsx-preserve.jsx": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot, <tgpu.Comp />)
			`,
			"/mangled.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot, tgpu.init_)
			`,
			"/dead-code.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot)
				if (false) console.log(tgpu.init)
			`,
			"/index.js": `
				export * as tgpu from './tgpu.js'
			`,
			"/tgpu.js": `
				export function slot() {}
				export function init() {}
				export function init_() {}
				export class Comp {}
			`,
		},
		entryPaths: []string{"/index-access.js", "/jsx-preserve.jsx", "/mangled.js", "/dead-code.js"},
		options: config.Options{
			Mode:         config.ModeBundle,
			OutputFormat: config.FormatESModule,
			AbsOutputDir: "/out",
			MangleProps:  regexp.MustCompile("_$"),
			JSX: config.JSXOptions{
				Parse:    true,
				Preserve: true,
			},
		},
	})
}

func TestReExportedNamespaceNestedReExports(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.fromClause, tgpu.fromStar, tgpu.fromExternal, tgpu.fromCJS, tgpu.missing)
			`,
			"/index.js": `
				export * as tgpu from './tgpu.js'
			`,
			"/tgpu.js": `
				export { fromClause } from './clause.js'
				export * from './star.js'
				export { fromExternal } from 'external'
				export { fromCJS } from './cjs.cjs'
				export function unused() {}
			`,
			"/clause.js": `
				export function fromClause() {}
				export function unusedClause() {}
			`,
			"/star.js": `
				export function fromStar() {}
				export function unusedStar() {}
			`,
			"/cjs.cjs": `
				exports.fromCJS = 1
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
			ExternalSettings: config.ExternalSettings{
				PreResolve: config.ExternalMatchers{
					Exact: map[string]bool{"external": true},
				},
			},
		},
		expectedCompileLog: `entry.js: WARNING: Import "missing" will always be undefined because there is no matching export in "tgpu.js"
`,
	})
}

func TestReExportedNamespaceMinified(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import { tgpu } from './index.js'
				console.log(tgpu.slot(), tgpu.CONST)
			`,
			"/index.js": `
				export * as tgpu from './tgpu.js'
			`,
			"/tgpu.js": `
				export function slot() {}
				export const CONST = 123
				export function init() {}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:              config.ModeBundle,
			OutputFormat:      config.FormatESModule,
			AbsOutputFile:     "/out.js",
			MinifySyntax:      true,
			MinifyIdentifiers: true,
		},
	})
}

func TestReExportedNamespaceThroughImportStar(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import * as lib from './index.js'
				console.log(lib.tgpu.slot, lib.tgpu.slot())
			`,
			"/index.js": `
				export * as tgpu from './tgpu.js'
			`,
			"/tgpu.js": `
				export function slot() {}
				export function init() {}
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
	})
}

func TestReExportedNamespaceMissing(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/entry.js": `
				import { b } from './index.js'
				console.log(b.a.x, b.a.missing, b.missing, b.a.missing.deep, b.a.missing)
				b.a.missing()
			`,
			"/index.js": `
				export * as b from './b.js'
			`,
			"/b.js": `
				export * as a from './a.js'
				export const unusedB = 1
			`,
			"/a.js": `
				export const x = 1
				export const unusedA = 2
			`,
		},
		entryPaths: []string{"/entry.js"},
		options: config.Options{
			Mode:          config.ModeBundle,
			OutputFormat:  config.FormatESModule,
			AbsOutputFile: "/out.js",
		},
		expectedCompileLog: `entry.js: WARNING: Import "missing" will always be undefined because there is no matching export in "a.js"
entry.js: WARNING: Import "missing" will always be undefined because there is no matching export in "b.js"
`,
	})
}

func TestReExportedNamespaceDeepMutationAndIndex(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/assign.js": `
				import { b } from './index.js'
				console.log(b.a.x, b.a.y)
				b.a.x = 1
			`,
			"/index-access.js": `
				import { b } from './index.js'
				console.log(b.a.x, b.a['y'])
			`,
			"/escape.js": `
				import { b } from './index.js'
				console.log(b.a.x, b.a)
			`,
			"/index.js": `
				export * as b from './b.js'
			`,
			"/b.js": `
				export * as a from './a.js'
				export const unusedB = 1
			`,
			"/a.js": `
				export const x = 1
				export const y = 2
			`,
		},
		entryPaths: []string{"/assign.js", "/index-access.js", "/escape.js"},
		options: config.Options{
			Mode:         config.ModeBundle,
			OutputFormat: config.FormatESModule,
			AbsOutputDir: "/out",
		},
	})
}

func TestReExportedNamespaceReviewRegressions(t *testing.T) {
	importstar_suite.expectBundled(t, bundled{
		files: map[string]string{
			"/dotted-key.js": `
				import { ns } from './index.js'
				console.log(ns.foo, ns['a.x'], ns['no.such'], ns.inner['x'])
			`,
			"/jsx-nested.jsx": `
				import { ns } from './index.js'
				console.log(ns.foo, <ns.inner.Comp />)
			`,
			"/dead-writes.js": `
				import { ns } from './index.js'
				console.log(ns.foo, ns.missing)
				if (0) { delete ns.foo; ns.missing = 1; ns.inner.x++ }
			`,
			"/warning-loc.js": `
				import { ns } from './index.js'
				console.log(ns.inner.nope.y)
			`,
			"/optional.js": `
				import { ns } from './index.js'
				console.log(ns.inner?.x, ns.inner.x)
			`,
			"/pow.js": `
				import { ns } from './index.js'
				import * as star from './ns.js'
				console.log(ns.nope ** 2, star.nope ** 2)
			`,
			"/index.js": `
				export * as ns from './ns.js'
			`,
			"/ns.js": `
				export let foo = 1
				let v = 'dotted'
				export { v as 'a.x' }
				export * as inner from './inner.js'
			`,
			"/inner.js": `
				export let x = 'x'
				export function Comp() {}
				export const unused = 1
			`,
		},
		entryPaths: []string{"/dotted-key.js", "/jsx-nested.jsx", "/dead-writes.js", "/warning-loc.js", "/optional.js", "/pow.js"},
		options: config.Options{
			Mode:         config.ModeBundle,
			OutputFormat: config.FormatESModule,
			AbsOutputDir: "/out",
			JSX: config.JSXOptions{
				Parse:    true,
				Preserve: true,
			},
		},
		expectedCompileLog: `dead-writes.js: WARNING: Import "missing" will always be undefined because there is no matching export in "ns.js"
pow.js: WARNING: Import "nope" will always be undefined because there is no matching export in "ns.js"
pow.js: WARNING: Import "nope" will always be undefined because there is no matching export in "ns.js"
warning-loc.js: WARNING: Import "nope" will always be undefined because there is no matching export in "inner.js"
`,
	})
}
