package main

type Product_link_wire struct {
	Id             int    `orm:"id"`
	MaterialStatus string `orm:"material_status"` // 材质状态
	Weight         string `orm:"weight"`          // 单重量
	TotalArea      string `orm:"total_area"`      // 总面积(㎡)
	Flag           int    `orm:"flag"`
	CreatedAt      string `orm:"created_at"`
	UpdatedAt      string `orm:"updated_at"`
	DeletedAt      string `orm:"deleted_at"`
}

type Send_return struct {
	Id               int     `orm:"id"`
	SendId           int     `orm:"send_id"`
	PackingId        int     `orm:"packing_id"`
	ProductId        int     `orm:"product_id"`
	PackingProductId int     `orm:"packing_product_id"`
	ProjectId        int     `orm:"project_id"`
	Cuid             int     `orm:"cuid"`
	Count            float64 `orm:"count"`      // 退货数量
	Reason           string  `orm:"reason"`     // 退货原因
	Remark           string  `orm:"remark"`     // 其他描述
	Attachment       string  `orm:"attachment"` // 照片
	RepairId         int     `orm:"repair_id"`  // 补货对应id 暂时不用
	Status           int     `orm:"status"`     // 发起中 1接收
	UseCount         float64 `orm:"use_count"`  // 接收数量
	Flag             int     `orm:"flag"`
	CreatedAt        string  `orm:"created_at"`
	UpdatedAt        string  `orm:"updated_at"`
	DeletedAt        string  `orm:"deleted_at"`
	ReturnType       string  `orm:"return_type"` // 退货类型
	CompanyId        int     `orm:"company_id"`  // gongsi id

	MaterialName string `orm:"material_name"` // 材料名称
	ReplenishId  int    `orm:"replenish_id"`  // 补货id
	IsReplenish  int    `orm:"is_replenish"`  // 是否补货
}

type Send struct {
	Id                int     `orm:"id"`
	SendNo            string  `orm:"send_no"`            // 订单编号
	Count             float64 `orm:"count"`              // 发货总数
	ActualReceiver    string  `orm:"actual_receiver"`    // 签收人
	Address           string  `orm:"address"`            // 收货地址
	ReceiveAttachment string  `orm:"receive_attachment"` // 收货附件
	ReceiveDate       string  `orm:"receive_date"`       // 收货时间
	ReceiveCount      float64 `orm:"receive_count"`      // 收货总数量
	ReceiveRemark     string  `orm:"receive_remark"`     // 收货备注
	Remark            string  `orm:"remark"`             // 备注
	CompanyId         int     `orm:"company_id"`         // 仓库id
	ProjectId         int     `orm:"project_id"`
	ExpressNo         string  `orm:"express_no"` // 快递编号
	Flag              int     `orm:"flag"`
	CreatedAt         string  `orm:"created_at"`
	UpdatedAt         string  `orm:"updated_at"`
	DeletedAt         string  `orm:"deleted_at"`
	Status            int     `orm:"status"`       // 0未签收 1签收
	ReturnCount       int     `orm:"return_count"` // 退货数量
}

type Users struct {
	Id        int    `orm:"id"`
	Cuid      int    `orm:"cuid"`       // ucenter用户id
	Username  string `orm:"username"`   // 账号 - 暂时不用
	Password  string `orm:"password"`   // 密码 - 暂时不用
	Nickname  string `orm:"nickname"`   // 昵称
	Mobile    string `orm:"mobile"`     // 手机号
	MUserKey  string `orm:"m_user_key"` // 用户Key
	Flag      int    `orm:"flag"`
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
	DeletedAt string `orm:"deleted_at"`
	Avatar    string `orm:"avatar"`
	DUserKey  string `orm:"d_user_key"` // 公装用户key
}

