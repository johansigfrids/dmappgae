package dmapp

import (
	"appengine"
	"appengine/datastore"
	"strconv"
	"strings"
)

func getMonster(c appengine.Context, encodedKey string) (*Monster, error) {
	var monster = new(Monster)
	key, err := datastore.DecodeKey(encodedKey)
	if err != nil {
		return nil, err
	}
	err = datastore.Get(c, key, monster)
	if err != nil {
		return nil, err
	}
	monster.encodedKey = encodedKey
	return monster, nil
}

func getAllMonsters(c appengine.Context) ([]Monster, error) {
	query := datastore.NewQuery("Monster").
		Project("Name", "Level", "Role", "Size", "Origin", "Type", "XP").
		Order("Level")
	var monsters []Monster
	keys, err := query.GetAll(c, &monsters)
	for i, _ := range monsters {
		monsters[i].encodedKey = keys[i].Encode()
	}
	return monsters, err
}

func saveMonster(c appengine.Context, monster *Monster) (string, error) {
	key := datastore.NewIncompleteKey(c, "Monster", nil)
	key2, err := datastore.Put(c, key, monster)
	return key2.Encode(), err
}

func calcAbilityMod(ability int, level int) int {
	return (ability-10)/2 + level/2
}

func makeProp(key string, value interface{}, noindex bool, multiple bool) datastore.Property {
	return datastore.Property{
		Name:     key,
		Value:    value,
		NoIndex:  noindex,
		Multiple: multiple,
	}
}

