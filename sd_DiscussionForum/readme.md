## Functional reqs
- Forum (reddit, quora, ...)
- Features
    - user logins to post, vote, comment
    - post fields
        - title
        - tags
        - body
    - user can comment
    - comments are flat
    - user can delete his own post/comment
    - user can upvote/downvote post/comment
        - user can have at max 1 upvote per post/comment
        - user can have at max 1 downvote per post/comment
    - most popular posts in last 24 hrs
    - show number of upvotes/downvotes of every post/comment

## Non-functional reqs
    - Scalability
    - Latency 
    - Availability
    - Fault tolerance

## Draft design
- post/comment has `author_id` that refs to `user_id` in `users` table
- comments has `post_id` that refs to a post in `posts` table
- `votes` table:
```js
// DESIGN 1: 
// pros: db foreign key
// cons: null value
{
    id: 'uuid',
    post_id: 'uuid NULL'
    comment_id: 'uuid NULL'
    type: 'UP DOWN'
}
// DESIGN 2: separate tables
{
    id: 'uuid'
    post_id: 'uuid'
    type: 'UP DOWN'
}
{
    id: 'uuid'
    comment_id: 'uuid'
    type: 'UP DOWN'
}
// DESIGN 3: polymorphism
// pros: easy to scale by adding types
// cons: no foreign key, need app logic to handle populate object_id
{
    id: 'uuid'
    object_id: 'uuid NOT NULL'
    object_type: 'comment/post'
    type: 'UP DOWN'
}
```
- `tags` table:
```js
// DESIGN 1: tags && tags mapping
// TABLE: tags
{
    value: 'TEXT PRIMARY KEY',
}
// TABLE: object_tags
{
    tag_id: 'uuid',
    object_id: 'uuid NOT NULL',
    object_type: 'comment/post'
}
```

## Q&A
- Why login uses POST instead of GET?