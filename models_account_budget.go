package account_budget

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

func init() {
	h.AccountBudgetPost().DeclareModel()

	h.AccountBudgetPost().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:   "Name",
			Required: true,
		},
		"AccountIds": models.Many2ManyField{
			RelationModel:    h.AccountAccount(),
			M2MLinkModelName: "",
			M2MOurField:      "",
			M2MTheirField:    "",
			String:           "Accounts",
			Filter:           q.Deprecated().Equals(False),
		},
		"CrossoveredBudgetLine": models.One2ManyField{
			RelationModel: h.CrossoveredBudgetLines(),
			ReverseFK:     "",
			String:        "Budget Lines",
		},
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
			Required:      true,
			Default:       func(env models.Environment) interface{} { return env["res.company"]._company_default_get() },
		},
	})
	h.CrossoveredBudget().DeclareModel()

	h.CrossoveredBudget().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:   "Budget Name",
			Required: true,
			//states={'done': [('readonly', True)]}
		},
		"CreatingUserId": models.Many2OneField{
			RelationModel: h.User(),
			String:        "Responsible",
			Default:       func(env models.Environment) interface{} { return env.Uid() },
		},
		"DateFrom": models.DateField{
			String:   "Start Date",
			Required: true,
			//states={'done': [('readonly', True)]}
		},
		"DateTo": models.DateField{
			String:   "End Date",
			Required: true,
			//states={'done': [('readonly', True)]}
		},
		"State": models.SelectionField{
			Selection: types.Selection{
				"draft":    "Draft",
				"cancel":   "Cancelled",
				"confirm":  "Confirmed",
				"validate": "Validated",
				"done":     "Done",
			},
			String:   "Status",
			Default:  models.DefaultValue("draft"),
			Index:    true,
			Required: true,
			ReadOnly: true,
			NoCopy:   true,
			//track_visibility='always'
		},
		"CrossoveredBudgetLine": models.One2ManyField{
			RelationModel: h.CrossoveredBudgetLines(),
			ReverseFK:     "",
			String:        "Budget Lines",
			//states={'done': [('readonly', True)]}
			NoCopy: false,
		},
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
			Required:      true,
			Default:       func(env models.Environment) interface{} { return env["res.company"]._company_default_get() },
		},
	})
	h.CrossoveredBudget().Methods().ActionBudgetConfirm().DeclareMethod(
		`ActionBudgetConfirm`,
		func(rs m.CrossoveredBudgetSet) {
			//        self.write({'state': 'confirm'})
		})
	h.CrossoveredBudget().Methods().ActionBudgetDraft().DeclareMethod(
		`ActionBudgetDraft`,
		func(rs m.CrossoveredBudgetSet) {
			//        self.write({'state': 'draft'})
		})
	h.CrossoveredBudget().Methods().ActionBudgetValidate().DeclareMethod(
		`ActionBudgetValidate`,
		func(rs m.CrossoveredBudgetSet) {
			//        self.write({'state': 'validate'})
		})
	h.CrossoveredBudget().Methods().ActionBudgetCancel().DeclareMethod(
		`ActionBudgetCancel`,
		func(rs m.CrossoveredBudgetSet) {
			//        self.write({'state': 'cancel'})
		})
	h.CrossoveredBudget().Methods().ActionBudgetDone().DeclareMethod(
		`ActionBudgetDone`,
		func(rs m.CrossoveredBudgetSet) {
			//        self.write({'state': 'done'})
		})
	h.CrossoveredBudgetLines().DeclareModel()

	h.CrossoveredBudgetLines().AddFields(map[string]models.FieldDefinition{
		"CrossoveredBudgetId": models.Many2OneField{
			RelationModel: h.CrossoveredBudget(),
			String:        "Budget",
			OnDelete:      `cascade`,
			Index:         true,
			Required:      true,
		},
		"AnalyticAccountId": models.Many2OneField{
			RelationModel: h.AccountAnalyticAccount(),
			String:        "Analytic Account",
		},
		"GeneralBudgetId": models.Many2OneField{
			RelationModel: h.AccountBudgetPost(),
			String:        "Budgetary Position",
			Required:      true,
		},
		"DateFrom": models.DateField{
			String:   "Start Date",
			Required: true,
		},
		"DateTo": models.DateField{
			String:   "End Date",
			Required: true,
		},
		"PaidDate": models.DateField{
			String: "Paid Date",
		},
		"PlannedAmount": models.FloatField{
			String:   "Planned Amount",
			Required: true,
			//digits=0
		},
		"PracticalAmount": models.FloatField{
			Compute: h.CrossoveredBudgetLines().Methods().ComputePracticalAmount(),
			String:  "Practical Amount",
			//digits=0
		},
		"TheoriticalAmount": models.FloatField{
			Compute: h.CrossoveredBudgetLines().Methods().ComputeTheoriticalAmount(),
			String:  "Theoretical Amount",
			//digits=0
		},
		"Percentage": models.FloatField{
			Compute: h.CrossoveredBudgetLines().Methods().ComputePercentage(),
			String:  "Achievement",
		},
		"CompanyId": models.Many2OneField{
			Related:       `CrossoveredBudgetId.CompanyId`,
			RelationModel: h.Company(),
			String:        "Company",
			Stored:        true,
			ReadOnly:      true,
		},
	})
	h.CrossoveredBudgetLines().Methods().ComputePracticalAmount().DeclareMethod(
		`ComputePracticalAmount`,
		func(rs h.CrossoveredBudgetLinesSet) h.CrossoveredBudgetLinesData {
			//        for line in self:
			//            result = 0.0
			//            acc_ids = line.general_budget_id.account_ids.ids
			//            if not acc_ids:
			//                raise UserError(_("The Budget '%s' has no accounts!") %
			//                                ustr(line.general_budget_id.name))
			//            date_to = self.env.context.get('wizard_date_to') or line.date_to
			//            date_from = self.env.context.get(
			//                'wizard_date_from') or line.date_from
			//            if line.analytic_account_id.id:
			//                self.env.cr.execute("""
			//                    SELECT SUM(amount)
			//                    FROM account_analytic_line
			//                    WHERE account_id=%s
			//                        AND (date between to_date(%s,'yyyy-mm-dd') AND to_date(%s,'yyyy-mm-dd'))
			//                        AND general_account_id=ANY(%s)""",
			//                                    (line.analytic_account_id.id, date_from, date_to, acc_ids))
			//                result = self.env.cr.fetchone()[0] or 0.0
			//            line.practical_amount = result
		})
	h.CrossoveredBudgetLines().Methods().ComputeTheoriticalAmount().DeclareMethod(
		`ComputeTheoriticalAmount`,
		func(rs h.CrossoveredBudgetLinesSet) h.CrossoveredBudgetLinesData {
			//        today = fields.Datetime.now()
			//        for line in self:
			//            # Used for the report
			//
			//            if self.env.context.get('wizard_date_from') and self.env.context.get('wizard_date_to'):
			//                date_from = fields.Datetime.from_string(
			//                    self.env.context.get('wizard_date_from'))
			//                date_to = fields.Datetime.from_string(
			//                    self.env.context.get('wizard_date_to'))
			//                if date_from < fields.Datetime.from_string(line.date_from):
			//                    date_from = fields.Datetime.from_string(line.date_from)
			//                elif date_from > fields.Datetime.from_string(line.date_to):
			//                    date_from = False
			//
			//                if date_to > fields.Datetime.from_string(line.date_to):
			//                    date_to = fields.Datetime.from_string(line.date_to)
			//                elif date_to < fields.Datetime.from_string(line.date_from):
			//                    date_to = False
			//
			//                theo_amt = 0.00
			//                if date_from and date_to:
			//                    line_timedelta = fields.Datetime.from_string(
			//                        line.date_to) - fields.Datetime.from_string(line.date_from)
			//                    elapsed_timedelta = date_to - date_from
			//                    if elapsed_timedelta.days > 0:
			//                        theo_amt = (elapsed_timedelta.total_seconds(
			//                        ) / line_timedelta.total_seconds()) * line.planned_amount
			//            else:
			//                if line.paid_date:
			//                    if fields.Datetime.from_string(line.date_to) <= fields.Datetime.from_string(line.paid_date):
			//                        theo_amt = 0.00
			//                    else:
			//                        theo_amt = line.planned_amount
			//                else:
			//                    line_timedelta = fields.Datetime.from_string(
			//                        line.date_to) - fields.Datetime.from_string(line.date_from)
			//                    elapsed_timedelta = fields.Datetime.from_string(
			//                        today) - (fields.Datetime.from_string(line.date_from))
			//
			//                    if elapsed_timedelta.days < 0:
			//                        # If the budget line has not started yet, theoretical amount should be zero
			//                        theo_amt = 0.00
			//                    elif line_timedelta.days > 0 and fields.Datetime.from_string(today) < fields.Datetime.from_string(line.date_to):
			//                        # If today is between the budget line date_from and date_to
			//                        theo_amt = (elapsed_timedelta.total_seconds(
			//                        ) / line_timedelta.total_seconds()) * line.planned_amount
			//                    else:
			//                        theo_amt = line.planned_amount
			//
			//            line.theoritical_amount = theo_amt
		})
	h.CrossoveredBudgetLines().Methods().ComputePercentage().DeclareMethod(
		`ComputePercentage`,
		func(rs h.CrossoveredBudgetLinesSet) h.CrossoveredBudgetLinesData {
			//        for line in self:
			//            if line.theoritical_amount != 0.00:
			//                line.percentage = float(
			//                    (line.practical_amount or 0.0) / line.theoritical_amount) * 100
			//            else:
			//                line.percentage = 0.00
		})
}
