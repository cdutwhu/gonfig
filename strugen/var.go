package strugen

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/str"
)

var (
	fSln             = fmt.Sprintln
	fSf              = fmt.Sprintf
	fPln             = fmt.Println
	fEf              = fmt.Errorf
	sIndex           = strings.Index
	sSplit           = strings.Split
	sHasPrefix       = strings.HasPrefix
	sHasSuffix       = strings.HasSuffix
	sTrim            = strings.Trim
	sTrimLeft        = strings.TrimLeft
	sContains        = strings.Contains
	sCount           = strings.Count
	sReplaceAll      = strings.ReplaceAll
	sToUpper         = strings.ToUpper
	ucIsUpper        = unicode.IsUpper
	rxMustCompile    = regexp.MustCompile
	failP1OnErr      = fn.FailP1OnErr
	failP1OnErrWhen  = fn.FailP1OnErrWhen
	warnP1OnErr      = fn.WarnP1OnErr
	warnP1OnErrWhen  = fn.WarnP1OnErrWhen
	warner           = fn.Warner
	enableWarnDetail = fn.EnableWarnDetail
	rmTailFromFirst  = str.RmTailFromFirst
	rmHeadToFirst    = str.RmHeadToFirst
	hasAnyPrefix     = str.HasAnyPrefix
	splitLn          = str.SplitLn
	mustAppendFile   = io.MustAppendFile
	isNumeric        = judge.IsNumeric
)

var (
	selfDev = false
)
