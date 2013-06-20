package dmapp

type Monster struct {
	MonsterID string
	Name      string
	Level     int
	Role      string
	Size      string
	Origin    string
	Type      string
	XP        int

	Health          int
	CurrentHealth   int
	TemporaryHealth int
	BloodiedHealth  int

	InitiativeBonus int

	ArmorClass int
	Fortitude  int
	Reflex     int
	Will       int

	Senses []string
	Speed  string

	Immune     []string
	Resist     []string
	Vulnerable []string

	SavingThrows int
	ActionPoints int

	Traits           []Trait
	StandardActions  []StandardAction
	MoveActions      []MoveAction
	MinorActions     []MinorAction
	TriggeredActions []TriggeredAction

	Acrobatics    int
	Arcana        int
	Athletics     int
	Bluff         int
	Diplomacy     int
	Dungeoneering int
	Endurence     int
	Heal          int
	History       int
	Insight       int
	Intimidate    int
	Nature        int
	Perception    int
	Religion      int
	Stealth       int
	Streetwise    int
	Thievery      int

	Strength        int
	StrengthMod     int
	Constitution    int
	ConstitutionMod int
	Dexterity       int
	DexterityMod    int
	Intelligence    int
	IntelligenceMod int
	Wisdom          int
	WisdomMod       int
	Charisma        int
	CharismaMod     int

	Alignment string
	Languages []string
	Equipment []string
}

type Trait struct {
	Name     string
	Keywords []string
	Range    string
	Effect   string
}

type StandardAction struct {
	Name     string
	Keywords []string
	Usage    string
	Recharge []int

	Uses        int
	CurrentUses int
	UsesPer     string

	Attacks []Attack

	Hits []Hit

	Miss string

	Effect string
}

type MoveAction struct {
	Name     string
	Keywords []string
	Usage    string

	Uses        int
	CurrentUses int
	UsesPer     string

	Requirement string
	Effect      string
}

type MinorAction struct {
	Name     string
	Keywords []string
	Usage    string
	Recharge []int

	Uses        int
	CurrentUses int
	UsesPer     string

	Attacks []Attack

	Hits []Hit

	Miss string

	Effect string
}

type TriggeredAction struct {
	Name     string
	Keywords []string
	Usage    string
	Recharge []int

	Uses        int
	CurrentUses int
	UsesPer     string

	Trigger  string
	Reaction string

	Attacks []Attack

	Hits []Hit

	Miss string

	Effect string
}

type Attack struct {
	Range       string
	Targets     string
	AttackBonus int
	Versus      string
	AttackInfo  string
}

type Hit struct {
	DieCount    int
	DieSides    int
	DamageBonus int
	HitInfo     string
}
