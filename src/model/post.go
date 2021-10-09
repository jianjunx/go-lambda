package model

type PostAttributes struct {
	Id               int      `json:"id" dynamo:"id"`                                 // 文档编号
	Slug             string   `json:"slug" dynamo:"slug"`                             // 文档路径
	Title            string   `json:"title" dynamo:"title"`                           // 标题
	BookId           int      `json:"book_id" dynamo:"book_id"`                       // 仓库编号，就是 repo_id
	UserId           int      `json:"user_id" dynamo:"user_id"`                       // 用户/团队编号
	Format           string   `json:"format" dynamo:"format"`                         // 描述了正文的格式 [lake , markdown]
	Body             string   `json:"body" dynamo:"body"`                             // 正文 Markdown 源代码
	BodyHtml         string   `json:"body_html" dynamo:"body_html"`                   // 转换过后的正文 HTML
	CreatorId        int      `json:"creator_id" dynamo:"creator_id"`                 // 文档创建人 User Id
	Public           int      `json:"public" dynamo:"public"`                         // 公开级别 [0 - 私密, 1 - 公开]
	Status           int      `json:"status" dynamo:"-"`                              // 状态 [0 - 草稿, 1 - 发布]
	ContentUpdatedAt string   `json:"content_updated_at" dynamo:"content_updated_at"` // 文档内容更新时间
	DeletedAt        string   `json:"deleted_at" dynamo:"deleted_at"`                 // 删除时间，未删除为 null
	CreatedAt        string   `json:"created_at" dynamo:"created_at"`                 // 创建时间
	UpdatedAt        string   `json:"updated_at" dynamo:"updated_at"`                 // 更新时间
	WordCount        int      `json:"word_count" dynamo:"word_count"`                 // 字数
	Path             string   `json:"path" dynamo:"path"`                             // 文档的完整访问路径（不包括域名）
	ActionType       string   `json:"action_type" dynamo:"-"`                         // 值有 publish - 发布、 update - 更新、 delete - 删除
	Publish          bool     `json:"publish" dynamo:"-"`                             // 文档是否为第一次发布，第一次发布时为 true
	User             PostUser `json:"user" dynamo:"user"`                             //用户信息
	Book             PostBook `json:"book" dynamo:"book"`                             // Book 可以当成分类
}

type PostUser struct {
	Id          int    `json:"id" dynamo:"id"`
	Type        string `json:"type" dynamo:"type"`
	Login       string `json:"login" dynamo:"login"`
	Name        string `json:"name" dynamo:"name"`
	Description string `json:"description" dynamo:"description"`
	AvatarUrl   string `json:"avatar_url" dynamo:"avatar_url"`
}
type PostBook struct {
	Id          int    `json:"id" dynamo:"id"`
	Type        string `json:"type" dynamo:"type"`
	Slug        string `json:"slug" dynamo:"slug"`
	Name        string `json:"name" dynamo:"name"`
	UserId      int    `json:"user_id" dynamo:"user_id"`
	Description string `json:"description" dynamo:"description"`
	ItemsCount  int    `json:"items_count" dynamo:"items_count"` // 文章数量
}

type PostSearchParam struct {
	Page int
	Size int
	Book int
}
