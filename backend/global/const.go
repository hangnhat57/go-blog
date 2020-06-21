package global

const (
	dburl       = "mongodb+srv://gopher:123456a@cluster0-esppk.gcp.mongodb.net/go-blog?retryWrites=true&w=majority"
	dbname      = "go-blog"
	performance = 100
)

var (
	jwtSecret = []byte("blogSecret")
)
