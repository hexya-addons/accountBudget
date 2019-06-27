package account_budget

import (
	"github.com/hexya-addons/base"
	"github.com/hexya-erp/pool/h"
)

//vars

var ()

//rights
func init() {
	h.CrossoveredBudget().Methods().Load().AllowGroup(GroupAccountManager)
	h.AccountBudgetPost().Methods().Load().AllowGroup(GroupAccountManager)
	h.AccountBudgetPost().Methods().AllowAllToGroup(GroupAccountUser)
	h.CrossoveredBudget().Methods().AllowAllToGroup(GroupAccountUser)
	h.CrossoveredBudgetLines().Methods().AllowAllToGroup(GroupAccountUser)
	h.CrossoveredBudgetLines().Methods().Load().AllowGroup(base.GroupUser)
	h.CrossoveredBudgetLines().Methods().Write().AllowGroup(base.GroupUser)
	h.CrossoveredBudgetLines().Methods().Create().AllowGroup(base.GroupUser)
}
