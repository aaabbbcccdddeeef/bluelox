// Code generated by gen-ast-types. DO NOT EDIT.

package ast

import (
	"errors"
	"github.com/nanmu42/bluelox/token"
)

type Expression interface {
	Accept(visitor ExprVisitor) (result interface{}, err error)
}

type Statement interface {
	Accept(visitor StmtVisitor) (err error)
}

type ExprVisitor interface {
	VisitAssignExpr(v *AssignExpr) (result interface{}, err error)
	VisitBinaryExpr(v *BinaryExpr) (result interface{}, err error)
	VisitCallExpr(v *CallExpr) (result interface{}, err error)
	VisitGetExpr(v *GetExpr) (result interface{}, err error)
	VisitGroupingExpr(v *GroupingExpr) (result interface{}, err error)
	VisitLiteralExpr(v *LiteralExpr) (result interface{}, err error)
	VisitLogicalExpr(v *LogicalExpr) (result interface{}, err error)
	VisitSetExpr(v *SetExpr) (result interface{}, err error)
	VisitThisExpr(v *ThisExpr) (result interface{}, err error)
	VisitUnaryExpr(v *UnaryExpr) (result interface{}, err error)
	VisitVariableExpr(v *VariableExpr) (result interface{}, err error)
}

type StubExprVisitor struct{}

var _ ExprVisitor = StubExprVisitor{}

func (s StubExprVisitor) VisitAssignExpr(_ *AssignExpr) (interface{}, error) {
	return nil, errors.New("visit func for AssignExpr is not implemented")
}

func (s StubExprVisitor) VisitBinaryExpr(_ *BinaryExpr) (interface{}, error) {
	return nil, errors.New("visit func for BinaryExpr is not implemented")
}

func (s StubExprVisitor) VisitCallExpr(_ *CallExpr) (interface{}, error) {
	return nil, errors.New("visit func for CallExpr is not implemented")
}

func (s StubExprVisitor) VisitGetExpr(_ *GetExpr) (interface{}, error) {
	return nil, errors.New("visit func for GetExpr is not implemented")
}

func (s StubExprVisitor) VisitGroupingExpr(_ *GroupingExpr) (interface{}, error) {
	return nil, errors.New("visit func for GroupingExpr is not implemented")
}

func (s StubExprVisitor) VisitLiteralExpr(_ *LiteralExpr) (interface{}, error) {
	return nil, errors.New("visit func for LiteralExpr is not implemented")
}

func (s StubExprVisitor) VisitLogicalExpr(_ *LogicalExpr) (interface{}, error) {
	return nil, errors.New("visit func for LogicalExpr is not implemented")
}

func (s StubExprVisitor) VisitSetExpr(_ *SetExpr) (interface{}, error) {
	return nil, errors.New("visit func for SetExpr is not implemented")
}

func (s StubExprVisitor) VisitThisExpr(_ *ThisExpr) (interface{}, error) {
	return nil, errors.New("visit func for ThisExpr is not implemented")
}

func (s StubExprVisitor) VisitUnaryExpr(_ *UnaryExpr) (interface{}, error) {
	return nil, errors.New("visit func for UnaryExpr is not implemented")
}

func (s StubExprVisitor) VisitVariableExpr(_ *VariableExpr) (interface{}, error) {
	return nil, errors.New("visit func for VariableExpr is not implemented")
}

type AssignExpr struct {
	Name  *token.Token
	Value Expression
}

var _ Expression = (*AssignExpr)(nil)

func (b *AssignExpr) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.VisitAssignExpr(b)
}

type BinaryExpr struct {
	Left     Expression
	Operator *token.Token
	Right    Expression
}

var _ Expression = (*BinaryExpr)(nil)

func (b *BinaryExpr) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.VisitBinaryExpr(b)
}

type CallExpr struct {
	Callee    Expression
	Paren     *token.Token
	Arguments []Expression
}

var _ Expression = (*CallExpr)(nil)

func (b *CallExpr) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.VisitCallExpr(b)
}

type GetExpr struct {
	Object Expression
	Name   *token.Token
}

var _ Expression = (*GetExpr)(nil)

func (b *GetExpr) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.VisitGetExpr(b)
}

type GroupingExpr struct {
	Expr Expression
}

var _ Expression = (*GroupingExpr)(nil)

func (b *GroupingExpr) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.VisitGroupingExpr(b)
}

type LiteralExpr struct {
	Value interface{}
}

var _ Expression = (*LiteralExpr)(nil)

func (b *LiteralExpr) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.VisitLiteralExpr(b)
}

type LogicalExpr struct {
	Left     Expression
	Operator *token.Token
	Right    Expression
}

var _ Expression = (*LogicalExpr)(nil)

func (b *LogicalExpr) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.VisitLogicalExpr(b)
}

type SetExpr struct {
	Object Expression
	Name   *token.Token
	Value  Expression
}

var _ Expression = (*SetExpr)(nil)

func (b *SetExpr) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.VisitSetExpr(b)
}

type ThisExpr struct {
	Keyword *token.Token
}

