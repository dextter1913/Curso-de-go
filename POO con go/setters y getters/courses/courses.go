package courses

type course struct {
	name    string
	price   float64
	isFree  bool
	userIds []uint
	classes map[uint]string
}

func New() *course {
	return &course{}
}

func (c *course) SetName(name string) {
	c.name = name
}

func (c *course) SetPrice(price float64) {
	c.price = price
}

func (c *course) SetIsFree(isFree bool) {
	c.isFree = isFree
}

func (c *course) SetUserIds(userIds ...int) {
	for _, userId := range userIds {
		c.userIds = append(c.userIds, uint(userId))
	}
}

// func (c *course) SetClasses(classes map[uint]string) {
// 	c.classes = classes
// }

// haciendo el mapa dinamico
func (c *course) SetClasses(classes ...string) {
	c.classes = make(map[uint]string)
	for i, v := range classes {
		c.classes[uint(i)] = v
	}
}

func (c *course) GetName() string {
	return c.name
}

func (c *course) GetPrice() float64 {
	return c.price
}

func (c *course) GetIsFree() bool {
	return c.isFree
}

func (c *course) GetUserIds() []uint {
	return c.userIds
}

func (c *course) GetClasses() map[uint]string {
	return c.classes
}
