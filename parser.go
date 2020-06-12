package main



type Parser interface {
    func kindOf() SyntaxKind
    func nextOf() SyntaxKind
    func returnOf() SyntaxKind
}



type SyntaxKind struct {
    tokentype string

}


func (p Parser) parseComment() {


}

func (p Parser) kindOf(t Token) SyntaxKind {

    var kind SyntaxKind

    if t.lexeme == "var" {
	kind.tokentype = "keyword"
    } else if t.lexeme == "const" {
        kind.tokentype = "keyword"
    } else if t.lexeme == "func" {
        kind.tokentype = "keyword"
    } else if t.lexeme == "type" {
        kind.tokentype = "keyword"
    } else if t.lexeme == "interface" {
        kind.tokentype = "keyword"
    } else if t.lexeme == "struct" {
        kind.tokentype = "keyword"
    } else if t.lexeme == "import" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "package" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "if" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "else" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "switch" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "goto" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "case" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "break" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "chan" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "continue" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "default" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "defer" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "fallthrough" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "for" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "go" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "map" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "range" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "return" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "select" {
       kind.tokentype = "keyword"
    } else if t.lexeme == "type" {
       kind.tokentype = "keyword"
    }

    return kind

}

func (p Parser) nextOf(t Token) SyntaxKind {


}

func (p Parser) returnOf(t Token) SyntaxKind {
    if p.kindOf(t) !=  {
        return nextOf(t)
    } else {
        panic("Error parsing current Token")
    }
}




func (p Parser) parse(s Scanner) {
        token := s.nextToken()
	return result(token)
}
