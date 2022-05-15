package beanacount

import (
	"bufio"
	"fmt"
	"github.com/clstb/phi/ledger/internal/config"
	"github.com/shopspring/decimal"
	"os"
	"strconv"
	"time"
)

func (l *Ledger) UpdateLedger(providers []Provider, accounts []Account, transactions []TinkTransaction) {
	providersById := map[string]string{}
	for _, provider := range providers {
		providersById[provider.FinancialInstitutionId] = provider.DisplayName
	}

	opensByTinkId := l.Opens().ByTinkId()
	for _, account := range accounts {
		_, ok := opensByTinkId[account.ID]
		if ok {
			continue
		}

		*l = append(*l, Open{
			Date: "1970-01-01",
			Account: fmt.Sprintf(
				"Assets:%s:%s",
				providersById[account.FinancialInstitutionId],
				account.Name,
			),
			Metadata: []Metadata{
				{
					Key:   "tink_id",
					Value: strconv.Quote(account.ID),
				},
			},
		})
	}
	opensByTinkId = l.Opens().ByTinkId()

	transactionsByTinkId := l.Transactions().ByTinkId()
	for _, transaction := range transactions {
		_, ok := transactionsByTinkId[transaction.ID]
		if ok {
			continue
		}

		amount := AmountType{
			Decimal: decimal.New(
				transaction.Amount.Value.UnscaledValue,
				transaction.Amount.Value.Scale*-1,
			),
			Currency: transaction.Amount.CurrencyCode,
		}
		var balanceAccount string
		if amount.IsNegative() {
			balanceAccount = "Expenses:Unassigned"
		} else {
			balanceAccount = "Income:Unassigned"
		}

		date, _ := time.Parse("2006-01-02", transaction.Dates.Booked)

		*l = append(*l, Transaction{
			Date:      date,
			Type:      "*",
			Payee:     transaction.Reference,
			Narration: transaction.Descriptions,
			Metadata: []Metadata{
				{
					Key:   "tink_id",
					Value: strconv.Quote(transaction.ID),
				},
			},
			Postings: []Posting{
				{
					Account: balanceAccount,
					Units:   amount.Neg(),
				},
				{
					Account: opensByTinkId[transaction.AccountID].Account,
					Units:   amount,
				},
			},
		})
	}
}

func (l *Ledger) PersistLedger(username string) error {
	filePath := fmt.Sprintf("%s/%s/transactions.beancount", config.DataDirPath, username)
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	datawriter := bufio.NewWriter(file)
	defer datawriter.Flush()
	defer file.Close()

	for _, i := range *l {
		_, _ = datawriter.WriteString(i.String() + "\n")
	}

	return nil
}
