package providers

import "log"


type Provider interface {
	Register()
	Boot()
}

type App struct {
	providers []Provider
}

func NewApp() *App {
	return &App{
		providers: []Provider{},
	}
}

func (app *App) Register(provider Provider) {
	app.providers = append(app.providers, provider)
}

func (app *App) Boot() {
	for _, p := range app.providers {
		p.Register()
	}
	for _, p := range app.providers {
		if err := p.Boot(); err != nil {
			log.Println(err)
		}
	}
}