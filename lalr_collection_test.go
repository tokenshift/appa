package appa

import "os"
import "testing"

func Test_CreatingExampleGrammarLALRCollection(t *testing.T) {
	g, start := createExpressionGrammar()

	out, err := os.Create("lalr_collection_test.dot")
	if err != nil {
		t.Error(err)
		return
	}

	defer out.Close()

	g.WriteLALRCollection(start, out)

	// Verify output manually.
}