type Product struct {
	Id                 int     `orm:"id"`
	MaterialName       string  `orm:"material_name"`       // 材料名称
	BlankingAttachment string  `orm:"blanking_attachment"` // 下料附件信息(与码里公装关联)
	Attachment         string  `orm:"attachment"`          // 附件
	InstallMap         string  `orm:"install_map"`         // 安装示意图
	Price              float64 `orm:"price"`               // 价格
	Count              float64 `orm:"count"`               // 数量
	ContractCount      float64 `orm:"contract_count"`      // 与供应商签的合同数量(来源码里公装)
	PackCount          float64 `orm:"pack_count"`          // 打包数量
	SendCount          float64 `orm:"send_count"`          // 发货数量
	ReturnCount        float64 `orm:"return_count"`        // 退货数量
	ReceiveCount       float64 `orm:"receive_count"`       // 签收数量
	Unit               string  `orm:"unit"`                // 单位
	ProjectId          int     `orm:"project_id"`
	ProjectName        string  `orm:"project_name"`
	ReplenishmentFlag  int     `orm:"replenishment_flag"` // 是否补货产品
	ProductSubFlag     int     `orm:"product_sub_flag"`   // 是否有子部件
	ConfigData         string  `orm:"config_data"`        // 自定义字段信息
	AppendAttachment   string  `orm:"append_attachment"`  // 附加的资源库信息
	ProjectAdditional  string  `orm:"project_additional"` // 项目补充信息
	Remark             string  `orm:"remark"`             // 备注
	Length             float64 `orm:"length"`             // 长
	Width              float64 `orm:"width"`              // 宽
	Height             float64 `orm:"height"`             // 厚
	Location           string  `orm:"location"`           // 安装位置
	Standard           string  `orm:"standard"`           // 规格
	ArriveDate         string  `orm:"arrive_date"`        // 到货时间
	Cuid               int     `orm:"cuid"`
	CompanyId          int     `orm:"company_id"`
	SupplyCycle        int     `orm:"supply_cycle"` // 供货周期
	Flag               int     `orm:"flag"`
	CreatedAt          string  `orm:"created_at"`
	UpdatedAt          string  `orm:"updated_at"`
	DeletedAt          string  `orm:"deleted_at"`
	MaterialId         int     `orm:"material_id"` // 材料单id
	PlatformKey        string  `orm:"platform_key"`
	PlatformUid        string  `orm:"platform_uid"`
	PlatformId         string  `orm:"platform_id"`
	ContractId         int     `orm:"contract_id"`
	SendReturnId       int     `orm:"send_return_id"` // 退货id
}

type Product_class struct {
	Id        int    `orm:"id"`
	ClassName string `orm:"class_name"` // 材料类型名称
	Desc      string `orm:"desc"`       // 描述
	CatsId    int    `orm:"cats_id"`
	CompanyId int    `orm:"company_id"`
	Cuid      int    `orm:"cuid"`
	Flag      int    `orm:"flag"`
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
	DeletedAt string `orm:"deleted_at"`
}

type Project struct {
	Id                int     `orm:"id"`
	ProjectName       string  `orm:"project_name"` // 项目名称
	State             int     `orm:"state"`        // 0在建中 1已完成
	Remark            string  `orm:"remark"`       // 备注
	Cuid              int     `orm:"cuid"`
	CompanyId         int     `orm:"company_id"`
	AppendAttachment  string  `orm:"append_attachment"`  // 附加信息
	ReceiverMembers   string  `orm:"receiver_members"`   // 收货人ids
	BindState         int     `orm:"bind_state"`         // 绑定工程账号的状态,手动新建的项目默认已绑定.未绑定,待处理,已绑定
	BindType          int     `orm:"bind_type"`          // 绑定的项目属于哪个系统(项目版，企业版，还是新建的) 新建,项目版,企业版,私有化定制
	DataOrigin        string  `orm:"data_origin"`        // 绑定工程的账号名称的数据来源(例如:码里公装，中建深装)
	ProjectAccount    string  `orm:"project_account"`    // 绑定工程的账号名称
	SupplierAccountid int     `orm:"supplier_accountid"` // 绑定工程供应商id
	ProjectAccountid  int     `orm:"project_accountid"`  // 绑定工程账号(项目)的id
	ContractMoney     float64 `orm:"contract_money"`     // 合同总金额
	ReceivedMoney     float64 `orm:"received_money"`     // 已回款总金额
	ReceiptMoney      float64 `orm:"receipt_money"`      // 已开票总金额
	Flag              int     `orm:"flag"`
	CreatedAt         string  `orm:"created_at"`
	UpdatedAt         string  `orm:"updated_at"`
	DeletedAt         string  `orm:"deleted_at"`
	Status            int     `orm:"status"` // 状态 1已接收 如果是自建的会自动设置1
	PlatformKey       string  `orm:"platform_key"`
	PlatformUid       string  `orm:"platform_uid"`     // 平台用户id
	PlatformId        string  `orm:"platform_id"`      // 平台id
	IsPlatform        int     `orm:"is_platform"`      // 是否是三方同步
	ReceiveTime       string  `orm:"receive_time"`     // 接收时间
	ReceiverAddress   string  `orm:"receiver_address"` // 收货人地址
}

