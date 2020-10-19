package lib

/*
The grammar for the lexer/parser

---
program: command

command: command AND command | command OR command | env exec args

env: env,env | k=v
exec: IDENT
args: args args | IDENT

AND: &&
OR: ||

k: IDENT
v: IDENT | command
---
*/
