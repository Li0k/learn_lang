import requests
import codecs
from bs4 import BeautifulSoup

URL = 'http://movie.douban.com/top250/'

def download_page(url):
    # print(requests.get(url))
    headers = {
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36'
    }
    data = requests.get(URL, headers=headers).content
    return data


def parse_html(html):
    soup = BeautifulSoup(html,"html.parser")
    movie_list = soup.find('ol', class_='grid_view')
    # print(movie_list)

    movie_name_list = []


    for movie in movie_list.find_all('li'):
        detail = movie.find('div', class_='hd')
        movie_name = detail.find('span', class_='title').getText()
        print (movie_name)
        movie_name_list.append(movie_name)

    next_page = soup.find('span',class_='next').find('a')
    if next_page:
        return movie_name_list, URL + next_page['href']
    return movie_name_list, None
        

def main():
    url = URL

    with codecs.open('movies', 'wb', encoding='utf-8') as fp:
        while url:
            html = download_page(url)
            movies, url = parse_html(html)
            fp.write(u'{movies}\n'.format(movies='\n'.join(movies)))

if __name__ == '__main__':
    main()
