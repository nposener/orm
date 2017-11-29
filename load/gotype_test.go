package load

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		g                 GoType
		wantString        string
		wantType          string
		wantImportPath    string
		wantNonPointer    string
		wantExtTypeName   string
		wantExtNonPointer string
		wantPackage       string
		wantIsPointer     bool
	}{
		{
			g:                 newGoType("int64"),
			wantType:          "int64",
			wantString:        "int64",
			wantNonPointer:    "int64",
			wantExtTypeName:   "int64",
			wantExtNonPointer: "int64",
		},
		{
			g:                 newGoType("*int64"),
			wantType:          "*int64",
			wantString:        "*int64",
			wantNonPointer:    "int64",
			wantExtTypeName:   "*int64",
			wantExtNonPointer: "int64",
			wantIsPointer:     true,
		},
		{
			g:                 newGoType("github.com/posener/pkg.Struct"),
			wantType:          "Struct",
			wantImportPath:    "github.com/posener/pkg",
			wantString:        "github.com/posener/pkg.Struct",
			wantNonPointer:    "Struct",
			wantExtTypeName:   "pkg.Struct",
			wantExtNonPointer: "pkg.Struct",
			wantPackage:       "pkg",
		},
		{
			g:                 newGoType("*github.com/posener/pkg.Struct"),
			wantType:          "*Struct",
			wantImportPath:    "github.com/posener/pkg",
			wantString:        "*github.com/posener/pkg.Struct",
			wantNonPointer:    "Struct",
			wantExtTypeName:   "*pkg.Struct",
			wantExtNonPointer: "pkg.Struct",
			wantPackage:       "pkg",
			wantIsPointer:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.g.String(), func(t *testing.T) {
			assert.Equal(t, tt.wantString, tt.g.String())
			assert.Equal(t, tt.wantNonPointer, tt.g.NonPointer())
			assert.Equal(t, tt.wantExtTypeName, tt.g.ExtTypeName())
			assert.Equal(t, tt.wantPackage, tt.g.Package())
			assert.Equal(t, tt.wantImportPath, tt.g.ImportPath)
			assert.Equal(t, tt.wantExtNonPointer, tt.g.ExtNonPointer())
			assert.Equal(t, tt.wantIsPointer, tt.g.IsPointer())
		})
	}

}