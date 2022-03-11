package SQL

import (
	"database/sql"
	_ "database/sql/driver"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Sex       bool   `json:"sex"`
	Status    bool   `json:"status"`
	Title     string `json:"title"`
}

func InitMysql() *sql.DB {
	DB, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Fatal(err)
	}
	return DB
}

// Query // QueryRow 查询单条数据示例
//func QueryRow(db *sql.DB, userid int) {
//	sqlStr := "select id, name, age from bubble where id=?"
//	var u User
//	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
//	err := db.QueryRow(sqlStr, userid).Scan(&u.ID, &u.Title, &u.Status) //单行查询用QueryRow
//	//rows,err := db.Query(sqlStr, "xys")//多行查询用Query
//	if err != nil {
//		fmt.Printf("scan failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("ID:%d Title:%s Status:%d\n", u.ID, u.Title, u.Status)
//}
//func Query(db *sql.DB) {
//	sqlStr := "select * from bubble;"
//	var u User
//	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
//	//err := db.QueryRow(sqlStr, userid).Scan(&u.ID, &u.Title, &u.Status) //单行查询用QueryRow
//	rows, err := db.Query(sqlStr, "xys") //多行查询用Query
//	if err != nil {
//		fmt.Printf("scan failed, err:%v\n", err)
//		return
//	}
//	for rows.Next() {
//		err=rows.Scan(&u.ID, &u.Title, &u.Status)
//		if err != nil {
//			log.Fatal(err)
//		}
//		log.Printf("get data, ID: %d, Title: %s, Status: %d", u.ID, u.Title, u.Status)
//	}
//
//	fmt.Printf("ID:%d Title:%s Status:%d\n", u.ID, u.Title, u.Status)
//}

func CreatAccount(db *sql.DB, user *User) uint8 {
	//对
	Password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("bcrypt.GenerateFromPassword 加密失败！")
		//加密失败返回150
		return 150
	}
	_, err = db.Exec(`INSERT INTO user ( name, password,telephone,sex,email) VALUES ( ?, ?,?,?,?);`, user.Name, Password, user.Telephone, user.Sex, user.Email)
	//defer rows.Close()
	if err != nil {
		log.Printf("insert data error: %v\n", err)
		//用户存在返回100
		return 100
	}
	//var result int
	//rows.Scan(&result)
	log.Printf("insert is Ok!")
	return 200
}

//func DeleteDB(db *sql.DB, ID string) {
//	_, err := db.Exec(`delete from bubble where id=?;`, ID)
//	//defer rows.Close()
//	if err != nil {
//		log.Fatalf("Delete data error: %v\n", err)
//	}
//	//var result int
//	//rows.Scan(&result)
//	log.Printf("Delete Ok!")
//}
//
//func FindDB(db *sql.DB, todoList []User) []User {
//	sqlStr := "select * from bubble "
//	var u User
//	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
//	//err := db.QueryRow(sqlStr, userid).Scan(&u.ID, &u.Title, &u.Status) //单行查询用QueryRow
//	rows, err := db.Query(sqlStr) //多行查询用Query
//	if err != nil {
//		fmt.Printf("scan failed, err:%v\n", err)
//		return make([]User, 0)
//	}
//	for rows.Next() {
//		err := rows.Scan(&u.ID, &u.Title, &u.Status)
//		if err != nil {
//			log.Fatal(err)
//		}
//		todoList = append(todoList, u)
//
//		log.Printf("get data, ID: %d, Title: %s, Status: %d", u.ID, u.Title, u.Status)
//	}
//	return todoList
//}

func TodoIsOK(user *User) (bool, string) {
	//匹配电子邮箱
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(user.Email) {
		return false, "邮箱格式错误"
	}
	//8位长度密码
	pattern = `^[a-zA-Z]\w{5,17}$` //匹配密码
	reg = regexp.MustCompile(pattern)
	if !reg.MatchString(user.Password) {
		return false, "密码格式错误"
	}
	log.Printf("%v", len(user.Name))
	if len(user.Name) <= 3 || len(user.Name) >= 32 {
		return false, "昵称太长！"
	}
	return true, "格式正确!"
}
