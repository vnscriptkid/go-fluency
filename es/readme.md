## ES ops
- Check connection
`curl -XGET http://localhost:9200`
- Create mapping
`curl -H 'Content-Type: application/json' -XPUT http://localhost:9200/shakespeare --data-binary @shakes-mapping.json`
- Feed data to mapping
`curl -H 'Content-Type: application/json' -XPOST http://localhost:9200/shakespeare/_bulk --data-binary @shakespeare_7.0.json`
- Delete an index
`curl -H 'Content-Type: application/json' -XDELETE http://localhost:9200/shakespeare`
- Create an index
```bash
curl -H 'Content-Type: application/json' -XPUT http://localhost:9200/shakespeare -d '{
    "mappings": {
        "properties": {
            "id": {"type": "integer"},
            "year": {"type": "date"},
            "genre": {"type": "keyword"},
            "title": {"type": "text", "analyzer": "english"}
        }
    }
}'
```

## Concepts
- Document
- Indices
- Inverted index
- TF/IDF
- Scaling
- Mappings
    - Field types
    - Field index
    - Field analyzer
- Analyzers
    - Character filters
    - Tokenizer
    - Token filter

- Analyzer options
    - Standard
    - Simple
    - Whitespace
    - Language
- Update doc using _version
- Concurrency
    - Optimistic concurrency control: 
        - using `_seq_no` + retries 
        - Manual way: `?if_seq_no=7&if_primary_term=1`
        - Rely on ES: `?retry_con_conflicts=5`
- Types
    - Keyword: exact match
    - Text: analyzers applied -> partial match
- Parent-children
    - `has_child`
    - `has_parent`
- Flattened data type
    - Why? Avoid mapping explosion
    - Keyword -> exact match

- Mappings
    - Explicit: predefined 
        - Potential exception
        - _settings: `"index.mapping.ignore_malformed": true` 
    - Dynamic -> mapping explosion
        - Solution: dead-letter queue

- Searching
    - `query > match > title : "hello"`
    - filters vs queries
        - filters: yes/no -> faster
        - queries: relevance
- Fuziness
- Partial matching
    - prefix (year): specific case of wildcard
    - wildcard
    - regex
- Search as you type
    - Query time
        - `match_phrase_prefix`
        - `slop`
    - Ngram
        - Create analyzer
    - Completion suggesters
    - `search_as_you_type` type

- Agg
    - ages: most frequently occurring values
        
![image](https://user-images.githubusercontent.com/28957748/215276148-336df223-6884-4624-9f5a-89b2539225c9.png)

![image](https://user-images.githubusercontent.com/28957748/215276112-4e07cd39-af44-40a4-8bd3-d92f8b512afb.png)

- Pagination
    - from (0-based idx), size
    - Have upperbound not to kill perf

- Sort
    - number
    - text field: use raw sub field of type keyword

## Datasets
- https://files.grouplens.org/datasets/movielens/ml-latest-small-README.html
