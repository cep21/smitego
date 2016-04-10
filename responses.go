package smitego

import (
	"fmt"
)

type createSessionResp struct {
	RetMsg    string `json:"ret_msg"`
	SessionID string `json:"session_id"`
	Timestamp string `json:"timestamp"`
}

// DataUsed is the JSON object returned by GetDataUsed
type DataUsed struct {
	ActiveSessions     int    `json:"Active_Session"`
	ConcurrentSessions int    `json:"Concurrent_Sessions"`
	RequestLimitDaily  int    `json:"Request_Limit_Daily"`
	SessionCap         int    `json:"Session_Cap"`
	SessionTimeLimit   int    `json:"Session_Time_Limit"`
	TotalRequestsToday int    `json:"Total_Requests_Today"`
	TotalSessionsToday int    `json:"Total_Sessions_Today"`
	RetMsg             string `json:"ret_msg"`
}

// String returns all the DataUsed info
func (d *DataUsed) String() string {
	return fmt.Sprintf("Active: %d Concurrent: %d Limit: %d Session Cap: %d Time Limit: %d Today's requests: %d Today's session %d RetMsg %s", d.ActiveSessions, d.ConcurrentSessions, d.RequestLimitDaily, d.SessionCap, d.SessionTimeLimit, d.TotalRequestsToday, d.TotalSessionsToday, d.RetMsg)
}

// God is a smite God
type God struct {
	Ability1                   string             `json:"Ability1"`
	Ability2                   string             `json:"Ability2"`
	Ability3                   string             `json:"Ability3"`
	Ability4                   string             `json:"Ability4"`
	Ability5                   string             `json:"Ability5"`
	AbilityID1                 int                `json:"AbilityId1"`
	AbilityID2                 int                `json:"AbilityId2"`
	AbilityID3                 int                `json:"AbilityId3"`
	AbilityID4                 int                `json:"AbilityId4"`
	AbilityID5                 int                `json:"AbilityId5"`
	AbilityInfo1               Ability            `json:"Ability_1"`
	AbilityInfo2               Ability            `json:"Ability_2"`
	AbilityInfo3               Ability            `json:"Ability_3"`
	AbilityInfo4               Ability            `json:"Ability_4"`
	AbilityInfo5               Ability            `json:"Ability_5"`
	AttackSpeed                float64            `json:"AttackSpeed"`
	AttackSpeedPerLevel        float64            `json:"AttackSpeedPerLevel"`
	Cons                       string             `json:"Cons"`
	HP5PerLevel                float64            `json:"HP5PerLevel"`
	Health                     int                `json:"Health"`
	HealthPerFive              int                `json:"HealthPerFive"`
	HealthPerLevel             int                `json:"HealthPerLevel"`
	Lore                       string             `json:"Lore"`
	MP5PerLevel                float64            `json:"MP5PerLevel"`
	MagicProtection            int                `json:"MagicProtection"`
	MagicProtectionPerLevel    float64            `json:"MagicProtectionPerLevel"`
	MagicalPower               int                `json:"MagicalPower"`
	MagicalPowerPerLevel       float64            `json:"MagicalPowerPerLevel"`
	Mana                       int                `json:"Mana"`
	ManaPerFive                float64            `json:"ManaPerFive"`
	ManaPerLevel               float64            `json:"ManaPerLevel"`
	Name                       string             `json:"Name"`
	OnFreeRotation             string             `json:"OnFreeRotation"`
	Pantheon                   string             `json:"Pantheon"`
	PhysicalPower              int                `json:"PhysicalPower"`
	PhysicalPowerPerLevel      float64            `json:"PhysicalPowerPerLevel"`
	PhysicalProtection         int                `json:"PhysicalProtection"`
	PhysicalProtectionPerLevel float64            `json:"PhysicalProtectionPerLevel"`
	Pros                       string             `json:"Pros"`
	Roles                      string             `json:"Roles"`
	Speed                      int                `json:"Speed"`
	Title                      string             `json:"Title"`
	Type                       string             `json:"Type"`
	AbilityDescription1        AbilityDescription `json:"abilityDescription1"`
	AbilityDescription2        AbilityDescription `json:"abilityDescription2"`
	AbilityDescription3        AbilityDescription `json:"abilityDescription3"`
	AbilityDescription4        AbilityDescription `json:"abilityDescription4"`
	AbilityDescription5        AbilityDescription `json:"abilityDescription5"`
	BasicAttack                AbilityDescription `json:"basicAttack"`
	GodAbilityURL1             string             `json:"godAbility1_URL"`
	GodAbilityURL2             string             `json:"godAbility1_URL"`
	GodAbilityURL3             string             `json:"godAbility1_URL"`
	GodAbilityURL4             string             `json:"godAbility1_URL"`
	GodAbilityURL5             string             `json:"godAbility1_URL"`
	GodCardURL                 string             `json:"godCard_URL"`
	GodIconURL                 string             `json:"godIcon_URL"`
	ID                         int                `json:"id"`
	LatestGod                  string             `json:"latestGod"`
	RetMsg                     string             `json:"ret_msg"`
}

