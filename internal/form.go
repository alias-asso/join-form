package internal

import (
	"log"

	"github.com/charmbracelet/huh"
)

var (
	name     string
	username string
	password string
	teams    []string
	more     string
)

func DisplayForm() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Quel est ton nom ?").Value(&name),
			huh.NewInput().Title("Quel nom d'utilisateur veux-tu utiliser ?").Value(&username),
			huh.NewInput().Title("Quel mot-de passe veux-tu utiliser ?").EchoMode(huh.EchoModePassword).Value(&username),
		),
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Choisis ta team (tu peux en séléctionner plusieurs).").
				Options(
					huh.NewOption("Animation", "animation"),
					huh.NewOption("Communication", "communication"),
					huh.NewOption("Développement", "developpement"),
					huh.NewOption("Logistique", "logistique"),
					huh.NewOption("Secrétariat", "secretariat"),
					huh.NewOption("Trésorerie", "tresorerie"),
				).
				Value(&teams),
		),
		huh.NewGroup(
			huh.NewText().Title("Quelque chose à ajouter ?").Value(&more),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
}
