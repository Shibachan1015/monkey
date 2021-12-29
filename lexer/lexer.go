package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  //入力における現在の位置（現在の文字を指し示す）
	readPosition int  //これから読み込む位置（現在の文字の次）
	ch           byte //現在検査中の文字
}

func New(input string) *Lexer {  //New関数
	l := &Lexer{input: input}
	l.readChar()
	return l
}
// p32
func (l *Lexer) readChar() {    //メソッド
	if l.readPosition >= len(l.input) {
		l.ch = 0                        // 入力の終端に到達したらl.chを０にする。“この値はASCIIコードの"NUL"文字に対応していて、私たちは「まだ何も読み込んでいない」あるいは「ファイルの終わり」を表すために使う。　
	} else {
		l.ch = l.input[l.readPosition]  // l.readPositionは常に次に読もうとしている場所を指す
	}
	l.position = l.readPosition  // l.positionは常に最後に読んだ場所を指す
	l.readPosition += 1
}
// p33
func (l *Lexer) NextToken() token.Token { // token.Token型のメソッド？
	var tok token.Token  // 変数 tok をtoken.Token型で宣言i

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token .Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok.Literal = ""       //空文字を挿入　token.goのstructのフィールドから
		tok.Type = token.EOF
	// p38 l.chが認識された文字ではないときに識別子かどうかを点検できるようにdefault分岐をswitch分に追加
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL,l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace(){
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter (l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 	ヘルパー関数　与えられた引数が英字かどうかを判定する p38
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
