package model

import (
	"reflect"
	"time"
)

type User struct {
	Id       int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"` // id, 全局唯一
	Username string    `gorm:"column:username" json:"username"`                // 用户名
	Nickname string    `gorm:"column:nickname" json:"nickname"`                // 昵称
	Email    string    `gorm:"column:email" json:"email"`                      // 邮箱
	Phone    string    `gorm:"column:phone" json:"phone"`                      // 电话
	Ctime    time.Time `gorm:"column:ctime" json:"ctime"`                      // 创建时间
	Mtime    time.Time `gorm:"column:mtime" json:"mtime"`                      // 更新时间
}

func (t *User) TableName() string {
	return "user"
}

func (t User) Builder() *UserBuilder {
	return &UserBuilder{user: &User{}}
}

type UserBuilder struct {
	user *User
}

func (t *UserBuilder) Build() *User {
	return t.user
}
func (t *UserBuilder) Flush() *UserBuilder {
	t.user = &User{}
	return t
}

func (t *UserBuilder) Id(id int64) *UserBuilder {
	t.user.Id = id
	return t
}
func (t *UserBuilder) Username(username string) *UserBuilder {
	t.user.Username = username
	return t
}
func (t *UserBuilder) Nickname(nickname string) *UserBuilder {
	t.user.Nickname = nickname
	return t
}
func (t *UserBuilder) Email(email string) *UserBuilder {
	t.user.Email = email
	return t
}
func (t *UserBuilder) Phone(phone string) *UserBuilder {
	t.user.Phone = phone
	return t
}
func (t *UserBuilder) Ctime(ctime time.Time) *UserBuilder {
	t.user.Ctime = ctime
	return t
}
func (t *UserBuilder) Mtime(mtime time.Time) *UserBuilder {
	t.user.Mtime = mtime
	return t
}

type UserCondition struct {
	Where  string
	Order  string
	Offset int64
	Limit  int64
	Params []interface{}
}

func (t *UserCondition) AndIdEq(id int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id = ?"
	} else {
		t.Where += " id = ?"
	}
	t.Params = append(t.Params, id)
	return t
}

func (t *UserCondition) AndIdBt(start int64, end int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id between ? and ?"
	} else {
		t.Where += " id between ? and ?"
	}
	t.Params = append(t.Params, start)
	t.Params = append(t.Params, end)
	return t
}

func (t *UserCondition) AndIdGt(id int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id > ?"
	} else {
		t.Where += " id > ?"
	}
	t.Params = append(t.Params, id)
	return t
}

func (t *UserCondition) AndIdLt(id int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id < ?"
	} else {
		t.Where += " id < ?"
	}
	t.Params = append(t.Params, id)
	return t
}

func (t *UserCondition) AndIdGtAndEq(id int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id >= ?"
	} else {
		t.Where += " id >= ?"
	}
	t.Params = append(t.Params, id)
	return t
}

func (t *UserCondition) AndIdLtAndEq(id int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id <= ?"
	} else {
		t.Where += " id <= ?"
	}
	t.Params = append(t.Params, id)
	return t
}

func (t *UserCondition) AndIdLike(id int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id like ?"
	} else {
		t.Where += " id like ?"
	}
	t.Params = append(t.Params, id)
	return t
}

func (t *UserCondition) AndIdIn(ids []int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id in(?)"
	} else {
		t.Where += " id in(?)"
	}
	t.Params = append(t.Params, ids)
	return t
}

func (t *UserCondition) AndIdNotIn(ids []int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id not in(?)"
	} else {
		t.Where += " id not in(?)"
	}
	t.Params = append(t.Params, ids)
	return t
}
func (t *UserCondition) AndIdNotLike(id int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id not like ?"
	} else {
		t.Where += " id not like ?"
	}
	t.Params = append(t.Params, id)
	return t
}
func (t *UserCondition) AndIdNotBt(start int64, end int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id not between ? and ?"
	} else {
		t.Where += " id not between ? and ?"
	}
	t.Params = append(t.Params, start, end)
	return t
}
func (t *UserCondition) AndIdNotEq(id int64) *UserCondition {
	if t.Where != "" {
		t.Where += " and id != ?"
	} else {
		t.Where += " id != ?"
	}
	t.Params = append(t.Params, id)
	return t
}

