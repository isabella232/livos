package impl

import (
	"testing"

	"github.com/dedis/livos/voting"
	"github.com/stretchr/testify/require"
)

var VoteList = make(map[string]*VotingInstance)
var VoteSystem = NewVotingSystem(nil, VoteList)

//Creation of a empty list of choces (for history)
var histoChoice = make([]voting.Choice, 0)

var userNoemien, err = VoteSystem.NewUser("Noemien", make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice)

func TestCreationUserNoemien(t *testing.T) {
	require.Equal(t, err, nil, "Cannot create VotingConfig")
}

var userGuillaume, err1 = VoteSystem.NewUser("Guillaume", make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice)

func TestCreationUserGuillaume(t *testing.T) {
	require.Equal(t, err1, nil, "Cannot create user")
}

var userEtienne, err2 = VoteSystem.NewUser("Etienne", make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice)

func TestCreationUserEtienne(t *testing.T) {
	require.Equal(t, err2, nil, "Cannot create user")
}

var userJoseph, err3 = VoteSystem.NewUser("Joseph", make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice)

func TestCreationUserJoseph(t *testing.T) {
	require.Equal(t, err3, nil, "Cannot create user")
}

var voters = []*voting.User{&userNoemien, &userGuillaume, &userEtienne, &userJoseph}
var candidats = make([]string, 3)

//var votes = make(map[string]voting.Choice)

func TestVotingSystemCreate(t *testing.T) {
	voteConfig, err := NewVotingConfig(voters, "TestVotingTitle", "Quick description", candidats)
	require.Equal(t, err, nil, "Cannot create VotingConfig")

	VoteSystem.CreateAndAdd("Session01", voteConfig, "open")
	id := VoteSystem.VotingInstancesList["Session01"].Id
	require.Equal(t, id, "Session01", "The id of the votingInstance just created is incorrect, got: %s, want %s.", id, "Session01")

	status := VoteSystem.VotingInstancesList["Session01"].Status
	require.Equal(t, status, "open", "The status of the votingInstance just created is incorrect, got: %s, want %s.", status, "open")

	config := VoteSystem.VotingInstancesList["Session01"].Config
	require.Equal(t, config.Title, "TestVotingTitle", "The config title of the votingInstance just created is incorrect, got: %s, want %s.", config.Title, "TestVotingTitle")

	require.Equal(t, config.Description, "Quick description", "The config description of the votingInstance just created is incorrect, got: %s, want %s.", config.Description, "Quick description")
}

func TestSetStatus(t *testing.T) {

	voteConfig, err := NewVotingConfig(voters, "TestVotingTitle", "Quick description", candidats)
	require.Equal(t, err, nil, "Creation of votingConfig is incorrect.")

	VoteSystem.CreateAndAdd("Session01", voteConfig, "open")
	addVoteInst := VoteSystem.VotingInstancesList["Session01"]

	s := "close"
	addVoteInst.SetStatus(s)
	require.Equal(t, addVoteInst.Status, s, "Status incorrect. Was: %s, should be: %s", addVoteInst.Status, s)
}

func TestCloseVoting(t *testing.T) {
	voteConfig, err := NewVotingConfig(voters, "TestVotingTitle", "Quick description", candidats)
	require.Equal(t, err, nil, "Creation of votingConfig is incorrect.")

	s := "close"
	VoteSystem.CreateAndAdd("Session01", voteConfig, "open")
	addVoteInst := VoteSystem.VotingInstancesList["Session01"]
	addVoteInst.CloseVoting()
	require.Equal(t, addVoteInst.Status, s, "Status incorrect. Was: %s, should be: %s", addVoteInst.Status, s)

}

func TestGetResults(t *testing.T) {
	voteConfig, err := NewVotingConfig(voters, "TestVotingTitle", "Quick description", candidats)
	require.Equal(t, err, nil, "Creation of votingConfig is incorrect.")

	vi, err := VoteSystem.CreateAndAdd("Session01", voteConfig, "open")
	require.Equal(t, err, nil, "Creation of votingInstance is incorrect.")

	yesChoice := make(map[string]voting.Liquid)
	noChoice := make(map[string]voting.Liquid)
	midChoice := make(map[string]voting.Liquid)

	liq100, err := NewLiquid(100)
	require.Equal(t, err, nil, "Creation of liquid is incorrect.")

	liq50, err := NewLiquid(50)
	require.Equal(t, err, nil, "Creation of liquid is incorrect.")

	liqid0, err := NewLiquid(0)
	require.Equal(t, err, nil, "Creation of liquid is incorrect.")

	yesChoice["yes"] = liq100
	yesChoice["no"] = liqid0
	noChoice["no"] = liq100
	noChoice["yes"] = liqid0
	midChoice["no"] = liq50
	midChoice["yes"] = liq50
	choiceGuillaume, errG := NewChoice(noChoice)
	choiceEtienne, errE := NewChoice(midChoice)
	choiceNoemien, errN := NewChoice(yesChoice)
	require.Equal(t, errN, nil, "Creation of the choice is incorrect.")

	require.Equal(t, errE, nil, "Creation of the choice is incorrect.")

	require.Equal(t, errG, nil, "Creation of the choice is incorrect.")

	err = vi.SetVote(&userGuillaume, choiceGuillaume)
	require.Equal(t, err, nil, "Impossible to cast a vote, negative voting Power.")

	err = vi.SetVote(&userEtienne, choiceEtienne)
	require.Equal(t, err, nil, "Impossible to cast a vote, negative voting Power.")

	err = vi.SetVote(&userNoemien, choiceNoemien)
	require.Equal(t, err, nil, "Impossible to cast a vote, negative voting Power.")

	propYes := vi.GetResults()["yes"]
	require.Equal(t, propYes, 50., "Yes proportion is incorrect, got: %f, want: %f.", propYes, 50.)

	propNo := vi.GetResults()["no"]
	require.Equal(t, propYes, 50., "No proportion is incorrect, got: %f, want: %f.", propNo, 50.)

}

func TestCastVotes(t *testing.T) {

	voteConfig, err := NewVotingConfig(voters, "TestVotingTitle", "Quick description", candidats)
	require.Equal(t, err, nil, "Creation of votingConfig is incorrect.")

	vi, err := VoteSystem.CreateAndAdd("Session01", voteConfig, "open")
	require.Equal(t, err, nil, "Creation of votingInstance is incorrect.")

	yesChoice := make(map[string]voting.Liquid)

	liq100, err := NewLiquid(100)
	require.Equal(t, err, nil, "Creation of liquid is incorrect.")

	liqid0, err := NewLiquid(0)
	require.Equal(t, err, nil, "Creation of liquid is incorrect.")

	yesChoice["yes"] = liq100
	yesChoice["no"] = liqid0

	choiceJoseph, errN := NewChoice(yesChoice)
	require.Equal(t, errN, nil, "Creation of the choice is incorrect.")

	err = vi.SetVote(&userJoseph, choiceJoseph)
	require.Equal(t, err, nil, "Impossible to set a vote, negative voting Power.")

	require.Equal(t, userJoseph.HistoryOfChoice[0].VoteValue["yes"].Percentage, 100., "Proportion in yes is incorrect. Was: %f, should be %f", userJoseph.HistoryOfChoice[0].VoteValue["yes"].Percentage, 100.)
}
