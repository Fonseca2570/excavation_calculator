package hotspots

import "excavation_calculator/materials"

type Hotspot struct {
	Name                 string
	Level                int
	MaterialsWithDecimal []MaterialWithDecimal
	Total                float64
}

type MaterialWithDecimal struct {
	Material materials.Material
	Decimal  float64
}

var AllHotspots = []Hotspot{
	VenatorRemains,
	LegionaryRemains,
	CastraDebris,
	LodgeBarStorage,
	LodgeArtStorage,
	AdministratuDebris,
	CultistFootlocker,
	SacrificialAltar,
	ProdromoiRemains,
	DisDungeonDebris,
	PraesidioRemains,
	MonocerosRemains,
	AmphitheatreDebris,
	CeramicsStudioDebris,
	CarceremDebris,
	StadioDebris,
	InfernalArt,
	ShakrothRemains,
	DominionGamesPodium,
	IkovianMemorial,
	OikosStudioDebris,
	KharidetChapelDebris,
	GladiatorialGoblinRemains,
	KeshikGer,
	AnimalTrophies,
	PontifexRemains,
	CrucibleStandsDebris,
	TailoryDebris,
	GoblinDormDebris,
	OikosFishingHutRemnants,
	WeaponsResearchDebris,
	OrcusAltar,
	DisOverspill,
	BigHighWarGodShrine,
	VaranusaurRemains,
	GravitronResearchDebris,
	AcropolisDebris,
	ArmariumDebris,
	YubiuskAnimalPen,
	KeshikTowerDebris,
	DragonkinReliquary,
	GoblinTraineeRemains,
	ByzrothRemains,
	DestroyedGolem,
	DragonkinCoffin,
	IcyeneWeaponRack,
	CulinarumDebris,
	KyzajChampionsBoudoir,
	AutopsyTable,
	ExperimentWorkbench,
	KeshikWeaponRack,
	HellfireForge,
	WarforgeScrapPile,
	StockpiledArt,
	AughraRemains,
	AncientMagickMunitions,
	MokshaDevice,
	BibliothekeDebris,
	ChthonianTrophies,
	WarforgeWeaponRack,
	FlightResearchDebris,
	AetheriumForge,
	XoloMine,
	PraetorianRemains,
	BandossSanctumDebris,
	TsutsarothRemains,
	OptimatoiRemains,
	WarTableDebris,
	HowlsWorkshopDebris,
	MakeshiftPieOven,
	XoloRemains,
	SaurthenDebris,
}

var VenatorRemains = Hotspot{
	Name: "Venator remains",
	Level: 5,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  5,
		},
		{
			Material: materials.ZarosianInsignia,
			Decimal:  5,
		},
	},
}

var LegionaryRemains = Hotspot{
	Name: "Legionary Remains",
	Level: 12,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  3,
		},
		{
			Material: materials.SamiteSilk,
			Decimal:  3,
		},
		{
			Material: materials.ImperialSteel,
			Decimal:  4,
		},
	},
}

var CastraDebris = Hotspot{
	Name: "Castra debris",
	Level: 17,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.WhiteOak,
			Decimal:  4,
		},
		{
			Material: materials.SamiteSilk,
			Decimal:  2,
		},
		{
			Material: materials.ThirdAgeIron,
			Decimal:  2,
		},
		{
			Material: materials.ZarosianInsignia,
			Decimal:  2,
		},
	},
}

var LodgeBarStorage = Hotspot{
	Name: "Lodge bar storage",
	Level: 20,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Orthenglass,
			Decimal:  3,
		},
		{
			Material: materials.Goldrune,
			Decimal:  4,
		},
		{
			Material: materials.ThirdAgeIron,
			Decimal:  3,
		},
	},
}

var LodgeArtStorage = Hotspot{
	Name: "Lodge art storage",
	Level: 24,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.SamiteSilk,
			Decimal:  2,
		},
		{
			Material: materials.WhiteOak,
			Decimal:  2,
		},
		{
			Material: materials.Vellum,
			Decimal:  2,
		},
		{
			Material: materials.CadmiumRed,
			Decimal:  4,
		},
	},
}

