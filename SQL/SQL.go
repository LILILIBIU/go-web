package SQL

import (
	_ "database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

type User struct {
	ID        uint        `json:"-"`
	Name      string      `json:"name"`
	Password  string      `json:"password"`
	Email     string      `json:"email"`
	Telephone string      `json:"telephone"`
	Sex       bool        `json:"sex"`
	Token     string      `json:"token"`
	Status    bool        `json:"status"`
	LiveID    string      `json:"liveid"`
	Title     string      `json:"title"`
	Ch        chan string `json:"-"`
}

var DB *sqlx.DB

// Init 初始化MySQL连接
func Init() (err error) {
	username := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.ip")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", username, password, host, port, database)
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(5)
	return
}

// Close 关闭MySQL连接
func Close() {
	_ = DB.Close()
}

func Query(name string) *User {
	u := User{}
	sqlStr := fmt.Sprintf("select * from user where name='%s'", name)
	//sqlStr = fmt.Sprint(sqlStr, name)
	//name = "'" + name + "'"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := DB.Get(&u, sqlStr)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}
	fmt.Printf("id:%v name:%v password:%v\n", u.ID, u.Name, u.Password)
	return &u
}

//func QueryID(name uint) *User {
//	u := User{}
//	sqlStr := fmt.Sprintf("select id, name, password from user where id=%s", name)
//	err := DB.Get(&u, sqlStr)
//	if err != nil {
//		fmt.Printf("scan failed, err:%v\n", err)
//	}
//	fmt.Printf("id:%v name:%v password:%v\n", u.ID, u.Name, u.Password)
//	return &u
//}

func CreatAccount(user *User) uint8 {
	//对
	Password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("bcrypt.GenerateFromPassword 加密失败！")
		//加密失败返回150
		return 150
	}
	_, err = DB.Exec(`INSERT INTO user ( name, password,telephone,sex,email) VALUES ( ?, ?,?,?,?);`, user.Name, Password, user.Telephone, user.Sex, user.Email)
	//defer rows.Close()
	if err != nil {
		log.Printf("insert data error: %v\n", err)
		//用户存在返回100
		return 100
	}
	//var result int
	//rows.Scan(&result)
	log.Printf("insert is Ok!")
	//DB.Close()
	return 200
}

//func DeleteDB(DB *sql.DB, ID string) {
//	_, err := DB.Exec(`delete from bubble where id=?;`, ID)
//	//defer rows.Close()
//	if err != nil {
//		log.Fatalf("Delete data error: %v\n", err)
//	}
//	//var result int
//	//rows.Scan(&result)
//	log.Printf("Delete Ok!")
//}

func ListIsOK(user *User, flag bool) (bool, string) {
	//匹配电子邮箱
	if flag {
		pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
		reg := regexp.MustCompile(pattern)
		if !reg.MatchString(user.Email) {
			return false, "邮箱格式错误"
		}
	}
	//8位长度密码
	password := `^[a-zA-Z]\w{5,17}$` //匹配密码
	reg := regexp.MustCompile(password)
	if !reg.MatchString(user.Password) {
		return false, "密码格式错误"
	}
	log.Printf("%v", len(user.Name))
	if len(user.Name) <= 3 || len(user.Name) >= 32 {
		return false, "昵称太长！"
	}
	return true, "格式正确!"
}
