package casbin_test

import (
	"fmt"
	casbin "gorm-casbin"
	"testing"
)

func TestNewEnforcer(t *testing.T) {
	enforcer, err := casbin.NewEnforcer(&casbin.Options{
		Model:    "./model.conf",
		Debug:    true,
		Enable:   true,
		Autoload: true,
		Table:    "casbin_policy_test",
		Database: "root:123456@tcp(192.168.8.173:3306)/backstage?charset=utf8mb4&parseTime=True&loc=Local",
	})
	if err != nil {
		t.Fatal(err)
	}

	// add a permission node for role
	ok, err := enforcer.AddPolicy("role_1", "node_1")
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		fmt.Println("success: add a permission node for role")
	} else {
		fmt.Println("failed: add a permission node for role")
	}

	// batch add permission nodes for roles
	ok, err = enforcer.AddPolicies([][]string{
		{"role_2", "node_2"},
		{"role_3", "node_3"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		fmt.Println("success: batch add permission nodes for roles")
	} else {
		fmt.Println("failed: batch add permission nodes for roles")
	}

	// add a role for user
	ok, err = enforcer.AddGroupingPolicy("user_1", "role_1")
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		fmt.Println("success: add a role for user")
	} else {
		fmt.Println("failed: add a role for user")
	}

	// batch add roles for users
	ok, err = enforcer.AddGroupingPolicies([][]string{
		{"user_2", "role_2"},
		{"user_3", "role_3"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		fmt.Println("success: batch add roles for users")
	} else {
		fmt.Println("failed: batch add roles for users")
	}

	// check role_1 policy
	if ok = enforcer.HasPolicy("role_1", "node_1"); ok {
		fmt.Println("role_1 is allowed access node_1")
	} else {
		fmt.Println("role_1 is not allowed access node_1")
	}

	// check role_1 policy
	if ok = enforcer.HasPolicy("role_1", "node_2"); ok {
		fmt.Println("role_1 is allowed access node_2")
	} else {
		fmt.Println("role_1 is not allowed access node_2")
	}

	// check user_1 policy
	if ok = enforcer.HasGroupingPolicy("user_1", "role_1"); ok {
		fmt.Println("user_1 has role_1")
	} else {
		fmt.Println("user_1 has not role_1")
	}

	// check user_1 policy
	if ok = enforcer.HasGroupingPolicy("user_1", "role_2"); ok {
		fmt.Println("user_1 has role_2")
	} else {
		fmt.Println("user_1 has not role_2")
	}

	// check access permission of user_1
	if ok, _ = enforcer.Enforce("user_1", "node_1"); ok {
		fmt.Println("user_1 is allowed access node_1")
	} else {
		fmt.Println("user_1 is not allowed access node_1")
	}

	// check access permission of user_1
	if ok, _ = enforcer.Enforce("user_1", "node_2"); ok {
		fmt.Println("user_1 is allowed access node_2")
	} else {
		fmt.Println("user_1 is not allowed access node_2")
	}

	// remove a policy
	ok, err = enforcer.RemovePolicy("role_1", "node_1")
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		fmt.Println("success: remove a policy")
	} else {
		fmt.Println("failed: remove a policy")
	}

	// get all policies
	policies := enforcer.GetPolicy()
	fmt.Println()
	fmt.Println("all policies:")
	fmt.Println(policies)

	// batch remove policies
	ok, err = enforcer.RemovePolicies([][]string{
		{"role_2", "node_2"},
		{"role_3", "node_3"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		fmt.Println("success: batch remove policies")
	} else {
		fmt.Println("failed: batch remove policies")
	}

	// get all policies
	policies = enforcer.GetPolicy()
	fmt.Println()
	fmt.Println("all policies:")
	fmt.Println(policies)

	// remove a grouping policy
	ok, err = enforcer.RemoveGroupingPolicy("user_1", "role_1")
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		fmt.Println("success: remove a grouping policy")
	} else {
		fmt.Println("failed: remove a grouping policy")
	}

	// get all grouping policies
	policies = enforcer.GetGroupingPolicy()
	fmt.Println()
	fmt.Println("all grouping policies:")
	fmt.Println(policies)

	// batch remove grouping policies
	ok, err = enforcer.RemoveGroupingPolicies([][]string{
		{"user_2", "role_2"},
		{"user_3", "role_3"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		fmt.Println("success: batch remove grouping policies")
	} else {
		fmt.Println("failed: batch remove grouping policies")
	}

	// get all grouping policies
	policies = enforcer.GetGroupingPolicy()
	fmt.Println()
	fmt.Println("all grouping policies:")
	fmt.Println(policies)
	fmt.Println()

	// check role_1 policy
	if ok = enforcer.HasPolicy("role_1", "node_1"); ok {
		fmt.Println("role_1 is allowed access node_1")
	} else {
		fmt.Println("role_1 is not allowed access node_1")
	}

	// check role_1 policy
	if ok = enforcer.HasPolicy("role_1", "node_2"); ok {
		fmt.Println("role_1 is allowed access node_2")
	} else {
		fmt.Println("role_1 is not allowed access node_2")
	}

	// check user_1 policy
	if ok = enforcer.HasGroupingPolicy("user_1", "role_1"); ok {
		fmt.Println("user_1 has role_1")
	} else {
		fmt.Println("user_1 has not role_1")
	}

	// check user_1 policy
	if ok = enforcer.HasGroupingPolicy("user_1", "role_2"); ok {
		fmt.Println("user_1 has role_2")
	} else {
		fmt.Println("user_1 has not role_2")
	}
}
