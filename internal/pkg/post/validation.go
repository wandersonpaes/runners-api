package post

import (
	"errors"
	"strings"
)

func (post *Posts) prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()
	return nil

}

func (post *Posts) validate() error {
	if post.Title == "" {
		return errors.New("the title is required and canot be empty")
	}
	if post.PostText == "" {
		return errors.New("the post text is required and canot be empty")
	}

	return nil
}

func (post *Posts) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.PostText = strings.TrimSpace(post.PostText)
}