func (m *Monster) Save(c chan<- datastore.Property) error {
	c <- makeProp("Name", m.Name, false, false)
	c <- makeProp("Level", int64(m.Level), false, false)
	c <- makeProp("Role", m.Role, false, false)
	c <- makeProp("Size", m.Size, false, false)
	c <- makeProp("Origin", m.Origin, false, false)
	c <- makeProp("Type", m.Type, false, false)
	for _, keyword := range m.Keywords {
		c <- makeProp("Keywords", keyword, true, true)
	}
	c <- makeProp("XP", int64(m.XP), false, false)
	c <- makeProp("Health", int64(m.Health), true, false)
	c <- makeProp("InitiativeBonus", int64(m.InitiativeBonus), true, false)
	c <- makeProp("ArmorClass", int64(m.ArmorClass), true, false)
	c <- makeProp("Fortitude", int64(m.Fortitude), true, false)
	c <- makeProp("Reflex", int64(m.Reflex), true, false)
	c <- makeProp("Will", int64(m.Will), true, false)
	for _, sense := range m.Senses {
		c <- makeProp("Senses", sense, true, true)
	}
	c <- makeProp("Speed", m.Speed, true, false)
	for _, immune := range m.Immune {
		c <- makeProp("Immune", immune, true, true)
	}
	for _, resist := range m.Resist {
		c <- makeProp("Resist", resist, true, true)
	}
	for _, vulnerable := range m.Vulnerable {
		c <- makeProp("Vulnerable", vulnerable, true, true)
	}
	c <- makeProp("SavingThrows", int64(m.SavingThrows), true, false)
	c <- makeProp("ActionPoints", int64(m.ActionPoints), true, false)
	numTraits := strconv.Itoa(len(m.Traits))
	for i, t := range m.Traits {
		traitNum := strconv.Itoa(i)
		c <- makeProp("Traits."+numTraits+"."+traitNum+".Name", t.Name, true, false)
		for _, keyword := range t.Keywords {
			c <- makeProp("Traits."+numTraits+"."+traitNum+".Keywords", keyword, true, true)
		}
		c <- makeProp("Traits."+numTraits+"."+traitNum+".Range", t.Range, true, false)
		c <- makeProp("Traits."+numTraits+"."+traitNum+".Effect", t.Name, true, false)
	}
	numSAs := strconv.Itoa(len(m.StandardActions))
	for i, sa := range m.StandardActions {
		saNum := strconv.Itoa(i)
		c <- makeProp("StandardActions."+numSAs+"."+saNum+".Name", sa.Name, true, false)
		for _, keyword := range sa.Keywords {
			c <- makeProp("StandardActions."+numSAs+"."+saNum+".Keywords", keyword, true, true)
		}
		c <- makeProp("StandardActions."+numSAs+"."+saNum+".Usage", sa.Usage, true, false)
		for _, recharge := range sa.Recharge {
			c <- makeProp("StandardActions."+numSAs+"."+saNum+".Recharge", int64(recharge), true, true)
		}
		c <- makeProp("StandardActions."+numSAs+"."+saNum+".Uses", int64(sa.Uses), true, false)
		c <- makeProp("StandardActions."+numSAs+"."+saNum+".UsesPer", sa.UsesPer, true, false)
		numAtts := strconv.Itoa(len(sa.Attacks))
		for j, attack := range sa.Attacks {
			attNum := strconv.Itoa(j)
			c <- makeProp("StandardActions."+numSAs+"."+saNum+".Attacks."+numAtts+"."+attNum+".Range", attack.Range, true, false)
			c <- makeProp("StandardActions."+numSAs+"."+saNum+".Attacks."+numAtts+"."+attNum+".Targets", attack.Targets, true, false)
			c <- makeProp("StandardActions."+numSAs+"."+saNum+".Attacks."+numAtts+"."+attNum+".AttackBonus", int64(attack.AttackBonus), true, false)
			c <- makeProp("StandardActions."+numSAs+"."+saNum+".Attacks."+numAtts+"."+attNum+".Versus", attack.Versus, true, false)
			c <- makeProp("StandardActions."+numSAs+"."+saNum+".Attacks."+numAtts+"."+attNum+".AttackInfo", attack.AttackInfo, true, false)
		}
		numHits := strconv.Itoa(len(sa.Hits))
		for j, hit := range sa.Hits {
			hitNum := strconv.Itoa(j)
			c <- makeProp("StandardActions."+numSAs+"."+saNum+".Hits."+numHits+"."+hitNum+".DieCount", int64(hit.DieCount), true, false)
			c <- makeProp("StandardActions."+numSAs+"."+saNum+".Hits."+numHits+"."+hitNum+".DieSides", int64(hit.DieSides), true, false)
			c <- makeProp("StandardActions."+numSAs+"."+saNum+".Hits."+numHits+"."+hitNum+".DamageBonus", int64(hit.DamageBonus), true, false)
			c <- makeProp("StandardActions."+numSAs+"."+saNum+".Hits."+numHits+"."+hitNum+".HitInfo", hit.HitInfo, true, false)
		}
		c <- makeProp("StandardActions."+numSAs+"."+saNum+".Miss", sa.Miss, true, false)
		c <- makeProp("StandardActions."+numSAs+"."+saNum+".Effect", sa.Effect, true, false)
	}
	numMoves := strconv.Itoa(len(m.MoveActions))
	for i, move := range m.MoveActions {
		moveNum := strconv.Itoa(i)
		c <- makeProp("MoveActions."+numMoves+"."+moveNum+".Name", move.Name, true, false)
		for _, keyword := range move.Keywords {
			c <- makeProp("MoveActions."+numMoves+"."+moveNum+".Keywords", keyword, true, true)
		}
		c <- makeProp("MoveActions."+numMoves+"."+moveNum+".Usage", move.Usage, true, false)
		c <- makeProp("MoveActions."+numMoves+"."+moveNum+".Uses", int64(move.Uses), true, false)
		c <- makeProp("MoveActions."+numMoves+"."+moveNum+".CurrentUses", int64(move.CurrentUses), true, false)
		c <- makeProp("MoveActions."+numMoves+"."+moveNum+".UsesPer", move.UsesPer, true, false)
		c <- makeProp("MoveActions."+numMoves+"."+moveNum+".Requirement", move.Requirement, true, false)
		c <- makeProp("MoveActions."+numMoves+"."+moveNum+".Effect", move.Effect, true, false)
	}
	numMinos := strconv.Itoa(len(m.MinorActions))
	for i, mino := range m.MinorActions {
		minoNum := strconv.Itoa(i)
		c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Name", mino.Name, true, false)
		for _, keyword := range mino.Keywords {
			c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Keywords", keyword, true, true)
		}
		c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Usage", mino.Usage, true, false)
		for _, recharge := range mino.Recharge {
			c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Recharge", int64(recharge), true, true)
		}
		c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Uses", int64(mino.Uses), true, false)
		c <- makeProp("MinorActions."+numMinos+"."+minoNum+".UsesPer", mino.UsesPer, true, false)
		numAtts := strconv.Itoa(len(mino.Attacks))
		for j, attack := range mino.Attacks {
			attNum := strconv.Itoa(j)
			c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Attacks."+numAtts+"."+attNum+".Range", attack.Range, true, false)
			c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Attacks."+numAtts+"."+attNum+".Targets", attack.Targets, true, false)
			c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Attacks."+numAtts+"."+attNum+".AttackBonus", int64(attack.AttackBonus), true, false)
			c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Attacks."+numAtts+"."+attNum+".Versus", attack.Versus, true, false)
			c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Attacks."+numAtts+"."+attNum+".AttackInfo", attack.AttackInfo, true, false)
		}
		numHits := strconv.Itoa(len(mino.Hits))
		for j, hit := range mino.Hits {
			hitNum := strconv.Itoa(j)
			c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Hits."+numHits+"."+hitNum+".DieCount", int64(hit.DieCount), true, false)
			c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Hits."+numHits+"."+hitNum+".DieSides", int64(hit.DieSides), true, false)
			c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Hits."+numHits+"."+hitNum+".DamageBonus", int64(hit.DamageBonus), true, false)
			c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Hits."+numHits+"."+hitNum+".HitInfo", hit.HitInfo, true, false)
		}
		c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Miss", mino.Miss, true, false)
		c <- makeProp("MinorActions."+numMinos+"."+minoNum+".Effect", mino.Effect, true, false)
	}
	numTrigs := strconv.Itoa(len(m.TriggeredActions))
	for i, trig := range m.TriggeredActions {
		trigNum := strconv.Itoa(i)
		c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Name", trig.Name, true, false)
		for _, keyword := range trig.Keywords {
			c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Keywords", keyword, true, true)
		}
		c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Usage", trig.Usage, true, false)
		for _, recharge := range trig.Recharge {
			c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Recharge", int64(recharge), true, true)
		}
		c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Uses", int64(trig.Uses), true, false)
		c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".UsesPer", trig.UsesPer, true, false)
		c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Trigger", trig.Trigger, true, false)
		c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Reaction", trig.Reaction, true, false)
		numAtts := strconv.Itoa(len(trig.Attacks))
		for j, attack := range trig.Attacks {
			attNum := strconv.Itoa(j)
			c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Attacks."+numAtts+"."+attNum+".Range", attack.Range, true, false)
			c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Attacks."+numAtts+"."+attNum+".Targets", attack.Targets, true, false)
			c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Attacks."+numAtts+"."+attNum+".AttackBonus", int64(attack.AttackBonus), true, false)
			c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Attacks."+numAtts+"."+attNum+".Versus", attack.Versus, true, false)
			c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Attacks."+numAtts+"."+attNum+".AttackInfo", attack.AttackInfo, true, false)
		}
		numHits := strconv.Itoa(len(trig.Hits))
		for j, hit := range trig.Hits {
			hitNum := strconv.Itoa(j)
			c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Hits."+numHits+"."+hitNum+".DieCount", int64(hit.DieCount), true, false)
			c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Hits."+numHits+"."+hitNum+".DieSides", int64(hit.DieSides), true, false)
			c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Hits."+numHits+"."+hitNum+".DamageBonus", int64(hit.DamageBonus), true, false)
			c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Hits."+numHits+"."+hitNum+".HitInfo", hit.HitInfo, true, false)
		}
		c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Miss", trig.Miss, true, false)
		c <- makeProp("TriggeredActions."+numTrigs+"."+trigNum+".Effect", trig.Effect, true, false)
	}
	c <- makeProp("Acrobatics", int64(m.Acrobatics), true, false)
	c <- makeProp("Arcana", int64(m.Arcana), true, false)
	c <- makeProp("Athletics", int64(m.Athletics), true, false)
	c <- makeProp("Bluff", int64(m.Bluff), true, false)
	c <- makeProp("Diplomacy", int64(m.Diplomacy), true, false)
	c <- makeProp("Dungeoneering", int64(m.Dungeoneering), true, false)
	c <- makeProp("Endurance", int64(m.Endurance), true, false)
	c <- makeProp("Heal", int64(m.Heal), true, false)
	c <- makeProp("History", int64(m.History), true, false)
	c <- makeProp("Insight", int64(m.Insight), true, false)
	c <- makeProp("Intimidate", int64(m.Intimidate), true, false)
	c <- makeProp("Nature", int64(m.Nature), true, false)
	c <- makeProp("Perception", int64(m.Perception), true, false)
	c <- makeProp("Religion", int64(m.Religion), true, false)
	c <- makeProp("Stealth", int64(m.Stealth), true, false)
	c <- makeProp("Streetwise", int64(m.Streetwise), true, false)
	c <- makeProp("Thievery", int64(m.Thievery), true, false)
	c <- makeProp("Strength", int64(m.Strength), true, false)
	c <- makeProp("Constitution", int64(m.Constitution), true, false)
	c <- makeProp("Dexterity", int64(m.Dexterity), true, false)
	c <- makeProp("Intelligence", int64(m.Intelligence), true, false)
	c <- makeProp("Wisdom", int64(m.Wisdom), true, false)
	c <- makeProp("Charisma", int64(m.Charisma), true, false)
	c <- makeProp("Alignmment", m.Alignment, true, false)
	for _, language := range m.Languages {
		c <- makeProp("Languages", language, true, true)
	}
	for _, equipment := range m.Equipment {
		c <- makeProp("Equipment", equipment, true, true)
	}
	close(c)
	return nil
}

