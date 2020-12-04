package newsfeed


type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

// 一个对象
type Item struct {
	Title string	`json:"title"`
	Post string		`json:"post"`
}

// 将对象放到下面这个数组中
type Repo struct {
	Items []Item
}

// 获取repo的指针
func New() *Repo{
	return &Repo{
		Items:[]Item{},
	}
}

// 添加一个对象
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items,item)
}

// 获取整个repo items的列表
func (r *Repo) GetAll() []Item {
    return r.Items
}