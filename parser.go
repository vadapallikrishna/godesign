package main



type Parser interface {
    func kindOf(t Token) SyntaxKind
    func nextOf(t Token) SyntaxKind
    func iserror(s SyntaxKind) ErrorKind
    func check(s SyntaxKind) 
}


type SyntaxKind struct {
    tokentype string
    nexttoken string
}

type ErrorKind struct {
    expectedtype string
}



func (p Parser) kindOf(t Token) SyntaxKind {

    var kind SyntaxKind

    if t.lexeme == "var" {
	kind.tokentype = "keyword"
	kind.nexttoken = "identifier"
    } else if t.lexeme == "const" {
        kind.tokentype = "keyword"
	kind.nexttoken = "identifier"
    } else if t.lexeme == "func" {
        kind.tokentype = "keyword"
    } else if t.lexeme == "type" {
        kind.tokentype = "keyword"
	kind.tokentype = "identifier"
    } else if t.lexeme == "interface" {
        kind.tokentype = "keyword"
	kind.nexttoken = "{"
    } else if t.lexeme == "struct" {
        kind.tokentype = "keyword"
	kind.nexttoken = "{"
    } else if t.lexeme == "import" {
       kind.tokentype = "keyword"
       kind.nexttoken = "("
    } else if t.lexeme == "package" {
       kind.tokentype = "keyword"
       kind.nexttoken = "identifier"
    } else if t.lexeme == "if" {
       kind.tokentype = "keyword"
       kind.nexttoken = "identifier"
    } else if t.lexeme == "else" {
       kind.tokentype = "keyword"
       kind.nexttoken = "identifier"
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
     var kind, nextkind SyntaxKind

     kind := kindOf(t)
     if kind.tokentype == "keyword" {
         if kind.nexttoken == "identifier" {
	     nextkind.tokentype = "identifier"
	     return nextkind
	 }
     }


}

func (p Parser) iserror(s SyntaxKind) ErrorKind {

}

func (p Parser) check(s SyntaxKind)  {
    if s.tokentype != p.iserror(s) {
	return  
    } else {
        panic("Error parsing current Token")
    }
}
