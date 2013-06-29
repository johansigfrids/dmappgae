package dmapp

type Monster struct {
	encodedKey string `datastore:",noindex"`
	Name       string
	Level      int
	Role       string
	Size       string
	Origin     string
	Type       string
	Keywords   []string `datastore:",noindex"`
	XP         int

	Health          int `datastore:",noindex"`
	CurrentHealth   int `datastore: "-"`
	TemporaryHealth int `datastore: "-"`
	BloodiedHealth  int `datastore: "-"`

	InitiativeBonus int `datastore:",noindex"`

	ArmorClass int `datastore:",noindex"`
	Fortitude  int `datastore:",noindex"`
	Reflex     int `datastore:",noindex"`
	Will       int `datastore:",noindex"`

	Senses []string `datastore:",noindex"`
	Speed  string   `datastore:",noindex"`

	Immune     []string `datastore:",noindex"`
	Resist     []string `datastore:",noindex"`
	Vulnerable []string `datastore:",noindex"`

	SavingThrows int `datastore:",noindex"`
	ActionPoints int `datastore:",noindex"`

	Traits           []Trait           `datastore: "-"`
	StandardActions  []StandardAction  `datastore: "-"`
	MoveActions      []MoveAction      `datastore: "-"`
	MinorActions     []MinorAction     `datastore: "-"`
	TriggeredActions []TriggeredAction `datastore: "-"`

	Acrobatics    int `datastore:",noindex"`
	Arcana        int `datastore:",noindex"`
	Athletics     int `datastore:",noindex"`
	Bluff         int `datastore:",noindex"`
	Diplomacy     int `datastore:",noindex"`
	Dungeoneering int `datastore:",noindex"`
	Endurance     int `datastore:",noindex"`
	Heal          int `datastore:",noindex"`
	History       int `datastore:",noindex"`
	Insight       int `datastore:",noindex"`
	Intimidate    int `datastore:",noindex"`
	Nature        int `datastore:",noindex"`
	Perception    int `datastore:",noindex"`
	Religion      int `datastore:",noindex"`
	Stealth       int `datastore:",noindex"`
	Streetwise    int `datastore:",noindex"`
	Thievery      int `datastore:",noindex"`

	Strength        int `datastore:",noindex"`
	StrengthMod     int `datastore: "-"`
	Constitution    int `datastore:",noindex"`
	ConstitutionMod int `datastore: "-"`
	Dexterity       int `datastore:",noindex"`
	DexterityMod    int `datastore: "-"`
	Intelligence    int `datastore:",noindex"`
	IntelligenceMod int `datastore: "-"`
	Wisdom          int `datastore:",noindex"`
	WisdomMod       int `datastore: "-"`
	Charisma        int `datastore:",noindex"`
	CharismaMod     int `datastore: "-"`

	Alignment string   `datastore:",noindex"`
	Languages []string `datastore:",noindex"`
	Equipment []string `datastore:",noindex"`
}

type Trait struct {
	Name     string   `datastore:",noindex"`
	Keywords []string `datastore:",noindex"`
	Range    string   `datastore:",noindex"`
	Effect   string   `datastore:",noindex"`
}

type StandardAction struct {
	Name     string   `datastore:",noindex"`
	Keywords []string `datastore:",noindex"`
	Usage    string   `datastore:",noindex"`
	Recharge []int    `datastore:",noindex"`

	Uses        int    `datastore:",noindex"`
	CurrentUses int    `datastore:"-"`
	UsesPer     string `datastore:",noindex"`

	Attacks []Attack `datastore:",noindex"`
	Hits    []Hit    `datastore:",noindex"`

	Miss string `datastore:",noindex"`

	Effect string `datastore:",noindex"`
}

type MoveAction struct {
	Name     string   `datastore:",noindex"`
	Keywords []string `datastore:",noindex"`
	Usage    string   `datastore:",noindex"`

	Uses        int    `datastore:",noindex"`
	CurrentUses int    `datastore:"-"`
	UsesPer     string `datastore:",noindex"`

	Requirement string `datastore:",noindex"`
	Effect      string `datastore:",noindex"`
}

type MinorAction struct {
	Name     string   `datastore:",noindex"`
	Keywords []string `datastore:",noindex"`
	Usage    string   `datastore:",noindex"`
	Recharge []int    `datastore:",noindex"`

	Uses        int    `datastore:",noindex"`
	CurrentUses int    `datastore:"-"`
	UsesPer     string `datastore:",noindex"`

	Attacks []Attack `datastore:",noindex"`
	Hits    []Hit    `datastore:",noindex"`

	Miss string `datastore:",noindex"`

	Effect string `datastore:",noindex"`
}

type TriggeredAction struct {
	Name     string   `datastore:",noindex"`
	Keywords []string `datastore:",noindex"`
	Usage    string   `datastore:",noindex"`
	Recharge []int    `datastore:",noindex"`

	Uses        int    `datastore:",noindex"`
	CurrentUses int    `datastore:"-"`
	UsesPer     string `datastore:",noindex"`

	Trigger  string `datastore:",noindex"`
	Reaction string `datastore:",noindex"`

	Attacks []Attack `datastore:",noindex"`
	Hits    []Hit    `datastore:",noindex"`

	Miss string `datastore:",noindex"`

	Effect string `datastore:",noindex"`
}

type Attack struct {
	Range       string `datastore:",noindex"`
	Targets     string `datastore:",noindex"`
	AttackBonus int    `datastore:",noindex"`
	Versus      string `datastore:",noindex"`
	AttackInfo  string `datastore:",noindex"`
}

type Hit struct {
	DieCount    int    `datastore:",noindex"`
	DieSides    int    `datastore:",noindex"`
	DamageBonus int    `datastore:",noindex"`
	HitInfo     string `datastore:",noindex"`
}