type Company_users struct {
	Id        int    `orm:"id"`
	CompanyId int    `orm:"company_id"` // 企业id
	Cuid      int    `orm:"cuid"`
	IsMain    int    `orm:"is_main"`   // 是否主用户
	RoleId    int    `orm:"role_id"`   // 角色id
	RuleData  string `orm:"rule_data"` // 其他权限
	Status    int    `orm:"status"`    // 状态 0停用 1正常
	Flag      int    `orm:"flag"`
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
	DeletedAt string `orm:"deleted_at"`
	IsDefault int    `orm:"is_default"` // 默认企业
}

type Contract struct {
	Id                    int     `orm:"id"`
	ContractName          string  `orm:"contract_name"`            // 合同名
	ContractNo            string  `orm:"contract_no"`              // 合同编号
	UseTime               string  `orm:"use_time"`                 // 签订时间
	UseAddress            string  `orm:"use_address"`              // 签约地点
	Price                 float64 `orm:"price"`                    // 全部总金额
	AName                 string  `orm:"a_name"`                   // 甲方名
	ATel                  string  `orm:"a_tel"`                    // 甲方电话
	AEmail                string  `orm:"a_email"`                  // 甲方email
	BName                 string  `orm:"b_name"`                   // 乙方名
	BTel                  string  `orm:"b_tel"`                    // 乙方电话
	BEmail                string  `orm:"b_email"`                  // 乙方email
	ContractPrice         float64 `orm:"contract_price"`           // 合同金额
	Attachment            string  `orm:"attachment"`               // 合同附件
	ContractType          string  `orm:"contract_type"`            // 合同类型 供应商合同 框架协议
	ProjectId             int     `orm:"project_id"`               // 项目id
	StartDate             string  `orm:"start_date"`               // 合同开始时间
	EndDate               string  `orm:"end_date"`                 // 合同结束时间
	PayWay                string  `orm:"pay_way"`                  // 付款方式
	BreachItem            string  `orm:"breach_item"`              // 违约条款
	TotalContractTaxPrice float64 `orm:"total_contract_tax_price"` // 合同含税总价
	Remark                string  `orm:"remark"`                   // 备注
	ItemReceiptAmount     float64 `orm:"item_receipt_amount"`      // 已开进项发票总额
	InStorageAmount       float64 `orm:"in_storage_amount"`        // 合同入库材料总金额
	RequestAccount        float64 `orm:"request_account"`          // 总请款金额
	ReceiptAccount        float64 `orm:"receipt_account"`          // 已收发票金额
	PayAccount            float64 `orm:"pay_account"`              // 付款总金额
	HasR                  float64 `orm:"has_r"`                    // 已请总金额
	CompanyId             int     `orm:"company_id"`               // 公司id
	Cuid                  int     `orm:"cuid"`
	Flag                  int     `orm:"flag"`
	CreatedAt             string  `orm:"created_at"`
	UpdatedAt             string  `orm:"updated_at"`
	DeletedAt             string  `orm:"deleted_at"`
	PlatformKey           string  `orm:"platform_key"`
	PlatformUid           string  `orm:"platform_uid"` // 平台用户id
	PlatformId            string  `orm:"platform_id"`  // 平台id
	IsPlatform            int     `orm:"is_platform"`  // 是否是三方同步
	BindState             int     `orm:"bind_state"`   // 是否绑定三方平台
}

