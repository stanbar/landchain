package main

import (
	"flowteam/landchain/types"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// func CreateRequest(landID string, parties []types.Individual, intermediary types.Individual) error {
// 	return nil
// }

var requests map[string]types.SignedRequest
var lands = map[string]types.Land{
	"GD1G/00000123/1": {
		ID:      "GD1G/00000123/1",
		Address: "ul. Kolejowa 1, 00-000 Gdańsk",
		Owner: types.Owner{
			Share: 100,
			Individual: types.Individual{
				PESEL:     "12345678901",
				Name:      "Jan Kowalski",
				PublicKey: "1234567890123456789012345678901234567890123456789012345678901234"}},
		Area: 100,
		Mortgage: types.Mortgage{
			Creditor: "Bank",
			Amount:   1000000,
		},
	},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func validateRequest(request types.SignedRequest) error {
	return nil
}

func areOtherRequestsWaitingForPayment(req types.SignedRequest) bool {
	for _, request := range requests {
		if req.ID == request.ID && request.State == types.WaitingForPayment {
			return true
		}
	}
	return false
}
func thereExistsOtherRequestsWaitingForPaymentCreatedEarlierThanThisRequest(req types.SignedRequest) (error, bool) {
	reqCreationTime, err := time.Parse(time.RFC3339, req.CreationTime)
	if err != nil {
		return err, false
	}

	for _, request := range requests {
		aRequestCreationTime, err := time.Parse(time.RFC3339, request.CreationTime)
		if err != nil {
			return err, false
		}

		if req.ID == request.ID && request.State == types.WaitingForPayment && reqCreationTime.Before(aRequestCreationTime) {
			return nil, true
		}
	}
	return nil, false
}

func canFastForward(req types.SignedRequest) bool {
	currentNotarialAct := lands[req.LandID].NotarialAct
	return reflect.DeepEqual(currentNotarialAct, req.PreviousNotarialAct)
}

func getStateRequest(key string) types.SignedRequest {
	return requests[key]
}

func putStateRequest(key string, value types.SignedRequest) {
	requests[key] = value
}

func getStateLand(key string) types.Land {
	return lands[key]
}

func putStateLand(key string, value types.Land) {
	lands[key] = value
}

func SubmitRequest(request types.SignedRequest) error {
	if err := validateRequest(request); err != nil {
		return err
	}

	key := fmt.Sprintf("%s.%s", request.LandID, randStringRunes(16))
	request.ID = key
	if canFastForward(request) {
		TerminateRequest(request, types.Accepted)
	} else {
		if areOtherRequestsWaitingForPayment(request) {
			request.State = types.Blocked
			putStateRequest(request.ID, request)
		} else {
			request.State = types.WaitingForPayment
			putStateRequest(request.ID, request)
		}
		request.State = types.WaitingForApproval
	}
	return nil
}

func Pay(request types.SignedRequest) error {
	err, result := thereExistsOtherRequestsWaitingForPaymentCreatedEarlierThanThisRequest(request)
	if err != nil {
		return err
	}

	if result {
		request.State = types.Blocked
		putStateRequest(request.ID, request)
	} else {
		request.State = types.WaitingForApproval
		putStateRequest(request.ID, request)
	}
	return nil
}

func ApproveRequest(request types.SignedRequest) error {
	return TerminateRequest(request, types.Accepted)
}

func RejectRequest(request types.SignedRequest) error {
	return TerminateRequest(request, types.Rejected)
}

func TerminateRequest(request types.SignedRequest, requestType types.State) error {
	if requestType == types.Accepted {
		request.State = types.Accepted
		putStateRequest(request.ID, request)
		currentLand := getStateLand(request.LandID)
		currentLand.NotarialAct = request.NewNotarialAct
		currentLand.Owner = request.NewOwner
		putStateLand(request.LandID, currentLand)
	} else if requestType == types.Rejected {
		request.State = types.Rejected
		putStateRequest(request.ID, request)
	}

	request.State = types.Rejected
	putStateRequest(request.ID, request)
	return nil
}

func GetAllRequestsWaitingForApproval() []types.SignedRequest {
	var result []types.SignedRequest
	for _, request := range requests {
		if request.State == types.WaitingForApproval {
			result = append(result, request)
		}
	}
	return result
}
