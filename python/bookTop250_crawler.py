import requests
import codecs
from bs4 import BeautifulSoup

URL = 'https://book.douban.com/top250'


def download_page(url):
    headers = {
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36'
    }
    return requests.get(url, headers=headers).content


def parse_html(html):
    soup = BeautifulSoup(html, "html.parser")
    book_list = soup.find(id='content').find('div', class_='indent')

    book_name_list = []

    for book in book_list.find_all('table'):
        print(book)
        detail = book.find('div', class_='pl2')
        book_name = detail.find('a')['title']
        book_name_list.append(book_name)

    next_page = soup.find('span', class_='next').find('a')
    if next_page:
        return book_name_list, next_page['href']
    return book_name_list, None


def main():
    url = URL

    with codecs.open('books', 'wb', encoding='utf-8') as fp:
        while url:
            print(url)
            html = download_page(url)
            books, url = parse_html(html)
            fp.write(u'{books}\n'.format(books='\n'.join(books)))


if __name__ == '__main__':
    main()
