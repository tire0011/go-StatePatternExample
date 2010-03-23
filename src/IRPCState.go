package IRPCgame

type IState interface {
	Move(*Context) string
	Attack(*Context) string
	Stop(*Context) string
	Run(*Context) string
	Panic(*Context) string
	CalmDown(*Context) string
	Text() string
	Identify() string
}
