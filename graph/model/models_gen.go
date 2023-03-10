// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateTodoError interface {
	IsCreateTodoError()
}

type CreateUserError interface {
	IsCreateUserError()
}

type Error interface {
	IsError()
	GetMessage() string
}

type Node interface {
	IsNode()
	GetID() string
}

type UpdateTodoError interface {
	IsUpdateTodoError()
}

type CreateTodoInput struct {
	UserID  string `json:"userId"`
	Content string `json:"content"`
}

type CreateTodoPayload struct {
	Todo   *Todo             `json:"todo"`
	Errors []CreateTodoError `json:"errors"`
}

type CreateUserInput struct {
	Name string `json:"name"`
}

type CreateUserPayload struct {
	User   *User             `json:"user"`
	Errors []CreateUserError `json:"errors"`
}

type PageInfo struct {
	HasNextPage bool    `json:"hasNextPage"`
	EndCursor   *string `json:"endCursor"`
}

type Todo struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
	Owner   *User  `json:"owner"`
}

func (Todo) IsNode()            {}
func (this Todo) GetID() string { return this.ID }

type TodoConnection struct {
	Edges      []*TodoEdge `json:"edges"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

type TodoEdge struct {
	Node   *Todo  `json:"node"`
	Cursor string `json:"cursor"`
}

type TodoNotFoundError struct {
	Message string `json:"message"`
}

func (TodoNotFoundError) IsUpdateTodoError() {}

func (TodoNotFoundError) IsError()                {}
func (this TodoNotFoundError) GetMessage() string { return this.Message }

type UnknownUserError struct {
	Message string `json:"message"`
	UserID  string `json:"userId"`
}

func (UnknownUserError) IsCreateTodoError() {}

func (UnknownUserError) IsError()                {}
func (this UnknownUserError) GetMessage() string { return this.Message }

type UpdateTodoInput struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type UpdateTodoPayload struct {
	Todo   *Todo             `json:"todo"`
	Errors []UpdateTodoError `json:"errors"`
}

type User struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Todos *TodoConnection `json:"todos"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }

type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor string `json:"cursor"`
}

type UsernameAlreadyExistsError struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

func (UsernameAlreadyExistsError) IsCreateUserError() {}

func (UsernameAlreadyExistsError) IsError()                {}
func (this UsernameAlreadyExistsError) GetMessage() string { return this.Message }
