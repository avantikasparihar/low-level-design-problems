package internal

var CommentsMgr CommentsManager

type Comment struct {
	Id        int
	PostId    int
	UserId    int
	Content   string
	Replies   []int
	Upvotes   int
	Downvotes int
	observers []Observer
}

func (c *Comment) register(obv Observer) {
	c.observers = append(c.observers, obv)
}

func (c *Comment) notifyAll() {
	for _, obv := range c.observers {
		obv.UpdateVotes(c.Upvotes - c.Downvotes)
	}
}

type CommentsManager interface {
	AddComment()
	ReplyComment()
	DeleteComment()
	ContentManager
	VotesManager
}
