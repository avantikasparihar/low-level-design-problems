package main

import . "github.com/avantikasparihar/low-level-design-problems/stack-overflow/internal"

func main() {
	PostsMgr.CreatePost()
	CommentsMgr.AddComment()
	CommentsMgr.ReplyComment()
}