type Receiver_users struct {
	Id          int    `orm:"id"`
	CompanyId   int    `orm:"company_id"`
	ContractId  int    `orm:"contract_id"`
	Cuid        int    `orm:"cuid"`
	PlatformKey string `orm:"platform_key"`
	PlatformUid string `orm:"platform_uid"`
	Flag        int    `orm:"flag"`
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
	DeletedAt   string `orm:"deleted_at"`
	ProjectId   int    `orm:"project_id"` // 项目id
	Nickname    string `orm:"nickname"`   // 姓名
	Mobile      string `orm:"mobile"`
}

type Ma struct {
	Id                        int     `orm:"id"`
	CreateDate                string  `orm:"create_date"`
	ModifyDate                string  `orm:"modify_date"`
	BindOpenid                string  `orm:"bind_openid"`
	Birthday                  string  `orm:"birthday"`
	DeleteFlag                int     `orm:"delete_flag"`
	HeadImage                 string  `orm:"head_image"`
	InvitationCode            string  `orm:"invitation_code"`
	LoginDate                 string  `orm:"login_date"`
	Mobile                    string  `orm:"mobile"`
	Money                     float64 `orm:"money"`
	MyinvitationCode          string  `orm:"myinvitation_code"`
	Nickname                  string  `orm:"nickname"`
	Password                  string  `orm:"password"`
	Score                     int     `orm:"score"`
	CompanyName               string  `orm:"company_name"`
	ActivateProlocutor        int     `orm:"activate_prolocutor"`
	MemberType                int     `orm:"member_type"`
	Point                     float64 `orm:"point"`
	RequireType               int     `orm:"require_type"`
	CustomPoint               float64 `orm:"custom_point"`
	PopularizeMember          int     `orm:"popularize_member"`
	WxQrCodeUrl               string  `orm:"wx_qr_code_url"`
	Address                   string  `orm:"address"`
	Telephone                 string  `orm:"telephone"`
	TopContacts               string  `orm:"top_contacts"`
	Department                string  `orm:"department"`
	Admin                     int     `orm:"admin"`
	JobNumber                 string  `orm:"job_number"`
	Bind                      int     `orm:"bind"`
	Sources                   string  `orm:"sources"`
	Textarea                  string  `orm:"textarea"`
	ExtendFields              string  `orm:"extend_fields"`
	Attendance                int     `orm:"attendance"`
	IdNo                      string  `orm:"id_no"`
	WorkerType                string  `orm:"worker_type"`
	FaceImageUrl              string  `orm:"face_image_url"`
	WorkerPlatformReceiveFlag int     `orm:"worker_platform_receive_flag"`
	Locations                 string  `orm:"locations"`
	IdCardfmurl               string  `orm:"id_cardfmurl"`
	IdCardRecognition         string  `orm:"id_card_recognition"`
	IdCardzmurl               string  `orm:"id_cardzmurl"`
	HoursPerDay               int     `orm:"hours_per_day"`
	Remarks                   string  `orm:"remarks"`
	WorkingState              int     `orm:"working_state"`
	WorkingStateApproveDate   string  `orm:"working_state_approve_date"`
	Duties                    int     `orm:"duties"`
	MemberGroup               int     `orm:"member_group"`
	MemberGroupId             int     `orm:"member_group_id"`
	FaceInfo                  string  `orm:"face_info"`
	AcceptBindDate            string  `orm:"accept_bind_date"`
}

type Packing_product struct {
	Id            int     `orm:"id"`
	PackingId     int     `orm:"packing_id"`     // 打包id
	CompanyId     int     `orm:"company_id"`     // 企业id
	OrderReturnid int     `orm:"order_returnid"` // 订单退货详情id
	ProductId     int     `orm:"product_id"`
	MaterialId    int     `orm:"material_id"`
	Count         float64 `orm:"count"`         // 打包数量
	ReturnCount   float64 `orm:"return_count"`  // 退货数量
	MaterialName  string  `orm:"material_name"` // 产品名称
	Flag          int     `orm:"flag"`
	CreatedAt     string  `orm:"created_at"`
	UpdatedAt     string  `orm:"updated_at"`
	DeletedAt     string  `orm:"deleted_at"`
	ContractId    int     `orm:"contract_id"` // 合同id
	ProjectId     int     `orm:"project_id"`
	DepositoryId  int     `orm:"depository_id"` // 仓库id
	Status        int     `orm:"status"`        // 0已打包 1已发货 4已收货 已验收
	ReceiveCount  float64 `orm:"receive_count"` // 签收数量
	ReceiveTime   string  `orm:"receive_time"`  // 签收时间
	IsReceive     int     `orm:"is_receive"`    // 是否签收
}

