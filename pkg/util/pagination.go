package util

import "github.com/devanfer02/ratemyubprof/internal/dto"

func GetPagination(items, limit, page uint) (dto.PaginationResponse) {
	var (
		response = dto.PaginationResponse{}
	)

  if page == 0 || limit == 0 {
    return response 
  }

	if limit > 0 {
		response.TotalPages = (items + limit - 1) / limit
	}

	if page < response.TotalPages {
		response.Next = page + 1
	} else {
		response.Next = response.TotalPages
	}

	if page > 1 {
		response.Prev = page - 1
	} else {
		response.Prev = 1
	}

	response.Current = page
	response.TotalItems = items

	return response

}

