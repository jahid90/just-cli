package parser

import (
	"fmt"
)

// Node Represents a node in the AST
type Node interface {
	eval()
	Print(level int) string
}

// CommandNode Represents a node consisting of a complete command
type CommandNode struct {
	andCommand *AndCommandNode
	orCommand  *OrCommandNode
	command    *CommandExpressionNode
}

// Print print
func (cn *CommandNode) Print(level int) {
	prefix := getPrefix(level)

	if cn == nil {
		fmt.Println(prefix + "CommandNode = nil")
		return
	}

	fmt.Println(prefix + "CommandNode")
	cn.andCommand.Print(level + 1)
	cn.orCommand.Print(level + 1)
	cn.command.Print(level + 1)
}

// AndCommandNode Reprecents a pair of commands joined via an AND operator
type AndCommandNode struct {
	left  *CommandNode
	right *CommandNode
}

// Print print
func (acn *AndCommandNode) Print(level int) {
	prefix := getPrefix(level)

	if acn == nil {
		fmt.Println(prefix + "AndCommandNode = nil")
		return
	}

	fmt.Println(prefix + "AndCommandNode")
	acn.left.Print(level + 1)
	acn.right.Print(level + 1)
}

// OrCommandNode Represents a pair of commands joined via an OR operator
type OrCommandNode struct {
	left  *CommandNode
	right *CommandNode
}

// Print print
func (ocn *OrCommandNode) Print(level int) {
	prefix := getPrefix(level)

	if ocn == nil {
		fmt.Println(prefix + "OrCommandNode = nil")
		return
	}

	fmt.Println(prefix + "OrCommandNode")
	ocn.left.Print(level + 1)
	ocn.right.Print(level + 1)
}

// CommandExpressionNode node
type CommandExpressionNode struct {
	env  *EnvNode
	exec *ExecNode
	args *ArgNode
}

// Print print
func (cen *CommandExpressionNode) Print(level int) {
	prefix := getPrefix(level)

	if cen == nil {
		fmt.Println(prefix + "CommandExpressionNode = nil")
		return
	}

	fmt.Println(prefix + "CommandExpressionNode")
	cen.env.Print(level + 1)
	cen.exec.Print(level + 1)
	cen.args.Print(level + 1)
}

// EnvNode Represents a node containing env variables
type EnvNode struct {
	multi *MultiEnvNode
}

// Print print
func (en *EnvNode) Print(level int) {
	prefix := getPrefix(level)

	if en == nil {
		fmt.Println(prefix + "EnvNode = nil")
		return
	}

	fmt.Println(prefix + "EnvNode")
	en.multi.Print(level + 1)
}

// MultiEnvNode node
type MultiEnvNode struct {
	keyVal *KeyValNode
	env    *EnvNode
}

// Print print
func (men *MultiEnvNode) Print(level int) {
	prefix := getPrefix(level)

	if men == nil {
		fmt.Println(prefix + "MultiEnvNode = nil")
		return
	}

	fmt.Println(prefix + "MultiEnvNode")
	men.keyVal.Print(level + 1)
	men.env.Print(level + 1)
}

// KeyValNode node
type KeyValNode struct {
	key   *IdentNode
	value *EnvValNode
}

// Print print
func (kvn *KeyValNode) Print(level int) {
	prefix := getPrefix(level)

	if kvn == nil {
		fmt.Println(prefix + "MultiEnvNode = nil")
		return
	}

	fmt.Println(prefix + "KeyValueNode")
	kvn.key.Print(level + 1)
	kvn.value.Print(level + 1)
}

// EnvValNode Represents the value component of an environment variable pair
type EnvValNode struct {
	command *CommandNode
	value   *IdentNode
}

// Print print
func (evn *EnvValNode) Print(level int) {
	prefix := getPrefix(level)

	if evn == nil {
		fmt.Println(prefix + "EnvValNode = nil")
		return
	}

	fmt.Println(prefix + "EnvValNode")
	evn.value.Print(level + 1)
	evn.command.Print(level + 1)
}

// ExecNode Represents a node consisting of an executable
type ExecNode struct {
	value *IdentNode
}

// Print print
func (en *ExecNode) Print(level int) {
	prefix := getPrefix(level)

	if en == nil {
		fmt.Println(prefix + "ExecNode = nil")
		return
	}

	fmt.Println(prefix + "ExecNode")
	en.value.Print(level + 1)
}

// ArgNode Represents a node consisting of arguments to be passed to an executable
type ArgNode struct {
	multi *MultiArgNode
}

// Print print
func (an *ArgNode) Print(level int) {
	prefix := getPrefix(level)

	if an == nil {
		fmt.Println(prefix + "ArgNode = nil")
		return
	}

	fmt.Println(prefix + "ArgNode")
	an.multi.Print(level + 1)
}

// MultiArgNode node
type MultiArgNode struct {
	arg     *ArgNode
	command *CommandNode
	value   *IdentNode
}

// Print print
func (man *MultiArgNode) Print(level int) {
	prefix := getPrefix(level)

	if man == nil {
		fmt.Println(prefix + "MultiArgNode = nil")
		return
	}

	fmt.Println(prefix + "MultiArgNode")
	man.value.Print(level + 1)
	man.command.Print(level + 1)
	man.arg.Print(level + 1)
}

// IdentNode node
type IdentNode struct {
	value string
}

// Print print
func (in *IdentNode) Print(level int) {
	prefix := getPrefix(level)

	if in == nil {
		fmt.Println(prefix + "IdentNode = nil")
		return
	}

	fmt.Println(prefix + "IdentNode = " + in.value)
}

func getPrefix(level int) string {
	var prefix string
	for i := 0; i < level; i++ {
		prefix += "  "
	}

	return prefix
}
