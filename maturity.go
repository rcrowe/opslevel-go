package opslevel

type ServiceMaturity struct {
	Name           string
	MaturityReport struct {
		CategoryBreakdown []struct {
			Category struct {
				Name string
			}
			Level struct {
				Name string
			}
		}
		OverallLevel struct {
			Name string
		}
	}
}

func (s *ServiceMaturity) GetValues(fields ...string) []string {
	var output []string
	for _, field := range fields {
		if field == "Name" {
			output = append(output, s.Name)
			continue
		}
		if field == "Overall" {
			output = append(output, s.MaturityReport.OverallLevel.Name)
			continue
		}
		for _, breakdown := range s.MaturityReport.CategoryBreakdown {
			if field == breakdown.Category.Name {
				output = append(output, breakdown.Level.Name)
				break
			}
		}
	}
	return output
}

func (c *Client) ListServicesMaturity() ([]ServiceMaturity, error) {
	var q struct {
		Account struct {
			Services struct {
				Nodes    []ServiceMaturity
				PageInfo PageInfo
			} `graphql:"services(after: $after, first: $first)"`
		}
	}
	v := c.InitialPageVariables()

	var output []ServiceMaturity
	if err := c.Query(&q, v); err != nil {
		return nil, err
	}
	output = append(output, q.Account.Services.Nodes...)
	for q.Account.Services.PageInfo.HasNextPage {
		v["after"] = q.Account.Services.PageInfo.End
		if err := c.Query(&q, v); err != nil {
			return nil, err
		}
		output = append(output, q.Account.Services.Nodes...)
	}

	return output, nil
}
