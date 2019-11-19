package php

//#include "tree_sitter/parser.h"
//TSLanguage *tree_sitter_php();
import "C"
import (
	"unsafe"

	sitter "github.com/smacker/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_php())
	return sitter.NewLanguage(ptr)
}
