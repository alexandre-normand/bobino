package main

import (
	"github.com/alexandre-normand/bobino"
	"github.com/zchee/color"
)

func main() {
	nom := bobino.AskForInput("Quel est ton nom?")
	color.Magenta("🤝 Enchanté, %s!", nom)

	age := bobino.AskForNumber("Quel est ton âge?")
	color.Green("🤯 Tu as %d ans!", age)

	economie := bobino.AskForDecimalNumber("Combien d'argent as-tu économisé cette semaine?")
	color.Yellow("🤑 Tu as économisé $%.2f!", economie)

	animal := bobino.AskWithChoice("Quel est ton animal préféré?", "chat", "chien", "hibou", "cheval", "léopard", "corbeau")
	color.Cyan("Moi aussi mon animal préféré c'est les %ss!", animal)
}