type Packing struct {
	Id           int     `orm:"id"`
	PackingName  string  `orm:"packing_name"` // 包装名称
	SerialNo     string  `orm:"serial_no"`    // 包装编号
	Count        int     `orm:"count"`        // 产品总数
	ReturnCount  int     `orm:"return_count"` // 包装下退货数量
	Remark       string  `orm:"remark"`       // 描述
	CompanyId    int     `orm:"company_id"`
	ProductId    int     `orm:"product_id"`
	MaterialId   int     `orm:"material_id"`
	Flag         int     `orm:"flag"`
	CreatedAt    string  `orm:"created_at"`
	UpdatedAt    string  `orm:"updated_at"`
	DeletedAt    string  `orm:"deleted_at"`
	ContractId   int     `orm:"contract_id"` // 合同id
	ProjectId    int     `orm:"project_id"`
	DepositoryId int     `orm:"depository_id"` // 仓库id
	Status       int     `orm:"status"`        // 0已打包 1已发货 4已收货 已验收
	SendId       int     `orm:"send_id"`
	ReceiveCount float64 `orm:"receive_count"`
}

type Product_cats struct {
	Id         int    `orm:"id"`
	CatName    string `orm:"cat_name"`    // 大类名
	ModelName  string `orm:"model_name"`  // 数据模型名称
	ModelTable string `orm:"model_table"` // 数据表名 全名
	Desc       string `orm:"desc"`
	IsShow     int    `orm:"is_show"`
	Flag       int    `orm:"flag"`
	CreatedAt  string `orm:"created_at"`
	UpdatedAt  string `orm:"updated_at"`
	DeletedAt  string `orm:"deleted_at"`
}

type Material struct {
	Id             int     `orm:"id"`
	Name           string  `orm:"name"`             // 材料单名称
	TotalAmount    float64 `orm:"total_amount"`     // 下料总额（不含税）
	TotalTaxAmount float64 `orm:"total_tax_amount"` // 下料总额（含税）
	DesignNo       string  `orm:"design_no"`        // 设计订单号
	ApplyNo        string  `orm:"apply_no"`         // 下料单号
	Remark         string  `orm:"remark"`           // 备注
	CreateType     int     `orm:"create_type"`      // 创建类型 新建,    采购计划生成
	Type           int     `orm:"type"`             // 类型    内装材料,    幕墙面材,    幕墙辅材,    幕墙线材
	PlatformKey    string  `orm:"platform_key"`     // 平台key
	PlatformUid    string  `orm:"platform_uid"`     // 平台用户id
	PlatformId     string  `orm:"platform_id"`      // 平台id  或者对照订单号
	Flag           int     `orm:"flag"`
	CreatedAt      string  `orm:"created_at"`
	UpdatedAt      string  `orm:"updated_at"`
	ProjectId      int     `orm:"project_id"`  // 项目id
	CompanyId      int     `orm:"company_id"`  // 公司id
	ContractId     int     `orm:"contract_id"` // 合同id
	BeginTime      string  `orm:"begin_time"`  // 项目开始时间
}

type Material_link struct {
	Id         int     `orm:"id"`
	MaterialId int     `orm:"material_id"` // 材料单id
	ProductId  int     `orm:"product_id"`
	CompanyId  int     `orm:"company_id"`
	Count      float64 `orm:"count"`
	Flag       int     `orm:"flag"`
	CreatedAt  string  `orm:"created_at"`
	UpdatedAt  string  `orm:"updated_at"`
	DeletedAt  string  `orm:"deleted_at"`
	ProjectId  int     `orm:"project_id"`
	Price      float64 `orm:"price"`
}

