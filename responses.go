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

// LanguageCode controls what language a response is in
type LanguageCode int

var (
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
