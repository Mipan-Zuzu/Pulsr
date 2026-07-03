package checker

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"gorm.io/gorm"
// )
// var results []map[string]interface{}

// func SupabaseProjectAll(authToken string, db *gorm.DB) (results, error)  {
// 	req, err := http.Get(fmt.Sprintf("%s/projects", os.Getenv("SUPABASE_URL")))
// 	if err != nil {
// 		return nil,err
// 	}
	
// 	req.Header.Set("Authorization", authToken)
// 		return resu
// }