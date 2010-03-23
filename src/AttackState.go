package IRPCgame

import "rand"

type AttackState struct {

}

func (attack *AttackState) Move(context *Context) string {
	return "you need to stop attacking first"
}

func (attack *AttackState) Attack(context *Context) string {
	return "you attack the darkness for " + string(rand.Intn(20)) + "damage"
}

func (attack *AttackState) Stop(context *Context) string {
	context.State = new(RestingState)
	return "you are calm down and come to rest"
}

func (attack *AttackState) Run(context *Context) string {
	context.State = new(MoveState)
	return "you run away from the fray"
}

func (attack *AttackState) Panic(context *Context) string {
	context.State = new(PanickState)
	return context.State.Text()
}

func (attack *AttackState) CalmDown(context *Context) string {
	context.State = new(RestingState)
	return context.State.Text()
}

func (attack *AttackState) Text() string {
	return "you start attacking the darkness"
}
	
func (attack *AttackState) Identify() string {
	return "I am the attack state"
}
