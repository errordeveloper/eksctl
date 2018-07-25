package cloudformation

// AWSBudgetsBudget_BudgetData AWS CloudFormation Resource (AWS::Budgets::Budget.BudgetData)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html
type AWSBudgetsBudget_BudgetData struct {

	// BudgetLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgetlimit
	BudgetLimit *AWSBudgetsBudget_Spend `json:"BudgetLimit,omitempty"`

	// BudgetName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgetname
	BudgetName string `json:"BudgetName,omitempty"`

	// BudgetType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgettype
	BudgetType string `json:"BudgetType,omitempty"`

	// CostFilters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-costfilters
	CostFilters interface{} `json:"CostFilters,omitempty"`

	// CostTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-costtypes
	CostTypes *AWSBudgetsBudget_CostTypes `json:"CostTypes,omitempty"`

	// TimePeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-timeperiod
	TimePeriod *AWSBudgetsBudget_TimePeriod `json:"TimePeriod,omitempty"`

	// TimeUnit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-timeunit
	TimeUnit string `json:"TimeUnit,omitempty"`
}

type UntypedAWSBudgetsBudget_BudgetData struct {

	// BudgetLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgetlimit
	BudgetLimit *UntypedAWSBudgetsBudget_Spend `json:"BudgetLimit,omitempty"`

	// BudgetName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgetname
	BudgetName interface{} `json:"BudgetName,omitempty"`

	// BudgetType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-budgettype
	BudgetType interface{} `json:"BudgetType,omitempty"`

	// CostFilters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-costfilters
	CostFilters interface{} `json:"CostFilters,omitempty"`

	// CostTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-costtypes
	CostTypes *UntypedAWSBudgetsBudget_CostTypes `json:"CostTypes,omitempty"`

	// TimePeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-timeperiod
	TimePeriod *UntypedAWSBudgetsBudget_TimePeriod `json:"TimePeriod,omitempty"`

	// TimeUnit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-budgetdata.html#cfn-budgets-budget-budgetdata-timeunit
	TimeUnit interface{} `json:"TimeUnit,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBudgetsBudget_BudgetData) AWSCloudFormationType() string {
	return "AWS::Budgets::Budget.BudgetData"
}
