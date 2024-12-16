package errs

import "errors"

var (
	ErrInvalidInputParameters       = errors.New("您输入的参数不合法！")
	ErrUsernameLengthOutOfRange     = errors.New("用户名长度超出要求范围，长度应在6 ~ 16位！")
	ErrPasswordLengthOutOfRange     = errors.New("密码长度超出要求范围，长度应在8 ~ 24位！")
	ErrInviteCodeLengthOutOfRange   = errors.New("邀请码长度超出要求范围，长度应在 2 ~ 24位！")
	ErrUserRegistrationFailed       = errors.New("用户注册失败！")
	ErrGetCurrentUserDetailsFailed  = errors.New("获取当前用户详细信息失败！")
	ErrLogoutFailed                 = errors.New("退出登陆失败！")
	ErrUserLoginFailed              = errors.New("用户登陆失败！")
	ErrFileSizeInvalid              = errors.New("文件大小非法！")
	ErrFileExtensionNotSupported    = errors.New("文件后缀格式不支持！")
	ErrGoalCharacterCountOutOfRange = errors.New("分析目标的字符数应在 2 ~ 255间！")
	ErrTableNameLengthOutOfRange    = errors.New("表名的字符数应在 2 ~ 128 间！")
	ErrUnsupportedChartType         = errors.New("不支持该图表类型！")
	ErrFindMyChartFailed            = errors.New("查找我的图表失败！")
	ErrInvalidPageSize              = errors.New("页面大小不合法！")
	ErrInvalidPageParameters        = errors.New("当前页面参数不合法！")
	ErrDeleteMyChartFailed          = errors.New("删除我的图表失败！")
	ErrGenerateChartFailed          = errors.New("生成图表失败！")
	ErrGetExampleChartFailed        = errors.New("获取示例图表失败！")
	ErrOperateOtherUserChart        = errors.New("不允许操作其他人的图表！")
	ErrFindNotExistChart            = errors.New("图表不存在！")
	ErrUpdateChartFailed            = errors.New("修改图表失败！")
	ErrShareChartFailed             = errors.New("生成分享图表链接失败！")
)
