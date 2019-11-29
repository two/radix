package radix

type IRadix interface {
	Insert(string)
}

type Radix struct {
	end      bool
	val      string
	children map[string]*Radix
}

func NewRadix() Radix {
	return Radix{}
}

// Insert put s into a radix tree
func (r *Radix) Insert(s string) {
	node := r
	var index int
	var matchKey string

	if node.children == nil {
		node.children = make(map[string]*Radix)
	}

	for matchKey = range node.children {
		index = prefixIndex(matchKey, s)
		if index > 0 {
			break
		}
	}

	// add new leaf
	if index == 0 {
		node.children[s] = &Radix{
			end: true,
			val: s,
		}
		return
	}

	// equal
	if matchKey == s {
		node.children[s].end = true
		return
	}

	// split
	if index == len(s) {
		// split a new node
		n := node.children[matchKey]
		n.val = matchKey[index:]

		// remote old node
		delete(node.children, matchKey)

		// add new node with split node
		node.children[s] = &Radix{
			end: true,
			val: s,
			children: map[string]*Radix{
				n.val: n,
			},
		}
		return
	}

	// add next node
	if index == len(matchKey) {
		node = node.children[matchKey]
		node.Insert(s[index:])
		return
	}

	// different suffix
	n := node.children[matchKey]
	n.val = matchKey[index:]

	// remote old node
	delete(node.children, matchKey)

	// add new node with two split node
	node.children[matchKey[:index]] = &Radix{
		val: matchKey[:index],
		children: map[string]*Radix{
			n.val: n,
			s[index:]: &Radix{
				end: true,
				val: s[index:],
			},
		},
	}
	return

}

func prefixIndex(s1, s2 string) int {
	if s1 == s2 {
		return len(s1)
	}

	var index int
	l1 := len(s1)
	l2 := len(s2)
	mixLen := l1
	if mixLen > l2 {
		mixLen = l2
	}

	for i := 0; i < mixLen; i++ {
		if s1[i] == s2[i] {
			index++
			continue
		}
		return index
	}
	return index
}
