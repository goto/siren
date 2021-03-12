package rules

import (
	"context"
	"encoding/json"
	"fmt"
	cortexClient "github.com/grafana/cortex-tools/pkg/client"
	"github.com/grafana/cortex-tools/pkg/rules/rwrulefmt"
	"github.com/odpf/siren/templates"
	"github.com/prometheus/prometheus/pkg/rulefmt"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

const (
	namePrefix = "siren_api"
)

type variable struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type Variables struct {
	Variables []variable `json:"variables"`
}

// Repository talks to the store to read or insert data
type Repository struct {
	db     *gorm.DB
	client *cortexClient.CortexClient
}

// NewRepository returns repository struct
func NewRepository(db *gorm.DB) *Repository {
	cfg := cortexClient.Config{
		Address:         "http://localhost:8080",
		UseLegacyRoutes: true,
	}
	client, err := cortexClient.New(cfg)
	if err != nil {
		return nil
	}
	return &Repository{db: db, client: client}
}

func (r Repository) Migrate() error {
	err := r.db.AutoMigrate(&Rule{})
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) Upsert(rule *Rule) (*Rule, error) {
	rule.Name = fmt.Sprintf("%s_%s_%s_%s_%s", namePrefix,
		rule.Entity, rule.Namespace, rule.GroupName, rule.Template)
	var existingRule Rule
	var rulesWithinGroup []Rule
	err := r.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Where(fmt.Sprintf("name = '%s'", rule.Name)).Find(&existingRule)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			result = tx.Create(rule)
		} else {
			result = tx.Where("id = ?", existingRule.ID).Updates(rule)
		}
		if result.Error != nil {
			return result.Error
		}
		result = tx.Where(fmt.Sprintf("name = '%s'", rule.Name)).Find(&existingRule)
		result = tx.Where(fmt.Sprintf("namespace = '%s' AND entity = '%s' AND group_name = '%s'",
			rule.Namespace, rule.Entity, rule.GroupName)).Find(&rulesWithinGroup)
		if result.Error != nil {
			return result.Error
		}
		renderedBodyForThisGroup := ""
		for i := 0; i < len(rulesWithinGroup); i++ {
			if rulesWithinGroup[i].Status == "disabled" {
				continue
			}
			inputValue := make(map[string]string)
			var variables []variable
			jsonBlob := []byte(rulesWithinGroup[i].Variables)
			_ = json.Unmarshal(jsonBlob, &variables)
			for _, v := range variables {
				inputValue[v.Name] = v.Value
			}
			service := templates.NewService(r.db)
			renderedBody, err := service.Render(rulesWithinGroup[i].Template, inputValue)
			if err != nil {
				return nil
			}
			renderedBodyForThisGroup += renderedBody
		}
		ctx := cortexClient.NewContextWithTenantId(context.Background(), rule.Entity)
		if renderedBodyForThisGroup == "" {
			//all alerts disabled in this group so we can delete the rule group
			err := r.client.DeleteRuleGroup(ctx, rule.Namespace, rule.GroupName)
			if err != nil {
				return err
			}
			return nil
		}
		var ruleNodes []rulefmt.RuleNode
		err := yaml.Unmarshal([]byte(renderedBodyForThisGroup), &ruleNodes)
		if err != nil {
			fmt.Println(err)
		}
		y := rwrulefmt.RuleGroup{
			RuleGroup: rulefmt.RuleGroup{
				Name:  rule.GroupName,
				Rules: ruleNodes,
			},
		}
		err = r.client.CreateRuleGroup(ctx, rule.Namespace, y)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &existingRule, err
}

func (r Repository) Get(namespace string, entity string, groupName string, status string, template string) ([]Rule, error) {
	var rules []Rule
	selectQuery := `SELECT * from rules`
	selectQueryWithWhereClause := `SELECT * from rules WHERE `
	var filterConditions []string
	if namespace != "" {
		filterConditions = append(filterConditions, fmt.Sprintf("namespace = '%s' ", namespace))
	}
	if entity != "" {
		filterConditions = append(filterConditions, fmt.Sprintf("entity = '%s' ", entity))
	}
	if groupName != "" {
		filterConditions = append(filterConditions, fmt.Sprintf("group_name = '%s' ", groupName))
	}
	if status != "" {
		filterConditions = append(filterConditions, fmt.Sprintf("status = '%s' ", status))
	}
	if template != "" {
		filterConditions = append(filterConditions, fmt.Sprintf("template = '%s' ", template))
	}
	var finalSelectQuery string
	if len(filterConditions) == 0 {
		finalSelectQuery = selectQuery
	} else {
		finalSelectQuery = selectQueryWithWhereClause
		for i := 0; i < len(filterConditions); i++ {
			if i == 0 {
				finalSelectQuery += filterConditions[i]
			} else {
				finalSelectQuery += " AND " + filterConditions[i]
			}
		}
	}
	result := r.db.Raw(finalSelectQuery).Scan(&rules)
	if result.Error != nil {
		return nil, result.Error
	}
	return rules, nil
}