func (g *God) String() string {
	return g.Name
}

// Match describes a ranked match
type Match struct {
	AwayTeamClanID  int    `json:"away_team_clan_id"`
	AwayTeamName    string `json:"away_team_name"`
	AwayTeamTagName string `json:"away_team_tagname"`

	HomeTeamClanID  int    `json:"home_team_clan_id"`
	HomeTeamName    string `json:"home_team_name"`
	HomeTeamTagName string `json:"home_team_tagname"`

	MapInstanceID string `json:"map_instance_id"`
	MatchDate     string `json:"match_date"`
	MatchNumber   string `json:"match_number"`
	MatchStatus   string `json:"match_status"`

	MatchupID string `json:"matchup_id"`

	Region            string `json:"region"`
	RetMsg            string `json:"ret_msg"`
	TournamentName    string `json:"tournament_name"`
	WinningTeamClanID int    `json:"winning_team_clan_id"`
}

func (m *Match) String() string {
	return fmt.Sprintf("%s vs %s", m.HomeTeamName, m.AwayTeamTagName)
}

// Ability is a god ability
type Ability struct {
	Description AbilityDescription `json:"Description"`
	ID          int                `json:"Id"`
	Summary     string             `json:"summary"`
	URL         string             `json:"URL"`
}

// AbilityDescription is ... honestly a useless abstraction for them
type AbilityDescription struct {
	ItemDescription ItemDescription `json:"itemDescription"`
}

// ItemDescription gives specifics about ability or item info
type ItemDescription struct {
	Cooldown             string     `json:"cooldown"`
	Cost                 string     `json:"cost"`
	Description          string     `json:"description"`
	MenuItems            []MenuItem `json:"menuitems"`
	RankItems            []MenuItem `json:"rankitems"`
	SecondaryDescription string     `json:"secondaryDescription"`
}

// MenuItem describes a strength of an ability
type MenuItem struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

// OldMatchDetails is no longer needed, but returns a limited set of match info
type OldMatchDetails struct {
	Ban1               string `json:"Ban1"`
	Ban2               string `json:"Ban2"`
	EntryDatetime      string `json:"Entry_Datetime"`
	Match              int    `json:"Match"`
	MatchTime          int    `json:"Match_Time"`
	OfflineSpectators  int    `json:"Offline_Spectators"`
	RealtimeSpectators int    `json:"Realtime_Spectators"`
	RecordingEnded     string `json:"Recording_Ended"`
	RecordingStarted   string `json:"Recording_Started"`
	Team1AvgLevel      int    `json:"Team1_AvgLevel"`
	Team1Gold          int    `json:"Team1_Gold"`
	Team1Kills         int    `json:"Team1_Kills"`
	Team1Score         int    `json:"Team1_Score"`
	Team2AvgLevel      int    `json:"Team2_AvgLevel"`
	Team2Gold          int    `json:"Team2_Gold"`
	Team2Kills         int    `json:"Team2_Kills"`
	Team2Score         int    `json:"Team2_Score"`
	WinningTeam        int    `json:"Winning_Team"`
	RetMsg             string `json:"ret_msg"`
}

// Player is a smite player/account
type Player struct {
	AccountID string `json:"account_id"`
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
	PlayerID  string `json:"player_id"`
	RetMsg    string `json:"ret_msg"`
}

func (s *Player) String() string {
	return s.Name
}

type RankedInfo struct {
	Leaves int `json:"Leaves"`
	Losses int `json:"Losses"`
	Name string `json:"Conquest"`
	Points int `json:"Points"`
	PrevRank int `json:"PrevRank"`
	Rank int `json:"Rank"`
	Rank_Stat_Conquest string `json:"Rank_Stat_Conquest"`
	Rank_Stat_Duel string `json:"Rank_Stat_Duel"`
	Rank_Stat_Joust string `json:"Rank_Stat_Joust"`
	Season int `json:"Season"`
	Tier int `json:"Tier"`
	Trend int `json:"Trend"`
	Wins int `json:"Wins"`
	PlayerId int `json:"player_id"`
	RetMsg string `json:"ret_msg"`
}

