package information

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/gookit/color"
	"time"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

var apis = []model.SysApi{
	{global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/base/login", "用户登录", "base", "POST"},
	{global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/register", "用户注册", "user", "POST"},
	{global.GVA_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/createApi", "创建api", "api", "POST"},
	{global.GVA_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiList", "获取api列表", "api", "POST"},
	{global.GVA_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiById", "获取api详细信息", "api", "POST"},
	{global.GVA_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/deleteApi", "删除Api", "api", "POST"},
	{global.GVA_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/updateApi", "更新Api", "api", "POST"},
	{global.GVA_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getAllApis", "获取所有api", "api", "POST"},
	{global.GVA_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/createAuthority", "创建角色", "authority", "POST"},
	{global.GVA_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/deleteAuthority", "删除角色", "authority", "POST"},
	{global.GVA_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/getAuthorityList", "获取角色列表", "authority", "POST"},
	{global.GVA_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenu", "获取菜单树", "menu", "POST"},
	{global.GVA_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuList", "分页获取基础menu列表", "menu", "POST"},
	{global.GVA_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addBaseMenu", "新增菜单", "menu", "POST"},
	{global.GVA_MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuTree", "获取用户动态路由", "menu", "POST"},
	{global.GVA_MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addMenuAuthority", "增加menu和角色关联关系", "menu", "POST"},
	{global.GVA_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuAuthority", "获取指定角色menu", "menu", "POST"},
	{global.GVA_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/deleteBaseMenu", "删除菜单", "menu", "POST"},
	{global.GVA_MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/updateBaseMenu", "更新菜单", "menu", "POST"},
	{global.GVA_MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuById", "根据id获取菜单", "menu", "POST"},
	{global.GVA_MODEL{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/changePassword", "修改密码", "user", "POST"},
	{global.GVA_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/getUserList", "获取用户列表", "user", "POST"},
	{global.GVA_MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserAuthority", "修改用户角色", "user", "POST"},
	{global.GVA_MODEL{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/upload", "文件上传示例", "fileUploadAndDownload", "POST"},
	{global.GVA_MODEL{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/getFileList", "获取上传文件列表", "fileUploadAndDownload", "POST"},
	{global.GVA_MODEL{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/updateCasbin", "更改角色api权限", "casbin", "POST"},
	{global.GVA_MODEL{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/getPolicyPathByAuthorityId", "获取权限列表", "casbin", "POST"},
	{global.GVA_MODEL{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/deleteFile", "删除文件", "fileUploadAndDownload", "POST"},
	{global.GVA_MODEL{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/jwt/jsonInBlacklist", "jwt加入黑名单", "jwt", "POST"},
	{global.GVA_MODEL{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/setDataAuthority", "设置角色资源权限", "authority", "POST"},
	{global.GVA_MODEL{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getSystemConfig", "获取配置文件内容", "system", "POST"},
	{global.GVA_MODEL{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/setSystemConfig", "设置配置文件内容", "system", "POST"},
	{global.GVA_MODEL{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "创建客户", "customer", "POST"},
	{global.GVA_MODEL{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "更新客户", "customer", "PUT"},
	{global.GVA_MODEL{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "删除客户", "customer", "DELETE"},
	{global.GVA_MODEL{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "获取单一客户", "customer", "GET"},
	{global.GVA_MODEL{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customerList", "获取客户列表", "customer", "GET"},
	{global.GVA_MODEL{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/casbinTest/:pathParam", "RESTFUL模式测试", "casbin", "GET"},
	{global.GVA_MODEL{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/createTemp", "自动化代码", "autoCode", "POST"},
	{global.GVA_MODEL{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/updateAuthority", "更新角色信息", "authority", "PUT"},
	{global.GVA_MODEL{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/copyAuthority", "拷贝角色", "authority", "POST"},
	{global.GVA_MODEL{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/deleteUser", "删除用户", "user", "DELETE"},
	{global.GVA_MODEL{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/createSysDictionaryDetail", "新增字典内容", "sysDictionaryDetail", "POST"},
	{global.GVA_MODEL{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/deleteSysDictionaryDetail", "删除字典内容", "sysDictionaryDetail", "DELETE"},
	{global.GVA_MODEL{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/updateSysDictionaryDetail", "更新字典内容", "sysDictionaryDetail", "PUT"},
	{global.GVA_MODEL{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/findSysDictionaryDetail", "根据ID获取字典内容", "sysDictionaryDetail", "GET"},
	{global.GVA_MODEL{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/getSysDictionaryDetailList", "获取字典内容列表", "sysDictionaryDetail", "GET"},
	{global.GVA_MODEL{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/createSysDictionary", "新增字典", "sysDictionary", "POST"},
	{global.GVA_MODEL{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/deleteSysDictionary", "删除字典", "sysDictionary", "DELETE"},
	{global.GVA_MODEL{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/updateSysDictionary", "更新字典", "sysDictionary", "PUT"},
	{global.GVA_MODEL{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/findSysDictionary", "根据ID获取字典", "sysDictionary", "GET"},
	{global.GVA_MODEL{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/getSysDictionaryList", "获取字典列表", "sysDictionary", "GET"},
	{global.GVA_MODEL{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/createSysOperationRecord", "新增操作记录", "sysOperationRecord", "POST"},
	{global.GVA_MODEL{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecord", "删除操作记录", "sysOperationRecord", "DELETE"},
	{global.GVA_MODEL{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/findSysOperationRecord", "根据ID获取操作记录", "sysOperationRecord", "GET"},
	{global.GVA_MODEL{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/getSysOperationRecordList", "获取操作记录列表", "sysOperationRecord", "GET"},
	{global.GVA_MODEL{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getTables", "获取数据库表", "autoCode", "GET"},
	{global.GVA_MODEL{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getDB", "获取所有数据库", "autoCode", "GET"},
	{global.GVA_MODEL{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getColumn", "获取所选table的所有字段", "autoCode", "GET"},
	{global.GVA_MODEL{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecordByIds", "批量删除操作历史", "sysOperationRecord", "DELETE"},
	{global.GVA_MODEL{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/upload", "插件版分片上传", "simpleUploader", "POST"},
	{global.GVA_MODEL{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/checkFileMd5", "文件完整度验证", "simpleUploader", "GET"},
	{global.GVA_MODEL{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/mergeFileMd5", "上传完成合并文件", "simpleUploader", "GET"},
	{global.GVA_MODEL{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserInfo", "设置用户信息", "user", "PUT"},
	{global.GVA_MODEL{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getServerInfo", "获取服务器信息", "system", "POST"},
	{global.GVA_MODEL{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/email/emailTest", "发送测试邮件", "email", "POST"},
	{global.GVA_MODEL{ID: 68, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/createWorkflowProcess", "新建工作流", "workflowProcess", "POST"},
	{global.GVA_MODEL{ID: 69, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/deleteWorkflowProcess", "删除工作流", "workflowProcess", "DELETE"},
	{global.GVA_MODEL{ID: 70, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/deleteWorkflowProcessByIds", "批量删除工作流", "workflowProcess", "DELETE"},
	{global.GVA_MODEL{ID: 71, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/updateWorkflowProcess", "更新工作流", "workflowProcess", "PUT"},
	{global.GVA_MODEL{ID: 72, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/findWorkflowProcess", "根据ID获取工作流", "workflowProcess", "GET"},
	{global.GVA_MODEL{ID: 73, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getWorkflowProcessList", "获取工作流", "workflowProcess", "GET"},
	{global.GVA_MODEL{ID: 74, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/findWorkflowStep", "获取工作流步骤", "workflowProcess", "GET"},
	{global.GVA_MODEL{ID: 75, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/startWorkflow", "启动工作流", "workflowProcess", "POST"},
	{global.GVA_MODEL{ID: 76, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getMyStated", "获取我发起的工作流", "workflowProcess", "GET"},
	{global.GVA_MODEL{ID: 77, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getMyNeed", "获取我的待办", "workflowProcess", "GET"},
	{global.GVA_MODEL{ID: 78, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getWorkflowMoveByID", "根据id获取当前节点详情和历史", "workflowProcess", "GET"},
	{global.GVA_MODEL{ID: 79, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/completeWorkflowMove", "提交工作流", "workflowProcess", "POST"},
	{global.GVA_MODEL{ID: 80, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/preview", "预览自动化代码", "autoCode", "POST"},
	{global.GVA_MODEL{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/importExcel", "预览自动化代码", "autoCode", "POST"},
	{global.GVA_MODEL{ID: 82, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/loadExcel", "预览自动化代码", "autoCode", "POST"},
	{global.GVA_MODEL{ID: 83, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/exportExcel", "预览自动化代码", "autoCode", "POST"},
	{global.GVA_MODEL{ID: 84, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/downloadTemplate", "预览自动化代码", "autoCode", "POST"},
	{global.GVA_MODEL{ID: 10000, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/logTemplate/getRealtimeResult", "获取实时模版", "logTemplate", "GET"},
	{global.GVA_MODEL{ID: 10001, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/logTemplate/createLogTemplate", "新增日志模版", "logTemplate", "POST"},
	{global.GVA_MODEL{ID: 10002, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/logTemplate/deleteLogTemplate", "删除日志模版", "logTemplate", "DELETE"},
	{global.GVA_MODEL{ID: 10003, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/logTemplate/deleteLogTemplateByIds", "批量删除日志模版", "logTemplate", "DELETE"},
	{global.GVA_MODEL{ID: 10004, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/logTemplate/getLogTemplateList", "获取日志模版列表", "logTemplate", "GET"},
	{global.GVA_MODEL{ID: 10005, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/app/app", "创建app", "app", "POST"},
	{global.GVA_MODEL{ID: 10006, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/app/app", "删除app", "app", "DELETE"},
	{global.GVA_MODEL{ID: 10007, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/app/appList", "获取app列表", "app", "GET"},
	{global.GVA_MODEL{ID: 10008, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/app/app", "更新app", "app", "PUT"},
	{global.GVA_MODEL{ID: 10009, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/setAppAuthority", "设置角色App权限", "authority", "POST"},
	{global.GVA_MODEL{ID: 10010, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/templateAlarm/createTemplateAlarmStrategy", "新增模版报警策略", "templateAlarm", "POST"},
	{global.GVA_MODEL{ID: 10011, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/templateAlarm/deleteTemplateAlarmStrategy", "删除模版报警策略", "templateAlarm", "DELETE"},
	{global.GVA_MODEL{ID: 10012, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/templateAlarm/deleteTemplateAlarmStrategyByIds", "批量删除模版报警策略", "templateAlarm", "DELETE"},
	{global.GVA_MODEL{ID: 10013, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/templateAlarm/updateTemplateAlarmStrategy", "更新模版报警策略", "templateAlarm", "PUT"},
	{global.GVA_MODEL{ID: 10014, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/templateAlarm/findTemplateAlarmStrategy", "根据ID获取模版报警策略", "templateAlarm", "GET"},
	{global.GVA_MODEL{ID: 10015, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/templateAlarm/getTemplateAlarmStrategyList", "获取模版报警策略列表", "templateAlarm", "GET"},
	{global.GVA_MODEL{ID: 10016, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/levelAlarm/createLevelAlarmStrategy", "新增级别报警策略", "levelAlarm", "POST"},
	{global.GVA_MODEL{ID: 10017, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/levelAlarm/deleteLevelAlarmStrategy", "删除级别报警策略", "levelAlarm", "DELETE"},
	{global.GVA_MODEL{ID: 10018, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/levelAlarm/deleteLevelAlarmStrategyByIds", "批量删除级别报警策略", "levelAlarm", "DELETE"},
	{global.GVA_MODEL{ID: 10019, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/levelAlarm/updateLevelAlarmStrategy", "更新级别报警策略", "levelAlarm", "PUT"},
	{global.GVA_MODEL{ID: 10020, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/levelAlarm/findLevelAlarmStrategy", "根据ID获取级别报警策略", "levelAlarm", "GET"},
	{global.GVA_MODEL{ID: 10021, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/levelAlarm/getLevelAlarmStrategyList", "获取级别报警策略列表", "levelAlarm", "GET"},
	{global.GVA_MODEL{ID: 10022, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/regexAlarm/createRegexAlarmStrategy", "创建正则报警", "regexAlarm", "POST"},
	{global.GVA_MODEL{ID: 10023, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/regexAlarm/deleteRegexAlarmStrategy", "删除正则报警", "regexAlarm", "DELETE"},
	{global.GVA_MODEL{ID: 10024, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/regexAlarm/deleteRegexAlarmStrategyByIds", "批量删除正则报警", "regexAlarm", "DELETE"},
	{global.GVA_MODEL{ID: 10025, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/regexAlarm/updateRegexAlarmStrategy", "更新正则报警", "regexAlarm", "PUT"},
	{global.GVA_MODEL{ID: 10026, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/regexAlarm/findRegexAlarmStrategy", "用id查询正则报警", "regexAlarm", "GET"},
	{global.GVA_MODEL{ID: 10027, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/regexAlarm/getRegexAlarmStrategyList", "分页获取正则报警列表", "regexAlarm", "GET"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_apis 表数据初始化
func (a *api) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 67}).Find(&[]model.SysApi{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_apis 表初始数据成功!")
		return nil
	})
}
