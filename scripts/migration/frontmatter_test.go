package main

import (
	"reflect"
	"testing"
)

func TestUpdateTitle(t *testing.T) {
	testCases := []struct {
		fm       FrontMatter
		body     string
		wantFm   FrontMatter
		wantBody string
	}{
		{
			fm:       FrontMatter{},
			body:     "# My Page Title\nPage content",
			wantFm:   FrontMatter{Title: "My Page Title"},
			wantBody: "\nPage content",
		},
		{
			fm:       FrontMatter{RedirectTo: "/other"},
			body:     "# Title\nContent",
			wantFm:   FrontMatter{RedirectTo: "/other"},
			wantBody: "# Title\nContent",
		},
		{
			fm:       FrontMatter{},
			body:     "No title\nJust content",
			wantFm:   FrontMatter{},
			wantBody: "No title\nJust content",
		},
		{
			fm:       FrontMatter{},
			body:     "# My Page Title\nPage content\n# Another title",
			wantFm:   FrontMatter{Title: "My Page Title"},
			wantBody: "\nPage content\n# Another title",
		},
	}

	for _, tc := range testCases {
		gotFm, gotBody := updateTitles(tc.fm, tc.body, "file.md")
		if !reflect.DeepEqual(gotFm, tc.wantFm) {
			t.Errorf("got fm %v, want %v", gotFm, tc.wantFm)
		}
		if gotBody != tc.wantBody {
			t.Errorf("got body %q, want %q", gotBody, tc.wantBody)
		}
	}
}