type Platform_users struct {
	Id          int    `orm:"id"`
	Cuid        int    `orm:"cuid"`         // 用户中心用户id
	PlatformKey string `orm:"platform_key"` // 平台key
	PlatformUid string `orm:"platform_uid"` // 平台用户id
	Flag        int    `orm:"flag"`         // 删除标识
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
}

type Pr struct {
	Id           int     `orm:"id"`
	CompanyId    int     `orm:"company_id"`
	ProjectId    int     `orm:"project_id"`
	Price        float64 `orm:"price"`     // 实际金额
	TplPrice     float64 `orm:"tpl_price"` // 模板金额 输入的
	Type         int     `orm:"type"`      // 0进度款 1预付款 2结算款 3质保金
	Desc         string  `orm:"desc"`      // 描述
	Count        float64 `orm:"count"`     // 数量
	Cuid         int     `orm:"cuid"`
	Status       int     `orm:"status"` // 0正常 1接收 2审批通过 -1未通过
	PlatformKey  string  `orm:"platform_key"`
	PlatformUid  string  `orm:"platform_uid"`
	PlatformId   string  `orm:"platform_id"`
	IsPlatform   int     `orm:"is_platform"`  // 是否平台审核
	PlatformMsg  string  `orm:"platform_msg"` // 平台审核消息
	PrNo         string  `orm:"pr_no"`        // 请款唯一编号
	Flag         int     `orm:"flag"`
	CreatedAt    string  `orm:"created_at"`
	UpdatedAt    string  `orm:"updated_at"`
	DeletedAt    string  `orm:"deleted_at"`
	ApprovalTime string  `orm:"approval_time"` // 审批时间
}

type Product_link_auxiliary struct {
	Id               int    `orm:"id"`
	SurfaceTreatment string `orm:"surface_treatment"` // 表面处理
	Color            string `orm:"color"`             // 颜色
	Area             string `orm:"area"`              // 单面积(㎡/片)
	TotalCount       string `orm:"total_count"`       // 总面积(㎡)
	Flag             int    `orm:"flag"`
	CreatedAt        string `orm:"created_at"`
	UpdatedAt        string `orm:"updated_at"`
	DeletedAt        string `orm:"deleted_at"`
}

type Product_link_surface struct {
	Id               int     `orm:"id"`
	W1               float64 `orm:"w1"`
	W2               float64 `orm:"w2"`
	W3               float64 `orm:"w3"`
	W4               float64 `orm:"w4"`
	W5               float64 `orm:"w5"`
	W6               float64 `orm:"w6"`
	W7               float64 `orm:"w7"`
	W8               float64 `orm:"w8"`
	W9               float64 `orm:"w9"`
	H1               float64 `orm:"h1"`
	H2               float64 `orm:"h2"`
	H3               float64 `orm:"h3"`
	H4               float64 `orm:"h4"`
	H5               float64 `orm:"h5"`
	H6               float64 `orm:"h6"`
	H7               float64 `orm:"h7"`
	H8               float64 `orm:"h8"`
	H9               float64 `orm:"h9"`
	L1               float64 `orm:"l1"`
	L2               float64 `orm:"l2"`
	L3               float64 `orm:"l3"`
	L4               float64 `orm:"l4"`
	L5               float64 `orm:"l5"`
	L6               float64 `orm:"l6"`
	L7               float64 `orm:"l7"`
	L8               float64 `orm:"l8"`
	L9               float64 `orm:"l9"`
	WSize            int     `orm:"w_size"`            // 宽度数量
	HSize            int     `orm:"h_size"`            // 高度数量
	LSize            int     `orm:"l_size"`            // 长度数量
	SurfaceTreatment string  `orm:"surface_treatment"` // 表面处理
	Color            string  `orm:"color"`             // 颜色
	Area             string  `orm:"area"`              // 单面积
	TotalCount       string  `orm:"total_count"`       // 总面积
	Flag             int     `orm:"flag"`
	CreatedAt        string  `orm:"created_at"`
	UpdatedAt        string  `orm:"updated_at"`
	DeletedAt        string  `orm:"deleted_at"`
}

