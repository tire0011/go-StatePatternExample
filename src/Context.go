package IRPCgame

type Context struct {
	State IState
}

func (context *Context)Request(key string) {

	var result string
	switch(key){
		case "m": 
			result = context.State.Move(context)
			break
		case "a":
			result = context.State.Attack(context)
			break
		case "s":
			result = context.State.Stop(context)
			break
		case "r":
			result = context.State.Run(context)
			break	
		case "p":
			result = context.State.Panic(context)
			break	
		case "c":
			result = context.State.CalmDown(context)
			break	
		case "e":
			println("thank you for playing the RPC game")
			break	
		default:
			println("there exit not a action with that key, maybe you want to programm it to me")		
		}
	
	println(result)
}
	