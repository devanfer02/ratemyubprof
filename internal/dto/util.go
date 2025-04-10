package dto

func (pageQuery *PaginationQuery) SetDefaultValue() {
	if pageQuery.Limit == 0 {
		pageQuery.Limit = 10
	}
	if pageQuery.Page == 0 {
		pageQuery.Page = 1
	}	
}