func (t *UserCondition) AndUsernameEq(username string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username = ?"
	} else {
		t.Where += " username = ?"
	}
	t.Params = append(t.Params, username)
	return t
}

func (t *UserCondition) AndUsernameBt(start string, end string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username between ? and ?"
	} else {
		t.Where += " username between ? and ?"
	}
	t.Params = append(t.Params, start)
	t.Params = append(t.Params, end)
	return t
}

func (t *UserCondition) AndUsernameGt(username string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username > ?"
	} else {
		t.Where += " username > ?"
	}
	t.Params = append(t.Params, username)
	return t
}

func (t *UserCondition) AndUsernameLt(username string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username < ?"
	} else {
		t.Where += " username < ?"
	}
	t.Params = append(t.Params, username)
	return t
}

func (t *UserCondition) AndUsernameGtAndEq(username string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username >= ?"
	} else {
		t.Where += " username >= ?"
	}
	t.Params = append(t.Params, username)
	return t
}

func (t *UserCondition) AndUsernameLtAndEq(username string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username <= ?"
	} else {
		t.Where += " username <= ?"
	}
	t.Params = append(t.Params, username)
	return t
}

func (t *UserCondition) AndUsernameLike(username string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username like ?"
	} else {
		t.Where += " username like ?"
	}
	t.Params = append(t.Params, username)
	return t
}

func (t *UserCondition) AndUsernameIn(usernames []string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username in(?)"
	} else {
		t.Where += " username in(?)"
	}
	t.Params = append(t.Params, usernames)
	return t
}

func (t *UserCondition) AndUsernameNotIn(usernames []string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username not in(?)"
	} else {
		t.Where += " username not in(?)"
	}
	t.Params = append(t.Params, usernames)
	return t
}
func (t *UserCondition) AndUsernameNotLike(username string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username not like ?"
	} else {
		t.Where += " username not like ?"
	}
	t.Params = append(t.Params, username)
	return t
}
func (t *UserCondition) AndUsernameNotBt(start string, end string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username not between ? and ?"
	} else {
		t.Where += " username not between ? and ?"
	}
	t.Params = append(t.Params, start, end)
	return t
}
func (t *UserCondition) AndUsernameNotEq(username string) *UserCondition {
	if t.Where != "" {
		t.Where += " and username != ?"
	} else {
		t.Where += " username != ?"
	}
	t.Params = append(t.Params, username)
	return t
}

func (t *UserCondition) AndNicknameEq(nickname string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname = ?"
	} else {
		t.Where += " nickname = ?"
	}
	t.Params = append(t.Params, nickname)
	return t
}

func (t *UserCondition) AndNicknameBt(start string, end string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname between ? and ?"
	} else {
		t.Where += " nickname between ? and ?"
	}
	t.Params = append(t.Params, start)
	t.Params = append(t.Params, end)
	return t
}

func (t *UserCondition) AndNicknameGt(nickname string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname > ?"
	} else {
		t.Where += " nickname > ?"
	}
	t.Params = append(t.Params, nickname)
	return t
}

func (t *UserCondition) AndNicknameLt(nickname string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname < ?"
	} else {
		t.Where += " nickname < ?"
	}
	t.Params = append(t.Params, nickname)
	return t
}

func (t *UserCondition) AndNicknameGtAndEq(nickname string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname >= ?"
	} else {
		t.Where += " nickname >= ?"
	}
	t.Params = append(t.Params, nickname)
	return t
}

func (t *UserCondition) AndNicknameLtAndEq(nickname string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname <= ?"
	} else {
		t.Where += " nickname <= ?"
	}
	t.Params = append(t.Params, nickname)
	return t
}

func (t *UserCondition) AndNicknameLike(nickname string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname like ?"
	} else {
		t.Where += " nickname like ?"
	}
	t.Params = append(t.Params, nickname)
	return t
}

