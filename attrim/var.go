package attrim

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/str"
)

var (
	fPln            = fmt.Println
	sJoin           = strings.Join
	sHasPrefix      = strings.HasPrefix
	sHasSuffix      = strings.HasSuffix
	sTrim           = strings.Trim
	sTrimLeft       = strings.TrimLeft
	sContains       = strings.Contains
	sReplaceAll     = strings.ReplaceAll
	failP1OnErr     = fn.FailP1OnErr
	rmTailFromFirst = str.RmTailFromFirst
	splitLn         = str.SplitLn
	mustWriteFile   = io.MustWriteFile
)
