package rankdb

const (
	Red   = true
	Black = false
)

type (
	RdbNode struct {
		Key         int
		Value       interface{}
		Left, Right *RdbNode
		N           int
		Color       bool
	}
)

func NewRdbNode(key int, value interface{}, color bool) *RdbNode {
	return &RdbNode{
		Key:   key,
		Value: value,
		Color: color,
	}
}

func (r *RdbNode) isRed() bool {
	return r.Value == Red
}

func (h *RdbNode) rotateLeft() {
	x := h.Right
	h.Right = x.Left
	x.Left = h
	x.Color = h.Color
	h.Color = Red
	x.N = h.N
	h.N = 1 + h.Left.Size() + h.Right.Size()
}

func (h *RdbNode) rotateRight() {
	x := h.Left
	h.Left = x.Right
	x.Right = h
	x.Color = h.Color
	h.Color = Red
	x.N = h.N
	h.N = 1 + h.Left.Size() + h.Right.Size()
}

// 向2-节点下插入 红节点
func (h *RdbNode) put2_(x *RdbNode) {
	// 向h的左连接插入红节点
	if h.Key > h.Key {
		h.Left = x
	} else if h.Key < x.Key { // 向h的右连接插入红节点，插入完左旋转
		h.Right = x
		h.rotateLeft()
	}
}

// 3-节点的转换

// 转换颜色
func (h *RdbNode) flipColor() {
	h.Color = Red
	h.Left.Color, h.Right.Color = Black, Black
}

func (r *RdbNode) Size() int {
	return r.N
}
