package materials

type Material struct {
	Id    int
	Name  string
	Url   string
	Value int
}

var AllMaterials = []Material{
	ThirdAgeIron,
	SamiteSilk,
	WhiteOak,
	Goldrune,
	Orthenglass,
	Vellum,
	LeatherScraps,
	AnimalFurs,
	FossilisedBone,
	Soapstone,
	StormguardSteel,
	WingsOfWar,
	ArmadyleanYellow,
	AetheriumAlloy,
	Quintessence,
	MalachiteGreen,
	YubiuskClay,
	MarkOfTheKyzaj,
	VulcanisedRubber,
	WarforgedBronze,
	Felt,
	DragonMetal,
	Orgone,
	CarbonBlack,
	CompassRose,
	Keramos,
	WhiteMarble,
	CobaltBlue,
	EverlightSilvthril,
	StarOfSaradomin,
	CadmiumRed,
	ChaoticBrimstone,
	Demonhide,
	HellfireMetal,
	EyeOfDagon,
	ZarosianInsignia,
	ImperialSteel,
	TyrianPurple,
	AncientVis,
	BloodOfOrcus,
}

// Normal materials

var ThirdAgeIron = Material{
	Id:   1,
	Name: "Third Age Iron",
	Url:  "https://runescape.wiki/w/Third_Age_iron",
}

var SamiteSilk = Material{
	Id:   2,
	Name: "Samite Silk",
	Url:  "https://runescape.wiki/w/Samite_silk",
}

var WhiteOak = Material{
	Id:   3,
	Name: "White Oak",
	Url:  "https://runescape.wiki/w/White_oak",
}

var Goldrune = Material{
	Id:   4,
	Name: "Goldrune",
	Url:  "https://runescape.wiki/w/Goldrune",
}

var Orthenglass = Material{
	Id:   5,
	Name: "Orthenglass",
	Url:  "https://runescape.wiki/w/Orthenglass",
}

var Vellum = Material{
	Id:   6,
	Name: "Vellum",
	Url:  "https://runescape.wiki/w/Vellum",
}

var LeatherScraps = Material{
	Id:   7,
	Name: "Leather scraps",
	Url:  "https://runescape.wiki/w/Leather_scraps",
}

var AnimalFurs = Material{
	Id:   8,
	Name: "Animal Furs",
	Url:  "https://runescape.wiki/w/Animal_furs",
}

var FossilisedBone = Material{
	Id:   9,
	Name: "Fossilised Bone",
	Url:  "https://runescape.wiki/w/Fossilised_bone",
}

var Soapstone = Material{
	Id:   10,
	Name: "Soapstone",
	Url:  "https://runescape.wiki/w/Soapstone",
}

// Armadyl

var StormguardSteel = Material{
	Id:   11,
	Name: "Stormguard Steel",
	Url:  "https://runescape.wiki/w/Stormguard_steel",
}

var WingsOfWar = Material{
	Id:   12,
	Name: "Wings of War",
	Url:  "https://runescape.wiki/w/Wings_of_War",
}

var ArmadyleanYellow = Material{
	Id:   13,
	Name: "Armadylean yellow",
	Url:  "https://runescape.wiki/w/Armadylean_yellow",
}

var AetheriumAlloy = Material{
	Id:   14,
	Name: "Aetherium Alloy",
	Url:  "https://runescape.wiki/w/Aetherium_alloy",
}

var Quintessence = Material{
	Id:   15,
	Name: "Quintessence",
	Url:  "https://runescape.wiki/w/Quintessence",
}

//Bandos

var MalachiteGreen = Material{
	Id:   16,
	Name: "Malachite Green",
	Url:  "https://runescape.wiki/w/Malachite_green",
}

var MarkOfTheKyzaj = Material{
	Id:   17,
	Name: "Mark of the Kyzaj",
	Url:  "https://runescape.wiki/w/Mark_of_the_Kyzaj",
}

var VulcanisedRubber = Material{
	Id:   18,
	Name: "Vulcanised Rubber",
	Url:  "https://runescape.wiki/w/Vulcanised_rubber",
}

var WarforgedBronze = Material{
	Id:   19,
	Name: "Warforged Bronze",
	Url:  "https://runescape.wiki/w/Warforged_bronze",
}

var YubiuskClay = Material{
	Id:   20,
	Name: "Yu'biusk clay",
	Url:  "https://runescape.wiki/w/Yu%27biusk_clay",
}

// DragonKin

var Felt = Material{
	Id:   21,
	Name: "Felt",
	Url:  "https://runescape.wiki/w/Felt",
}

var DragonMetal = Material{
	Id:   22,
	Name: "Dragon Metal",
	Url:  "https://runescape.wiki/w/Dragon_metal",
}

var Orgone = Material{
	Id:   23,
	Name: "Orgone",
	Url:  "https://runescape.wiki/w/Orgone",
}

var CarbonBlack = Material{
	Id:   24,
	Name: "Carbon Black",
	Url:  "https://runescape.wiki/w/Carbon_black",
}

var CompassRose = Material{
	Id:   25,
	Name: "Compass Rose",
	Url:  "https://runescape.wiki/w/Compass_rose",
}

// Saradomin

var Keramos = Material{
	Id:   26,
	Name: "Keramos",
	Url:  "https://runescape.wiki/w/Keramos",
}

var WhiteMarble = Material{
	Id:   27,
	Name: "White Marble",
	Url:  "https://runescape.wiki/w/White_marble",
}

var CobaltBlue = Material{
	Id:   28,
	Name: "Cobalt Blue",
	Url:  "https://runescape.wiki/w/Cobalt_blue",
}

var EverlightSilvthril = Material{
	Id:   29,
	Name: "Everlight silvthril",
	Url:  "https://runescape.wiki/w/Everlight_silvthril",
}

var StarOfSaradomin = Material{
	Id:   30,
	Name: "Star of Saradomin",
	Url:  "https://runescape.wiki/w/Star_of_Saradomin",
}

// Zamorak

var CadmiumRed = Material{
	Id:   31,
	Name: "Cadmium Red",
	Url:  "https://runescape.wiki/w/Cadmium_red",
}

var ChaoticBrimstone = Material{
	Id:   32,
	Name: "Chaotic brimstone",
	Url:  "https://runescape.wiki/w/Chaotic_brimstone",
}

var Demonhide = Material{
	Id:   33,
	Name: "Demonhide",
	Url:  "https://runescape.wiki/w/Demonhide",
}

var EyeOfDagon = Material{
	Id:   34,
	Name: "Eye of Dagon",
	Url:  "https://runescape.wiki/w/Eye_of_Dagon",
}

var HellfireMetal = Material{
	Id:   35,
	Name: "Hellfire Metal",
	Url:  "https://runescape.wiki/w/Hellfire_metal",
}

// Zarosian

var ZarosianInsignia = Material{
	Id:   36,
	Name: "Zarosian Insignia",
	Url:  "https://runescape.wiki/w/Zarosian_insignia",
}

var ImperialSteel = Material{
	Id:   37,
	Name: "Imperial Steel",
	Url:  "https://runescape.wiki/w/Imperial_steel",
}

var AncientVis = Material{
	Id:   38,
	Name: "Ancient Vis",
	Url:  "https://runescape.wiki/w/Ancient_vis",
}

var TyrianPurple = Material{
	Id:   39,
	Name: "Tyrian Purple",
	Url:  "https://runescape.wiki/w/Tyrian_purple",
}

var BloodOfOrcus = Material{
	Id:   40,
	Name: "Blood of Orcus",
	Url:  "https://runescape.wiki/w/Blood_of_Orcus",
}