type Company struct {
	Id         int    `orm:"id"`
	Cuid       int    `orm:"cuid"`         // 注册人cuid
	Name       string `orm:"name"`         // 企业名称
	Mobile     string `orm:"mobile"`       // 企业手机号
	Tel        string `orm:"tel"`          // 电话号
	Address    string `orm:"address"`      // 企业地址
	Desc       string `orm:"desc"`         // 描述
	AuthPics   string `orm:"auth_pics"`    // 资质 多图
	VipLevel   int    `orm:"vip_level"`    // 企业购买等级
	VipEndTime int    `orm:"vip_end_time"` // 到期时间
	Status     int    `orm:"status"`       // 状态 0 停业  1营业 -1停用
	Flag       int    `orm:"flag"`
	CreatedAt  string `orm:"created_at"`
	UpdatedAt  string `orm:"updated_at"`
	DeletedAt  string `orm:"deleted_at"`
	CompanyKey string `orm:"company_key"` // 企业Key自动生成
}

type Platform_company struct {
	Id          int    `orm:"id"`
	CompanyId   int    `orm:"company_id"` // 企业id
	CompanyKey  string `orm:"company_key"`
	PlatformKey string `orm:"platform_key"` // 平台key
	PlatformUid string `orm:"platform_uid"` // 平台用户id
	Flag        int    `orm:"flag"`
	CreatedAt   string `orm:"created_at"`
	UpdatedAt   string `orm:"updated_at"`
	DeletedAt   string `orm:"deleted_at"`
	DataOrigin  string `orm:"data_origin"`
	Opt         string `orm:"opt"`
	SupplierId  string `orm:"supplier_id"`
}

type Product_sub struct {
	Id                int     `orm:"id"`
	MaterialName      string  `orm:"material_name"`      // 材料名称
	Attachment        string  `orm:"attachment"`         // 附件
	InstallMap        string  `orm:"install_map"`        // 安装示意图
	Price             float64 `orm:"price"`              // 价格
	Count             float64 `orm:"count"`              // 数量
	Unit              string  `orm:"unit"`               // 单位
	ReplenishmentFlag int     `orm:"replenishment_flag"` // 是否补货产品
	ConfigData        string  `orm:"config_data"`        // 自定义字段信息
	AppendAttachment  string  `orm:"append_attachment"`  // 附加的资源库信息
	Remark            string  `orm:"remark"`             // 备注
	Length            float64 `orm:"length"`             // 长
	Width             float64 `orm:"width"`              // 宽
	Height            float64 `orm:"height"`             // 厚
	Location          string  `orm:"location"`           // 安装位置
	PackCount         float64 `orm:"pack_count"`         // 打包数量
	SendCount         float64 `orm:"send_count"`         // 发货数量
	ReturnCount       float64 `orm:"return_count"`       // 退货数量
	ReceiveCount      float64 `orm:"receive_count"`      // 签收数量
	ProjectName       string  `orm:"project_name"`
	ProjectId         int     `orm:"project_id"`
	MaterialId        int     `orm:"material_id"` // 材料单id
	ProductId         int     `orm:"product_id"`
	PlatformKey       string  `orm:"platform_key"`
	PlatformUid       string  `orm:"platform_uid"`
	PlatformId        string  `orm:"platform_id"`
	ContractId        int     `orm:"contract_id"`
	Flag              int     `orm:"flag"`
	CreatedAt         string  `orm:"created_at"`
	UpdatedAt         string  `orm:"updated_at"`
	DeletedAt         string  `orm:"deleted_at"`
}