var AdministratuDebris = Hotspot{
	Level: 25,
	Name: "Administratum debris",
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.AncientVis,
			Decimal:  3,
		},
		{
			Material: materials.TyrianPurple,
			Decimal:  3,
		},
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.ZarosianInsignia,
			Decimal:  2,
		},
	},
}

var CultistFootlocker = Hotspot{
	Name: "Cultist footlocker",
	Level: 29,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.LeatherScraps,
			Decimal:  2,
		},
		{
			Material: materials.ChaoticBrimstone,
			Decimal:  3,
		},
		{
			Material: materials.Demonhide,
			Decimal:  3,
		},
		{
			Material: materials.ThirdAgeIron,
			Decimal:  2,
		},
	},
}

var SacrificialAltar = Hotspot{
	Name: "Sacrificial Altar",
	Level: 36,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.CadmiumRed,
			Decimal:  2,
		},
		{
			Material: materials.EyeOfDagon,
			Decimal:  3,
		},
		{
			Material: materials.HellfireMetal,
			Decimal:  3,
		},
	},
}

var ProdromoiRemains = Hotspot{
	Name: "Prodromoi remains",
	Level: 42,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  4,
		},
		{
			Material: materials.Keramos,
			Decimal:  3,
		},
		{
			Material: materials.WhiteMarble,
			Decimal:  3,
		},
	},
}

var DisDungeonDebris = Hotspot{
	Name: "Dis dungeon debris",
	Level: 45,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  3,
		},
		{
			Material: materials.ChaoticBrimstone,
			Decimal:  2,
		},
		{
			Material: materials.HellfireMetal,
			Decimal:  2,
		},
		{
			Material: materials.EyeOfDagon,
			Decimal:  3,
		},
	},
}

var PraesidioRemains = Hotspot{
	Name: "Praesidio Remains",
	Level: 47,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  2,
		},
		{
			Material: materials.ImperialSteel,
			Decimal:  2,
		},
		{
			Material: materials.AncientVis,
			Decimal:  3,
		},
		{
			Material: materials.Goldrune,
			Decimal:  3,
		},
	},
}

var MonocerosRemains = Hotspot{
	Name: "Monoceros remains",
	Level: 48,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.LeatherScraps,
			Decimal:  3,
		},
		{
			Material: materials.CobaltBlue,
			Decimal:  4,
		},
		{
			Material: materials.Keramos,
			Decimal:  3,
		},
	},
}

var AmphitheatreDebris = Hotspot{
	Name: "Amphitheatre debris",
	Level: 51,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.WhiteOak,
			Decimal:  2,
		},
		{
			Material: materials.EverlightSilvthril,
			Decimal:  3,
		},
		{
			Material: materials.StarOfSaradomin,
			Decimal:  3,
		},
	},
}

var CeramicsStudioDebris = Hotspot{
	Name: "Ceramics studio debris",
	Level: 56,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Goldrune,
			Decimal:  5,
		},
		{
			Material: materials.WhiteMarble,
			Decimal:  5,
		},
	},
}

var CarceremDebris = Hotspot{
	Name: "Carcerem debris",
	Level: 58,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.Vellum,
			Decimal:  2,
		},
		{
			Material: materials.BloodOfOrcus,
			Decimal:  3,
		},
		{
			Material: materials.AncientVis,
			Decimal:  3,
		},
	},
}

var StadioDebris = Hotspot{
	Name: "Stadio debris",
	Level: 61,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.SamiteSilk,
			Decimal:  2,
		},
		{
			Material: materials.ThirdAgeIron,
			Decimal:  2,
		},
		{
			Material: materials.StarOfSaradomin,
			Decimal:  3,
		},
		{
			Material: materials.Keramos,
			Decimal:  3,
		},
	},
}

var InfernalArt = Hotspot{
	Name: "Infernal Art",
	Level: 65,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.SamiteSilk,
			Decimal:  2,
		},
		{
			Material: materials.WhiteOak,
			Decimal:  2,
		},
		{
			Material: materials.CadmiumRed,
			Decimal:  3,
		},
		{
			Material: materials.Goldrune,
			Decimal:  3,
		},
	},
}

