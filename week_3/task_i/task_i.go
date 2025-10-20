package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value    string
	left     *Node
	right    *Node
	width    int
	height   int
	oper_pos int
}

type Parser struct {
	input string
	pos   int
}

func (p *Parser) WatchNext() byte {
	if p.pos >= len(p.input) {
		return 0
	}
	return p.input[p.pos]
}

func (p *Parser) GetNext() byte {
	if p.pos >= len(p.input) {
		return 0
	}
	ch := p.input[p.pos]
	p.pos++
	return ch
}

func (p *Parser) ParseSumSub() *Node {
	node := p.ParseMulDiv()

	for {
		ch := p.WatchNext()
		if ch != '+' && ch != '-' {
			break
		}

		p.GetNext()
		right := p.ParseMulDiv()
		node = &Node{
			value: string(ch),
			left:  node,
			right: right,
		}
	}
	return node
}

func (p *Parser) ParseMulDiv() *Node {
	node := p.ParsePower()

	for {
		ch := p.WatchNext()
		if ch != '*' && ch != '/' {
			break
		}

		p.GetNext()
		right := p.ParsePower()
		node = &Node{
			value: string(ch),
			left:  node,
			right: right,
		}
	}
	return node
}

func (p *Parser) ParsePower() *Node {
	node := p.ParseElement()

	if p.WatchNext() == '^' {
		p.GetNext()
		right := p.ParsePower()
		node = &Node{
			value: "^",
			left:  node,
			right: right,
		}
	}
	return node
}

func (p *Parser) ParseElement() *Node {
	ch := p.WatchNext()
	if ch == '(' {
		p.GetNext()
		node := p.ParseSumSub()
		if p.WatchNext() == ')' {
			p.GetNext()
		}
		return node
	} else if ch >= 'a' && ch <= 'z' {
		return &Node{value: string(p.GetNext())}
	}
	return nil
}

func CalculateSizes(node *Node) {
	if node == nil {
		return
	}

	CalculateSizes(node.left)
	CalculateSizes(node.right)

	if node.left == nil && node.right == nil {
		node.width = 1
		node.height = 1
		node.oper_pos = 0
	} else {
		left_width, left_height := 0, 0
		right_width, right_height := 0, 0

		if node.left != nil {
			left_width = node.left.width
			left_height = node.left.height
		}

		if node.right != nil {
			right_width = node.right.width
			right_height = node.right.height
		}

		node.height = max(left_height, right_height) + 2
		node.width = left_width + right_width + 5
		node.oper_pos = left_width + 2
	}
}

func PaintTree(node *Node, field *[][]byte, row, col int) {
	if node == nil {
		return
	}

	if node.left == nil && node.right == nil {
		(*field)[row][col] = node.value[0]
		return
	}

	left_width := 0
	left_oper, right_oper := 0, 0

	if node.left != nil {
		left_width = node.left.width
		left_oper = node.left.oper_pos
	}
	if node.right != nil {
		right_oper = node.right.oper_pos
	}

	left_start := col
	right_start := col + left_width + 5
	left_vert := left_start + left_oper
	right_vert := right_start + right_oper

	oper_col := col + node.oper_pos

	for i := left_vert + 1; i < oper_col-1; i++ {
		(*field)[row][i] = '-'
	}

	for i := oper_col + 2; i < right_vert; i++ {
		(*field)[row][i] = '-'
	}

	(*field)[row][oper_col-1] = '['
	(*field)[row][oper_col] = node.value[0]
	(*field)[row][oper_col+1] = ']'

	if node.left != nil {
		(*field)[row][left_vert] = '.'
		(*field)[row+1][left_vert] = '|'
	}
	if node.right != nil {
		(*field)[row][right_vert] = '.'
		(*field)[row+1][right_vert] = '|'
	}

	if node.left != nil {
		PaintTree(node.left, field, row+2, left_start)
	}
	if node.right != nil {
		PaintTree(node.right, field, row+2, right_start)
	}
}

func PrintTree(node *Node) {
	if node == nil {
		return
	}

	CalculateSizes(node)

	field := make([][]byte, node.height)
	for i := range field {
		field[i] = make([]byte, node.width)
		for j := range field[i] {
			field[i][j] = ' '
		}
	}

	PaintTree(node, &field, 0, 0)

	for _, line := range field {
		for _, ch := range line {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		parser := &Parser{input: input, pos: 0}
		tree := parser.ParseSumSub()
		PrintTree(tree)
	}
}
