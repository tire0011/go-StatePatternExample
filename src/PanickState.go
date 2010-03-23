package IRPCgame

type PanickState struct {

}

func (panick *PanickState) Move(context *Context) string {
	return "you move around randomly in a blind panic"
}

func (panick *PanickState) Attack(context *Context) string {
	return "you start attacking the darkness, but keep on missing"
}

func (panick *PanickState) Stop(context *Context) string {
	context.State = new(MoveState)
	return "you start relaxing, but keep on moving"
}

func (panick *PanickState) Run(context *Context) string {
	return "you run around in your panic"
}

func (panick *PanickState) Panic(context *Context) string {
	return "you are already in a panic"
}

func (panick *PanickState) CalmDown(context *Context) string {
	context.State = new(RestingState)
	return "you stand still and relax"
}

func (panick *PanickState) Text() string {
	return "you start panicking and begin seeing thinks"
}
	
func (panick *PanickState) Identify() string {
	return "I am the panick state"	
}