type GetPlayerResponse struct {
	AvatarURL string `json:"Avatar_URL"`
	CreatedDatetime string `json:"Created_Datetime"`
	Id int `json:"Id"`
	LastLoginDatetime string `json:"Last_Login_Datetime"`
	Leaves int `json:"Leaves"`
	Level int `json:"Level"`
	Losses int `json:"Losses"`
	MasteryLevel int `json:"MasteryLevel"`
	Name string `json:"Name"`
	RankStatConquest int `json:"Rank_Stat_Conquest"`
	RankStatDuel int `json:"Rank_Stat_Duel"`
	RankStatJoust int `json:"Rank_Stat_Joust"`
	RankedConquest RankedInfo `json:"RankedConquest"`
	RankedDuel RankedInfo `json:"RankedDuel"`
	RankedJoust RankedInfo `json:"RankedJoust"`
	TeamId int `json:"TeamId"`
	TeamName string `json:"TeamName"`
	Tier_Conquest int `json:"Tier_Conquest"`
	Tier_Duel int `json:"Tier_Duel"`
	Tier_Joust int `json:"Tier_Joust"`
	Total_Achievements int `json:"Total_Achievements"`
	Total_Worshippers int `json:"Total_Worshippers"`
	Wins int `json:"Wins"`
	RetMsg int `json:"ret_msg"`
}

// RecommendedItem is item stats
type RecommendedItem struct {
	Category        string `json:"Category"`
	Item            string `json:"Item"`
	Role            string `json:"Role"`
	CategoryValueID int    `json:"category_value_id"`
	GodID           int    `json:"god_id"`
	GodName         string `json:"god_name"`
	IconID          int    `json:"icon_id"`
	ItemID          int    `json:"item_id"`
	RetMsg          string `json:"ret_msg"`
	RoleValueID     int    `json:"role_value_id"`
}

// Item describes a smite item
type Item struct {
	ChildItemID     int             `json:"ChildItemId"`
	DeviceName      string          `json:"DeviceName"`
	IconID          int             `json:"IconId"`
	ItemDescription ItemDescription `json:"ItemDescription"`
	ItemID          int             `json:"ItemId"`
	ItemTier        int             `json:"ItemTier"`
	Price           int             `json:"Price"`
	RootItemID      int             `json:"RootItemId"`
	ShortDesc       string          `json:"ShortDesc"`
	StartingItem    bool            `json:"StartingItem"`
	Type            string          `json:"Type"`
	ItemIconURL     string          `json:"itemIcon_URL"`
	RetMsg          string          `json:"ret_msg"`
}

func (i *Item) String() string {
	return i.DeviceName
}

