package rule

import (
	"fmt"
	"go/ast"
	"go/types"
	"sync"

	"github.com/mgechev/revive/lint"
)

const (
	ruleUTAMessagePanic   = "type assertion will panic if not matched"
	ruleUTAMessageIgnored = "type assertion result ignored"
)

// UncheckedTypeAssertionRule lints missing or ignored `ok`-value in danymic type casts.
type UncheckedTypeAssertionRule struct {
	sync.Mutex
	acceptIgnoredAssertionResult bool
}

func (u *UncheckedTypeAssertionRule) configure(arguments lint.Arguments) {
	u.Lock()
	defer u.Unlock()

	if len(arguments) == 0 {
		return
	}

	args, ok := arguments[0].(map[string]any)
	if !ok {
		panic("Unable to get arguments. Expected object of key-value-pairs.")
	}

	for k, v := range args {
		switch k {
		case "acceptIgnoredAssertionResult":
			u.acceptIgnoredAssertionResult, _ = v.(bool)
		default:
			panic(fmt.Sprintf("Unknown argument: %s", k))
		}
	}
}

// Apply applies the rule to given file.
func (u *UncheckedTypeAssertionRule) Apply(file *lint.File, args lint.Arguments) []lint.Failure {
	u.configure(args)

	var failures []lint.Failure

	walker := &lintUnchekedTypeAssertion{
		pkg: file.Pkg,
		onFailure: func(failure lint.Failure) {
			failures = append(failures, failure)
		},
		acceptIgnoredTypeAssertionResult: u.acceptIgnoredAssertionResult,
	}

	file.Pkg.TypeCheck()
	ast.Walk(walker, file.AST)

	return failures
}

// Name returns the rule name.
func (*UncheckedTypeAssertionRule) Name() string {
	return "unchecked-type-assertion"
}

type lintUnchekedTypeAssertion struct {
	pkg                              *lint.Package
	onFailure                        func(lint.Failure)
	acceptIgnoredTypeAssertionResult bool
}

func (w *lintUnchekedTypeAssertion) isBoolType(e ast.Expr) bool {
	t := w.pkg.TypeOf(e)
	ut, ok := t.Underlying().(*types.Basic)
	return (ok && ut != nil) && ut.Info()&types.IsBoolean == 0
}

func isTypeSwitch(e *ast.TypeAssertExpr) bool {
	return e.Type == nil
}

func (w *lintUnchekedTypeAssertion) requireNoTypeAssert(expr ast.Expr) {
	e, ok := expr.(*ast.TypeAssertExpr)
	if ok && !isTypeSwitch(e) {
		w.addFailure(e, ruleUTAMessagePanic)
	}
}

func (w *lintUnchekedTypeAssertion) handleSwitch(n *ast.SwitchStmt) {
	w.requireNoTypeAssert(n.Tag)
}

func (w *lintUnchekedTypeAssertion) handleAssignment(n *ast.AssignStmt) {
	if len(n.Rhs) == 0 {
		return
	}

	e, ok := n.Rhs[0].(*ast.TypeAssertExpr)
	if !ok || e == nil {
		return
	}

	if isTypeSwitch(e) {
		return
	}

	if len(n.Lhs) == 1 {
		w.addFailure(e, ruleUTAMessagePanic)
	}

	if !w.acceptIgnoredTypeAssertionResult && len(n.Lhs) == 2 && !w.isBoolType(n.Lhs[1]) {
		w.addFailure(e, ruleUTAMessageIgnored)
	}
}

// handles "return foo(.*bar)" - one of them is enough to fail as golang does not forward the type cast tuples in return statements
func (w *lintUnchekedTypeAssertion) handleReturn(n *ast.ReturnStmt) {
	for _, r := range n.Results {
		w.requireNoTypeAssert(r)
	}
}

func (w *lintUnchekedTypeAssertion) handleRange(n *ast.RangeStmt) {
	w.requireNoTypeAssert(n.X)
}

func (w *lintUnchekedTypeAssertion) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.RangeStmt:
		w.handleRange(n)
	case *ast.SwitchStmt:
		w.handleSwitch(n)
	case *ast.ReturnStmt:
		w.handleReturn(n)
	case *ast.AssignStmt:
		w.handleAssignment(n)
	}

	return w
}

func (w *lintUnchekedTypeAssertion) addFailure(n *ast.TypeAssertExpr, why string) {
	s := fmt.Sprintf("type cast result is unchecked in %v - %s", gofmt(n), why)
	w.onFailure(lint.Failure{
		Category:   "bad practice",
		Confidence: 1,
		Node:       n,
		Failure:    s,
	})
}