package treeops

import (
	"testing"
)

var explodeTest = []expressionScenario{
	{
		document:   `{a: mike}`,
		expression: `explode(.a)`,
		expected: []string{
			"D0, P[], (doc)::{a: mike}\n",
		},
	},
	{
		document:   `{f : {a: &a cat, b: *a}}`,
		expression: `explode(.f)`,
		expected: []string{
			"D0, P[], (doc)::{f: {a: cat, b: cat}}\n",
		},
	},
	{
		document:   `{f : {a: &a cat, *a: b}}`,
		expression: `explode(.f)`,
		expected: []string{
			"D0, P[], (doc)::{f: {a: cat, cat: b}}\n",
		},
	},
	{
		document:   mergeDocSample,
		expression: `.foo* | explode(.) | (. style="flow")`,
		expected: []string{
			"D0, P[foo], (!!map)::{a: foo_a, thing: foo_thing, c: foo_c}\n",
			"D0, P[foobarList], (!!map)::{b: bar_b, a: foo_a, thing: bar_thing, c: foobarList_c}\n",
			"D0, P[foobar], (!!map)::{c: foo_c, a: foo_a, thing: foobar_thing}\n",
		},
	},
	{
		document:   mergeDocSample,
		expression: `.foo* | explode(explode(.)) | (. style="flow")`,
		expected: []string{
			"D0, P[foo], (!!map)::{a: foo_a, thing: foo_thing, c: foo_c}\n",
			"D0, P[foobarList], (!!map)::{b: bar_b, a: foo_a, thing: bar_thing, c: foobarList_c}\n",
			"D0, P[foobar], (!!map)::{c: foo_c, a: foo_a, thing: foobar_thing}\n",
		},
	},
	{
		document:   `{f : {a: &a cat, b: &b {f: *a}, *a: *b}}`,
		expression: `explode(.f)`,
		expected: []string{
			"D0, P[], (doc)::{f: {a: cat, b: {f: cat}, cat: {f: cat}}}\n",
		},
	},
}

func TestExplodeOperatorScenarios(t *testing.T) {
	for _, tt := range explodeTest {
		testScenario(t, &tt)
	}
}