func (m *Monster) Load(c <-chan datastore.Property) error {
	for prop := range c {
		keys := strings.Split(prop.Name, ".")
		switch keys[0] {
		case "Name":
			m.Name = prop.Value.(string)
		case "Level":
			m.Level = int(prop.Value.(int64))
		case "Role":
			m.Role = prop.Value.(string)
		case "Size":
			m.Size = prop.Value.(string)
		case "Origin":
			m.Origin = prop.Value.(string)
		case "Type":
			m.Type = prop.Value.(string)
		case "Keywords":
			if m.Keywords == nil {
				m.Keywords = make([]string, 0, 0)
			}
			m.Keywords = append(m.Keywords, prop.Value.(string))
		case "XP":
			m.XP = int(prop.Value.(int64))
		case "Health":
			m.Health = int(prop.Value.(int64))
		case "InitiativeBonus":
			m.InitiativeBonus = int(prop.Value.(int64))
		case "ArmorClass":
			m.ArmorClass = int(prop.Value.(int64))
		case "Fortitude":
			m.Fortitude = int(prop.Value.(int64))
		case "Reflex":
			m.Reflex = int(prop.Value.(int64))
		case "Will":
			m.Will = int(prop.Value.(int64))
		case "Senses":
			if m.Senses == nil {
				m.Senses = make([]string, 0, 0)
			}
			m.Senses = append(m.Keywords, prop.Value.(string))
		case "Speed":
			m.Speed = prop.Value.(string)
		case "Immune":
			if m.Immune == nil {
				m.Immune = make([]string, 0, 0)
			}
			m.Immune = append(m.Immune, prop.Value.(string))
		case "Resist":
			if m.Resist == nil {
				m.Resist = make([]string, 0, 0)
			}
			m.Resist = append(m.Resist, prop.Value.(string))
		case "Vulnerable":
			if m.Vulnerable == nil {
				m.Vulnerable = make([]string, 0, 0)
			}
			m.Vulnerable = append(m.Vulnerable, prop.Value.(string))
		case "SavingThrows":
			m.SavingThrows = int(prop.Value.(int64))
		case "ActionPoints":
			m.ActionPoints = int(prop.Value.(int64))
		case "Traits":
			numTraits, err := strconv.ParseInt(keys[1], 10, 64)
			traitNum, err := strconv.ParseInt(keys[2], 10, 64)
			if err != nil {
				return err
			}
			if m.Traits == nil {
				m.Traits = make([]Trait, numTraits)
			}
			switch keys[3] {
			case "Name":
				m.Traits[traitNum].Name = prop.Value.(string)
			case "Keywords":
				if m.Traits[traitNum].Keywords == nil {
					m.Traits[traitNum].Keywords = make([]string, 0, 0)
				}
				m.Traits[traitNum].Keywords = append(m.Traits[traitNum].Keywords, prop.Value.(string))
			case "Range":
				m.Traits[traitNum].Range = prop.Value.(string)
			case "Effect":
				m.Traits[traitNum].Effect = prop.Value.(string)
			}
		case "StandardActions":
			numSAs, err := strconv.ParseInt(keys[1], 10, 64)
			saNum, err := strconv.ParseInt(keys[2], 10, 64)
			if err != nil {
				return err
			}
			if m.StandardActions == nil {
				m.StandardActions = make([]StandardAction, numSAs, numSAs)
			}
			switch keys[3] {
			case "Name":
				m.StandardActions[saNum].Name = prop.Value.(string)
			case "Keywords":
				if m.StandardActions[saNum].Keywords == nil {
					m.StandardActions[saNum].Keywords = make([]string, 0, 0)
				}
				m.StandardActions[saNum].Keywords = append(m.StandardActions[saNum].Keywords, prop.Value.(string))
			case "Usage":
				m.StandardActions[saNum].Usage = prop.Value.(string)
			case "Recharge":
				if m.StandardActions[saNum].Recharge == nil {
					m.StandardActions[saNum].Recharge = make([]int, 0, 0)
				}
				m.StandardActions[saNum].Recharge = append(m.StandardActions[saNum].Recharge, int(prop.Value.(int64)))
			case "Uses":
				m.StandardActions[saNum].Uses = int(prop.Value.(int64))
			case "UsesPer":
				m.StandardActions[saNum].UsesPer = prop.Value.(string)
			case "Attacks":
				numAtts, err := strconv.ParseInt(keys[4], 10, 64)
				attNum, err := strconv.ParseInt(keys[5], 10, 64)
				if err != nil {
					return err
				}
				if m.StandardActions[saNum].Attacks == nil {
					m.StandardActions[saNum].Attacks = make([]Attack, numAtts, numAtts)
				}
				switch keys[6] {
				case "Range":
					m.StandardActions[saNum].Attacks[attNum].Range = prop.Value.(string)
				case "Targets":
					m.StandardActions[saNum].Attacks[attNum].Targets = prop.Value.(string)
				case "AttackBonus":
					m.StandardActions[saNum].Attacks[attNum].AttackBonus = int(prop.Value.(int64))
				case "Versus":
					m.StandardActions[saNum].Attacks[attNum].Versus = prop.Value.(string)
				case "AttackInfo":
					m.StandardActions[saNum].Attacks[attNum].AttackInfo = prop.Value.(string)
				}
			case "Hits":
				numHits, err := strconv.ParseInt(keys[4], 10, 64)
				hitNum, err := strconv.ParseInt(keys[5], 10, 64)
				if err != nil {
					return err
				}
				if m.StandardActions[saNum].Hits == nil {
					m.StandardActions[saNum].Hits = make([]Hit, numHits, numHits)
				}
				switch keys[6] {
				case "DieCount":
					m.StandardActions[saNum].Hits[hitNum].DieCount = int(prop.Value.(int64))
				case "DieSides":
					m.StandardActions[saNum].Hits[hitNum].DieSides = int(prop.Value.(int64))
				case "DamageBonus":
					m.StandardActions[saNum].Hits[hitNum].DamageBonus = int(prop.Value.(int64))
				case "HitInfo":
					m.StandardActions[saNum].Hits[hitNum].HitInfo = prop.Value.(string)
				}
			case "Miss":
				m.StandardActions[saNum].Miss = prop.Value.(string)
			case "Effect":
				m.StandardActions[saNum].Effect = prop.Value.(string)
			}
		case "MoveActions":
			numMoves, err := strconv.ParseInt(keys[1], 10, 64)
			moveNum, err := strconv.ParseInt(keys[2], 10, 64)
			if err != nil {
				return err
			}
			if m.MoveActions == nil {
				m.MoveActions = make([]MoveAction, numMoves, numMoves)
			}
			switch keys[3] {
			case "Name":
				m.MoveActions[moveNum].Name = prop.Value.(string)
			case "Keywords":
				if m.MoveActions[moveNum].Keywords == nil {
					m.MoveActions[moveNum].Keywords = make([]string, 0, 0)
				}
				m.MoveActions[moveNum].Keywords = append(m.MoveActions[moveNum].Keywords, prop.Value.(string))
			case "Usage":
				m.MoveActions[moveNum].Usage = prop.Value.(string)
			case "Uses":
				m.MoveActions[moveNum].Uses = int(prop.Value.(int64))
			case "CurrentUses":
				m.MoveActions[moveNum].CurrentUses = int(prop.Value.(int64))
			case "UsesPer":
				m.MoveActions[moveNum].UsesPer = prop.Value.(string)
			case "Requirement":
				m.MoveActions[moveNum].Requirement = prop.Value.(string)
			}
		case "MinorActions":
			numMinors, err := strconv.ParseInt(keys[1], 10, 64)
			minorNum, err := strconv.ParseInt(keys[2], 10, 64)
			if err != nil {
				return err
			}
			if m.MinorActions == nil {
				m.MinorActions = make([]MinorAction, numMinors, numMinors)
			}
			switch keys[3] {
			case "Name":
				m.MinorActions[minorNum].Name = prop.Value.(string)
			case "Keywords":
				if m.MinorActions[minorNum].Keywords == nil {
					m.MinorActions[minorNum].Keywords = make([]string, 0, 0)
				}
				m.MinorActions[minorNum].Keywords = append(m.MinorActions[minorNum].Keywords, prop.Value.(string))
			case "Usage":
				m.MinorActions[minorNum].Usage = prop.Value.(string)
			case "Recharge":
				if m.MinorActions[minorNum].Recharge == nil {
					m.MinorActions[minorNum].Recharge = make([]int, 0, 0)
				}
				m.MinorActions[minorNum].Recharge = append(m.MinorActions[minorNum].Recharge, int(prop.Value.(int64)))
			case "Uses":
				m.MinorActions[minorNum].Uses = int(prop.Value.(int64))
			case "UsesPer":
				m.MinorActions[minorNum].UsesPer = prop.Value.(string)
			case "Attacks":
				numAtts, err := strconv.ParseInt(keys[4], 10, 64)
				attNum, err := strconv.ParseInt(keys[5], 10, 64)
				if err != nil {
					return err
				}
				if m.MinorActions[minorNum].Attacks == nil {
					m.MinorActions[minorNum].Attacks = make([]Attack, numAtts, numAtts)
				}
				switch keys[6] {
				case "Range":
					m.MinorActions[minorNum].Attacks[attNum].Range = prop.Value.(string)
				case "Targets":
					m.MinorActions[minorNum].Attacks[attNum].Targets = prop.Value.(string)
				case "AttackBonus":
					m.MinorActions[minorNum].Attacks[attNum].AttackBonus = int(prop.Value.(int64))
				case "Versus":
					m.MinorActions[minorNum].Attacks[attNum].Versus = prop.Value.(string)
				case "AttackInfo":
					m.MinorActions[minorNum].Attacks[attNum].AttackInfo = prop.Value.(string)
				}
			case "Hits":
				numHits, err := strconv.ParseInt(keys[4], 10, 64)
				hitNum, err := strconv.ParseInt(keys[5], 10, 64)
				if err != nil {
					return err
				}
				if m.MinorActions[minorNum].Hits == nil {
					m.MinorActions[minorNum].Hits = make([]Hit, numHits, numHits)
				}
				switch keys[6] {
				case "DieCount":
					m.MinorActions[minorNum].Hits[hitNum].DieCount = int(prop.Value.(int64))
				case "DieSides":
					m.MinorActions[minorNum].Hits[hitNum].DieSides = int(prop.Value.(int64))
				case "DamageBonus":
					m.MinorActions[minorNum].Hits[hitNum].DamageBonus = int(prop.Value.(int64))
				case "HitInfo":
					m.MinorActions[minorNum].Hits[hitNum].HitInfo = prop.Value.(string)
				}
			case "Miss":
				m.MinorActions[minorNum].Miss = prop.Value.(string)
			case "Effect":
				m.MinorActions[minorNum].Effect = prop.Value.(string)
			}
		case "TriggeredActions":
			numTrigs, err := strconv.ParseInt(keys[1], 10, 64)
			trigNum, err := strconv.ParseInt(keys[2], 10, 64)
			if err != nil {
				return err
			}
			if m.TriggeredActions == nil {
				m.TriggeredActions = make([]TriggeredAction, numTrigs, numTrigs)
			}
			switch keys[3] {
			case "Name":
				m.TriggeredActions[trigNum].Name = prop.Value.(string)
			case "Keywords":
				if m.TriggeredActions[trigNum].Keywords == nil {
					m.TriggeredActions[trigNum].Keywords = make([]string, 0, 0)
				}
				m.TriggeredActions[trigNum].Keywords = append(m.TriggeredActions[trigNum].Keywords, prop.Value.(string))
			case "Usage":
				m.TriggeredActions[trigNum].Usage = prop.Value.(string)
			case "Recharge":
				if m.TriggeredActions[trigNum].Recharge == nil {
					m.TriggeredActions[trigNum].Recharge = make([]int, 0, 0)
				}
				m.TriggeredActions[trigNum].Recharge = append(m.TriggeredActions[trigNum].Recharge, int(prop.Value.(int64)))
			case "Uses":
				m.TriggeredActions[trigNum].Uses = int(prop.Value.(int64))
			case "UsesPer":
				m.TriggeredActions[trigNum].UsesPer = prop.Value.(string)
			case "Trigger":
				m.TriggeredActions[trigNum].Trigger = prop.Value.(string)
			case "Reaction":
				m.TriggeredActions[trigNum].Reaction = prop.Value.(string)
			case "Attacks":
				numAtts, err := strconv.ParseInt(keys[4], 10, 64)
				attNum, err := strconv.ParseInt(keys[5], 10, 64)
				if err != nil {
					return err
				}
				if m.TriggeredActions[trigNum].Attacks == nil {
					m.TriggeredActions[trigNum].Attacks = make([]Attack, numAtts, numAtts)
				}
				switch keys[6] {
				case "Range":
					m.TriggeredActions[trigNum].Attacks[attNum].Range = prop.Value.(string)
				case "Targets":
					m.TriggeredActions[trigNum].Attacks[attNum].Targets = prop.Value.(string)
				case "AttackBonus":
					m.TriggeredActions[trigNum].Attacks[attNum].AttackBonus = int(prop.Value.(int64))
				case "Versus":
					m.TriggeredActions[trigNum].Attacks[attNum].Versus = prop.Value.(string)
				case "AttackInfo":
					m.TriggeredActions[trigNum].Attacks[attNum].AttackInfo = prop.Value.(string)
				}
			case "Hits":
				numHits, err := strconv.ParseInt(keys[4], 10, 64)
				hitNum, err := strconv.ParseInt(keys[5], 10, 64)
				if err != nil {
					return err
				}
				if m.TriggeredActions[trigNum].Hits == nil {
					m.TriggeredActions[trigNum].Hits = make([]Hit, numHits, numHits)
				}
				switch keys[6] {
				case "DieCount":
					m.TriggeredActions[trigNum].Hits[hitNum].DieCount = int(prop.Value.(int64))
				case "DieSides":
					m.TriggeredActions[trigNum].Hits[hitNum].DieSides = int(prop.Value.(int64))
				case "DamageBonus":
					m.TriggeredActions[trigNum].Hits[hitNum].DamageBonus = int(prop.Value.(int64))
				case "HitInfo":
					m.TriggeredActions[trigNum].Hits[hitNum].HitInfo = prop.Value.(string)
				}
			case "Miss":
				m.TriggeredActions[trigNum].Miss = prop.Value.(string)
			case "Effect":
				m.TriggeredActions[trigNum].Effect = prop.Value.(string)
			}
		case "Acrobatics":
			m.Acrobatics = int(prop.Value.(int64))
		case "Arcana":
			m.Arcana = int(prop.Value.(int64))
		case "Athletics":
			m.Athletics = int(prop.Value.(int64))
		case "Bluff":
			m.Bluff = int(prop.Value.(int64))
		case "Diplomacy":
			m.Diplomacy = int(prop.Value.(int64))
		case "Dungeoneering":
			m.Dungeoneering = int(prop.Value.(int64))
		case "Endurance":
			m.Endurance = int(prop.Value.(int64))
		case "Heal":
			m.Heal = int(prop.Value.(int64))
		case "History":
			m.History = int(prop.Value.(int64))
		case "Insight":
			m.Insight = int(prop.Value.(int64))
		case "Intimidate":
			m.Intimidate = int(prop.Value.(int64))
		case "Nature":
			m.Nature = int(prop.Value.(int64))
		case "Perception":
			m.Perception = int(prop.Value.(int64))
		case "Religion":
			m.Religion = int(prop.Value.(int64))
		case "Stealth":
			m.Stealth = int(prop.Value.(int64))
		case "Streetwise":
			m.Streetwise = int(prop.Value.(int64))
		case "Thievery":
			m.Thievery = int(prop.Value.(int64))
		case "Strength":
			m.Strength = int(prop.Value.(int64))
		case "Constitution":
			m.Constitution = int(prop.Value.(int64))
		case "Dexterity":
			m.Dexterity = int(prop.Value.(int64))
		case "Intelligence":
			m.Intelligence = int(prop.Value.(int64))
		case "Wisdom":
			m.Wisdom = int(prop.Value.(int64))
		case "Charisma":
			m.Charisma = int(prop.Value.(int64))
		case "Alignment":
			m.Alignment = prop.Value.(string)
		case "Languages":
			if m.Languages == nil {
				m.Languages = make([]string, 0, 0)
			}
			m.Languages = append(m.Languages, prop.Value.(string))
		case "Equipment":
			if m.Equipment == nil {
				m.Equipment = make([]string, 0, 0)
			}
			m.Equipment = append(m.Equipment, prop.Value.(string))
		}
	}
	m.CurrentHealth = m.Health
	m.BloodiedHealth = m.Health / 2
	m.StrengthMod = calcAbilityMod(m.Strength, m.Level)
	m.ConstitutionMod = calcAbilityMod(m.Constitution, m.Level)
	m.DexterityMod = calcAbilityMod(m.Dexterity, m.Level)
	m.IntelligenceMod = calcAbilityMod(m.Intelligence, m.Level)
	m.WisdomMod = calcAbilityMod(m.Wisdom, m.Level)
	m.CharismaMod = calcAbilityMod(m.Charisma, m.Level)
	for i := range m.StandardActions {
		m.StandardActions[i].CurrentUses = m.StandardActions[i].Uses
	}
	for i := range m.MinorActions {
		m.MinorActions[i].CurrentUses = m.MinorActions[i].Uses
	}
	for i := range m.TriggeredActions {
		m.TriggeredActions[i].CurrentUses = m.TriggeredActions[i].Uses
	}
	return nil
}