var ShakrothRemains = Hotspot{
	Name: "Shakroth Remains",
	Level: 68,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  2,
		},
		{
			Material: materials.LeatherScraps,
			Decimal:  2,
		},
		{
			Material: materials.ChaoticBrimstone,
			Decimal:  3,
		},
		{
			Material: materials.HellfireMetal,
			Decimal:  3,
		},
	},
}

var DominionGamesPodium = Hotspot{
	Name: "Dominion Games podium",
	Level: 69,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.Orthenglass,
			Decimal:  2,
		},
		{
			Material: materials.EverlightSilvthril,
			Decimal:  3,
		},
		{
			Material: materials.StarOfSaradomin,
			Decimal:  3,
		},
	},
}

var IkovianMemorial = Hotspot{
	Name: "Ikovian memorial",
	Level: 70,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  3,
		},
		{
			Material: materials.WhiteOak,
			Decimal:  3,
		},
		{
			Material: materials.StormguardSteel,
			Decimal:  2,
		},
		{
			Material: materials.WingsOfWar,
			Decimal:  2,
		},
	},
}

var OikosStudioDebris = Hotspot{
	Name: "Oikos studio debris",
	Level: 72,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Keramos,
			Decimal:  2,
		},
		{
			Material: materials.WhiteMarble,
			Decimal:  2,
		},
		{
			Material: materials.CobaltBlue,
			Decimal:  3,
		},
		{
			Material: materials.EverlightSilvthril,
			Decimal:  3,
		},
	},
}

var KharidetChapelDebris = Hotspot{
	Name: "Kharid-et chapel debris",
	Level: 74,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.WhiteOak,
			Decimal:  2,
		},
		{
			Material: materials.ThirdAgeIron,
			Decimal:  2,
		},
		{
			Material: materials.Goldrune,
			Decimal:  3,
		},
		{
			Material: materials.TyrianPurple,
			Decimal:  3,
		},
	},
}

var GladiatorialGoblinRemains = Hotspot{
	Name: "Gladiatorial goblin remains",
	Level: 76,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.MalachiteGreen,
			Decimal:  2,
		},
		{
			Material: materials.WarforgedBronze,
			Decimal:  2,
		},
		{
			Material: materials.MarkOfTheKyzaj,
			Decimal:  3,
		},
		{
			Material: materials.VulcanisedRubber,
			Decimal:  3,
		},
	},
}

var KeshikGer = Hotspot{
	Name: "Keshik ger",
	Level: 76,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.LeatherScraps,
			Decimal:  2,
		},
		{
			Material: materials.AnimalFurs,
			Decimal:  2,
		},
		{
			Material: materials.ArmadyleanYellow,
			Decimal:  3,
		},
		{
			Material: materials.WhiteOak,
			Decimal:  3,
		},
	},
}

var AnimalTrophies = Hotspot{
	Name: "Animal trophies",
	Level: 81,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Orthenglass,
			Decimal:  2,
		},
		{
			Material: materials.ChaoticBrimstone,
			Decimal:  2,
		},
		{
			Material: materials.AnimalFurs,
			Decimal:  3,
		},
		{
			Material: materials.CadmiumRed,
			Decimal:  3,
		},
	},
}

var PontifexRemains = Hotspot{
	Name: "Pontifex remains",
	Level: 81,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.ImperialSteel,
			Decimal:  2,
		},
		{
			Material: materials.ZarosianInsignia,
			Decimal:  3,
		},
		{
			Material: materials.AncientVis,
			Decimal:  3,
		},
	},
}

var CrucibleStandsDebris = Hotspot{
	Name: "Crucible stands debris",
	Level: 81,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.FossilisedBone,
			Decimal:  4,
		},
		{
			Material: materials.MarkOfTheKyzaj,
			Decimal:  3,
		},
		{
			Material: materials.WarforgedBronze,
			Decimal:  3,
		},
	},
}

