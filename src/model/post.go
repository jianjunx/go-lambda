package model

type PostAttributes struct {
	Id               int      `json:"id" dynamodbav:"id"`                                 // 文档编号
	Slug             string   `json:"slug" dynamodbav:"slug"`                             // 文档路径
	Title            string   `json:"title" dynamodbav:"title"`                           // 标题
	BookId           int      `json:"book_id" dynamodbav:"book_id"`                       // 仓库编号，就是 repo_id
	UserId           int      `json:"user_id" dynamodbav:"user_id"`                       // 用户/团队编号
	Format           string   `json:"format" dynamodbav:"format"`                         // 描述了正文的格式 [lake , markdown]
	Body             string   `json:"body" dynamodbav:"body"`                             // 正文 Markdown 源代码
	BodyHtml         string   `json:"body_html" dynamodbav:"body_html"`                   // 转换过后的正文 HTML
	CreatorId        int      `json:"creator_id" dynamodbav:"creator_id"`                 // 文档创建人 User Id
	Public           int      `json:"public" dynamodbav:"public"`                         // 公开级别 [0 - 私密, 1 - 公开]
	Status           int      `json:"status" dynamodbav:"status"`                         // 状态 [0 - 草稿, 1 - 发布]
	ContentUpdatedAt string   `json:"content_updated_at" dynamodbav:"content_updated_at"` // 文档内容更新时间
	DeletedAt        string   `json:"deleted_at" dynamodbav:"deleted_at"`                 // 删除时间，未删除为 null
	CreatedAt        string   `json:"created_at" dynamodbav:"created_at"`                 // 创建时间
	UpdatedAt        string   `json:"updated_at" dynamodbav:"updated_at"`                 // 更新时间
	WordCount        int      `json:"word_count" dynamodbav:"word_count"`                 // 字数
	Path             string   `json:"path" dynamodbav:"path"`                             // 文档的完整访问路径（不包括域名）
	ActionType       string   `json:"action_type" dynamodbav:"action_type"`               // 值有 publish - 发布、 update - 更新、 delete - 删除
	Publish          bool     `json:"publish" dynamodbav:"publish"`                       // 文档是否为第一次发布，第一次发布时为 true
	User             PostUser `json:"user" dynamodbav:"user"`
	Book             PostBook `json:"book" dynamodbav:"book"`
}

type PostUser struct {
	Id          int    `json:"id" dynamodbav:"id"`
	Type        string `json:"type" dynamodbav:"type"`
	Login       string `json:"login" dynamodbav:"login"`
	Name        string `json:"name" dynamodbav:"name"`
	Description string `json:"description" dynamodbav:"description"`
	AvatarUrl   string `json:"avatar_url" dynamodbav:"avatar_url"`
}
type PostBook struct {
	Id          int    `json:"id" dynamodbav:"id"`
	Type        string `json:"type" dynamodbav:"type"`
	Slug        string `json:"slug" dynamodbav:"slug"`
	Name        string `json:"name" dynamodbav:"name"`
	UserId      int    `json:"user_id" dynamodbav:"user_id"`
	Description string `json:"description" dynamodbav:"description"`
	ItemsCount  int    `json:"items_count" dynamodbav:"items_count"` // 文章数量
}
