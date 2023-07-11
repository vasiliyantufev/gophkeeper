package table

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/vasiliyantufev/gophkeeper/internal/client/model"
	"github.com/vasiliyantufev/gophkeeper/internal/client/storage/layouts"
	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/service"
)

//const ColId = 0
//const ColName = 1
//const ColDescription = 2
//const ColText = 3
//const ColTblText = 5
//const ColTblCard = 9

func SearchByColumn(slice [][]string, targetColumn int, targetValue string) bool {
	for i := 1; i < len(slice) && len(slice) > 1; i++ {
		if slice[i][targetColumn] == targetValue {
			return true
		}
	}
	return false
}

//func GetIndex(slice [][]string, targetColumn int, targetValue string) (index int) {
//	for index = 1; index < len(slice) && len(slice) > 1; index++ {
//		if slice[index][targetColumn] == targetValue {
//			return index
//		}
//	}
//	return 0
//}

func RemoveRow(slice [][]string, indexRow int) [][]string {
	return append(slice[:indexRow], slice[indexRow+1:]...)
}

func UpdateRowLoginPassword(login, password string, slice [][]string, indexRow int) [][]string {
	indexColLogin := 2
	indexColPassword := 3
	indexColUpdateAt := 5
	slice[indexRow][indexColLogin] = login
	slice[indexRow][indexColPassword] = password
	slice[indexRow][indexColUpdateAt] = time.Now().Format(layouts.LayoutDateAndTime.ToString())
	return slice
}

func UpdateRowText(text string, slice [][]string, indexRow int) [][]string {
	indexColText := 2
	indexColUpdateAt := 4
	slice[indexRow][indexColText] = text
	slice[indexRow][indexColUpdateAt] = time.Now().Format(layouts.LayoutDateAndTime.ToString())
	return slice
}

func UpdateRowCard(paymentSystem, number, holder, cvc, endDate string, slice [][]string, indexRow int) [][]string {
	indexColPaymentSystem := 2
	indexColNumber := 3
	indexColHolder := 4
	indexColCvc := 5
	indexColEndDate := 6
	indexColUpdateAt := 8
	slice[indexRow][indexColPaymentSystem] = paymentSystem
	slice[indexRow][indexColNumber] = number
	slice[indexRow][indexColHolder] = holder
	slice[indexRow][indexColEndDate] = endDate
	slice[indexRow][indexColCvc] = cvc
	slice[indexRow][indexColUpdateAt] = time.Now().Format(layouts.LayoutDateAndTime.ToString())
	return slice
}

//func AppendText(node *grpc.Text, dataTblText *[][]string, plaintext string) {
//	created, _ := service.ConvertTimestampToTime(node.CreatedAt)
//	updated, _ := service.ConvertTimestampToTime(node.UpdatedAt)
//	if node.Key == string(variables.Name) {
//		row := []string{strconv.Itoa(int(node.Id)), node.Value, "", plaintext, created.Format(layouts.LayoutDateAndTime.ToString()), updated.Format(layouts.LayoutDateAndTime.ToString())}
//		*dataTblText = append(*dataTblText, row)
//	} else if node.Key == string(variables.Description) {
//		row := []string{strconv.Itoa(int(node.Id)), "", node.Value, plaintext, created.Format(layouts.LayoutDateAndTime.ToString()), updated.Format(layouts.LayoutDateAndTime.ToString())}
//		*dataTblText = append(*dataTblText, row)
//	}
//}
//
//func UpdateText(node *grpc.Text, dataTblText *[][]string, index int) {
//	if node.Key == string(variables.Name) {
//		(*dataTblText)[index][ColName] = node.Value
//	} else if node.Key == string(variables.Description) {
//		(*dataTblText)[index][ColDescription] = node.Value
//	}
//}

//func AppendCard(node *grpc.Card, dataTblCard *[][]string, jsonCard model.Card) {
//	created, _ := service.ConvertTimestampToTime(node.CreatedAt)
//	updated, _ := service.ConvertTimestampToTime(node.UpdatedAt)
//	if node.Key == string(variables.Name) {
//		row := []string{strconv.Itoa(int(node.Id)), node.Value, "", jsonCard.PaymentSystem, jsonCard.Number,
//			jsonCard.Holder, strconv.Itoa(jsonCard.CVC), jsonCard.EndDate.Format(layouts.LayoutDate.ToString()),
//			created.Format(layouts.LayoutDateAndTime.ToString()), updated.Format(layouts.LayoutDateAndTime.ToString())}
//		*dataTblCard = append(*dataTblCard, row)
//	} else if node.Key == string(variables.Description) {
//		row := []string{strconv.Itoa(int(node.Id)), "", node.Value, jsonCard.PaymentSystem, jsonCard.Number,
//			jsonCard.Holder, strconv.Itoa(jsonCard.CVC), jsonCard.EndDate.Format(layouts.LayoutDate.ToString()),
//			created.Format(layouts.LayoutDateAndTime.ToString()), updated.Format(layouts.LayoutDateAndTime.ToString())}
//		*dataTblCard = append(*dataTblCard, row)
//	}
//}
//
//func UpdateCard(node *grpc.Card, dataTblCard *[][]string, index int) {
//	if node.Key == string(variables.Name) {
//		(*dataTblCard)[index][ColName] = node.Value
//	} else if node.Key == string(variables.Description) {
//		(*dataTblCard)[index][ColDescription] = node.Value
//	}
//}