var TailoryDebris = Hotspot{
	Name: "Tailory debris",
	Level: 81,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.AnimalFurs,
			Decimal:  2,
		},
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.ArmadyleanYellow,
			Decimal:  3,
		},
		{
			Material: materials.SamiteSilk,
			Decimal:  3,
		},
	},
}

var GoblinDormDebris = Hotspot{
	Name: "Goblin dorm debris",
	Level: 83,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.WhiteOak,
			Decimal:  2,
		},
		{
			Material: materials.VulcanisedRubber,
			Decimal:  2,
		},
		{
			Material: materials.MalachiteGreen,
			Decimal:  3,
		},
		{
			Material: materials.YubiuskClay,
			Decimal:  3,
		},
	},
}

var OikosFishingHutRemnants = Hotspot{
	Name: "Goblin dorm debris",
	Level: 84,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.EverlightSilvthril,
			Decimal:  2,
		},
		{
			Material: materials.StarOfSaradomin,
			Decimal:  2,
		},
		{
			Material: materials.ThirdAgeIron,
			Decimal:  3,
		},
		{
			Material: materials.Goldrune,
			Decimal:  3,
		},
	},
}

var WeaponsResearchDebris = Hotspot{
	Name: "Weapons research debris",
	Level: 85,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Orthenglass,
			Decimal:  2,
		},
		{
			Material: materials.WingsOfWar,
			Decimal:  2,
		},
		{
			Material: materials.StormguardSteel,
			Decimal:  3,
		},
		{
			Material: materials.AetheriumAlloy,
			Decimal:  3,
		},
	},
}

var OrcusAltar = Hotspot{
	Name: "Orcus Altar",
	Level: 86,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ImperialSteel,
			Decimal:  2,
		},
		{
			Material: materials.AncientVis,
			Decimal:  2,
		},
		{
			Material: materials.Vellum,
			Decimal:  3,
		},
		{
			Material: materials.BloodOfOrcus,
			Decimal:  3,
		},
	},
}

var DisOverspill = Hotspot{
	Name: "Dis overspill",
	Level: 89,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  2,
		},
		{
			Material: materials.CadmiumRed,
			Decimal:  2,
		},
		{
			Material: materials.EyeOfDagon,
			Decimal:  3,
		},
		{
			Material: materials.HellfireMetal,
			Decimal:  3,
		},
	},
}

var BigHighWarGodShrine = Hotspot{
	Name: "Big High War God shrine",
	Level: 89,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.SamiteSilk,
			Decimal:  2,
		},
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.MalachiteGreen,
			Decimal:  3,
		},
		{
			Material: materials.MarkOfTheKyzaj,
			Decimal:  3,
		},
	},
}

var VaranusaurRemains = Hotspot{
	Name: "Varanusaur remains",
	Level: 90,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.CompassRose,
			Decimal:  3,
		},
		{
			Material: materials.Felt,
			Decimal:  3,
		},
		{
			Material: materials.Goldrune,
			Decimal:  4,
		},
	},
}

var GravitronResearchDebris = Hotspot{
	Name: "Gravitron research debris",
	Level: 91,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Orthenglass,
			Decimal:  2,
		},
		{
			Material: materials.LeatherScraps,
			Decimal:  2,
		},
		{
			Material: materials.StormguardSteel,
			Decimal:  3,
		},
		{
			Material: materials.Quintessence,
			Decimal:  3,
		},
	},
}

var AcropolisDebris = Hotspot{
	Name: "Acropolis debris",
	Level: 92,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.EverlightSilvthril,
			Decimal:  2,
		},
		{
			Material: materials.Keramos,
			Decimal:  3,
		},
		{
			Material: materials.WhiteMarble,
			Decimal:  3,
		},
	},
}

var ArmariumDebris = Hotspot{
	Name: "Armarium debris",
	Level: 93,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  4,
		},
		{
			Material: materials.ZarosianInsignia,
			Decimal:  4,
		},
		{
			Material: materials.ImperialSteel,
			Decimal:  2,
		},
	},
}