var _ Expression = (*ThisExpr)(nil)

func (b *ThisExpr) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.VisitThisExpr(b)
}

type UnaryExpr struct {
	Operator *token.Token
	Right    Expression
}

var _ Expression = (*UnaryExpr)(nil)

func (b *UnaryExpr) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.VisitUnaryExpr(b)
}

type VariableExpr struct {
	Name *token.Token
}

var _ Expression = (*VariableExpr)(nil)

func (b *VariableExpr) Accept(visitor ExprVisitor) (result interface{}, err error) {
	return visitor.VisitVariableExpr(b)
}

type StmtVisitor interface {
	VisitBlockStmt(v *BlockStmt) (err error)
	VisitClassStmt(v *ClassStmt) (err error)
	VisitExprStmt(v *ExprStmt) (err error)
	VisitFunctionStmt(v *FunctionStmt) (err error)
	VisitIfStmt(v *IfStmt) (err error)
	VisitPrintStmt(v *PrintStmt) (err error)
	VisitReturnStmt(v *ReturnStmt) (err error)
	VisitVarStmt(v *VarStmt) (err error)
	VisitWhileStmt(v *WhileStmt) (err error)
}

type StubStmtVisitor struct{}

var _ ExprVisitor = StubExprVisitor{}

func (s StubExprVisitor) VisitBlockStmt(_ *BlockStmt) error {
	return errors.New("visit func for BlockStmt is not implemented")
}

func (s StubExprVisitor) VisitClassStmt(_ *ClassStmt) error {
	return errors.New("visit func for ClassStmt is not implemented")
}

func (s StubExprVisitor) VisitExprStmt(_ *ExprStmt) error {
	return errors.New("visit func for ExprStmt is not implemented")
}

func (s StubExprVisitor) VisitFunctionStmt(_ *FunctionStmt) error {
	return errors.New("visit func for FunctionStmt is not implemented")
}

func (s StubExprVisitor) VisitIfStmt(_ *IfStmt) error {
	return errors.New("visit func for IfStmt is not implemented")
}

func (s StubExprVisitor) VisitPrintStmt(_ *PrintStmt) error {
	return errors.New("visit func for PrintStmt is not implemented")
}

func (s StubExprVisitor) VisitReturnStmt(_ *ReturnStmt) error {
	return errors.New("visit func for ReturnStmt is not implemented")
}

func (s StubExprVisitor) VisitVarStmt(_ *VarStmt) error {
	return errors.New("visit func for VarStmt is not implemented")
}

func (s StubExprVisitor) VisitWhileStmt(_ *WhileStmt) error {
	return errors.New("visit func for WhileStmt is not implemented")
}

type BlockStmt struct {
	Stmts []Statement
}

var _ Statement = (*BlockStmt)(nil)

func (b *BlockStmt) Accept(visitor StmtVisitor) (err error) {
	return visitor.VisitBlockStmt(b)
}

type ClassStmt struct {
	Name    *token.Token
	Methods []*FunctionStmt
}

var _ Statement = (*ClassStmt)(nil)

func (b *ClassStmt) Accept(visitor StmtVisitor) (err error) {
	return visitor.VisitClassStmt(b)
}

type ExprStmt struct {
	Expr Expression
}

var _ Statement = (*ExprStmt)(nil)

func (b *ExprStmt) Accept(visitor StmtVisitor) (err error) {
	return visitor.VisitExprStmt(b)
}

type FunctionStmt struct {
	Name   *token.Token
	Params []*token.Token
	Body   []Statement
}

var _ Statement = (*FunctionStmt)(nil)

func (b *FunctionStmt) Accept(visitor StmtVisitor) (err error) {
	return visitor.VisitFunctionStmt(b)
}

type IfStmt struct {
	Condition  Expression
	ThenBranch Statement
	ElseBranch Statement
}

var _ Statement = (*IfStmt)(nil)

func (b *IfStmt) Accept(visitor StmtVisitor) (err error) {
	return visitor.VisitIfStmt(b)
}

type PrintStmt struct {
	Expr Expression
}

var _ Statement = (*PrintStmt)(nil)

func (b *PrintStmt) Accept(visitor StmtVisitor) (err error) {
	return visitor.VisitPrintStmt(b)
}

type ReturnStmt struct {
	Keyword *token.Token
	Value   Expression
}

var _ Statement = (*ReturnStmt)(nil)

func (b *ReturnStmt) Accept(visitor StmtVisitor) (err error) {
	return visitor.VisitReturnStmt(b)
}

type VarStmt struct {
	Name        *token.Token
	Initializer Expression
}

var _ Statement = (*VarStmt)(nil)

func (b *VarStmt) Accept(visitor StmtVisitor) (err error) {
	return visitor.VisitVarStmt(b)
}

type WhileStmt struct {
	Condition Expression
	Body      Statement
}

var _ Statement = (*WhileStmt)(nil)

func (b *WhileStmt) Accept(visitor StmtVisitor) (err error) {
	return visitor.VisitWhileStmt(b)
}
