// Copyright 2016, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package planbuilder

import (
	"errors"
	"strconv"

	"github.com/youtube/vitess/go/vt/sqlparser"
)

func processHaving(having *sqlparser.Where, symbolTable *SymbolTable) error {
	if having == nil {
		return nil
	}
	for _, filter := range splitAndExpression(nil, having.Expr) {
		var routeBuilder *RouteBuilder
		err := sqlparser.Walk(func(node sqlparser.SQLNode) (kontinue bool, err error) {
			switch node := node.(type) {
			case *sqlparser.Subquery:
				// TODO(sougou): better error.
				return false, errors.New("subqueries not supported")
			case *sqlparser.ColName:
				newRoute, _ := symbolTable.FindColumn(node, nil, true)
				if newRoute != nil {
					if routeBuilder == nil {
						routeBuilder = newRoute
					} else if routeBuilder != newRoute {
						// TODO(sougou): better error.
						return false, errors.New("having clause is too complex")
					}
				}
			}
			return true, nil
		}, filter)
		if err != nil {
			return err
		}
		if routeBuilder == nil {
			routeBuilder = symbolTable.FirstRoute
		}
		routeBuilder.Select.AddHaving(filter)
	}
	return nil
}

func processOrderBy(orderBy sqlparser.OrderBy, symbolTable *SymbolTable) error {
	if orderBy == nil {
		return nil
	}
	routeNumber := 0
	for _, order := range orderBy {
		var route *RouteBuilder
		switch node := order.Expr.(type) {
		case *sqlparser.ColName:
			route, _ = symbolTable.FindColumn(node, nil, true)
		case sqlparser.NumVal:
			num, err := strconv.ParseInt(string(node), 0, 64)
			if err != nil {
				// TODO(sougou): better error.
				return errors.New("error parsing order by clause")
			}
			if num < 1 || num > int64(len(symbolTable.SelectSymbols)) {
				// TODO(sougou): better error.
				return errors.New("order by column number out of range")
			}
			route = symbolTable.SelectSymbols[num-1].Route
		default:
			// TODO(sougou): better error.
			return errors.New("order by clause is too complex")
		}
		if route == nil || route.Order() < routeNumber {
			// TODO(sougou): better error.
			return errors.New("order by clause is too complex")
		}
		routeNumber = route.Order()
		route.Select.OrderBy = append(route.Select.OrderBy, order)
	}
	return nil
}

func processLimit(limit *sqlparser.Limit, planBuilder PlanBuilder) error {
	if limit == nil {
		return nil
	}
	routeBuilder, ok := planBuilder.(*RouteBuilder)
	if !ok {
		return errors.New("query is too complex to allow limits")
	}
	if !routeBuilder.IsSingle() {
		return errors.New("query is too complex to allow limits")
	}
	routeBuilder.Select.Limit = limit
	return nil
}
