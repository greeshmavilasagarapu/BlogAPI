package main

import (
	"slices"
	"time"

	"github.com/google/uuid"
)

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

	authors := readallauthors()
	authors = append(authors, author)

	storeAuthors(authors)

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

	blogs := readallblogs()
	blogs = append(blogs, blog)

	storeBlogs(blogs)

	return blog
}
func readauthor(
	id uuid.UUID,
) Author {
	authors := readallauthors()
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
	blogs := readallblogs()
	for _, blog := range blogs {
		if blog.Id == id {
			return blog
		}
	}
	return Blog{}
}
func readallauthors() []Author {
	data := LoadData()
	return data.Authors
}
func readallblogs() []Blog {
	// time_tz, _ := time.LoadLocation("Asia/Kolkata")
	// c_time := time.Now().In(time_tz)
	// time_h := c_time.Format(time.RFC850)
	// println("readallblogs called at: " + time_h)
	data := LoadData()
	return data.Blogs
}
func updateauthor(
	id uuid.UUID,
	name string,
) Author {
	authors := readallauthors()
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
	blogs := readallblogs()
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
	authors := readallauthors()
	for i, author := range authors {
		if author.Id == id {
			authors = slices.Delete(authors, i, i+1)
			storeAuthors(authors)
			return true
		}
	}
	return false
}
func deleteblog(
	id uuid.UUID,
) bool {
	blogs := readallblogs()
	for i, blog := range blogs {
		if blog.Id == id {
			blogs = slices.Delete(blogs, i, i+1)
			storeBlogs(blogs)
			return true
		}
	}
	return false
}

func flushAuthors() {

	//authors = []Author{}
	storeAuthors([]Author{})
}

func flushBlogs() {
	//blogs = []Blog{}
	storeBlogs([]Blog{})
}
func readBlogsByAuthor(authorID uuid.UUID) []Blog {
	var authorBlogs []Blog
	blogs := readallblogs()
	for _, blog := range blogs {
		if blog.Author_id == authorID {
			authorBlogs = append(authorBlogs, blog)
		}
	}

	return authorBlogs
}
