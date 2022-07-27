package entity

type Product struct{
	ID		int
	Nama	string	
	Price	int
	Stok	int
}	

func (p Product) StatusStok() string{
	var status string
	
	if p.Stok < 3 {
		status = "Stok hampir habis!"
	}else if p.Stok < 10{
		status = "Stok terbatas1"
	}

	return status
}