func (t *UserCondition) AndNicknameIn(nicknames []string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname in(?)"
	} else {
		t.Where += " nickname in(?)"
	}
	t.Params = append(t.Params, nicknames)
	return t
}

func (t *UserCondition) AndNicknameNotIn(nicknames []string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname not in(?)"
	} else {
		t.Where += " nickname not in(?)"
	}
	t.Params = append(t.Params, nicknames)
	return t
}
func (t *UserCondition) AndNicknameNotLike(nickname string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname not like ?"
	} else {
		t.Where += " nickname not like ?"
	}
	t.Params = append(t.Params, nickname)
	return t
}
func (t *UserCondition) AndNicknameNotBt(start string, end string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname not between ? and ?"
	} else {
		t.Where += " nickname not between ? and ?"
	}
	t.Params = append(t.Params, start, end)
	return t
}
func (t *UserCondition) AndNicknameNotEq(nickname string) *UserCondition {
	if t.Where != "" {
		t.Where += " and nickname != ?"
	} else {
		t.Where += " nickname != ?"
	}
	t.Params = append(t.Params, nickname)
	return t
}

func (t *UserCondition) AndEmailEq(email string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email = ?"
	} else {
		t.Where += " email = ?"
	}
	t.Params = append(t.Params, email)
	return t
}

func (t *UserCondition) AndEmailBt(start string, end string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email between ? and ?"
	} else {
		t.Where += " email between ? and ?"
	}
	t.Params = append(t.Params, start)
	t.Params = append(t.Params, end)
	return t
}

func (t *UserCondition) AndEmailGt(email string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email > ?"
	} else {
		t.Where += " email > ?"
	}
	t.Params = append(t.Params, email)
	return t
}

func (t *UserCondition) AndEmailLt(email string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email < ?"
	} else {
		t.Where += " email < ?"
	}
	t.Params = append(t.Params, email)
	return t
}

func (t *UserCondition) AndEmailGtAndEq(email string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email >= ?"
	} else {
		t.Where += " email >= ?"
	}
	t.Params = append(t.Params, email)
	return t
}

func (t *UserCondition) AndEmailLtAndEq(email string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email <= ?"
	} else {
		t.Where += " email <= ?"
	}
	t.Params = append(t.Params, email)
	return t
}

func (t *UserCondition) AndEmailLike(email string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email like ?"
	} else {
		t.Where += " email like ?"
	}
	t.Params = append(t.Params, email)
	return t
}

func (t *UserCondition) AndEmailIn(emails []string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email in(?)"
	} else {
		t.Where += " email in(?)"
	}
	t.Params = append(t.Params, emails)
	return t
}

func (t *UserCondition) AndEmailNotIn(emails []string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email not in(?)"
	} else {
		t.Where += " email not in(?)"
	}
	t.Params = append(t.Params, emails)
	return t
}
func (t *UserCondition) AndEmailNotLike(email string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email not like ?"
	} else {
		t.Where += " email not like ?"
	}
	t.Params = append(t.Params, email)
	return t
}
func (t *UserCondition) AndEmailNotBt(start string, end string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email not between ? and ?"
	} else {
		t.Where += " email not between ? and ?"
	}
	t.Params = append(t.Params, start, end)
	return t
}
func (t *UserCondition) AndEmailNotEq(email string) *UserCondition {
	if t.Where != "" {
		t.Where += " and email != ?"
	} else {
		t.Where += " email != ?"
	}
	t.Params = append(t.Params, email)
	return t
}

func (t *UserCondition) AndPhoneEq(phone string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone = ?"
	} else {
		t.Where += " phone = ?"
	}
	t.Params = append(t.Params, phone)
	return t
}

func (t *UserCondition) AndPhoneBt(start string, end string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone between ? and ?"
	} else {
		t.Where += " phone between ? and ?"
	}
	t.Params = append(t.Params, start)
	t.Params = append(t.Params, end)
	return t
}

func (t *UserCondition) AndPhoneGt(phone string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone > ?"
	} else {
		t.Where += " phone > ?"
	}
	t.Params = append(t.Params, phone)
	return t
}