// MatchPlayerInfo describes detailed information about a single player in a match
type MatchPlayerInfo struct {
	AccountLevel int `json:"Account_Level"`
	ActiveID1    int `json:"ActiveId1"`
	ActiveID2    int `json:"ActiveId2"`
	Assists      int `json:"Assists"`

	Ban1   string `json:"Ban1"`
	Ban1Id int    `json:"Ban1Id"`

	Ban10   string `json:"Ban10"`
	Ban10Id int    `json:"Ban10Id"`

	Ban2   string `json:"Ban2"`
	Ban2Id int    `json:"Ban2Id"`

	Ban3   string `json:"Ban3"`
	Ban3Id int    `json:"Ban3Id"`

	Ban4   string `json:"Ban4"`
	Ban4Id int    `json:"Ban4Id"`

	Ban5   string `json:"Ban5"`
	Ban5Id int    `json:"Ban5Id"`

	Ban6   string `json:"Ban6"`
	Ban6Id int    `json:"Ban6Id"`

	Ban7   string `json:"Ban7"`
	Ban7Id int    `json:"Ban7Id"`

	Ban8   string `json:"Ban8"`
	Ban8Id int    `json:"Ban8Id"`

	Ban9   string `json:"Ban9"`
	Ban9Id int    `json:"Ban9Id"`

	CampsCleared int `json:"Camps_Cleared"`

	ConquestLosses int `json:"Conquest_Losses"`
	ConquestPoints int `json:"Conquest_Points"`
	ConquestTier   int `json:"Conquest_Tier"`
	ConquestWins   int `json:"Conquest_Wins"`

	DamageBot          int `json:"Damage_Bot"`
	DamageDoneMagical  int `json:"Damage_Done_Magical"`
	DamageDonePhysical int `json:"Damage_Done_Physical"`
	DamageMitigated    int `json:"Damage_Mitigated"`
	DamagePlayer       int `json:"Damage_Player"`
	DamageTaken        int `json:"Damage_Taken"`
	Deaths             int `json:"Deaths"`

	DuelLosses int `json:"Duel_Losses"`
	DuelPoints int `json:"Duel_Points"`
	DuelTier   int `json:"Duel_Tier"`
	DuelWins   int `json:"Duel_Wins"`

	EntryDatetime string `json:"Entry_Datetime"`

	FinalMatchLevel int `json:"Final_Match_Level"`

	FirstBanSide  string `json:"First_Ban_Side"`
	GodID         int    `json:"GodId"`
	GoldEarned    int    `json:"Gold_Earned"`
	GoldPerMinute int    `json:"Gold_Per_Minute"`

	Healing int `json:"Healing"`
	ItemID1 int `json:"ItemId1"`
	ItemID2 int `json:"ItemId2"`
	ItemID3 int `json:"ItemId3"`
	ItemID4 int `json:"ItemId4"`
	ItemID5 int `json:"ItemId5"`
	ItemID6 int `json:"ItemId6"`

	ItemActive1 string `json:"Item_Active_1"`
	ItemActive2 string `json:"Item_Active_2"`
	ItemActive3 string `json:"Item_Active_3"`

	ItemPurch1 string `json:"Item_Purch_1"`
	ItemPurch2 string `json:"Item_Purch_2"`
	ItemPurch3 string `json:"Item_Purch_3"`
	ItemPurch4 string `json:"Item_Purch_4"`
	ItemPurch5 string `json:"Item_Purch_5"`
	ItemPurch6 string `json:"Item_Purch_6"`

	JoustLosses int `json:"Joust_Losses"`
	JoustPoints int `json:"Joust_Points"`
	JoustTier   int `json:"Joust_Tier"`
	JoustWins   int `json:"Joust_Wins"`

	KillingSpree int `json:"Killing_Spree"`

	KillsBot int `json:"Kills_Bot"`

	KillsDouble int `json:"Kills_Double"`

	KillsFireGiant int `json:"Kills_Fire_Giant"`

	KillsFirstBlood int `json:"Kills_First_Blood"`

	KillsGoldFury int `json:"Kills_Gold_Fury"`

	KillsPenta           int    `json:"Kills_Penta"`
	KillsPhoenix         int    `json:"Kills_Phoenix"`
	KillsPlayer          int    `json:"Kills_Player"`
	KillsQuadra          int    `json:"Kills_Quadra"`
	KillsSiegeJuggernaut int    `json:"Kills_Siege_Juggernaut"`
	KillsSingle          int    `json:"Kills_Single"`
	KillsTriple          int    `json:"Kills_Triple"`
	KillsWildJuggernaut  int    `json:"Kills_Wild_Juggernaut"`
	MasteryLevel         int    `json:"Mastery_Level"`
	Match                int    `json:"Match"`
	Minutes              int    `json:"Minutes"`
	MultikillMax         int    `json:"Multi_kill_Max"`
	PartyID              int    `json:"PartyId"`
	RankStatConquest     int    `json:"Rank_Stat_Conquest"`
	RankStatDuel         int    `json:"Rank_Stat_Duel"`
	RankStatJoust        int    `json:"Rank_Stat_Joust"`
	ReferenceName        string `json:"Reference_Name"`
	Skin                 string `json:"Skin"`
	SkinID               int    `json:"SkinId"`
	StructureDamage      int    `json:"Structure_Damage"`
	Surrendered          string `json:"Surrendered"`
	Team1Score           int    `json:"Team1Score"`
	Team2Score           int    `json:"Team2Score"`
	TeamID               int    `json:"TeamId"`
	TeamName             string `json:"Team_Name"`
	TowersDestroyed      int    `json:"Towers_Destroyed"`
	WardsPlaced          int    `json:"Wards_Placed"`
	WinStatus            string `json:"Win_Status"`
	HasReplay            string `json:"hasReplay"`
	Name                 string `json:"name"`
	PlayerID             string `json:"playerId"`
	PlayerName           string `json:"playerName"`
	RetMsg               string `json:"ret_msg"`
}