var YubiuskAnimalPen = Hotspot{
	Name: "Yu'biusk animal pen",
	Level: 94,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.FossilisedBone,
			Decimal:  2,
		},
		{
			Material: materials.WarforgedBronze,
			Decimal:  2,
		},
		{
			Material: materials.VulcanisedRubber,
			Decimal:  3,
		},
		{
			Material: materials.YubiuskClay,
			Decimal:  3,
		},
	},
}

var KeshikTowerDebris = Hotspot{
	Name: "Keshik tower debris",
	Level: 95,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.WhiteOak,
			Decimal:  2,
		},
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.StormguardSteel,
			Decimal:  3,
		},
		{
			Material: materials.WingsOfWar,
			Decimal:  3,
		},
	},
}

var DragonkinReliquary = Hotspot{
	Name: "Dragonkin reliquary",
	Level: 96,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.DragonMetal,
			Decimal:  5,
		},
		{
			Material: materials.CompassRose,
			Decimal:  2.5,
		},
		{
			Material: materials.Goldrune,
			Decimal:  2.5,
		},
	},
}

var GoblinTraineeRemains = Hotspot{
	Name: "Goblin trainee remains",
	Level: 97,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  2,
		},
		{
			Material: materials.LeatherScraps,
			Decimal:  2,
		},
		{
			Material: materials.FossilisedBone,
			Decimal:  3,
		},
		{
			Material: materials.WarforgedBronze,
			Decimal:  3,
		},
	},
}

var ByzrothRemains = Hotspot{
	Name: "Byzroth remains",
	Level: 98,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.WhiteOak,
			Decimal:  2,
		},
		{
			Material: materials.Orthenglass,
			Decimal:  2,
		},
		{
			Material: materials.LeatherScraps,
			Decimal:  3,
		},
		{
			Material: materials.HellfireMetal,
			Decimal:  3,
		},
	},
}

var DestroyedGolem = Hotspot{
	Name: "Destroyed Golem",
	Level: 98,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Vellum,
			Decimal:  2,
		},
		{
			Material: materials.Soapstone,
			Decimal:  2,
		},
		{
			Material: materials.AetheriumAlloy,
			Decimal:  3,
		},
		{
			Material: materials.Quintessence,
			Decimal:  3,
		},
	},
}

var DragonkinCoffin = Hotspot{
	Name: "Dragonkin coffin",
	Level: 99,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Soapstone,
			Decimal:  1,
		},
		{
			Material: materials.CarbonBlack,
			Decimal:  2,
		},
		{
			Material: materials.CompassRose,
			Decimal:  2,
		},
		{
			Material: materials.Orgone,
			Decimal:  5,
		},
	},
}

var IcyeneWeaponRack = Hotspot{
	Name: "Icyene weapon rack",
	Level: 100,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.LeatherScraps,
			Decimal:  5,
		},
		{
			Material: materials.EverlightSilvthril,
			Decimal:  5,
		},
	},
}

var CulinarumDebris = Hotspot{
	Name: "Culinarum debris",
	Level: 100,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Vellum,
			Decimal:  2,
		},
		{
			Material: materials.BloodOfOrcus,
			Decimal:  2,
		},
		{
			Material: materials.AncientVis,
			Decimal:  3,
		},
		{
			Material: materials.ImperialSteel,
			Decimal:  3,
		},
	},
}

var KyzajChampionsBoudoir = Hotspot{
	Name: "Kyzaj champion's boudoir",
	Level: 100,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.WhiteOak,
			Decimal:  3,
		},
		{
			Material: materials.YubiuskClay,
			Decimal:  3,
		},
		{
			Material: materials.WarforgedBronze,
			Decimal:  4,
		},
	},
}

var AutopsyTable = Hotspot{
	Name: "Autopsy table",
	Level: 101,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Felt,
			Decimal:  5,
		},
		{
			Material: materials.DragonMetal,
			Decimal:  2.5,
		},
		{
			Material: materials.Orthenglass,
			Decimal:  2.5,
		},
	},
}

var ExperimentWorkbench = Hotspot{
	Name: "Experiment workbench",
	Level: 102,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.DragonMetal,
			Decimal:  5,
		},
		{
			Material: materials.Orgone,
			Decimal:  2.5,
		},
		{
			Material: materials.Orthenglass,
			Decimal:  2.5,
		},
	},
}

