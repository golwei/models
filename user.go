
type User struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age,text,年龄："`
	Sex   string
	Intro string `form:",textarea"`
}
