package group

type Group interface {
	Group(GroupQuery)
	GroupList(...GroupQuery)
}

type GroupQuery struct {
	ID int
}
