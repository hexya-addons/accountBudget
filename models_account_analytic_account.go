package account_budget

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.AccountAnalyticAccount().DeclareModel()

	h.AccountAnalyticAccount().AddFields(map[string]models.FieldDefinition{
		"CrossoveredBudgetLine": models.One2ManyField{
			RelationModel: h.CrossoveredBudgetLines(),
			ReverseFK:     "",
			String:        "Budget Lines",
		},
	})
}
