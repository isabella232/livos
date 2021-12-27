package simulation

import (
	"fmt"
	"io"
	"strconv"

	"github.com/dedis/livos/voting"
	"github.com/dedis/livos/voting/impl"
	"github.com/mazen160/go-random"
	"golang.org/x/xerrors"
)

// GenerateItemsGraphviz creates a graphviz representation of the items. One can
// generate a graphical representation with `dot -Tpdf graph.dot -o graph.pdf`
func Simulation_candidats(out io.Writer) {

	const InitialVotingPower = 100.

	var VoteList = make(map[string]voting.VotingInstance)
	var VoteSystem = impl.NewVotingSystem(nil, VoteList)
	var histoChoice = make([]voting.Choice, 0)

	var randomNumOfUser, err = random.IntRange(10, 11)
	if err != nil {
		xerrors.Errorf(err.Error())
	}

	//Random creating of a user and adds it to the list of voters
	var voters = make([]*voting.User, 0)
	for i := 0; i < randomNumOfUser; i++ {
		var chooseType, err1 = random.IntRange(1, 101)
		if err1 != nil {
			xerrors.Errorf(err.Error())
		}
		switch {
		case chooseType < 5:
			var user, err = VoteSystem.NewUser("user"+strconv.FormatInt(int64(i), 10), make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice, voting.YesVoter, nil)
			if err != nil {
				xerrors.Errorf(err.Error())
			}
			voters = append(voters, &user)
		case chooseType < 50:
			var user, err = VoteSystem.NewUser("user"+strconv.FormatInt(int64(i), 10), make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice, voting.IndecisiveVoter, nil)
			if err != nil {
				xerrors.Errorf(err.Error())
			}
			voters = append(voters, &user)
		case chooseType < 70:
			var user, err = VoteSystem.NewUser("user"+strconv.FormatInt(int64(i), 10), make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice, voting.ThresholdVoter, nil)
			if err != nil {
				xerrors.Errorf(err.Error())
			}
			voters = append(voters, &user)
		case chooseType < 80:
			var user, err = VoteSystem.NewUser("user"+strconv.FormatInt(int64(i), 10), make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice, voting.NonResponsibleVoter, nil)
			if err != nil {
				xerrors.Errorf(err.Error())
			}
			voters = append(voters, &user)
		case chooseType < 90:
			var user, err = VoteSystem.NewUser("user"+strconv.FormatInt(int64(i), 10), make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice, voting.ResponsibleVoter, nil)
			if err != nil {
				xerrors.Errorf(err.Error())
			}
			voters = append(voters, &user)
		default:
			var user, err = VoteSystem.NewUser("user"+strconv.FormatInt(int64(i), 10), make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice, voting.None, nil)
			if err != nil {
				xerrors.Errorf(err.Error())
			}
			voters = append(voters, &user)
		}

	}

	/*
		YesNumber := 10
		NoNumber := 10
		IndecisiveNumber := 10
		ThresholdNumber := 10
		NonResponsibleNumber := 10
		TotalNumber := NonResponsibleNumber + YesNumber + NoNumber + IndecisiveNumber + ThresholdNumber

		i := 0
		for i = 0; i < YesNumber; i++ {
			var user, err = VoteSystem.NewUser("user"+strconv.FormatInt(int64(i), 10), make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice, voting.YesVoter, make([]*voting.User, 0))
			if err != nil {
				xerrors.Errorf(err.Error())
			}
			voters = append(voters, &user)
		}
		for i = i; i < NoNumber+YesNumber; i++ {
			var user, err = VoteSystem.NewUser("user"+strconv.FormatInt(int64(i), 10), make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice, voting.NoVoter, make([]*voting.User, 0))
			if err != nil {
				xerrors.Errorf(err.Error())
			}
			voters = append(voters, &user)
		}
		for i = i; i < IndecisiveNumber+NoNumber+YesNumber; i++ {
			var user, err = VoteSystem.NewUser("user"+strconv.FormatInt(int64(i), 10), make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice, voting.IndecisiveVoter, make([]*voting.User, 0))
			if err != nil {
				xerrors.Errorf(err.Error())
			}
			voters = append(voters, &user)
		}
		for i = i; i < ThresholdNumber+IndecisiveNumber+NoNumber+YesNumber; i++ {
			var user, err = VoteSystem.NewUser("user"+strconv.FormatInt(int64(i), 10), make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice, voting.ThresholdVoter, make([]*voting.User, 0))
			if err != nil {
				xerrors.Errorf(err.Error())
			}
			voters = append(voters, &user)
		}
		for i = i; i < NonResponsibleNumber+ThresholdNumber+IndecisiveNumber+NoNumber+YesNumber; i++ {
			var user, err = VoteSystem.NewUser("user"+strconv.FormatInt(int64(i), 10), make(map[string]voting.Liquid), make(map[string]voting.Liquid), histoChoice, voting.NonResponsibleVoter, make([]*voting.User, 0))
			if err != nil {
				xerrors.Errorf(err.Error())
			}
			voters = append(voters, &user)
		}
	*/

	//candidats inputs
	var candidatTrump, _ = VoteSystem.NewCandidate("Trump")
	var candidatObama, _ = VoteSystem.NewCandidate("Obama")
	var candidatJeanMi, _ = VoteSystem.NewCandidate("JeanMi")
	var candidatMacron, _ = VoteSystem.NewCandidate("Macron")

	var candidats = []*voting.Candidate{&candidatObama, &candidatTrump, &candidatJeanMi, &candidatMacron}

	//empty list of votes
	//var votes = make(map[string]voting.Choice)

	//creation of votingConfig
	voteConfig, err := impl.NewVotingConfig(voters, "Simulation 1", "Who are you gonna elect as a President ?", candidats, "CandidateQuestion")
	if err != nil {
		fmt.Println(err.Error())
	}

	//creation of the voting instance
	VoteInstance, err := VoteSystem.CreateAndAdd("Simulation01", voteConfig, "open")
	if err != nil {
		fmt.Println(err.Error())
	}

	yesVote := func(user *voting.User, votingPower float64) {
		quantity := votingPower
		quantity_to_Vote, err := impl.NewLiquid(float64(quantity))
		if err != nil {
			fmt.Println(err.Error())
		}

		choiceTab := make(map[string]voting.Liquid)

		candidateChoice, err := random.IntRange(0, len(candidats))
		if err != nil {
			fmt.Println(err.Error())
		}

		choiceTab[candidats[candidateChoice].CandidateID] = quantity_to_Vote

		//MODIFY THE RESULTS NAME => IT HAS TO BE DYNAMIC COMPARED TO THE LIST OF CANDIDATES !!!
		// switch {
		// case candidateChoice == 0:
		// 	choiceTab["Trump"] = quantity_to_Vote
		// case candidateChoice == 1:
		// 	choiceTab["Obama"] = quantity_to_Vote
		// case candidateChoice == 2:
		// 	choiceTab["JeanMi"] = quantity_to_Vote
		// default:
		// 	choiceTab["Macron"] = quantity_to_Vote
		// }

		//create choice
		choice, err := impl.NewChoice(choiceTab)
		if err != nil {
			fmt.Println(err.Error())
		}

		//set the choice
		err = VoteInstance.SetVote(user, choice)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(user.UserID, " a voté pour ", quantity, "% ", "il était", user.TypeOfUser)
	}

	IndecisiveVote := func(user *voting.User, i int) {

		//Delegation action

		//random index creation (must NOT be == to index of current user)
		randomDelegateToIndex, err := random.IntRange(0, len(voters))
		if err != nil {
			fmt.Println(err.Error(), "fail to do randomDelegateToIndex first time")
		}
		for ok := true; ok; ok = (randomDelegateToIndex == i) {
			randomDelegateToIndex, err = random.IntRange(0, len(voters))
			if err != nil {
				fmt.Println(err.Error(), "fail to do randomDelegateToIndex")
			}
		}
		quantity_to_deleg, err := impl.NewLiquid(float64(user.VotingPower))
		if err != nil {
			fmt.Println(err.Error(), "fail to do quantity to deleg")
		}
		err = VoteInstance.DelegTo(user, voters[randomDelegateToIndex], quantity_to_deleg)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(user.UserID, " a delegué ", quantity_to_deleg, " à : ", voters[randomDelegateToIndex].UserID, "il était", user.TypeOfUser)
	}
	randomVote := func(user *voting.User, i int) {
		randomAction, err := random.IntRange(1, 3)
		if err != nil {
			fmt.Println(err.Error(), "fail to do randomAction")
		}

		if randomAction == 1 {
			//Delegation action

			//random index creation (must NOT be == to index of current user)
			randomDelegateToIndex, err := random.IntRange(0, len(voters))
			if err != nil {
				fmt.Println(err.Error(), "fail to do randomDelegateToIndex first time")
			}
			for ok := true; ok; ok = (randomDelegateToIndex == i) {
				randomDelegateToIndex, err = random.IntRange(0, len(voters))
				if err != nil {
					fmt.Println(err.Error(), "fail to do randomDelegateToIndex")
				}
			}
			randomQuantityToDelegate, err := random.IntRange(1, int(user.VotingPower/10)+1)
			if err != nil {
				fmt.Println(err.Error(), "fail to do randomQuantityToDelegate")
			}
			randomQuantityToDelegate *= 10
			quantity_to_deleg, err := impl.NewLiquid(float64(randomQuantityToDelegate))
			if err != nil {
				fmt.Println(err.Error(), "fail to do quantity to deleg")
			}
			err = VoteInstance.DelegTo(user, voters[randomDelegateToIndex], quantity_to_deleg)
			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println(user.UserID, " a delegué ", quantity_to_deleg, " à : ", voters[randomDelegateToIndex].UserID, "il était", user.TypeOfUser)

		} else if randomAction == 2 {
			//Vote action

			quantity := user.VotingPower
			yesVote(user, quantity)

		}
	}

	ThresholdVote := func(user *voting.User, i int, threshold int) {

		var thresholdComparator = 0.
		for i := range user.HistoryOfChoice {
			thresholdComparator += user.HistoryOfChoice[i].VoteValue["yes"].Percentage
			thresholdComparator += user.HistoryOfChoice[i].VoteValue["no"].Percentage
		}

		if thresholdComparator > float64(threshold) {
			//Delegation action
			IndecisiveVote(user, i)

		} else {
			//Vote action

			quantity := user.VotingPower
			yesVote(user, quantity)
		}
	}
	NonResponsibleVoter := func(user *voting.User, i int) {
		if len(user.HistoryOfChoice) == 0 {
			yesVote(user, InitialVotingPower)
		} else {
			//Delegation action
			IndecisiveVote(user, i)
		}
	}
	ResponsibleVoter := func(user *voting.User, i int) {
		randomAction, err := random.IntRange(1, 3)
		if err != nil {
			fmt.Println(err.Error(), "fail to do randomAction")
		}

		if len(user.HistoryOfChoice) != 0 {
			randomAction = 2
		} else if user.DelegatedTo != nil {
			randomAction = 1
		}

		if randomAction == 1 {
			//Delegation action
			IndecisiveVote(user, i)

		} else if randomAction == 2 {
			//Vote action

			quantity := user.VotingPower
			yesVote(user, quantity)
		}
	}

	for ok := true; ok; ok = VoteInstance.CheckVotingPowerOfVoters() {
		for i, user := range VoteInstance.GetConfig().Voters {

			if user.VotingPower > 0 {
				switch user.TypeOfUser {
				case voting.YesVoter:
					yesVote(user, user.VotingPower)
				case voting.IndecisiveVoter:
					IndecisiveVote(user, i)
				case voting.ThresholdVoter:
					var threshold = 600
					ThresholdVote(user, i, threshold)
				case voting.NonResponsibleVoter:
					NonResponsibleVoter(user, i)
				case voting.ResponsibleVoter:
					ResponsibleVoter(user, i)
				case voting.None:
					randomVote(user, i)
				}
			}
		}
	}

	counterYesVoter := 0
	counterIndecisiveVoter := 0
	counterThresholdVoter := 0
	counterNormalVoter := 0
	counterNonResponsibleVoter := 0
	counterResponsibleVoter := 0
	for _, user := range voters {
		//fmt.Println("Voting power of ", user.UserID, " = ", user.VotingPower, "il était de type", user.TypeOfUser)
		if user.TypeOfUser == "YesVoter" {
			counterYesVoter++
		} else if user.TypeOfUser == "IndecisiveVoter" {
			counterIndecisiveVoter++
		} else if user.TypeOfUser == "ThresholdVoter" {
			counterThresholdVoter++
		} else if user.TypeOfUser == "NonResponsibleVoter" {
			counterNonResponsibleVoter++
		} else if user.TypeOfUser == "ResponsibleVoter" {
			counterResponsibleVoter++
		} else {
			counterNormalVoter++
		}
	}
	fmt.Println("There is ", counterYesVoter, "yesVoter,", counterThresholdVoter, "Threshold Voter,", counterNonResponsibleVoter, "NonresponsibleVoter,", counterResponsibleVoter, "ResponsibleVoter,", counterIndecisiveVoter, "IndecisiveVoter and", counterNormalVoter, "normalVoter")

	VoteInstance.ConstructTextForGraphCandidates(out, VoteInstance.GetResults())

}
