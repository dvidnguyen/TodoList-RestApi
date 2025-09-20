package common
<<<<<<< HEAD

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit > 20 {
		p.Limit = 10
	}
}
=======
>>>>>>> a0d2b0f08439de8845809c52bb1225d72ff834a3
