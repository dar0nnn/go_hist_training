package main

import (
	"context"
	"fmt"
	mongoutils "mongo_hist/mongoutils"
	"reflect"
)

// func get_connection() {
// 	client, err := mongoutils.Create_client_with_conn()
// 	if err != nil {
// 		fmt.Println("fuck")
// 	}
// 	fmt.Println(mongoutils.Check_conn(client))
// 	err = client.Disconnect(context.TODO())
// }
func main() {
	client, err := mongoutils.Create_client_with_conn()
	if err != nil {
		fmt.Println("fuck")
	}
	fmt.Println(mongoutils.Check_conn(client))
	data := mongoutils.Get_data_for_hist(client, "5e3869f0294b4d34a1043e65")
	// for_print := data[0].project_id
	fmt.Println(len(data))
	fmt.Println(reflect.TypeOf(data))
	err = client.Disconnect(context.TODO())

}