// MatchPlayerDetails are limited information about a currently live match.
type MatchPlayerDetails struct {
	AccountLevel  int    `json:"Account_Level"`
	GodID         int    `json:"GodId"`
	GodName       string `json:"GodName"`
	MasteryLevel  int    `json:"Mastery_Level"`
	Match         int    `json:"Match"`
	Queue         string `json:"Queue"`
	SkinID        int    `json:"SkinId"`
	Tier          int    `json:"Tier"`
	PlayerCreated string `json:"playerCreated"`
	PlayerID      string `json:"playerId"`
	PlayerName    string `json:"playerName"`
	RetMsg        string `json:"ret_msg"`
	TaskForce     int    `json:"taskForce"`
	TierLosses    int    `json:"tierLosses"`
	TierWins      int    `json:"tierWins"`
}

// MatchQueueID is limited information about a running (or ran) match.
type MatchQueueID struct {
	ActiveFlag string `json:"Active_Flag"`
	Match      string `json:"Match"`
	RetMsg     string `json:"ret_msg"`
}

// LeaderboardPlayer is a result from getting tier leader boards
type LeaderboardPlayer struct {
	Leaves int `json:"Leaves"`
	Losses int `json:"Losses"`
	Name string `json:"Name"`
	Points int `json:"Points"`
	PrevRank int `json:"PrevRank"`
	Rank int `json:"Rank"`
	RankStatConquest string `json:"Rank_Stat_Conquest"`
	RankStatDuel string `json:"Rank_Stat_Duel"`
	RankStatJoust string `json:"Rank_Stat_Joust"`
	Season int `json:"Season"`
	Tier int `json:"Tier"`
	Trend int `json:"Trend"`
	Wins int `json:"Wins"`
	PlayerID string `json:"player_id"`
	RetMsg string `json:"ret_msg"`
}

type LeagueSeason struct {
	Complete bool `json:"complete"`
	ID int `json:"id"`
	Name string `json:"name"`
	RetMsg string `json:"ret_msg"`
}

type PlayerMatchHistory struct {
	ActiveId1 int `json:"ActiveId1"`
	ActiveId2 int `json:"ActiveId2"`
	Active_1 string `json:"Active_1"`
	Active_2 string `json:"Active_2"`
	Active_3 string `json:"Active_3"`
	Assists int `json:"Assists"`
	Ban1 string `json:"Ban1"`
	Ban1Id int `json:"Ban1Id"`

	Ban10 string `json:"Ban10"`
	Ban10Id int `json:"Ban10Id"`

	Ban2 string `json:"Ban2"`
	Ban2Id int `json:"Ban2Id"`

	Ban3 string `json:"Ban3"`
	Ban3Id int `json:"Ban3Id"`

	Ban4 string `json:"Ban4"`
	Ban4Id int `json:"Ban4Id"`

	Ban5 string `json:"Ban5"`
	Ban5Id int `json:"Ban5Id"`

	Ban6 string `json:"Ban6"`
	Ban6Id int `json:"Ban6Id"`

	Ban7 string `json:"Ban7"`
	Ban7Id int `json:"Ban7Id"`

	Ban8 string `json:"Ban8"`
	Ban8Id int `json:"Ban8Id"`

	Ban9 string `json:"Ban9"`
	Ban9Id int `json:"Ban9Id"`

	Creeps int `json:"Creeps"`
	Damage int `json:"Damage"`
	Damage_Bot int `json:"Damage_Bot"`
	Damage_Mitigated int `json:"Damage_Mitigated"`
	Damage_Structure int `json:"Damage_Structure"`
	Damage_Taken int `json:"Damage_Taken"`
	Deaths int `json:"Deaths"`
	First_Ban_Side string `json:"First_Ban_Side"`
	God string `json:"God"`
	GodId int `json:"GodId"`
	Gold int `json:"Gold"`
	Healing int `json:"Healing"`
	ItemId1 int `json:"ItemId1"`
	ItemId2 int `json:"ItemId2"`
	ItemId3 int `json:"ItemId3"`
	ItemId4 int `json:"ItemId4"`
	ItemId5 int `json:"ItemId5"`
	ItemId6 int `json:"ItemId6"`

	Item_1 string `json:"Item_1"`
	Item_2 string `json:"Item_2"`
	Item_3 string `json:"Item_3"`
	Item_4 string `json:"Item_4"`
	Item_5 string `json:"Item_5"`
	Item_6 string `json:"Item_6"`

	Killing_Spree int `json:"Killing_Spree"`
	Kills int `json:"Kills"`
	Level int `json:"Level"`
	Match int `json:"Match"`
	Match_Time string `json:"Match_Time"`
	Minutes int `json:"Minutes"`

	Multi_kill_Max int `json:"Multi_kill_Max"`
	Queue string `json:"Queue"`
	Skin string `json:"Skin"`
	SkinId int `json:"SkinId"`

	Surrendered string `json:"Surrendered"`
	Team1Score int `json:"Team1Score"`
	Team2Score int `json:"Team2Score"`
	Wards_Placed int `json:"Wards_Placed"`
	Win_Status string `json:"Win_Status"`
	PlayerName string `json:"playerName"`
	RetMsg string `json:"ret_msg"`
}