var KeshikWeaponRack = Hotspot{
	Name: "Keshik weapon rack",
	Level: 103,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.StormguardSteel,
			Decimal:  2,
		},
		{
			Material: materials.WingsOfWar,
			Decimal:  2,
		},
		{
			Material: materials.AetheriumAlloy,
			Decimal:  3,
		},
		{
			Material: materials.Quintessence,
			Decimal:  3,
		},
	},
}

var HellfireForge = Hotspot{
	Name: "Hellfire forge",
	Level: 104,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.CadmiumRed,
			Decimal:  2,
		},
		{
			Material: materials.ThirdAgeIron,
			Decimal:  2,
		},
		{
			Material: materials.EyeOfDagon,
			Decimal:  3,
		},
		{
			Material: materials.HellfireMetal,
			Decimal:  3,
		},
	},
}

var WarforgeScrapPile = Hotspot{
	Name: "Warforge scrap pile",
	Level: 104,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  5,
		},
		{
			Material: materials.WarforgedBronze,
			Decimal:  5,
		},
	},
}

var StockpiledArt = Hotspot{
	Name: "Stockpiled art",
	Level: 105,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.SamiteSilk,
			Decimal:  2,
		},
		{
			Material: materials.WhiteOak,
			Decimal:  2,
		},
		{
			Material: materials.Vellum,
			Decimal:  2,
		},
		{
			Material: materials.CobaltBlue,
			Decimal:  4,
		},
	},
}

var AughraRemains = Hotspot{
	Name: "Aughra remains",
	Level: 106,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Orgone,
			Decimal:  5,
		},
		{
			Material: materials.DragonMetal,
			Decimal:  2,
		},
		{
			Material: materials.CompassRose,
			Decimal:  2,
		},
		{
			Material: materials.CarbonBlack,
			Decimal:  1,
		},
	},
}

var AncientMagickMunitions = Hotspot{
	Name: "Ancient magick munitions",
	Level: 107,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Vellum,
			Decimal:  2,
		},
		{
			Material: materials.ImperialSteel,
			Decimal:  2,
		},
		{
			Material: materials.AncientVis,
			Decimal:  3,
		},
		{
			Material: materials.BloodOfOrcus,
			Decimal:  3,
		},
	},
}

var MokshaDevice = Hotspot{
	Name: "Moksha device",
	Level: 108,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Orgone,
			Decimal:  3,
		},
		{
			Material: materials.CarbonBlack,
			Decimal:  3,
		},
		{
			Material: materials.CompassRose,
			Decimal:  4,
		},
	},
}

var BibliothekeDebris = Hotspot{
	Name: "Bibliotheke debris",
	Level: 109,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.StarOfSaradomin,
			Decimal:  2,
		},
		{
			Material: materials.Vellum,
			Decimal:  3,
		},
		{
			Material: materials.WhiteMarble,
			Decimal:  3,
		},
	},
}

var ChthonianTrophies = Hotspot{
	Name: "Chthonian trophies",
	Level: 110,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ChaoticBrimstone,
			Decimal:  2,
		},
		{
			Material: materials.Orthenglass,
			Decimal:  2,
		},
		{
			Material: materials.WhiteOak,
			Decimal:  3,
		},
		{
			Material: materials.Demonhide,
			Decimal:  3,
		},
	},
}

var WarforgeWeaponRack = Hotspot{
	Name: "Warforge weapon rack",
	Level: 110,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.FossilisedBone,
			Decimal:  3,
		},
		{
			Material: materials.MalachiteGreen,
			Decimal:  3,
		},
		{
			Material: materials.WarforgedBronze,
			Decimal:  4,
		},
	},
}

var FlightResearchDebris = Hotspot{
	Name: "FlightResearchDebris",
	Level: 111,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.LeatherScraps,
			Decimal:  2,
		},
		{
			Material: materials.Orthenglass,
			Decimal:  2,
		},
		{
			Material: materials.SamiteSilk,
			Decimal:  3,
		},
		{
			Material: materials.ArmadyleanYellow,
			Decimal:  3,
		},
	},
}