func (t *UserCondition) AndPhoneLt(phone string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone < ?"
	} else {
		t.Where += " phone < ?"
	}
	t.Params = append(t.Params, phone)
	return t
}

func (t *UserCondition) AndPhoneGtAndEq(phone string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone >= ?"
	} else {
		t.Where += " phone >= ?"
	}
	t.Params = append(t.Params, phone)
	return t
}

func (t *UserCondition) AndPhoneLtAndEq(phone string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone <= ?"
	} else {
		t.Where += " phone <= ?"
	}
	t.Params = append(t.Params, phone)
	return t
}

func (t *UserCondition) AndPhoneLike(phone string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone like ?"
	} else {
		t.Where += " phone like ?"
	}
	t.Params = append(t.Params, phone)
	return t
}

func (t *UserCondition) AndPhoneIn(phones []string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone in(?)"
	} else {
		t.Where += " phone in(?)"
	}
	t.Params = append(t.Params, phones)
	return t
}

func (t *UserCondition) AndPhoneNotIn(phones []string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone not in(?)"
	} else {
		t.Where += " phone not in(?)"
	}
	t.Params = append(t.Params, phones)
	return t
}
func (t *UserCondition) AndPhoneNotLike(phone string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone not like ?"
	} else {
		t.Where += " phone not like ?"
	}
	t.Params = append(t.Params, phone)
	return t
}
func (t *UserCondition) AndPhoneNotBt(start string, end string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone not between ? and ?"
	} else {
		t.Where += " phone not between ? and ?"
	}
	t.Params = append(t.Params, start, end)
	return t
}
func (t *UserCondition) AndPhoneNotEq(phone string) *UserCondition {
	if t.Where != "" {
		t.Where += " and phone != ?"
	} else {
		t.Where += " phone != ?"
	}
	t.Params = append(t.Params, phone)
	return t
}

func (t *UserCondition) AndCtimeEq(ctime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime = ?"
	} else {
		t.Where += " ctime = ?"
	}
	t.Params = append(t.Params, ctime)
	return t
}

func (t *UserCondition) AndCtimeBt(start time.Time, end time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime between ? and ?"
	} else {
		t.Where += " ctime between ? and ?"
	}
	t.Params = append(t.Params, start)
	t.Params = append(t.Params, end)
	return t
}

func (t *UserCondition) AndCtimeGt(ctime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime > ?"
	} else {
		t.Where += " ctime > ?"
	}
	t.Params = append(t.Params, ctime)
	return t
}

func (t *UserCondition) AndCtimeLt(ctime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime < ?"
	} else {
		t.Where += " ctime < ?"
	}
	t.Params = append(t.Params, ctime)
	return t
}

func (t *UserCondition) AndCtimeGtAndEq(ctime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime >= ?"
	} else {
		t.Where += " ctime >= ?"
	}
	t.Params = append(t.Params, ctime)
	return t
}

func (t *UserCondition) AndCtimeLtAndEq(ctime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime <= ?"
	} else {
		t.Where += " ctime <= ?"
	}
	t.Params = append(t.Params, ctime)
	return t
}

func (t *UserCondition) AndCtimeLike(ctime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime like ?"
	} else {
		t.Where += " ctime like ?"
	}
	t.Params = append(t.Params, ctime)
	return t
}

func (t *UserCondition) AndCtimeIn(ctimes []time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime in(?)"
	} else {
		t.Where += " ctime in(?)"
	}
	t.Params = append(t.Params, ctimes)
	return t
}

func (t *UserCondition) AndCtimeNotIn(ctimes []time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime not in(?)"
	} else {
		t.Where += " ctime not in(?)"
	}
	t.Params = append(t.Params, ctimes)
	return t
}
func (t *UserCondition) AndCtimeNotLike(ctime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime not like ?"
	} else {
		t.Where += " ctime not like ?"
	}
	t.Params = append(t.Params, ctime)
	return t
}
func (t *UserCondition) AndCtimeNotBt(start time.Time, end time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime not between ? and ?"
	} else {
		t.Where += " ctime not between ? and ?"
	}
	t.Params = append(t.Params, start, end)
	return t
}
func (t *UserCondition) AndCtimeNotEq(ctime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and ctime != ?"
	} else {
		t.Where += " ctime != ?"
	}
	t.Params = append(t.Params, ctime)
	return t
}