type MOTDResponse struct {
	Description string `json:"description"`
	GameMode string `json:"gameMode"`
	MaxPlayers string `json:"maxPlayers"`
	Name string `json:"name"`
	RetMsg string `json:"ret_msg"`
	StartDateTime string `json:"startDateTime"`
	Team1GodsCSV string `json:"team1GodsCSV"`
	Team2GodsCSV string `json:"team2GodsCSV"`
	Title string `json:"title"`
}

// LanguageCode controls what language a response is in
type LanguageCode int

const (
	// English is the english language
	English LanguageCode = 1
	// German is the german language
	German = 2
	// French is the French language
	French = 3
	// Spanish is the Spanish language
	Spanish = 4
	// SpanishLA is the Spanish (Latin America) language
	SpanishLA = 5
	// Portuguese is the Portuguese language
	Portuguese = 10
	// Russian is the Russian language
	Russian = 11
	// Polish is the Polish language
	Polish = 12
	// Turkish is the Turkish language
	Turkish = 13
)

// Queue describes a type of smite match
type Queue int

const (
	// Conquest5v5 is the Conquest queue
	Conquest5v5 Queue = 423
	// NoviceQueue is unknown type
	NoviceQueue = 424
	// Conquest is unknown queue type
	Conquest = 426
	// Practice is a custom map
	Practice = 427
	// ConquestChallenge is an unknown type (Not sure what challenge means)
	ConquestChallenge = 429
	// ConquestRanked is ranked conquest
	ConquestRanked = 430
	// Domination is an unknown type
	Domination = 433
	// MOTD is a map of the day
	MOTD = 434 // (use with 465 to get all MOTD matches),
	// Arena map type
	Arena = 435
	// ArenaChallenge is unknown type
	ArenaChallenge = 438
	// DominationChallenge is unknown type
	DominationChallenge = 439
	// JoustRanked1v1RankedDuel is ranked duel
	JoustRanked1v1RankedDuel = 440
	// JoustChallenge is an unkonwn type
	JoustChallenge = 441
	// Assault is the assault type
	Assault = 445
	// AssaultChallenge is an unknown type
	AssaultChallenge = 446
	// Joust3v3 is normal joust
	Joust3v3 = 448
	// JoustRanked3v3 is ranked joust
	JoustRanked3v3 = 450
	// ConquestRanked2 is an unknown type.  But I think it's the new conquest (??)
	ConquestRanked2 = 451
	// ArenaRanked is an old map type, of ranked arena
	ArenaRanked = 452
	// MOTD2 is map of the day, but all map of the day types
	MOTD2 = 465 // (Supports “closing” the Queue by our platform; use with 434)
	// Clash is the clash map type
	Clash = 466
	// ClashChallenge is an unknown map type (Not sure what challenge means)
	ClashChallenge = 467
)

type Tier int

const (
	BronzeV Tier = 1
	BronzeIV = 2
	BronzeIII = 3
	BronzeII = 4
	BronzeI = 5
	SilverV = 6
	SilverIV = 7
	SilverIII = 8
	SilverII = 9
	SilverI = 10
	GoldV = 11
	GoldIV = 12
	GoldIII = 13
	GoldII = 14
	GoldI = 15
	PlatinumV = 16
	PlatinumIV = 17
	PlatinumIII = 18
	PlatinumII = 19
	PlatinumI = 20
	DiamondV = 21
	DiamondIV = 22
	DiamondIII = 23
	DiamondII = 24
	DiamondI = 25
	MastersI = 26
)
