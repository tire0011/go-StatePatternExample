package IRPCgame

type MoveState struct {

}

func (move *MoveState) Move(context *Context) string {
	return "you move around randomly"
}

func (move *MoveState) Attack(context *Context) string {
	return "you need to stop moving first"
}

func (move *MoveState) Stop(context *Context) string {
	context.State = new(RestingState)
	return "you stand still in a dark room"
}

func (move *MoveState) Run(context *Context) string {
	return "you run around in circles"
}

func (move *MoveState) Panic(context *Context) string {
	context.State = new(PanickState)
	return context.State.Text()
}

func (move *MoveState) CalmDown(context *Context) string {
	context.State = new(RestingState)
	return context.State.Text()
}

func (move *MoveState) Text() string {
	return "you start moving"
}
	
func (move *MoveState) Identify() string {
	return "I am the move state"
}