func (t *UserCondition) AndMtimeEq(mtime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime = ?"
	} else {
		t.Where += " mtime = ?"
	}
	t.Params = append(t.Params, mtime)
	return t
}

func (t *UserCondition) AndMtimeBt(start time.Time, end time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime between ? and ?"
	} else {
		t.Where += " mtime between ? and ?"
	}
	t.Params = append(t.Params, start)
	t.Params = append(t.Params, end)
	return t
}

func (t *UserCondition) AndMtimeGt(mtime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime > ?"
	} else {
		t.Where += " mtime > ?"
	}
	t.Params = append(t.Params, mtime)
	return t
}

func (t *UserCondition) AndMtimeLt(mtime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime < ?"
	} else {
		t.Where += " mtime < ?"
	}
	t.Params = append(t.Params, mtime)
	return t
}

func (t *UserCondition) AndMtimeGtAndEq(mtime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime >= ?"
	} else {
		t.Where += " mtime >= ?"
	}
	t.Params = append(t.Params, mtime)
	return t
}

func (t *UserCondition) AndMtimeLtAndEq(mtime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime <= ?"
	} else {
		t.Where += " mtime <= ?"
	}
	t.Params = append(t.Params, mtime)
	return t
}

func (t *UserCondition) AndMtimeLike(mtime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime like ?"
	} else {
		t.Where += " mtime like ?"
	}
	t.Params = append(t.Params, mtime)
	return t
}

func (t *UserCondition) AndMtimeIn(mtimes []time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime in(?)"
	} else {
		t.Where += " mtime in(?)"
	}
	t.Params = append(t.Params, mtimes)
	return t
}

func (t *UserCondition) AndMtimeNotIn(mtimes []time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime not in(?)"
	} else {
		t.Where += " mtime not in(?)"
	}
	t.Params = append(t.Params, mtimes)
	return t
}
func (t *UserCondition) AndMtimeNotLike(mtime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime not like ?"
	} else {
		t.Where += " mtime not like ?"
	}
	t.Params = append(t.Params, mtime)
	return t
}
func (t *UserCondition) AndMtimeNotBt(start time.Time, end time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime not between ? and ?"
	} else {
		t.Where += " mtime not between ? and ?"
	}
	t.Params = append(t.Params, start, end)
	return t
}
func (t *UserCondition) AndMtimeNotEq(mtime time.Time) *UserCondition {
	if t.Where != "" {
		t.Where += " and mtime != ?"
	} else {
		t.Where += " mtime != ?"
	}
	t.Params = append(t.Params, mtime)
	return t
}

func (t *UserCondition) SetOrderBy(order string) *UserCondition {
	t.Order = order
	return t
}
func (t *UserCondition) SetLimit(limit int64) *UserCondition {
	t.Limit = limit
	return t
}
func (t *UserCondition) setOffset(offset int64) *UserCondition {
	t.Offset = offset
	return t
}

var UserField = &struct {
	Id       string
	Username string
	Nickname string
	Email    string
	Phone    string
	Ctime    string
	Mtime    string
}{
	Id:       "Id",
	Username: "Username",
	Nickname: "Nickname",
	Email:    "Email",
	Phone:    "Phone",
	Ctime:    "Ctime",
	Mtime:    "Mtime",
}
var UserFieldColumn = map[string]string{
	"Id":       "id",
	"Username": "username",
	"Nickname": "nickname",
	"Email":    "email",
	"Phone":    "phone",
	"Ctime":    "ctime",
	"Mtime":    "mtime",
}

func (t *User) Selective(fields []string) (selective map[string]interface{}) {
	selective = make(map[string]interface{})
	for _, v := range fields {
		selective[UserFieldColumn[v]] = reflect.ValueOf(t).Elem().FieldByName(v).Interface()
	}
	return selective
}
