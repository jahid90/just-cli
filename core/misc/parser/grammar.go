package parser

/*
The grammar for the lexer/parser

---
program: command

command: andCommand | orCommand | commandExpression

andCommand: command '&&' command
orCommand: command '||' command
commandExpression: env exec args | env exec | exec args | exec

env: multiEnv
multiEnv: keyValEnv,env | keyValEnv

keyValEnv: IDENT'='envVal
envVal: IDENT | expression

exec: IDENT

args: multiArgs					(matches)
multiArgs: IDENT | expression | IDENT args | expression args
						 (IDENT expression args)
expression: '$('command')'
---
*/
