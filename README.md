# Logentries Go Client

Full fledged Go client for [Logentries](https://logentries.com/). The library supports CRUD operations for the following 
resources provided via the [Logentries REST Api](https://docs.logentries.com/docs/rest-api).

- [LogSets](https://docs.logentries.com/docs/logsets)
- [Logs](https://docs.logentries.com/docs/logs)
- [Tags](https://docs.logentries.com/docs/rest-tags)

The above resources are available in the client via its seamless easy-to-use interface and in a matter of few lines you
can have a working client ready to be used with Logentries.

# How to use the client?

Logentries Go Client is really easy to use. The client exposes multiple resources available in Logentries and
each of them offer create, read, update and delete (CRUD) operations.

Here is an example on how you can create a logentries client and query all the logsets under the account tight
to the API key which the client was configured with:

```

import (
	"github.com/dikhan/logentries_goclient"
)

func main() error {
	c := logentries_goclient.NewLogEntriesClient("LOGENTRIES_API_KEY")
	logSets, err := c.LogSets.GetLogSets()
	if err != nil {
	    return err
	}
	fmt.println(logSets)
}
```

## Contributing

- Fork it!
- Create your feature branch: git checkout -b my-new-feature
- Commit your changes: git commit -am 'Add some feature'
- Push to the branch: git push origin my-new-feature
- Submit a pull request :D

## Authors

Daniel I. Khan Ramiro

See also the list of [contributors](https://github.com/dikhan/logentries_goclient/graphs/contributors) who participated in this project.