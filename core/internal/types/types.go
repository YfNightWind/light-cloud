// Code generated by goctl. DO NOT EDIT.
package types

type UserRepositorySaveRequest struct {
	ParentId           int64  `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveResponse struct {
}

type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
}

type FileUploadResponse struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type UserRegisterRequest struct {
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 密码
	Email    string `json:"email""`   // 邮箱
	Code     string `json:"code"`     // 验证码
}

type UserRegisterResponse struct {
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserDetailRequest struct {
	Identity string `json:"identity""`
}

type UserDetailResponse struct {
	Name  string `json:"name"`
	Email string `json:"Email"`
}

type MailCodeSendRequest struct {
	Email string `json:"email"`
}

type MailCodeSendResponse struct {
}
