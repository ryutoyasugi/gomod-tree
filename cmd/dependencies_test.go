package cmd

import (
	"reflect"
	"testing"
)

func Test_newPackage(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Package
	}{
		{"main package", args{"gomod-tree"}, Package{"gomod-tree", ""}},
		{"dependency package", args{"github.com/spf13/cobra@v1.4.0"}, Package{"github.com/spf13/cobra", "v1.4.0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPackage(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackage_isMain(t *testing.T) {
	type fields struct {
		name    string
		version string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"main package", fields{"gomod-tree", ""}, true},
		{"dependency package", fields{"github.com/spf13/cobra", "v1.4.0"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Package{
				name:    tt.fields.name,
				version: tt.fields.version,
			}
			if got := p.isMain(); got != tt.want {
				t.Errorf("isMain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newPackagePair(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantParent Package
		wantChild  Package
	}{
		{
			"test",
			args{"github.com/spf13/cobra@v1.4.0 github.com/spf13/pflag@v1.0.5"},
			Package{"github.com/spf13/cobra", "v1.4.0"},
			Package{"github.com/spf13/pflag", "v1.0.5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParent, gotChild := newPackagePair(tt.args.s)
			if !reflect.DeepEqual(gotParent, tt.wantParent) {
				t.Errorf("newPackagePair() gotParent = %v, want %v", gotParent, tt.wantParent)
			}
			if !reflect.DeepEqual(gotChild, tt.wantChild) {
				t.Errorf("newPackagePair() gotChild = %v, want %v", gotChild, tt.wantChild)
			}
		})
	}
}
