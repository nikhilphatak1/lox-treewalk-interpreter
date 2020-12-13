package lox

// Unfinished

// AstPrinter print AST
type AstPrinter struct {}

// VisitAssignExpr visit
func (a *AstPrinter) VisitAssignExpr(expr *AssignExpr) interface{}

// VisitBinaryExpr visit
func (a *AstPrinter) VisitBinaryExpr(expr *BinaryExpr) interface{} {
	return a.parenthesize(expr.operator.lexeme, expr.left, expr.right)
}

// VisitCallExpr visit
func (a *AstPrinter) VisitCallExpr(expr *CallExpr) interface{}

// VisitGetExpr visit
func (a *AstPrinter) VisitGetExpr(expr *GetExpr) interface{}

// VisitGroupingExpr visit
func (a *AstPrinter) VisitGroupingExpr(expr *GroupingExpr) interface{} {
	return a.parenthesize("group", expr.expression)
}

// VisitLiteralExpr visit
func (a *AstPrinter) VisitLiteralExpr(expr *LiteralExpr) interface{} {
	if expr.value == nil {
		return "nil"
	}
	return expr.value
}

// VisitLogicalExpr visit
func (a *AstPrinter) VisitLogicalExpr(expr *LogicalExpr) interface{} {
	return a.parenthesize(expr.operator.lexeme, expr.left, expr.right)
}

// VisitSetExpr visit
func (a *AstPrinter) VisitSetExpr(expr *SetExpr) interface{} {

}

// VisitSuperExpr visit
func (a *AstPrinter) VisitSuperExpr(expr *SuperExpr) interface{}

// VisitThisExpr visit
func (a *AstPrinter) VisitThisExpr(expr *ThisExpr) interface{}

// VisitUnaryExpr visit
func (a *AstPrinter) VisitUnaryExpr(expr *UnaryExpr) interface{}

// VisitVariableExpr visit
func (a *AstPrinter) VisitVariableExpr(expr *VariableExpr) interface{}

// Print print expr
func (a *AstPrinter) Print(expr Expr) string {
	return expr.Accept(a).(string)
}

func (a *AstPrinter) parenthesize(name string, exprs ...Expr) string {
	output := "("
	for _, expr := range exprs {
		output = output + " " + expr.Accept(a).(string)
	}
	output = output + ")"
	return output
}