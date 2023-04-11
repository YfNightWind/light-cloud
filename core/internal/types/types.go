// Code generated by goctl. DO NOT EDIT.
package types

type UserShareListRequest struct {
}

type UserShareListResponse struct {
	List []*ShareBasicDetailResponse `json:"list"`
	Msg  string                      `json:"msg"`
	Code int                         `json:"code"`
}

type FileDownloadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
	Ext  string `json:"ext,optional"`
}

type FileDownloadResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data []byte `json:"data"`
}

type PublicRepositorySaveRequest struct {
	RepositoryIdentity string `json:"repositoryIdentity"`
	ParentId           int64  `json:"parentId"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type PublicRepositorySaveResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type PublicFileListRequest struct {
	Id   int64 `json:"id,optional"`
	Page int   `json:"page,optional"`
	Size int   `json:"size,optional"`
}

type PublicFileListResponse struct {
	Count int64         `json:"count"`
	List  []*PublicFile `json:"list"`
	Msg   string        `json:"msg"`
	Code  int           `json:"code"`
}

type PublicFile struct {
	Id                 int64  `json:"id"`
	ParentId           int64  `json:"parent_id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Owner              string `json:"owner"`
	Name               string `json:"name"`
	Size               int64  `json:"size"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	UpdatedAt          string `json:"updated_at"`
}

type PopularShareListRequest struct {
	ClickNum int `json:"click_num,optional"`
}

type PopularShareListResponse struct {
	List []*ShareBasicDetailResponse `json:"list"`
	Msg  string                      `json:"msg"`
	Code int                         `json:"code"`
}

type ShareStatisticsRequest struct {
}

type ShareStatisticsResponse struct {
	ShareCount int    `json:"share_count"`
	ClickNum   int    `json:"click_num"`
	Msg        string `json:"msg"`
	Code       int    `json:"code"`
}

type RegisterCountRequest struct {
}

type RegisterCountResponse struct {
	Count int64  `json:"count"`
	Msg   string `json:"msg"`
	Code  int    `json:"code"`
}

type UserUpdateRequest struct {
	Name     string `json:"name,optional"`
	Avatar   string `json:"avatar,optional"`
	Password string `json:"password,optional"`
	Email    string `json:"email,optional"`
}

type UserUpdateResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type FileUploadChunkCompleteRequest struct {
	Key        string      `json:"key"`
	UploadId   string      `json:"upload_id"`
	CosObjects []CosObject `json:"cos_objects"`
}

type CosObject struct {
	PartNumber int    `json:"part_number"`
	Etag       string `json:"etag"`
}

type FileUploadChunkCompleteResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type FileUploadChunkRequest struct {
}

type FileUploadChunkResponse struct {
	Etag string `json:"etag"` // md5
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type FileUploadPrepareRequest struct {
	Md5  string `json:"md5"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
}

type FileUploadPrepareResponse struct {
	Identity string `json:"identity"`
	UploadId string `json:"upload_id"`
	Key      string `json:"key"`
	Msg      string `json:"msg"`
	Code     int    `json:"code"`
}

type RefreshAuthorizationRequest struct {
}

type RefreshAuthorizationResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Msg          string `json:"msg"`
	Code         int    `json:"code"`
}

type ShareBasicSaveRequest struct {
	RepositoryIdentity string `json:"repository_identity"`
	ParentId           int64  `json:"parent_id"`
}

type ShareBasicSaveResponse struct {
	Identity string `json:"identity"`
	Msg      string `json:"msg"`
	Code     int    `json:"code"`
}

type ShareBasicDetailRequest struct {
	Identity string `json:"identity"`
}

type ShareBasicDetailResponse struct {
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Owner              string `json:"owner"`
	Avatar             string `json:"avatar"`
	Ext                string `json:"ext"`
	Size               int64  `json:"size"`
	Path               string `json:"path"`
	Msg                string `json:"msg"`
	Code               int    `json:"code"`
	ClickNum           int    `json:"click_num"`
	ExpiredTime        int    `json:"expired_time"`
	Desc               string `json:"desc"`
	UpdatedAt          string `json:"updated_at"`
}

type ShareBasicCreateRequest struct {
	UserRepositoryIdentity string `json:"user_repository_identity"`
	ExpiredTime            int    `json:"expired_time"`
	Desc                   string `json:"desc"`
}

type ShareBasicCreateResponse struct {
	Identity string `json:"identity"`
	Msg      string `json:"msg"`
	Code     int    `json:"code"`
}

type UserFileMoveRequest struct {
	Identity       string `json:"identity"`
	ParentIdentity string `json:"parent_identity"`
}

type UserFileMoveResponse struct {
	Msg string `json:"msg"`
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFileDeleteResponse struct {
	Msg string `json:"msg"`
}

type UserFolderCreateRequest struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type UserFolderCreateResponse struct {
	Identity string `json:"identity"`
	Msg      string `json:"msg"`
}

type UserFileNameUpdateRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileNameUpdateResponse struct {
	Identity string `json:"identity"`
	Msg      string `json:"msg"`
	Code     int    `json:"code"`
}

type UserFileListRequest struct {
	Id   int64 `json:"id,optional"`
	Page int   `json:"page,optional"`
	Size int   `json:"size,optional"`
}

type UserFileListResponse struct {
	Count       int64              `json:"count"`
	List        []*UserFile        `json:"list"`
	DeletedList []*DeletedUserFile `json:"deleted_list"`
	Msg         string             `json:"msg"`
	Code        int                `json:"code"`
}

type UserFile struct {
	Id                 int64  `json:"id"`
	ParentId           int64  `json:"parent_id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Size               int64  `json:"size"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	UpdatedAt          string `json:"updated_at"`
}

type DeletedUserFile struct {
	Id                 int64  `json:"id"`
	ParentId           int64  `json:"parent_id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Size               int64  `json:"size"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	DeletedAt          string `json:"deleted_at"`
}

type UserRepositorySaveRequest struct {
	ParentId           int64  `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
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
	Msg      string `json:"msg"`
	Code     int    `json:"code"`
}

type UserDetailRequest struct {
}

type UserDetailResponse struct {
	Name      string `json:"name"`
	Email     string `json:"Email"`
	Avatar    string `json:"avatar"`
	Msg       string `json:"msg"`
	Identity  string `json:"identity"`
	Capacity  int    `json:"capacity"`
	Code      int    `json:"code"`
	CreatedAt string `json:"created_at"`
}

type MailCodeSendRequest struct {
	Email string `json:"email"`
}

type MailCodeSendResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type UserRegisterRequest struct {
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 密码
	Email    string `json:"email""`   // 邮箱
	Code     string `json:"code"`     // 验证码
}

type UserRegisterResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Msg          string `json:"msg"`
	Code         int    `json:"code"`
}
