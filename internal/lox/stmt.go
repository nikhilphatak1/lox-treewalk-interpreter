package lox

// Stmt a valid lox expression
type Stmt interface {
	Accept(visitor StmtVisitor) interface{}
}

// StmtVisitor visitor for Exprs
type StmtVisitor interface {
	VisitBlockStmt(stmt *BlockStmt) interface{}
    VisitClassStmt(stmt *ClassStmt) interface{}
    VisitExpressionStmt(stmt *ExpressionStmt) interface{}
    VisitFunctionStmt(stmt *FunctionStmt) interface{}
    VisitIfStmt(stmt *IfStmt) interface{}
    VisitPrintStmt(stmt *PrintStmt) interface{}
    VisitReturnStmt(stmt *ReturnStmt) interface{}
    VisitVarStmt(stmt *VarStmt) interface{}
    VisitWhileStmt(stmt *WhileStmt) interface{}
}

// BlockStmt implements Stmt
type BlockStmt struct {
	statements []Stmt
}

// Accept accept visitor
func (s *BlockStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitBlockStmt(s)
}

// ClassStmt implements Stmt
type ClassStmt struct {
	name Token
	superclass VariableExpr
	methods []FunctionStmt
}

// Accept accept visitor
func (s *ClassStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitClassStmt(s)
}

// ExpressionStmt implements Stmt
type ExpressionStmt struct {
	expression Expr
}

// Accept accept visitor
func (s *ExpressionStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitExpressionStmt(s)
}

// FunctionStmt implements Stmt
type FunctionStmt struct {
	name Token
	params []Token
	body []Stmt
}

// Accept accept visitor
func (s *FunctionStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitFunctionStmt(s)
}

// IfStmt implements Stmt
type IfStmt struct {
	condition Expr
	thenBranch Stmt
	elseBranch Stmt
}

// Accept accept visitor
func (s *IfStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitIfStmt(s)
}

// PrintStmt implements Stmt
type PrintStmt struct {
	expression Expr
}

// Accept accept visitor
func (s *PrintStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitPrintStmt(s)
}

// ReturnStmt implements Stmt
type ReturnStmt struct {
	keyword Token
	value Expr
}

// Accept accept visitor
func (s *ReturnStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitReturnStmt(s)
}

// VarStmt implements Stmt
type VarStmt struct {
	name Token
	initializer Expr
}

// Accept accept visitor
func (s *VarStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitVarStmt(s)
}

// WhileStmt implements Stmt
type WhileStmt struct {
	condition Expr
	body Stmt
}

// Accept accept visitor
func (s *WhileStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitWhileStmt(s)
}
