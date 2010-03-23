package IRPCgame

type RestingState struct {

}

func (resting *RestingState) Move(context *Context) string {
	context.State = new(MoveState)
	return context.State.Text()
}

func (resting *RestingState) Attack(context *Context) string {
	context.State = new(AttackState)
	return context.State.Text()
}

func (resting *RestingState) Stop(context *Context) string {
	return "you already stopped!"
}

func (resting *RestingState) Run(context *Context) string {
	return "you can not run unless you are moving"
}

func (resting *RestingState) Panic(context *Context) string {
	context.State = new(PanickState)
	return context.State.Text()
}

func (resting *RestingState) CalmDown(context *Context) string {
	return "you are already relaxed"
}

func (resting *RestingState) Text() string {
	return "you fall down and sleep"
}
	
func (resting *RestingState) Identify() string {
	return "I am the resting state"
}
