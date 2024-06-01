package postgres

import (
	"fmt"
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/coffee"
	"reflect"
	"testing"
)

func TestSelectCoffee(t *testing.T) {

	db, _ := NewPostgresStorage()

	req := &pb.BuyCoffee{Name: "Latte", IsPaid: true}
	resp := &pb.PreparedCoffee{Name: "Latte", Volume: 40, Price: 23000}

	coffee := NewCoffeeRepo(db.Db)
	r, _ := coffee.SelectCoffee(req)
	fmt.Println(resp, r)

	//if (resp.Name != r.Name) || (resp.Volume != r.Volume) || (resp.Price != r.Price) {
	if !reflect.DeepEqual(r, resp) {
		t.Errorf("SelectCoffee(%v) returned %v\nwant %v", req, r, resp)
	}
}
