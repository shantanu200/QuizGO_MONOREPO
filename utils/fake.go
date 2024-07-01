package utils

import (
	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateRandomFakeUser() []interface{} {

	var users []interface{}

	for i := 1; i <= 100; i++ {
		user := map[string]interface{}{
			"name":     gofakeit.Name(),
			"email":    gofakeit.Email(),
			"password": gofakeit.Password(true, true, true, false, false, 10),
		}

		users = append(users, user)
	}
	return users
}

func CreateRandomStudentForLeaderBoard() ([]interface{}, error) {

	var students []interface{}
	_quiz, err := primitive.ObjectIDFromHex("662572a1b99bdc146e3e50e7")

	if err != nil {
		return nil, err
	}

	studentIds := []string{
		"6640df7d5ce9394683282657",
		"6640df7d5ce939468328264d",
		"6640df7d5ce9394683282630",
		"6640df7d5ce939468328262a",
		"6640df7d5ce9394683282642",
		"6640df7d5ce9394683282632",
		"6640df7d5ce939468328267c",
		"6640df7d5ce9394683282626",
		"6640df7d5ce9394683282645",
		"6640df7d5ce9394683282660",
		"6640df7d5ce9394683282647",
		"6640df7d5ce939468328267a",
		"6640df7d5ce9394683282654",
		"6640df7d5ce939468328264e",
		"6640df7d5ce9394683282652",
		"6640df7d5ce939468328266b",
		"6640df7d5ce9394683282665",
		"6640df7d5ce9394683282650",
		"6640df7d5ce9394683282635",
		"662b4730340095c79d3cd63c",
		"6640df7d5ce9394683282663",
		"6640df7d5ce9394683282664",
		"6640df7d5ce9394683282662",
		"6640df7d5ce939468328267e",
		"6640df7d5ce9394683282646",
		"6640df7d5ce939468328267d",
		"6640df7d5ce9394683282670",
		"6640df7d5ce939468328261e",
		"6640df7d5ce9394683282621",
		"6640df7d5ce9394683282667",
		"6640df7d5ce9394683282636",
		"6640df7d5ce9394683282639",
		"6640df7d5ce939468328266f",
		"6640df7d5ce939468328266e",
		"6640df7d5ce9394683282629",
		"6640df7d5ce939468328262f",
		"6640df7d5ce939468328265b",
		"6640df7d5ce939468328265f",
		"6640df7d5ce939468328266c",
		"6640df7d5ce939468328263e",
		"6640df7d5ce9394683282656",
		"6640df7d5ce9394683282653",
		"6640df7d5ce939468328262c",
		"6640df7d5ce939468328264b",
		"6640df7d5ce9394683282623",
		"6640df7d5ce939468328265a",
		"6640df7d5ce939468328267f",
		"6640df7d5ce939468328263a",
		"6640df7d5ce939468328265d",
		"6640df7d5ce9394683282659",
		"6640df7d5ce939468328267b",
		"6640df7d5ce9394683282640",
		"6640df7d5ce9394683282634",
		"6640df7d5ce9394683282676",
		"6640df7d5ce939468328262d",
		"6640df7d5ce939468328265c",
		"6640df7d5ce939468328264a",
		"6640df7d5ce9394683282679",
		"6640df7d5ce9394683282666",
		"6640df7d5ce939468328263d",
		"6640df7d5ce9394683282668",
		"6640df7d5ce9394683282638",
		"6640df7d5ce9394683282643",
		"6640df7d5ce939468328265e",
		"6640df7d5ce939468328266d",
		"6640df7d5ce9394683282661",
		"6640df7d5ce9394683282628",
		"6640df7d5ce9394683282631",
		"6640df7d5ce939468328261f",
		"6640df7d5ce9394683282648",
		"6640df7d5ce9394683282637",
		"6640df7d5ce9394683282673",
		"6640df7d5ce939468328264f",
		"6640df7d5ce9394683282669",
		"6640df7d5ce939468328266a",
		"6640df7d5ce9394683282655",
		"6640df7d5ce939468328262e",
		"6640df7d5ce939468328263c",
		"6640df7d5ce9394683282649",
		"6640df7d5ce9394683282641",
		"6640df7d5ce9394683282622",
		"6640df7d5ce9394683282674",
		"6640df7d5ce9394683282678",
		"6640df7d5ce939468328264c",
		"6640df7d5ce939468328262b",
		"6640df7d5ce9394683282620",
		"6640df7d5ce9394683282644",
		"6640df7d5ce9394683282625",
		"6640df7d5ce9394683282633",
		"6640df7d5ce939468328261c",
		"6640df7d5ce939468328263b",
		"6640df7d5ce9394683282677",
		"6640df7d5ce9394683282624",
		"6640df7d5ce9394683282672",
		"6640df698869e8d33911783a",
		"6640df7d5ce9394683282671",
		"6640df7d5ce9394683282675",
		"6640df7d5ce9394683282651",
		"6640df7d5ce939468328261d",
		"6640df7d5ce9394683282627",
	}

	for _, studentId := range studentIds {
		_user, err := primitive.ObjectIDFromHex(studentId)
		if err != nil {
			return nil, err
		}
		correctAnswers := gofakeit.Number(1, 15)
		wrongAnswers := 15 - correctAnswers
		totalAttempted := correctAnswers + wrongAnswers
		accuracy := float64(correctAnswers) / float64(totalAttempted)
		totalScore := correctAnswers*4 - wrongAnswers
		student := map[string]interface{}{
			"quizId":         _quiz,
			"userId":         _user,
			"isActive":       false,
			"isOver":         true,
			"correctAnswers": correctAnswers,
			"wrongAnswers":   wrongAnswers,
			"totalAttempted": totalAttempted,
			"accuracy":       accuracy,
			"totalScore":     totalScore,
		}

		students = append(students, student)
	}

	return students, nil
}