var AetheriumForge = Hotspot{
	Name: "Aetherium forge",
	Level: 112,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.Quintessence,
			Decimal:  2,
		},
		{
			Material: materials.WingsOfWar,
			Decimal:  3,
		},
		{
			Material: materials.AetheriumAlloy,
			Decimal:  3,
		},
	},
}

var XoloMine = Hotspot{
	Name: "Xolo mine",
	Level: 113,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Goldrune,
			Decimal:  3,
		},
		{
			Material: materials.Orgone,
			Decimal:  3,
		},
		{
			Material: materials.DragonMetal,
			Decimal:  4,
		},
	},
}

var PraetorianRemains = Hotspot{
	Name: "Praetorian remains",
	Level: 114,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.SamiteSilk,
			Decimal:  2,
		},
		{
			Material: materials.ImperialSteel,
			Decimal:  2,
		},
		{
			Material: materials.ZarosianInsignia,
			Decimal:  3,
		},
		{
			Material: materials.AncientVis,
			Decimal:  3,
		},
	},
}

var BandossSanctumDebris = Hotspot{
	Name: "Bandos's sanctum debris",
	Level: 115,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.AnimalFurs,
			Decimal:  2,
		},
		{
			Material: materials.YubiuskClay,
			Decimal:  2,
		},
		{
			Material: materials.FossilisedBone,
			Decimal:  3,
		},
		{
			Material: materials.VulcanisedRubber,
			Decimal:  3,
		},
	},
}

var TsutsarothRemains = Hotspot{
	Name: "Tsutsaroth remains",
	Level: 116,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.ThirdAgeIron,
			Decimal:  2,
		},
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.HellfireMetal,
			Decimal:  3,
		},
		{
			Material: materials.EyeOfDagon,
			Decimal:  3,
		},
	},
}

var OptimatoiRemains = Hotspot{
	Name: "Optimatoi remains",
	Level: 117,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.SamiteSilk,
			Decimal:  3,
		},
		{
			Material: materials.WhiteOak,
			Decimal:  3,
		},
		{
			Material: materials.EverlightSilvthril,
			Decimal:  4,
		},
	},
}

var WarTableDebris = Hotspot{
	Name: "War table debris",
	Level: 118,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Vellum,
			Decimal:  2,
		},
		{
			Material: materials.AncientVis,
			Decimal:  2,
		},
		{
			Material: materials.WhiteOak,
			Decimal:  3,
		},
		{
			Material: materials.TyrianPurple,
			Decimal:  3,
		},
	},
}

var HowlsWorkshopDebris = Hotspot{
	Name: "Howl's workshop debris",
	Level: 118,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Soapstone,
			Decimal:  2,
		},
		{
			Material: materials.ArmadyleanYellow,
			Decimal:  2,
		},
		{
			Material: materials.AetheriumAlloy,
			Decimal:  3,
		},
		{
			Material: materials.Quintessence,
			Decimal:  3,
		},
	},
}

var MakeshiftPieOven = Hotspot{
	Name: "Makeshift pie oven",
	Level: 119,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Soapstone,
			Decimal:  3,
		},
		{
			Material: materials.MalachiteGreen,
			Decimal:  3,
		},
		{
			Material: materials.YubiuskClay,
			Decimal:  4,
		},
	},
}

var XoloRemains = Hotspot{
	Name: "Xolo remains",
	Level: 119,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Felt,
			Decimal:  2,
		},
		{
			Material: materials.Goldrune,
			Decimal:  2,
		},
		{
			Material: materials.DragonMetal,
			Decimal:  2,
		},
		{
			Material: materials.Orgone,
			Decimal:  4,
		},
	},
}

var SaurthenDebris = Hotspot{
	Name: "Saurthen debris",
	Level: 120,
	MaterialsWithDecimal: []MaterialWithDecimal{
		{
			Material: materials.Orgone,
			Decimal:  2,
		},
		{
			Material: materials.DragonMetal,
			Decimal:  4,
		},
		{
			Material: materials.Goldrune,
			Decimal:  4,
		},
	},
}
