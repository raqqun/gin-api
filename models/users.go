package models

import (
    "errors"
    "time"
    // "log"
    "golang.org/x/crypto/bcrypt"

    "github.com/jinzhu/gorm"
)

type Users struct {
    ID          uint        `gorm:"primary_key" json:"ID"`
    Email       string      `gorm:"size:100;not null;unique" json:"email"`
    Nickname    string      `gorm:"size:255;not null" json:"nickname"`
    Password    string      `gorm:"size:100;not null;" json:"password,omitempty"`
    CreatedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
    getpassword bool
}


func (u *Users) AfterFind() (err error) {
    if !u.getpassword {
        u.Password = ""
    }

    return
}

func (u *Users) BeforeSave(scope *gorm.Scope) (err error) {
    hashedPassword, err := u.hashPassword(u.Password)
    if err != nil {
        return err
    }

    scope.SetColumn("Password", string(hashedPassword))

    return
}

func (u *Users) Save() (*Users, error) {

    err := db.Create(&u).Error

    if err != nil {
        return &Users{}, err
    }

    u.Password = ""

    return u, nil
}

func (u *Users) Update() (*Users, error) {

    err := db.Model(&u).Omit("id").Updates(&u).Take(&u).Error

    if gorm.IsRecordNotFoundError(err) {
        return &Users{}, errors.New("User Not Found")
    }

    u.Password = ""

    return u, err
}

func (u *Users) FindAll() (*[]Users, error) {

    users := []Users{}

    err := db.Model(&Users{}).Limit(100).Find(&users).Error

    if err != nil {
        return &[]Users{}, err
    }

    return &users, err
}

func (u *Users) FindOne(filter interface{}, shouldHavePassword bool) (*Users, error) {
    if shouldHavePassword {
        u.getpassword = true
    }

    err := db.Model(&Users{}).Where(filter).Take(&u).Error

    if gorm.IsRecordNotFoundError(err) {
        return &Users{}, errors.New("User Not Found")
    }

    return u, err
}

func (u *Users) FindByID(id uint) (*Users, error) {

    err := db.Model(&Users{}).Where("id = ?", id).Take(&u).Error

    if gorm.IsRecordNotFoundError(err) {
        return &Users{}, errors.New("User Not Found")
    }

    return u, err
}

func (u *Users) Delete(id uint) (uint, error) {

    err := db.Model(&Users{}).Where("id = ?", id).Delete(&Users{}).Error

    if err != nil {
        return 0, err
    }

    return 1, nil
}

func (u *Users) hashPassword(password string) ([]byte, error) {
    return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func (u *Users) VerifyPassword(hash string, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
