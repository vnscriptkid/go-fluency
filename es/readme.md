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

## Datasets
- https://files.grouplens.org/datasets/movielens/ml-latest-small-README.html