package internal

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/ssh"
)

var (
	name     string
	username string
	password string
	teams    []string
	more     string
)

func FormHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Bienvenue sur Augustine").
				Description("Salut, je suis *Galias* !\nJ'ai développé une technologie pour te parler à travers ce terminal.\nJe vais t'assister dans la création de ton **compte ALIAS**, notre serveur. Ce compte te sera utile pour\naccéder à nos différents services, comme par exemple Polybase."),
		),
		huh.NewGroup(
			huh.NewInput().Title("Quel est ton nom complet ?").Value(&name),
			huh.NewInput().Title("Quel nom d'utilisateur veux-tu utiliser ?").Value(&username),
			huh.NewInput().Title("Quel mot-de passe veux-tu utiliser ?").EchoMode(huh.EchoModePassword).Value(&username),
		),
		huh.NewGroup(
			huh.NewText().Title("Si tu as une clé ssh, tu peux copier ta clé publique ici.\n(laisse le champ vide si tu n'en a pas ou si tu ne sais ce que c'est)\n"),
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
			huh.NewText().Title("Quelque chose à ajouter ?\n").Value(&more),
		),
	)

	return Model{form, false}, []tea.ProgramOption{tea.WithAltScreen()}
}

type Model struct {
	form      *huh.Form // huh.Form is just a tea.Model
	submitted bool
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	return m, cmd
}

func (m Model) View() string {
	if m.form.State == huh.StateCompleted {
		return fmt.Sprintln("Merci ! (message à modifier)")
	}
	return m.form.View()
}
