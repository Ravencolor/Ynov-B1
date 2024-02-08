
class Book:
    def __init__(self, nb_pages, author, isbn, marked_Page):
        self.nb_pages = nb_pages
        self.author = author
        self.isbn = isbn
        self.marked_Page = marked_Page

    def mark_page(self, page):
        if page < self.nb_pages:
            self.marked_Page = page
        else:
            print("Page number is out of range")


class Library:
    def __init__(self, books):
        self.books = books
    def list_authors(self):
        authors = [book.author for book in self.books]
        return authors
    
    def sum_of_marked_pages(self):
        sum = 0
        for book in self.books:
            sum += book.marked_Page
        return sum

"""
class Library:
    def __init__(self, books):
        self.books = books
    def list_authors(self):
        authors = []
        for book in self.books:
            authors.append(book.author)
        return authors

    def sum_of_marked_pages(self):
        sum = 0
        for book in self.books:
            sum += book.marked_Page
        return sum

"""

book1 = Book(100, "Author1", "123456789", 50)
book2 = Book(200, "Author2", "987654321", 100)
book3 = Book(300, "Author3", "123456789", 150)

library = Library([book1, book2, book3])

print("\n")

print(library.list_authors())
print(library.sum_of_marked_pages())

book1.mark_page(101)

print("\n")

print(book1.marked_Page)
print(book2.marked_Page)
print(book3.marked_Page)

print("\n")

print(library.books[0].marked_Page)
print(library.books[1].marked_Page)
print(library.books[2].marked_Page)

