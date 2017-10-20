//line grammer.y:2
package parse

import __yyfmt__ "fmt"

//line grammer.y:2
import "github.com/wreulicke/gojg/ast"

//line grammer.y:8
type yySymType struct {
	yys    int
	ast    ast.AST
	values []ast.AST
	token  Token
}

const MINUS = 57346
const NUMBER = 57347
const FALSE = 57348
const NULL = 57349
const TRUE = 57350
const COLON = 57351
const COMMA = 57352
const STRING = 57353
const STRING_TEMPLATE = 57354
const TEMPLATE_BEGIN = 57355
const TEMPLATE_END = 57356
const OBJECT_BEGIN = 57357
const OBJECT_END = 57358
const BRACE_BEGIN = 57359
const BRACE_END = 57360
const ARRAY_BEGIN = 57361
const ARRAY_END = 57362
const BOOLEAN_PREFIX = 57363
const ID = 57364

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"MINUS",
	"NUMBER",
	"FALSE",
	"NULL",
	"TRUE",
	"COLON",
	"COMMA",
	"STRING",
	"STRING_TEMPLATE",
	"TEMPLATE_BEGIN",
	"TEMPLATE_END",
	"OBJECT_BEGIN",
	"OBJECT_END",
	"BRACE_BEGIN",
	"BRACE_END",
	"ARRAY_BEGIN",
	"ARRAY_END",
	"BOOLEAN_PREFIX",
	"ID",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammer.y:138

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 59

var yyAct = [...]int{

	11, 24, 27, 28, 2, 13, 14, 4, 5, 6,
	29, 22, 18, 19, 15, 33, 16, 25, 35, 20,
	17, 26, 12, 18, 19, 31, 38, 30, 23, 18,
	19, 34, 32, 21, 7, 9, 36, 37, 10, 25,
	39, 13, 14, 4, 5, 6, 8, 3, 18, 19,
	15, 1, 16, 0, 0, 0, 17, 0, 12,
}
var yyPact = [...]int{

	37, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 2, 28, -1000, -11, 12, 1, -1000, -1000,
	-12, -1000, 13, -1000, 9, 23, -1000, -5, 21, 0,
	-1000, -1000, 37, -1000, 37, -1000, 16, -1000, 18, -1000,
}
var yyPgo = [...]int{

	0, 51, 3, 47, 46, 38, 35, 0, 34, 2,
	1,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 7, 7, 8, 8, 4, 3, 6, 6, 10,
	10, 5, 5, 9, 9,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 1, 3, 4, 2, 3, 3,
	5, 2, 3, 1, 3,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, 6, 7, 8, -8, -4, -6,
	-5, -7, 21, 4, 5, 13, 15, 19, 11, 12,
	17, 5, 22, 16, -10, -7, 20, -9, -2, 22,
	14, 16, 9, 20, 10, 18, -2, -9, 10, -10,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 0, 0, 14, 0, 0, 0, 11, 12,
	0, 13, 0, 17, 0, 0, 21, 0, 23, 0,
	15, 18, 0, 22, 0, 16, 19, 24, 0, 20,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22,
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

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:33
		{
			yylex.(*Lexer).result = yyDollar[1].ast
			yyVAL.ast = yyDollar[1].ast
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:39
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:42
		{
			yyVAL.ast = &ast.BooleanNode{Value: false}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:45
		{
			yyVAL.ast = &ast.NullNode{}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:48
		{
			yyVAL.ast = &ast.BooleanNode{Value: true}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:51
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:54
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:57
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:60
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:63
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:68
		{
			yyVAL.ast = &ast.StringNode{Value: yyDollar[1].token.literal}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:71
		{
			yyVAL.ast = &ast.StringNode{ID: &ast.ID{yyDollar[1].token.literal}}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammer.y:76
		{
			lex := yylex.(*Lexer)
			num := lex.parseFloat(yyDollar[1].token.literal + yyDollar[2].token.literal)
			yyVAL.ast = &ast.NumberNode{Value: num}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:81
		{
			lex := yylex.(*Lexer)
			num := lex.parseFloat(yyDollar[1].token.literal)
			yyVAL.ast = &ast.NumberNode{Value: num}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:88
		{
			yyVAL.ast = &ast.RawValueTemplateNode{ID: &ast.ID{yyDollar[2].token.literal}}
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammer.y:93
		{
			yyVAL.ast = &ast.BooleanNode{ID: &ast.ID{yyDollar[3].token.literal}}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammer.y:98
		{
			yyVAL.ast = &ast.ObjectNode{Members: []ast.AST{}}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:102
		{
			yyVAL.ast = &ast.ObjectNode{Members: yyDollar[2].values}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:107
		{
			yyVAL.values = []ast.AST{&ast.MemberNode{Name: yyDollar[1].ast, Value: yyDollar[3].ast}}
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line grammer.y:110
		{
			size := len(yyDollar[5].values) + 1
			values := make([]ast.AST, 0, size)
			values = append(values, &ast.MemberNode{Name: yyDollar[1].ast, Value: yyDollar[3].ast})
			values = append(values, yyDollar[5].values...)
			yyVAL.values = values
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammer.y:119
		{
			yyVAL.ast = &ast.ArrayNode{Value: []ast.AST{}}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:122
		{
			yyVAL.ast = &ast.ArrayNode{Value: yyDollar[2].values}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:127
		{
			yyVAL.values = []ast.AST{yyDollar[1].ast}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:130
		{
			size := len(yyDollar[3].values) + 1
			values := make([]ast.AST, 0, size)
			values = append(values, yyDollar[1].ast)
			values = append(values, yyDollar[3].values...)
			yyVAL.values = values
		}
	}
	goto yystack /* stack new state and value */
}
