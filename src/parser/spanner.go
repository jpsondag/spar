// Code generated by goyacc -o src/parser/spanner.go src/parser/spanner.go.y. DO NOT EDIT.
//line src/parser/spanner.go.y:2
package parser

import __yyfmt__ "fmt"

//line src/parser/spanner.go.y:2
//line src/parser/spanner.go.y:5
type yySymType struct {
	yys       int
	empty     struct{}
	str       string
	strs      []string
	col       Column
	cols      []Column
	key       Key
	keys      []Key
	clstr     Cluster
	LastToken int
}

const PRIMARY = 57346
const KEY = 57347
const ASC = 57348
const DESC = 57349
const INTERLEAVE = 57350
const IN = 57351
const PARENT = 57352
const ARRAY = 57353
const OPTIONS = 57354
const NOT = 57355
const NULL = 57356
const ON = 57357
const DELETE = 57358
const CASCADE = 57359
const NO = 57360
const ACTION = 57361
const MAX = 57362
const UNIQUE = 57363
const NULL_FILTERED = 57364
const STORING = 57365
const true = 57366
const null = 57367
const allow_commit_timestamp = 57368
const CREATE = 57369
const ALTER = 57370
const DROP = 57371
const DATABASE = 57372
const TABLE = 57373
const INDEX = 57374
const BOOL = 57375
const INT64 = 57376
const FLOAT64 = 57377
const STRING = 57378
const BYTES = 57379
const DATE = 57380
const TIMESTAMP = 57381
const database_id = 57382
const table_name = 57383
const column_name = 57384
const index_name = 57385
const decimal_value = 57386
const hex_value = 57387

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"PRIMARY",
	"KEY",
	"ASC",
	"DESC",
	"INTERLEAVE",
	"IN",
	"PARENT",
	"ARRAY",
	"OPTIONS",
	"NOT",
	"NULL",
	"ON",
	"DELETE",
	"CASCADE",
	"NO",
	"ACTION",
	"MAX",
	"UNIQUE",
	"NULL_FILTERED",
	"STORING",
	"true",
	"null",
	"allow_commit_timestamp",
	"'('",
	"','",
	"')'",
	"';'",
	"CREATE",
	"ALTER",
	"DROP",
	"DATABASE",
	"TABLE",
	"INDEX",
	"BOOL",
	"INT64",
	"FLOAT64",
	"STRING",
	"BYTES",
	"DATE",
	"TIMESTAMP",
	"database_id",
	"table_name",
	"column_name",
	"index_name",
	"decimal_value",
	"hex_value",
	"'<'",
	"'>'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 110