type Contract_config struct {
	Id                int     `orm:"id"`
	OtherAmount       float64 `orm:"other_amount"`        // 合同其他费用
	OtherAmountRemark string  `orm:"other_amount_remark"` // 合同其他费用说明
	AutoFlag          int     `orm:"auto_flag"`           // 是否默认按照自动生成材料总额计算
	PrePayFlag        int     `orm:"pre_pay_flag"`        // 是否有预付款
	PrePayPercent     float64 `orm:"pre_pay_percent"`     // 预付款百分比
	PrePayAmount      float64 `orm:"pre_pay_amount"`      // 按照预付款金额
	PrePayRemark      string  `orm:"pre_pay_remark"`      // 预付款说明
	ProgressByMonth   int     `orm:"progress_by_month"`   // 按月结,否则按批结
	ProgressPercent   float64 `orm:"progress_percent"`    // 进度款百分比
	ProgressRemark    string  `orm:"progress_remark"`     // 进度款说明
	AccountFlag       int     `orm:"account_flag"`        // 是否有结算款
	AccountPercent    float64 `orm:"account_percent"`     // 结算款比例
	AccountRemark     string  `orm:"account_remark"`      // 结算款备注
	RetentionFlag     int     `orm:"retention_flag"`      // 是否有质保金
	RetentionPercent  float64 `orm:"retention_percent"`   // 质保金比例
	RetentionRemark   string  `orm:"retention_remark"`    // 质保金备注
	Flag              int     `orm:"flag"`
	CreatedAt         string  `orm:"created_at"`
	UpdatedAt         string  `orm:"updated_at"`
	DeletedAt         string  `orm:"deleted_at"`
	ContractId        int     `orm:"contract_id"`
}

type Platform struct {
	Id                 int    `orm:"id"`
	PlatformName       string `orm:"platform_name"` // 平台名称
	PlatformKey        string `orm:"platform_key"`  // 平台key 完全标识
	PlatformSk         string `orm:"platform_sk"`
	PlatformSecret     string `orm:"platform_secret"` // 平台secret
	Status             int    `orm:"status"`          // 状态 1运行 0暂停维护 -1禁用
	PayName            string `orm:"pay_name"`
	PayAk              string `orm:"pay_ak"`
	PaySk              string `orm:"pay_sk"`
	Ak                 string `orm:"ak"`
	Sk                 string `orm:"sk"`
	PayNotifyUrl       string `orm:"pay_notify_url"` // 支付回调
	PayNotifyFunc      string `orm:"pay_notify_func"`
	MessageCallbackUrl string `orm:"message_callback_url"` // 消息中心回调地址
	Flag               int    `orm:"flag"`                 // 删除标识
	CreatedAt          string `orm:"created_at"`
	UpdatedAt          string `orm:"updated_at"`
	DeletedAt          string `orm:"deleted_at"`
}

type Depository struct {
	Id        int    `orm:"id"`
	Name      string `orm:"name"`       // 仓库名称
	Desc      string `orm:"desc"`       // 描述
	Address   string `orm:"address"`    // 仓库地址
	CompanyId int    `orm:"company_id"` // 仓库id
	Flag      int    `orm:"flag"`
	Status    int    `orm:"status"` // 状态 0停用 1正常
	CreatedAt string `orm:"created_at"`
	UpdatedAt string `orm:"updated_at"`
	DeletedAt string `orm:"deleted_at"`
}

type Mb struct {
	Id                      int     `orm:"id"`
	CreateDate              string  `orm:"create_date"`
	ModifyDate              string  `orm:"modify_date"`
	AcceptBindDate          string  `orm:"accept_bind_date"`
	Attendance              int     `orm:"attendance"`
	Bind                    int     `orm:"bind"`
	BindOpenid              string  `orm:"bind_openid"`
	DeleteFlag              int     `orm:"delete_flag"`
	Duties                  int     `orm:"duties"`
	FaceImageUrl            string  `orm:"face_image_url"`
	HeadImage               string  `orm:"head_image"`
	IdCardfmurl             string  `orm:"id_cardfmurl"`
	IdCardzmurl             string  `orm:"id_cardzmurl"`
	IdNo                    string  `orm:"id_no"`
	MemberType              int     `orm:"member_type"`
	Mobile                  string  `orm:"mobile"`
	Money                   float64 `orm:"money"`
	Nickname                string  `orm:"nickname"`
	WorkerType              string  `orm:"worker_type"`
	WorkingState            int     `orm:"working_state"`
	WorkingStateApproveDate string  `orm:"working_state_approve_date"`
	Admin                   int     `orm:"admin"`
	AdminProject            int     `orm:"admin_project"`
	MemberGroup             int     `orm:"member_group"`
	Signature               string  `orm:"signature"`
	Unionid                 string  `orm:"unionid"`
	Clouduid                string  `orm:"clouduid"`
	Feature                 string  `orm:"feature"`
}
