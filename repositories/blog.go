package repositories

import (
	"slices"
	"time"

	"api/interfaces"

	"github.com/google/uuid"
)

func CreateAuthor(
	name string,
) interfaces.Author {
	id := uuid.New()
	created_at := time.Now()

	author := interfaces.Author{
		Id:         id,
		Name:       name,
		Created_at: created_at,
	}

	authors := ReadAllAuthors()
	authors = append(authors, author)

	storeAuthors(authors)

	return author
}
func CreateBlog(
	name, content string,
	author_id uuid.UUID,
) interfaces.Blog {
	id := uuid.New()
	current_time := time.Now()
	blog := interfaces.Blog{
		Id:         id,
		Name:       name,
		Content:    content,
		Author_id:  author_id,
		Created_at: current_time,
		Updated_at: current_time,
	}

	blogs := ReadAllBlogs()
	blogs = append(blogs, blog)

	storeBlogs(blogs)

	return blog
}
func ReadAuthor(
	id uuid.UUID,
) interfaces.Author {
	authors := ReadAllAuthors()
	for _, author := range authors {
		if author.Id == id {
			return author
		}
	}
	return interfaces.Author{}
}
func ReadBlog(
	id uuid.UUID,
) interfaces.Blog {
	blogs := ReadAllBlogs()
	for _, blog := range blogs {
		if blog.Id == id {
			return blog
		}
	}
	return interfaces.Blog{}
}
func ReadAllAuthors() []interfaces.Author {
	data := LoadData()
	return data.Authors
}
func ReadAllBlogs() []interfaces.Blog {
	// time_tz, _ := time.LoadLocation("Asia/Kolkata")
	// c_time := time.Now().In(time_tz)
	// time_h := c_time.Format(time.RFC850)
	// println("readallblogs called at: " + time_h)
	data := LoadData()
	return data.Blogs
}
func UpdateAuthor(
	id uuid.UUID,
	name string,
) interfaces.Author {
	authors := ReadAllAuthors()
	for i, author := range authors {
		if author.Id == id {
			authors[i].Name = name
			return authors[i]
		}
	}
	return interfaces.Author{}
}
func UpdateBlog(
	id uuid.UUID,
	name string,
	content string,
) interfaces.Blog {
	blogs := ReadAllBlogs()
	for i, blog := range blogs {
		if blog.Id == id {
			blogs[i].Name = name
			blogs[i].Content = content
			blogs[i].Updated_at = time.Now()
			return blogs[i]
		}
	}
	return interfaces.Blog{}
}
func deleteauthor(
	Id uuid.UUID,
) bool {
	authors := ReadAllAuthors()
	for i, author := range authors {
		if author.Id == Id {
			authors = slices.Delete(authors, i, i+1)
			storeAuthors(authors)
			return true
		}
	}
	return false
}
func deleteblog(
	Id uuid.UUID,
) bool {
	blogs := ReadAllBlogs()
	for i, blog := range blogs {
		if blog.Id == Id {
			blogs = slices.Delete(blogs, i, i+1)
			storeBlogs(blogs)
			return true
		}
	}
	return false
}

func flushAuthors() {

	//authors = []Author{}
	storeAuthors([]interfaces.Author{})
}

func flushBlogs() {
	//blogs = []Blog{}
	storeBlogs([]interfaces.Blog{})
}
func ReadBlogsByAuthor(authorID uuid.UUID) []interfaces.Blog {
	var authorBlogs []interfaces.Blog
	blogs := ReadAllBlogs()
	for _, blog := range blogs {
		if blog.Author_id == authorID {
			authorBlogs = append(authorBlogs, blog)
		}
	}

	return authorBlogs
}
func DeleteAuthor(Id uuid.UUID) bool {
	authors := ReadAllAuthors()
	for i, author := range authors {
		if author.Id == Id {
			authors = slices.Delete(authors, i, i+1)
			storeAuthors(authors)
			return true
		}
	}
	return false
}
func DeleteBlog(Id uuid.UUID) bool {
	blogs := ReadAllBlogs()
	for i, blog := range blogs {
		if blog.Id == Id {
			blogs = slices.Delete(blogs, i, i+1)
			storeBlogs(blogs)
			return true
		}
	}
	return false
}
func FlushAuthors() {
	storeAuthors([]interfaces.Author{})
}

func FlushBlogs() {
	storeBlogs([]interfaces.Blog{})
}
