package user_repository

import (
	"errors"
	"strings"

	"chat-platform-api.com/chat-platform-api/src/repository/common_repository"
	"chat-platform-api.com/chat-platform-api/src/tool/encrypt_tool"
	"chat-platform-api.com/chat-platform-api/src/tool/logging_tool"
	"chat-platform-api.com/chat-platform-api/src/type/entity/user_entity"
	"chat-platform-api.com/chat-platform-api/src/type/model/user_model"
	"chat-platform-api.com/chat-platform-api/src/variable/api_variable"
	"chat-platform-api.com/chat-platform-api/src/variable/auth_variable"
)

type UserRepository struct {
	common_repository.Repository
}

var (
	repository *UserRepository
)

func GetUserRepository() *UserRepository {
	if repository == nil {
		repository.InitRepository()
	}
	return repository
}

func (r *UserRepository) CreateUser(ipAddr string) (int, error) {
	var id int64
	var err error

	defer func() {
		if err != nil {
			logging_tool.PrintErrorLog("CreateUser", err.Error())
		}
	}()

	result, err := r.GetMasterDB().NamedQueryExecute(
		"INSERT INTO `user_entity` SET "+
			"`status` = 1, "+
			"`ip_addr` = :ipAddr",
		map[string]any{
			"ipAddr": ipAddr,
		},
	)
	if err != nil {
		err = errors.New(api_variable.INTERNAL_SERVER_ERROR)
		return 0, err
	}
	id, err = result.LastInsertId()
	return int(id), err
}

func (r *UserRepository) SetUserInformation(userData user_model.SetUserInformation) error {
	var err error

	defer func() {
		if err != nil {
			logging_tool.PrintErrorLog("CreateUser", err.Error())
		}
	}()

	query := "INSERT INTO `user_information` SET"
	dupQuery := "ON DUPLICATE KEY UPDATE "

	queryArr := []string{"userId = :userId"}
	dupQueryArr := []string{}

	if userData.Email != nil {
		queryArr = append(queryArr, ",`email` = :email")
		dupQueryArr = append(dupQueryArr, "email` = :email")
	}
	if userData.Password != nil {
		var str string
		str, err = encrypt_tool.Encrypt([]byte(*userData.Password), auth_variable.ENCRYPT_SECRET_KEY)
		if err != nil {
			return err
		}
		userData.Password = &str
		queryArr = append(queryArr, "`password` = :password")
		dupQueryArr = append(dupQueryArr, "`password` = :password")
	}
	if userData.Name != nil {
		queryArr = append(queryArr, "`name` = :name")
		dupQueryArr = append(dupQueryArr, "`name` = :name")
	}

	query = strings.Join(
		[]string{
			query,
			strings.Join(queryArr, " , "),
			dupQuery,
			strings.Join(dupQueryArr, " , "),
		},
		" ",
	)
	_, err = r.GetMasterDB().NamedQueryExecute(query, userData)
	return err
}

func (r *UserRepository) LoginCheck(email string, password string) (user_model.UserData, bool) {
	type User struct {
		user_entity.UserEntity
		user_entity.UserInformationEntity
		user_entity.UserOauthEntity
	}

	var userEntity []User

	encryptPassword, _ := encrypt_tool.Encrypt([]byte(password), auth_variable.ENCRYPT_SECRET_KEY)

	r.GetSlaveDB().QuerySelect(
		&userEntity,
		`SELECT 
			ui.user_id, 
			u.status, 
			u.ip_addr,
			ui.email,
			uo.oauth_id,
			uo.oauth_host,
			ui.name,
			ui.authentication,
			ui.auth_type
		FROM user AS u 
		INNER JOIN user_information AS ui 
		LEFT JOIN user_oauth AS uo 
		WHERE 
			ui.email = ? AND 
			ui.password = ?
		`,
		email,
		encryptPassword,
	)
	if len(userEntity) <= 0 {
		return user_model.UserData{}, false
	}

	user := userEntity[0]

	return user_model.UserData{
		UserId:         *user.UserInformationEntity.UserId,
		Status:         *user.Status,
		IpAddr:         *user.IpAddr,
		Authentication: *user.Authentication,
		Email:          user.Email,
		OauthId:        user.OauthId,
		OauthHost:      user.OauthHost,
		Name:           user.Name,
		AuthType:       user.AuthType,
	}, true

}

func (r *UserRepository) SetRefreshToken(userId int, token string, deviceId string, ipAddr string) error {
	_, err := r.GetMasterDB().NamedQueryExecute(
		"INSERT INTO `user_refresh_token` SET `userId` = :userId , `token` = :token , `deviceId` = :deviceId , ipAddr = :ipAddr "+
			"ON DUPLICATE KEY UPDATE `token` = :token, `deviceId` = :deviceId, `ipAddr` = :ipAddr",
		map[string]any{
			"userId":   userId,
			"token":    token,
			"deviceId": deviceId,
			"ipAddr":   ipAddr,
		},
	)

	if err != nil {
		return err
	}
	return nil
}
