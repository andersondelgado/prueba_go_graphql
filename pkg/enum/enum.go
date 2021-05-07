package enum

type Constant string

const (
	DbnameDefault               Constant = "dca2mvic9v6u8a"
	DialectDefault              Constant = "postgres"
	DbDefaultContext            Constant = "db_main_pg"
	RequestAuthorizationDefault Constant = "Authorization"
	GinContextKeyDefault        Constant = "GinContextKey"
	GinContextDBDefault         Constant = "GinContextDB"
	GinContextKeyAuthDefault    Constant = "GinContextKeyAuth"
	DefaultPort                 Constant = "8082"
)

type Router string

const (
	Api             Router = "/api"
	PlayGround      Router = "GraphQL"
	PlayGroundURL   Router = "/query"
	GraphqlURL      Router = "/query"
	GraphqlAuthURL  Router = "/query-auth"
	Users           Router = "users"
	UserElement     Router = "user-elements"
	UserElementType Router = "user-element-type"
	Profile         Router = "profiles"
	Roles           Router = "roles"
	Parameter       Router = "parameter"
	ParameterType   Router = "parameter-type"
	//Create          Router = "create"
	FindAll     Router = "findAll"
	SignIn      Router = "signIn"
	SignUp      Router = "signUp"
	EncryptTest Router = "encrypt"
)

type System string

const (
	DirViews System = "./views"
	DirFile  System = "./Images"
	PathFile System = "Images"
	Root     System = "/"
)

