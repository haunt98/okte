package main

type Account struct {
	Name string `json:"name"`
}

type Ticket struct {
	Name string `json:"name"`
}

type Cart struct {
	ID      int      `json:"id"`
	Account Account  `json:"account"`
	Tickets []Ticket `json:"tickets"`
	IsPaid  bool     `json:"is_paid"`
}

var acc1 = Account{
	Name: "van",
}

var acc2 = Account{
	Name: "nam",
}

var ticket1 = Ticket{
	Name: "god_of_war",
}

var ticket2 = Ticket{
	Name: "king_of_man",
}

var ticket3 = Ticket{
	Name: "dance_with_weed",
}

var cart1 = Cart{
	ID:      1,
	Account: acc1,
	Tickets: []Ticket{
		ticket1,
		ticket2,
	},
	IsPaid: false,
}

var cart2 = Cart{
	ID:      2,
	Account: acc2,
	Tickets: []Ticket{
		ticket3,
	},
	IsPaid: false,
}

var accounts = map[int]Account{
	1: acc1,
	2: acc2,
}

var tickets = map[int]Ticket{
	1: ticket1,
	2: ticket2,
}

var carts = map[int]Cart{
	1: cart1,
	2: cart2,
}
