package payment

import (
	"bwastartup/user"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	midtrans "github.com/veritrans/go-midtrans"
)

type service struct{
}

type Service interface{
	GetPaymentUrl(transaction Transaction, user user.User)(string, error)
}


func NewService() *service {
	return &service{}
}

func (s *service)GetPaymentUrl(transaction Transaction, user user.User)(string, error){
	err := godotenv.Load()
	midclient := midtrans.NewClient()
    midclient.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
    midclient.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
    midclient.APIEnvType = midtrans.Sandbox

    snapGateway := midtrans.SnapGateway{
        Client: midclient,
    }

	fmt.Println(user.Email)
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil{
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
