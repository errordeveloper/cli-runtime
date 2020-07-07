package kubernetes_test

import (
	"io/ioutil"
	"testing"

	. "k8s.io/cli-runtime/pkg/manifest"
)

func TestLoadAllItemsIntoFlattendList(t *testing.T) {
	for _, sample := range []struct {
		path        string
		expectedLen int
	}{
		{
			path:        "testdata/misc-sample-nested-list-1.json",
			expectedLen: 6,
		},
		{
			path:        "testdata/misc-sample-multidoc-nested-lists-1.yaml",
			expectedLen: 4,
		},
		{
			path:        "testdata/misc-sample-empty-list-1.json",
			expectedLen: 0,
		},
		{
			path:        "testdata/misc-sample-multidoc-empty-lists-1.yaml",
			expectedLen: 0,
		},
		{
			path:        "testdata/misc-sample-multidoc-empty-lists-2.yaml",
			expectedLen: 0,
		},
	} {
		data, err := ioutil.ReadFile(sample.path)
		if err != nil {
			t.Fatalf("unexpected error reading sample %q: %v", sample.path, err)
		}
		list, err := NewList(data)
		if err != nil {
			t.Fatalf("unexpected error converting data from %q into a list: %v", sample.path, err)
		}
		if list == nil {
			t.Fatalf("unexpected empty list from %q", sample.path)
		}
		if l := len(list.Items); l != sample.expectedLen {
			t.Fatalf("unexpected lenght of list from %q: expected %d, got %d", sample.path, sample.expectedLen, l)
		}
	}
}
