export GOPATH=/Users/mac/Desktop/Code/learn/go

go mod init
go mod download
go run main.go

####  Work with REST API, JSON
https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/

# Request
https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/


#### Redis
redis-cli 


###### Docs
- Gin: https://github.com/gin-gonic/gin
- Gorm: https://gorm.io/docs/
- Redis: https://github.com/go-redis/redis
- Elasticsearch: https://github.com/olivere/elastic/wiki
- Mongo: https://github.com/mongodb/mongo-go-driver
- Job: https://github.com/jasonlvhit/gocron

## More
- Job:
    Define - Use Redis: https://gobuffalo.io/en/docs/workers#how-to-use-background-tasks
    Use Redis: https://github.com/gocraft/work
    Use Redis: https://github.com/benmanns/
    
    
- Elasticsearch:
```
	_search := fmt.Sprintf(`
	{
		"query": {
			"prefix": {
				"name": {
					"value": "%s"
				}
			}
		}
	}
	`, q)
	
	searchResult, err := elasticClient.Search().
		Index(index).
		Source(_search).
		Pretty(true).
		Do(ctx)
```