var yyAct = [...]int{

	85, 68, 67, 66, 55, 46, 24, 28, 37, 30,
	31, 32, 33, 34, 35, 36, 101, 93, 53, 69,
	23, 96, 83, 47, 16, 15, 20, 6, 14, 98,
	97, 65, 56, 57, 30, 31, 32, 33, 34, 35,
	36, 11, 12, 72, 78, 72, 73, 82, 10, 9,
	8, 64, 25, 90, 59, 49, 26, 87, 63, 60,
	45, 44, 19, 21, 58, 18, 71, 102, 99, 100,
	94, 89, 38, 52, 79, 43, 77, 91, 70, 86,
	75, 76, 50, 62, 40, 2, 92, 7, 81, 84,
	41, 95, 80, 51, 5, 4, 3, 1, 17, 13,
	42, 88, 61, 48, 39, 74, 54, 29, 27, 22,
}
var yyPact = [...]int{

	-4, -4, -1000, 20, 19, 18, 7, -1000, -1000, -1000,
	-1000, -19, -21, 43, -1000, -1000, 35, -10, -1000, -26,
	-41, 23, 28, -3, 57, 80, -26, 62, -1000, -1000,
	-1000, -1000, -1000, 34, 33, -1000, -1000, -45, -22, 27,
	77, -1000, -1000, 59, -16, -16, -28, 32, -1000, 75,
	31, -1000, -1000, 22, -1000, -1000, -1000, -1000, 2, -48,
	-27, -1000, 69, -27, -1000, -1000, -1000, 17, -1000, 74,
	66, 15, -27, 24, -1000, -1000, -1000, -23, -1000, -1000,
	71, -1000, 30, 56, 25, -1000, 68, -29, -1000, 54,
	71, -24, 1, -1000, 51, -1000, -1000, -1000, -30, -1000,
	48, -1000, -1000,
}
var yyPgo = [...]int{

	0, 109, 63, 108, 7, 107, 18, 106, 105, 1,
	2, 104, 103, 102, 101, 100, 99, 98, 97, 85,
	96, 95, 94, 93, 92, 89, 88, 86, 0,
}
var yyR1 = [...]int{

	0, 18, 18, 19, 19, 19, 20, 21, 2, 2,
	2, 1, 11, 10, 10, 9, 8, 8, 8, 12,
	12, 13, 14, 14, 14, 3, 3, 4, 4, 4,
	4, 4, 4, 4, 6, 6, 5, 23, 15, 15,
	22, 16, 16, 17, 17, 24, 24, 26, 27, 27,
	25, 25, 25, 28, 7, 7,
}
var yyR2 = [...]int{

	0, 1, 2, 2, 2, 2, 3, 8, 0, 1,
	3, 4, 5, 1, 3, 2, 0, 1, 1, 0,
	2, 5, 0, 3, 4, 1, 1, 1, 1, 1,
	4, 4, 1, 1, 1, 1, 4, 0, 0, 2,
	12, 0, 1, 0, 1, 0, 1, 4, 1, 3,
	0, 1, 3, 3, 1, 1,
}
var yyChk = [...]int{

	-1000, -18, -19, -20, -21, -22, 31, -19, 30, 30,
	30, 34, 35, -16, 21, 44, 45, -17, 22, 27,
	36, -2, -1, 46, 47, 29, 28, -3, -4, -5,
	37, 38, 39, 40, 41, 42, 43, 11, 15, -11,
	4, -2, -15, 13, 27, 27, 50, 45, -12, 28,
	5, -23, 14, -6, -7, 20, 48, 49, -6, -4,
	27, -13, 8, 27, 29, 29, 51, -10, -9, 46,
	9, -10, 28, 29, -8, 6, 7, 10, 29, -9,
	-24, -26, 23, 45, -25, -28, 8, 27, -14, 15,
	28, 9, -27, 46, 16, -28, 45, 29, 28, 17,
	18, 46, 19,
}
var yyDef = [...]int{

	0, -2, 1, 0, 0, 0, 41, 2, 3, 4,
	5, 0, 0, 43, 42, 6, 0, 0, 44, 8,
	0, 0, 9, 0, 0, 0, 8, 38, 25, 26,
	27, 28, 29, 0, 0, 32, 33, 0, 0, 19,
	0, 10, 37, 0, 0, 0, 0, 0, 7, 0,
	0, 11, 39, 0, 34, 35, 54, 55, 0, 0,
	0, 20, 0, 0, 30, 31, 36, 0, 13, 16,
	0, 0, 0, 45, 15, 17, 18, 0, 12, 14,
	50, 46, 0, 22, 40, 51, 0, 0, 21, 0,
	0, 0, 0, 48, 0, 52, 53, 47, 0, 23,
	0, 49, 24,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	27, 29, 3, 3, 28, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 30,
	50, 3, 51,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 49,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:73
		{
			SetCreateDatabaseStatement(yylex, yyDollar[1].str, yyDollar[2].str, yyDollar[3].str)
		}
	case 7:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line src/parser/spanner.go.y:79
		{
			SetCreateTableStatement(yylex, yyDollar[1].str, yyDollar[2].str, yyDollar[3].str, yyDollar[5].cols, yyDollar[7].keys, yyDollar[8].clstr)
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:85
		{
			yyVAL.cols = make([]Column, 0, 0)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:89
		{
			yyVAL.cols = make([]Column, 0, 1)
			yyVAL.cols = append(yyVAL.cols, yyDollar[1].col)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:94
		{
			yyVAL.cols = append(yyDollar[3].cols, yyDollar[1].col)
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:100
		{
			yyVAL.col = Column{Name: yyDollar[1].str, Type: yyDollar[2].str, Nullability: yyDollar[3].str}
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line src/parser/spanner.go.y:106
		{
			yyVAL.keys = yyDollar[4].keys
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:112
		{
			yyVAL.keys = make([]Key, 0, 1)
			yyVAL.keys = append(yyVAL.keys, yyDollar[1].key)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:117
		{
			yyVAL.keys = append(yyDollar[1].keys, yyDollar[3].key)
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line src/parser/spanner.go.y:123
		{
			yyVAL.key = Key{Name: yyDollar[1].str, Order: yyDollar[2].str}
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:129
		{
			yyVAL.str = "ASC"
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:133
		{
			yyVAL.str = yyDollar[1].str
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:137
		{
			yyVAL.str = yyDollar[1].str
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:143
		{
			yyVAL.clstr = Cluster{}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line src/parser/spanner.go.y:147
		{
			yyVAL.clstr = yyDollar[2].clstr
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line src/parser/spanner.go.y:153
		{
			yyVAL.clstr = Cluster{TableName: yyDollar[4].str, OnDelete: yyDollar[5].str}
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:159
		{
			// default
			yyVAL.str = "NO ACTION"
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:164
		{
			yyVAL.str = yyDollar[3].str
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:168
		{
			yyVAL.str = yyDollar[3].str + " " + yyDollar[4].str
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:174
		{
			yyVAL.str = yyDollar[1].str
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:178
		{
			yyVAL.str = yyDollar[1].str
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:184
		{
			yyVAL.str = yyDollar[1].str
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:188
		{
			yyVAL.str = yyDollar[1].str
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:192
		{
			yyVAL.str = yyDollar[1].str
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:196
		{
			yyVAL.str = yyDollar[1].str + "(" + yyDollar[3].str + ")"
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:200
		{
			yyVAL.str = yyDollar[1].str + "(" + yyDollar[3].str + ")"
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:204
		{
			yyVAL.str = yyDollar[1].str
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:208
		{
			yyVAL.str = yyDollar[1].str
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:214
		{
			yyVAL.str = yyDollar[1].str
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:218
		{
			yyVAL.str = yyDollar[1].str
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:224
		{
			yyVAL.str = yyDollar[1].str + "(" + yyDollar[3].str + ")"
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:237
		{
			yyVAL.str = "NULL"
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line src/parser/spanner.go.y:241
		{
			yyVAL.str = yyDollar[1].str + " " + yyDollar[2].str
		}
	case 40:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line src/parser/spanner.go.y:247
		{
			// TODO Support storing_clause_opt, interleave_clause_list
			SetCreateIndexStatement(yylex, yyDollar[1].str, yyDollar[4].str, yyDollar[5].str, yyDollar[2].str, yyDollar[3].str, yyDollar[7].str, yyDollar[9].keys)
		}
	case 41:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:254
		{
			yyVAL.str = ""
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:258
		{
			yyVAL.str = yyDollar[1].str
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:264
		{
			yyVAL.str = ""
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:268
		{
			yyVAL.str = yyDollar[1].str
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:311
		{
			yyVAL.str = yyDollar[1].str
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:315
		{
			yyVAL.str = yyDollar[1].str
		}
	}
	goto yystack /* stack new state and value */
}
