package lox

// Expr a valid lox expression
type Expr interface {
	Accept(visitor ExprVisitor) interface{}
}

// ExprVisitor visitor for Exprs
type ExprVisitor interface {
	VisitAssignExpr(expr *AssignExpr) interface{}
    VisitBinaryExpr(expr *BinaryExpr) interface{}
    VisitCallExpr(expr *CallExpr) interface{}
    VisitGetExpr(expr *GetExpr) interface{}
    VisitGroupingExpr(expr *GroupingExpr) interface{}
    VisitLiteralExpr(expr *LiteralExpr) interface{}
    VisitLogicalExpr(expr *LogicalExpr) interface{}
    VisitSetExpr(expr *SetExpr) interface{}
    VisitSuperExpr(expr *SuperExpr) interface{}
    VisitThisExpr(expr *ThisExpr) interface{}
    VisitUnaryExpr(expr *UnaryExpr) interface{}
    VisitVariableExpr(expr *VariableExpr) interface{}
}

// AssignExpr implements Expr
type AssignExpr struct {
	name Token
	value Expr
}

// Accept accept visitor
func (e *AssignExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitAssignExpr(e)
}

// BinaryExpr implements Expr
type BinaryExpr struct {
	left Expr
	operator Token
	right Expr
}

// Accept accept visitor
func (e *BinaryExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitBinaryExpr(e)
}

// NewBinaryExpr make a BonaryExpr
func NewBinaryExpr(left Expr, operator Token, right Expr) *BinaryExpr {
	binaryExpr := BinaryExpr{}
	binaryExpr.left = left
	binaryExpr.right = right
	binaryExpr.operator = operator
	return &binaryExpr
}

// CallExpr implements Expr
type CallExpr struct {
	callee Expr
	paren Token
	arguments []Expr
}

// Accept accept visitor
func (e *CallExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitCallExpr(e)
}

// GetExpr implements Expr
type GetExpr struct {
	object Expr
	name Token
}

// Accept accept visitor
func (e *GetExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGetExpr(e)
}

// GroupingExpr implements Expr
type GroupingExpr struct {
	expression Expr
}

// Accept accept visitor
func (e *GroupingExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGroupingExpr(e)
}

// LiteralExpr implements Expr
type LiteralExpr struct {
	value interface{}
}

// Accept accept visitor
func (e *LiteralExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(e)
}

// NewLiteralExpr new literal expression
func NewLiteralExpr(literal interface{}) *LiteralExpr {
	literalExpr := LiteralExpr{}
	literalExpr.value = literal
	return &literalExpr
}

// LogicalExpr implements Expr
type LogicalExpr struct {
	left Expr
	operator Token
	right Expr
}

// Accept accept visitor
func (e *LogicalExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLogicalExpr(e)
}

// SetExpr implements Expr
type SetExpr struct {
	object Expr
	name Token
	value Expr
}

// Accept accept visitor
func (ae *SetExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitSetExpr(ae)
}

// SuperExpr implements Expr
type SuperExpr struct {
	keyword Token
	method Token
}

// Accept accept visitor
func (e *SuperExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitSuperExpr(e)
}

// ThisExpr implements Expr
type ThisExpr struct {
	keyword Token
}

// Accept accept visitor
func (e *ThisExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitThisExpr(e)
}

// UnaryExpr implements Expr
type UnaryExpr struct {
	operator Token
	right Expr
}

// Accept accept visitor
func (e *UnaryExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitUnaryExpr(e)
}

// NewUnaryExpr make a unary expr
func NewUnaryExpr(operator Token, right Expr) *UnaryExpr {
	unary := UnaryExpr{}
	unary.operator = operator
	unary.right = right
	return &unary
}

// VariableExpr implements Expr
type VariableExpr struct {
	name Token
}

// Accept accept visitor
func (e *VariableExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitVariableExpr(e)
}
