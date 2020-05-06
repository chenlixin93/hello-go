package models

import (
	"time"
)

type User struct {
	Id		  int
	Uuid	  string
	Name	  string
	Email	  string
	Password  string
	CreatedAt time.Time
}

// 为已存在的用户创建新的 session
func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values (?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, user.Email, user.Id, time.Now())

	stmtout, err := Db.Prepare("select id, uuid, email, user_id, created_at from sessions where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	//通过 QueryRow 查询数据，Scan过滤数据
	err = stmtout.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// 获取已存在的用户的 session
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)

	return
}

// 创建新的用户，并保存到数据库中
func (user *User) Create() (err error) {
	// 默认不会返回最新插入的 id
	statement := "insert into users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now())

	stmtout, err := Db.Prepare("select id, uuid, created_at from users where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	// 通过 QueryRow 查询数据，Scan过滤数据
	err = stmtout.QueryRow(uuid).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

// 删除用户
func (user *User) Delete() (err error) {
	statement := "delete from users where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

// 更新用户信息
func (user *User) Update() (err error) {
	statement := "update users set name = ?, email = ? where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Id)
	return
}

// 删除所有用户
func UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

// 获取所有用户，返回数据
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// 通过邮箱获取单个用户
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// 通过 UUID 获取单个用户
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = ?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	statement := "insert into threads (uuid, topic, user_id, created_at) values (?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, topic, user.Id, time.Now())

	stmtout, err := Db.Prepare("select id, uuid, topic, user_id, created_at from threads where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtout.QueryRow(uuid).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// Create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values (?, ?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, body, user.Id, conv.Id, time.Now())

	stmtout, err := Db.Prepare("select id, uuid, body, user_id, thread_id, created_at from posts where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtout.QueryRow(uuid).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}
