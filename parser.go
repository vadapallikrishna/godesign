package main

import "fmt"

type Parser interface {
    func kindOf(t Token) SyntaxKind
    func nextOf(t Token) SyntaxKind
    func isError(s SyntaxKind) error
    func check(s SyntaxKind) NodeKind
    func parse(s SyntaxKind) NodeKind

}

type SyntaxKind struct {

}

type ErrorKind struct {

}

type NodeKind struct {

}

func (p Parser) kindOf(t Token) SyntaxKind {
    var kind SyntaxKind
    return kind

}

func (p Parser) nextOf(t Token) SyntaxKind {
     var kind, nextkind SyntaxKind
}

func (p Parser) isError(s SyntaxKind) error {

}

func (p Parser) check(s SyntaxKind) NodeKind  {
    err := p.iserror(s)
    if err != nil {
        fmt.println(err)
    } else {
	return parse(s)
    }
}

func (p Parser) parse(s SyntaxKind) NodeKind {
    var r ResultNode
}

