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
}


func (u *Users) AfterFind() (err error) {
    u.Password = ""
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

func (u *Users) Save(db *gorm.DB) (*Users, error) {

    err := db.Create(&u).Error

    if err != nil {
        return &Users{}, err
    }

    u.Password = ""

    return u, nil
}

func (u *Users) Update(db *gorm.DB) (*Users, error) {

    err := db.Model(&u).Omit("id").Updates(&u).Take(&u).Error

    if err != nil {
        return nil, err
    }

    u.Password = ""

    return u, nil
}

func (u *Users) FindAll(db *gorm.DB) (*[]Users, error) {

    users := []Users{}

    err := db.Model(&Users{}).Limit(100).Find(&users).Error

    if err != nil {
        return &[]Users{}, err
    }

    return &users, err
}

func (u *Users) FindByID(db *gorm.DB, id uint) (*Users, error) {

    err := db.Model(&Users{}).Where("id = ?", id).Take(&u).Error

    if gorm.IsRecordNotFoundError(err) {
        return &Users{}, errors.New("User Not Found")
    }

    return u, err
}

func (u *Users) Delete(db *gorm.DB, id uint) (uint, error) {

    err := db.Model(&Users{}).Where("id = ?", id).Delete(&Users{}).Error

    if err != nil {
        return 0, err
    }

    return 1, nil
}

func (u *Users) hashPassword(password string) ([]byte, error) {
    return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func (u *Users) verifyPassword(hash string, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
