package main 

import (
	"fmt"
	"context"
	"log"
	"os"
	"github.com/shomali11/slacker"
)
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
func main(){
	os.Setenv("SLACK_BOT_TOKEN","xoxb-5137532802309-5140410324675-UmL9nwzQcqdlq8ejX24HeaOi")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A0541GJ4HQD-5125964668071-eca608ca3ea9437d90a8669ccd9eff2e763316799411728544e6d00b7c7e7bbf")

	bot:=slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvent())

	bot.Command("my YOB is <year>",&slacker.CommandDefinition{
		Description:"YOB calculator",
		Example:"my YOB is 2003",
		Handler:func(BotCtx slacker.BotContext,request slacker.Request,response slacker.ResponseWriter){
			year:=request.Param("year")
			yob,err:=strconv.Atoi(year)
			if err!= nil{
				println("error")
			}
			age :=2023-yob
			r:= fmt.Sprintf("age is %d",age)
			response.Reply(r)
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err:=bot.Listen(ctx)
	if err!= nil{
		log.fatal(err) 
	}				
}