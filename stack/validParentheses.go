package stack

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 1.左括号必须用相同类型的右括号闭合。
// 2.左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
// 示例1：
// 输入: "()"
// 输出: true
// 示例2：
// 输入: "()[]{}"
// 输出: true
// 示例3：
// 输入: "(]"
// 输出: false
// 示例4：
// 输入: "([)]"
// 输出: false
// 示例5：
// 输入: "{[]}"
// 输出: true
func isValid(s string) bool {
	n := len(s)
	if n&1 > 0 {
		return false
	}
	stack := []byte{}
	pairs := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || pairs[stack[len(stack)-1]] != s[i] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

func isValid2(s string) bool {
	stack := Stack{}
	for _, c := range s {
		if _, ok := parentheses[c]; ok {
			stack.Push(c)
		} else {
			e := stack.Pop()
			if e == nil {
				return false
			}
			if parentheses[e.(int32)] != c {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

var parentheses = map[int32]int32{'(': ')', '[': ']', '{': '}'}

type element struct {
	data interface{}
	next *element
}

type Stack struct {
	element *element
}

func (p *Stack) Pop() interface{} {
	if p.IsEmpty() {
		return nil
	}
	e := p.element.data
	p.element = p.element.next
	return e
}

func (p *Stack) Push(e interface{}) {
	node := &element{
		data: e,
		next: p.element,
	}
	p.element = node
}

func (p *Stack) IsEmpty() bool {
	return p.element == nil
}
