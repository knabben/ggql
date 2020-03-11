GraphQL Deep Explorer
---

![Docker Image CI](https://github.com/knabben/ggql/workflows/Docker%20Image%20CI/badge.svg)

This Github Action is responsible to tracking, version and compare different versions of your
GraphQL schema in a release timeline through GraphQL introspection analysis.

From the features we can enumerate:

* Automatic artifact launch.
* Easy GraphQL schema fetch with introspection.
* Integration with your actual Github Workflow.
* Free schema versioning through sqlite3 storage.

Standalone
---

It is possible to use the system in a standalone mode, using the binary.

### Direct endpoint

Generating the dump via a GraphQL endpoint:

```
$ ./gql scrape --url https://www.example.com/graphql 
```

### Existent dump

Some libraries like Graphene can dump the schema, so it can be used in the CI without
the necessity of running the service.

```
$ python manage.py graphql_schema
$ ./gql scrape --file schema.json 
```


Scraping schema - GH Action
---

[knabben/gql-pull](https://github.com/knabben/gql-pull)
