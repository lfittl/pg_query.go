// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 *		REASSIGN OWNED statement
 */
type ReassignOwnedStmt struct {
	Roles   List    `json:"roles"`
	Newrole *string `json:"newrole"`
}

func (node ReassignOwnedStmt) MarshalJSON() ([]byte, error) {
	type ReassignOwnedStmtMarshalAlias ReassignOwnedStmt
	return json.Marshal(map[string]interface{}{
		"ReassignOwnedStmt": (*ReassignOwnedStmtMarshalAlias)(&node),
	})
}

func (node *ReassignOwnedStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["roles"] != nil {
		node.Roles.Items, err = UnmarshalNodeArrayJSON(fields["roles"])
		if err != nil {
			return
		}
	}

	if fields["newrole"] != nil {
		err = json.Unmarshal(fields["newrole"], &node.Newrole)
		if err != nil {
			return
		}
	}

	return
}