//func AppendLoginPassword(node *grpc.LoginPassword, dataTblLoginPassword *[][]string, jsonLoginPassword model.LoginPassword) {
//	created, _ := service.ConvertTimestampToTime(node.CreatedAt)
//	updated, _ := service.ConvertTimestampToTime(node.UpdatedAt)
//	if node.Key == string(variables.Name) {
//		row := []string{strconv.Itoa(int(node.Id)), node.Value, "", jsonLoginPassword.Login, jsonLoginPassword.Password,
//			created.Format(layouts.LayoutDateAndTime.ToString()), updated.Format(layouts.LayoutDateAndTime.ToString())}
//		*dataTblLoginPassword = append(*dataTblLoginPassword, row)
//	} else if node.Key == string(variables.Description) {
//		row := []string{strconv.Itoa(int(node.Id)), "", node.Value, jsonLoginPassword.Login, jsonLoginPassword.Password,
//			created.Format(layouts.LayoutDateAndTime.ToString()), updated.Format(layouts.LayoutDateAndTime.ToString())}
//		*dataTblLoginPassword = append(*dataTblLoginPassword, row)
//	}
//}
//
//func UpdateLoginPassword(node *grpc.LoginPassword, dataTblLoginPassword *[][]string, index int) {
//	if node.Key == string(variables.Name) {
//		(*dataTblLoginPassword)[index][ColName] = node.Value
//	} else if node.Key == string(variables.Description) {
//		(*dataTblLoginPassword)[index][ColDescription] = node.Value
//	}
//}

//func DeleteColId(dataTblText *[][]string) {
//	for index := range *dataTblText {
//		(*dataTblText)[index] = (*dataTblText)[index][1:]
//	}
//}

// -------------------------------------------------
func AppendTextEntity(node *grpc.Entity, dataTblText *[][]string, plaintext string) error {
	created, err := service.ConvertTimestampToTime(node.CreatedAt)
	if err != nil {
		return err
	}
	updated, err := service.ConvertTimestampToTime(node.UpdatedAt)
	if err != nil {
		return err
	}
	var metadata model.MetadataEntity
	err = json.Unmarshal([]byte(node.Metadata), &metadata)
	if err != nil {
		return err
	}
	row := []string{metadata.Name, metadata.Description, plaintext, created.Format(layouts.LayoutDateAndTime.ToString()), updated.Format(layouts.LayoutDateAndTime.ToString())}
	*dataTblText = append(*dataTblText, row)
	return nil
}

func AppendLoginPasswordEntity(node *grpc.Entity, dataTblLoginPassword *[][]string, jsonLoginPassword model.LoginPassword) error {
	created, err := service.ConvertTimestampToTime(node.CreatedAt)
	if err != nil {
		return err
	}
	updated, err := service.ConvertTimestampToTime(node.UpdatedAt)
	if err != nil {
		return err
	}
	var metadata model.MetadataEntity
	err = json.Unmarshal([]byte(node.Metadata), &metadata)
	if err != nil {
		return err
	}
	row := []string{metadata.Name, metadata.Description, jsonLoginPassword.Login, jsonLoginPassword.Password,
		created.Format(layouts.LayoutDateAndTime.ToString()), updated.Format(layouts.LayoutDateAndTime.ToString())}
	*dataTblLoginPassword = append(*dataTblLoginPassword, row)
	return nil
}

func AppendCardEntity(node *grpc.Entity, dataTblCard *[][]string, jsonCard model.Card) error {
	created, err := service.ConvertTimestampToTime(node.CreatedAt)
	if err != nil {
		return err
	}
	updated, err := service.ConvertTimestampToTime(node.UpdatedAt)
	if err != nil {
		return err
	}
	var metadata model.MetadataEntity
	err = json.Unmarshal([]byte(node.Metadata), &metadata)
	if err != nil {
		return err
	}
	row := []string{metadata.Name, metadata.Description, jsonCard.PaymentSystem, jsonCard.Number,
		jsonCard.Holder, strconv.Itoa(jsonCard.CVC), jsonCard.EndDate.Format(layouts.LayoutDate.ToString()),
		created.Format(layouts.LayoutDateAndTime.ToString()), updated.Format(layouts.LayoutDateAndTime.ToString())}
	*dataTblCard = append(*dataTblCard, row)
	return nil
}

func AppendBinary(node *grpc.Binary, dataTblBinary *[][]string) error {
	created, err := service.ConvertTimestampToTime(node.CreatedAt)
	if err != nil {
		return err
	}
	row := []string{node.Name, created.Format(layouts.LayoutDateAndTime.ToString())}
	*dataTblBinary = append(*dataTblBinary, row)
	return nil
}
