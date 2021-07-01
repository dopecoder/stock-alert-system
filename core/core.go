package core

type Config struct {
	UpdatePostEndpoint string
	KiteToken          string
	KiteApiKey         string
	AngelApiKey        string
	AngelClientId      string
	AngelPassword      string
}

type GetTriggerOptions struct {
	Filter int
}

type TriggerSystemRepo interface {
	CreateTrigger(trigger *Trigger) error
	UpdateTrigger(trigger *Trigger) error
	GetTriggers(GetTriggerOptions) ([]Trigger, error)
	OnTrigger(trigger *Trigger) error
	DeleteTrigger(trigger *Trigger) error
}
