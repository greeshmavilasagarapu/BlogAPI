package main

import (
	"slices"
	"time"

	"github.com/google/uuid"
)

var ()

func createAuthor(
	name string,
) Author {
	id := uuid.New()
	created_at := time.Now()

	author := Author{
		Id:         id,
		Name:       name,
		Created_at: created_at,
	}

	authors = append(authors, author)

	return author
}
func createBlog(
	name, content string,
	author_id uuid.UUID,
) Blog {
	id := uuid.New()
	current_time := time.Now()
	blog := Blog{
		Id:         id,
		Name:       name,
		Content:    content,
		Author_id:  author_id,
		Created_at: current_time,
		Updated_at: current_time,
	}
	blogs = append(blogs, blog)
	return blog
}
func readauthor(
	id uuid.UUID,
) Author {
	for _, author := range authors {
		if author.Id == id {
			return author
		}
	}
	return Author{}
}
func readblog(
	id uuid.UUID,
) Blog {
	for _, blog := range blogs {
		if blog.Id == id {
			return blog
		}
	}
	return Blog{}
}
func readallauthors() []Author {
	return authors
}
func readallblogs() []Blog {
	time_tz, _ := time.LoadLocation("Asia/Kolkata")
	c_time := time.Now().In(time_tz)
	time_h := c_time.Format(time.RFC850)
	println("readallblogs called at: " + time_h)

	return blogs
}
func updateauthor(
	id uuid.UUID,
	name string,
) Author {
	for i, author := range authors {
		if author.Id == id {
			authors[i].Name = name
			return authors[i]
		}
	}
	return Author{}
}
func updateblog(
	id uuid.UUID,
	name string,
	content string,
) Blog {
	for i, blog := range blogs {
		if blog.Id == id {
			blogs[i].Name = name
			blogs[i].Content = content
			blogs[i].Updated_at = time.Now()
			return blogs[i]
		}
	}
	return Blog{}
}
func deleteauthor(
	id uuid.UUID,
) bool {
	for i, author := range authors {
		if author.Id == id {
			authors = slices.Delete(authors, i, i+1)
			return true
		}
	}
	return false
}
func deleteblog(
	id uuid.UUID,
) bool {
	for i, blog := range blogs {
		if blog.Id == id {
			blogs = slices.Delete(blogs, i, i+1)
			return true
		}
	}
	return false
}

func flushAuthors() {
	authors = []Author{}
}

func flushBlogs() {
	blogs = []Blog{}
}
func readBlogsByAuthor(authorID uuid.UUID) []Blog {
	var authorBlogs []Blog

	for _, blog := range blogs {
		if blog.Author_id == authorID {
			authorBlogs = append(authorBlogs, blog)
		}
	}

	return authorBlogs
}