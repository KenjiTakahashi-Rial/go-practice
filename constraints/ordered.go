package constraints

type Ordered interface {
	Less(other Ordered) bool
}

type OrderedInt int

func (o OrderedInt) Less(other Ordered) bool {
	return o < other.(OrderedInt)
}
