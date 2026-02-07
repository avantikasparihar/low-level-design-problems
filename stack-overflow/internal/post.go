package internal

var PostsMgr PostsManager

type Post struct {
	Id        int
	Title     string
	Content   string
	Comments  []int
	Upvotes   int
	Downvotes int
	Tags      []string
	observers []Observer
}

func (p *Post) register(obv Observer) {
	p.observers = append(p.observers, obv)
}

func (p *Post) notifyAll() {
	for _, obv := range p.observers {
		obv.UpdateVotes(p.Upvotes - p.Downvotes)
	}
}

type PostsManager interface {
	CreatePost()
	SearchPosts()
	DeletePosts()
	ContentManager
	VotesManager
}

type ContentManager interface {
	UpdateContent()
}

type VotesManager interface {
	Upvote()
	Downvote()
}
